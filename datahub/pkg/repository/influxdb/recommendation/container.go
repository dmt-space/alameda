package recommendation

import (
	"encoding/json"
	"fmt"
	"time"

	recommendation_entity "github.com/containers-ai/alameda/datahub/pkg/entity/influxdb/recommendation"
	"github.com/containers-ai/alameda/datahub/pkg/entity/influxdb/utils/enumconv"
	"github.com/containers-ai/alameda/datahub/pkg/repository/influxdb"
	"github.com/containers-ai/alameda/datahub/pkg/utils"
	"github.com/containers-ai/alameda/pkg/utils/log"
	datahub_v1alpha1 "github.com/containers-ai/api/alameda_api/v1alpha1/datahub"
	"github.com/golang/protobuf/ptypes/timestamp"
	influxdb_client "github.com/influxdata/influxdb/client/v2"
	"github.com/pkg/errors"
)

var (
	scope = log.RegisterScope("recommendation_db_measurement", "recommendation DB measurement", 0)
)

// ContainerRepository is used to operate node measurement of recommendation database
type ContainerRepository struct {
	influxDB *influxdb.InfluxDBRepository
}

// IsTag checks the column is tag or not
func (containerRepository *ContainerRepository) IsTag(column string) bool {
	for _, tag := range recommendation_entity.ContainerTags {
		if column == string(tag) {
			return true
		}
	}
	return false
}

// NewContainerRepository creates the ContainerRepository instance
func NewContainerRepository(influxDBCfg *influxdb.Config) *ContainerRepository {
	return &ContainerRepository{
		influxDB: &influxdb.InfluxDBRepository{
			Address:  influxDBCfg.Address,
			Username: influxDBCfg.Username,
			Password: influxDBCfg.Password,
		},
	}
}

// CreateContainerRecommendations add containers information container measurement
func (containerRepository *ContainerRepository) CreateContainerRecommendations(podRecommendations []*datahub_v1alpha1.PodRecommendation) error {
	points := []*influxdb_client.Point{}
	for _, podRecommendation := range podRecommendations {
		if podRecommendation.GetApplyRecommendationNow() {
			//TODO
		}

		podNS := podRecommendation.GetNamespacedName().GetNamespace()
		podName := podRecommendation.GetNamespacedName().GetName()
		containerRecommendations := podRecommendation.GetContainerRecommendations()
		topController := podRecommendation.GetTopController()

		for _, containerRecommendation := range containerRecommendations {
			tags := map[string]string{
				string(recommendation_entity.ContainerNamespace): podNS,
				string(recommendation_entity.ContainerPodName):   podName,
				string(recommendation_entity.ContainerName):      containerRecommendation.GetName(),
			}
			fields := map[string]interface{}{
				//TODO
				//string(recommendation_entity.ContainerPolicy):            "",
				string(recommendation_entity.ContainerTopControllerName): topController.GetNamespacedName().GetName(),
				string(recommendation_entity.ContainerTopControllerKind): enumconv.KindDisp[(topController.GetKind())],
			}
			initialLimitRecommendation := make(map[datahub_v1alpha1.MetricType]interface{})
			if containerRecommendation.GetInitialLimitRecommendations() != nil {
				for _, rec := range containerRecommendation.GetInitialLimitRecommendations() {
					// One and only one record in initial limit recommendation
					initialLimitRecommendation[rec.GetMetricType()] = rec.Data[0].NumValue
				}
			}
			initialRequestRecommendation := make(map[datahub_v1alpha1.MetricType]interface{})
			if containerRecommendation.GetInitialRequestRecommendations() != nil {
				for _, rec := range containerRecommendation.GetInitialRequestRecommendations() {
					// One and only one record in initial request recommendation
					initialRequestRecommendation[rec.GetMetricType()] = rec.Data[0].NumValue
				}
			}

			for _, metricData := range containerRecommendation.GetLimitRecommendations() {
				if data := metricData.GetData(); len(data) > 0 {
					for _, datum := range data {
						newFields := map[string]interface{}{}
						for key, value := range fields {
							newFields[key] = value
						}
						newFields[string(recommendation_entity.ContainerStartTime)] = datum.GetTime().GetSeconds()
						newFields[string(recommendation_entity.ContainerEndTime)] = datum.GetEndTime().GetSeconds()

						switch metricData.GetMetricType() {
						case datahub_v1alpha1.MetricType_CPU_USAGE_SECONDS_PERCENTAGE:
							if numVal, err := utils.StringToFloat64(datum.NumValue); err == nil {
								newFields[string(recommendation_entity.ContainerResourceLimitCPU)] = numVal
							}
							if value, ok := initialLimitRecommendation[datahub_v1alpha1.MetricType_CPU_USAGE_SECONDS_PERCENTAGE]; ok {
								if numVal, err := utils.StringToFloat64(value.(string)); err == nil {
									newFields[string(recommendation_entity.ContainerInitialResourceLimitCPU)] = numVal
								}
							} else {
								newFields[string(recommendation_entity.ContainerInitialResourceLimitCPU)] = 0
							}
						case datahub_v1alpha1.MetricType_MEMORY_USAGE_BYTES:
							if numVal, err := utils.StringToInt64(datum.NumValue); err == nil {
								newFields[string(recommendation_entity.ContainerResourceLimitMemory)] = numVal
							}
							if value, ok := initialLimitRecommendation[datahub_v1alpha1.MetricType_MEMORY_USAGE_BYTES]; ok {
								if numVal, err := utils.StringToInt64(value.(string)); err == nil {
									newFields[string(recommendation_entity.ContainerInitialResourceLimitMemory)] = numVal
								}
							} else {
								newFields[string(recommendation_entity.ContainerInitialResourceLimitMemory)] = 0
							}
						}

						if pt, err := influxdb_client.NewPoint(string(Container), tags, newFields, time.Unix(datum.GetTime().GetSeconds(), 0)); err == nil {
							points = append(points, pt)
						} else {
							scope.Error(err.Error())
						}
					}
				}
			}

			for _, metricData := range containerRecommendation.GetRequestRecommendations() {
				if data := metricData.GetData(); len(data) > 0 {
					for _, datum := range data {
						newFields := map[string]interface{}{}
						for key, value := range fields {
							newFields[key] = value
						}
						newFields[string(recommendation_entity.ContainerStartTime)] = datum.GetTime().GetSeconds()
						newFields[string(recommendation_entity.ContainerEndTime)] = datum.GetEndTime().GetSeconds()

						switch metricData.GetMetricType() {
						case datahub_v1alpha1.MetricType_CPU_USAGE_SECONDS_PERCENTAGE:
							if numVal, err := utils.StringToFloat64(datum.NumValue); err == nil {
								newFields[string(recommendation_entity.ContainerResourceRequestCPU)] = numVal
							}
							if value, ok := initialRequestRecommendation[datahub_v1alpha1.MetricType_CPU_USAGE_SECONDS_PERCENTAGE]; ok {
								if numVal, err := utils.StringToFloat64(value.(string)); err == nil {
									newFields[string(recommendation_entity.ContainerInitialResourceRequestCPU)] = numVal
								}
							} else {
								newFields[string(recommendation_entity.ContainerInitialResourceRequestCPU)] = 0
							}
						case datahub_v1alpha1.MetricType_MEMORY_USAGE_BYTES:
							if numVal, err := utils.StringToInt64(datum.NumValue); err == nil {
								newFields[string(recommendation_entity.ContainerResourceRequestMemory)] = numVal
							}
							if value, ok := initialRequestRecommendation[datahub_v1alpha1.MetricType_MEMORY_USAGE_BYTES]; ok {
								if numVal, err := utils.StringToInt64(value.(string)); err == nil {
									newFields[string(recommendation_entity.ContainerInitialResourceRequestMemory)] = numVal
								}
							} else {
								newFields[string(recommendation_entity.ContainerInitialResourceRequestMemory)] = 0
							}
						}
						if pt, err := influxdb_client.NewPoint(string(Container),
							tags, newFields,
							time.Unix(datum.GetTime().GetSeconds(), 0)); err == nil {
							points = append(points, pt)
						} else {
							scope.Error(err.Error())
						}
					}
				}
			}

			/*for _, metricData := range containerRecommendation.GetInitialLimitRecommendations() {
				if data := metricData.GetData(); len(data) > 0 {
					for _, datum := range data {
						newFields := map[string]interface{}{}
						for key, value := range fields {
							newFields[key] = value
						}
						newFields[string(recommendation_entity.ContainerStartTime)] = datum.GetTime().GetSeconds()
						newFields[string(recommendation_entity.ContainerEndTime)] = datum.GetEndTime().GetSeconds()

						switch metricData.GetMetricType() {
						case datahub_v1alpha1.MetricType_CPU_USAGE_SECONDS_PERCENTAGE:
							if numVal, err := utils.StringToFloat64(datum.NumValue); err == nil {
								newFields[string(recommendation_entity.ContainerInitialResourceLimitCPU)] = numVal
							}
						case datahub_v1alpha1.MetricType_MEMORY_USAGE_BYTES:
							if numVal, err := utils.StringToInt64(datum.NumValue); err == nil {
								newFields[string(recommendation_entity.ContainerInitialResourceLimitMemory)] = numVal
							}
						}
						if pt, err := influxdb_client.NewPoint(string(Container), tags, newFields, time.Unix(datum.GetTime().GetSeconds(), 0)); err == nil {
							points = append(points, pt)
						} else {
							scope.Error(err.Error())
						}
					}
				}
			}

			for _, metricData := range containerRecommendation.GetInitialRequestRecommendations() {
				if data := metricData.GetData(); len(data) > 0 {
					for _, datum := range data {
						newFields := map[string]interface{}{}
						for key, value := range fields {
							newFields[key] = value
						}
						newFields[string(recommendation_entity.ContainerStartTime)] = datum.GetTime().GetSeconds()
						newFields[string(recommendation_entity.ContainerEndTime)] = datum.GetEndTime().GetSeconds()

						switch metricData.GetMetricType() {
						case datahub_v1alpha1.MetricType_CPU_USAGE_SECONDS_PERCENTAGE:
							if numVal, err := utils.StringToFloat64(datum.NumValue); err == nil {
								newFields[string(recommendation_entity.ContainerInitialResourceRequestCPU)] = numVal
							}
						case datahub_v1alpha1.MetricType_MEMORY_USAGE_BYTES:
							if numVal, err := utils.StringToInt64(datum.NumValue); err == nil {
								newFields[string(recommendation_entity.ContainerInitialResourceRequestMemory)] = numVal
							}
						}
						if pt, err := influxdb_client.NewPoint(string(Container), tags, newFields, time.Unix(datum.GetTime().GetSeconds(), 0)); err == nil {
							points = append(points, pt)
						} else {
							scope.Error(err.Error())
						}
					}
				}
			}*/
		}
	}
	containerRepository.influxDB.WritePoints(points, influxdb_client.BatchPointsConfig{
		Database: string(influxdb.Recommendation),
	})
	return nil
}

// ListContainerRecommendations list container recommendations
func (containerRepository *ContainerRepository) ListContainerRecommendations(podNamespacedName *datahub_v1alpha1.NamespacedName,
	queryCondition *datahub_v1alpha1.QueryCondition,
	kind datahub_v1alpha1.Kind) ([]*datahub_v1alpha1.PodRecommendation, error) {

	podRecommendations := make([]*datahub_v1alpha1.PodRecommendation, 0)
	reqNS := podNamespacedName.GetNamespace()
	reqName := podNamespacedName.GetName()

	var (
		reqStartTime *timestamp.Timestamp
		reqEndTime   *timestamp.Timestamp
	)
	timeRange := queryCondition.GetTimeRange()
	if timeRange != nil {
		reqStartTime = timeRange.GetStartTime()
		reqEndTime = timeRange.GetEndTime()
	}

	whereStr := ""
	fieldToCompareRequestName := ""
	switch kind {
	case datahub_v1alpha1.Kind_POD:
		fieldToCompareRequestName = string(recommendation_entity.ContainerPodName)
	case datahub_v1alpha1.Kind_DEPLOYMENT:
		fieldToCompareRequestName = string(recommendation_entity.ContainerTopControllerName)
	case datahub_v1alpha1.Kind_DEPLOYMENTCONFIG:
		fieldToCompareRequestName = string(recommendation_entity.ContainerTopControllerName)
	default:
		return podRecommendations, errors.Errorf("no matching kind for Datahub Kind, received Kind: %s", datahub_v1alpha1.Kind_name[int32(kind)])
	}
	if reqNS != "" && reqName == "" {
		whereStr = fmt.Sprintf("WHERE \"%s\"='%s'", string(recommendation_entity.ContainerNamespace), reqNS)
	} else if reqNS == "" && reqName != "" {
		whereStr = fmt.Sprintf("WHERE \"%s\"='%s'", fieldToCompareRequestName, reqName)
	} else if reqNS != "" && reqName != "" {
		whereStr = fmt.Sprintf("WHERE \"%s\"='%s' AND \"%s\"='%s'", string(recommendation_entity.ContainerNamespace), reqNS, fieldToCompareRequestName, reqName)
	}

	timeConditionStr := ""
	if reqStartTime != nil && reqEndTime != nil {
		timeConditionStr = fmt.Sprintf("time >= %v AND time <= %v", utils.TimeStampToNanoSecond(reqStartTime), utils.TimeStampToNanoSecond(reqEndTime))
	} else if reqStartTime != nil && reqEndTime == nil {
		timeConditionStr = fmt.Sprintf("time >= %v", utils.TimeStampToNanoSecond(reqStartTime))
	} else if reqStartTime == nil && reqEndTime != nil {
		timeConditionStr = fmt.Sprintf("time <= %v", utils.TimeStampToNanoSecond(reqEndTime))
	}

	if whereStr == "" && timeConditionStr != "" {
		whereStr = fmt.Sprintf("WHERE %s", timeConditionStr)
	} else if whereStr != "" && timeConditionStr != "" {
		whereStr = fmt.Sprintf("%s AND %s", whereStr, timeConditionStr)
	}

	if kind != datahub_v1alpha1.Kind_POD {
		kindConditionStr := fmt.Sprintf("\"%s\"='%s'", string(recommendation_entity.ContainerTopControllerKind), enumconv.KindDisp[kind])
		if whereStr == "" {
			whereStr = fmt.Sprintf("WHERE %s", kindConditionStr)
		} else if whereStr != "" {
			whereStr = fmt.Sprintf("%s AND %s", whereStr, kindConditionStr)
		}
	}

	orderStr := containerRepository.buildOrderClause(queryCondition)
	limitStr := containerRepository.buildLimitClause(queryCondition)

	cmd := fmt.Sprintf("SELECT * FROM %s %s GROUP BY \"%s\",\"%s\",\"%s\" %s %s",
		string(Container), whereStr, recommendation_entity.ContainerName,
		recommendation_entity.ContainerNamespace, recommendation_entity.ContainerPodName, orderStr, limitStr)
	scope.Debugf(fmt.Sprintf("ListContainerRecommendations: %s", cmd))

	podRecommendations, err := containerRepository.queryRecommendation(cmd)
	if err != nil {
		return podRecommendations, err
	}

	return podRecommendations, nil

}

func (containerRepository *ContainerRepository) buildOrderClause(queryCondition *datahub_v1alpha1.QueryCondition) string {
	if queryCondition == nil {
		return "ORDER BY time ASC"
	}
	if queryCondition.GetOrder() == datahub_v1alpha1.QueryCondition_DESC {
		return "ORDER BY time DESC"
	} else if queryCondition.GetOrder() == datahub_v1alpha1.QueryCondition_ASC {
		return "ORDER BY time ASC"
	}
	return "ORDER BY time ASC"
}

func (containerRepository *ContainerRepository) buildLimitClause(queryCondition *datahub_v1alpha1.QueryCondition) string {
	if queryCondition == nil {
		return ""
	}
	limit := queryCondition.GetLimit()
	if queryCondition.GetLimit() > 0 {
		return fmt.Sprintf("LIMIT %v", limit)
	}
	return ""
}

func (c *ContainerRepository) ListAvailablePodRecommendations(in *datahub_v1alpha1.ListPodRecommendationsRequest) ([]*datahub_v1alpha1.PodRecommendation, error) {
	//podRecommendations := make([]*datahub_v1alpha1.PodRecommendation, 0)

	whereStrName := c.buildNameClause(in)
	whereStrKind := c.buildKindClause(in)
	whereStrTime := c.buildApplyTimeClause(in)

	whereStr := c.combineClause([]string{whereStrName, whereStrKind, whereStrTime})

	orderStr := c.buildOrderClause(in.QueryCondition)
	limitStr := c.buildLimitClause(in.QueryCondition)

	cmd := fmt.Sprintf("SELECT * FROM %s %s GROUP BY \"%s\",\"%s\",\"%s\" %s %s",
		string(Container), whereStr, recommendation_entity.ContainerName,
		recommendation_entity.ContainerNamespace, recommendation_entity.ContainerPodName, orderStr, limitStr)

	podRecommendations, err := c.queryRecommendation(cmd)
	if err != nil {
		return podRecommendations, err
	}

	return podRecommendations, nil
}

func (c *ContainerRepository) queryRecommendation(cmd string) ([]*datahub_v1alpha1.PodRecommendation, error) {
	podRecommendations := []*datahub_v1alpha1.PodRecommendation{}

	if results, err := c.influxDB.QueryDB(cmd, string(influxdb.Recommendation)); err == nil {
		for _, result := range results {
			//individual containers
			for _, ser := range result.Series {
				podName := ser.Tags[string(recommendation_entity.ContainerPodName)]
				podNS := ser.Tags[string(recommendation_entity.ContainerNamespace)]
				topControllerName := ""
				topControllerKind := datahub_v1alpha1.Kind_POD

				var startTime int64 = 0
				var endTime int64 = 0
				// per container time series data
				for _, val := range ser.Values {
					timeColIdx := utils.GetTimeIdxFromColumns(ser.Columns)
					timeObj, _ := utils.ParseTime(val[timeColIdx].(string))

					endTimeColIdx := utils.GetEndTimeIdxFromColumns(ser.Columns)
					ts, _ := val[endTimeColIdx].(json.Number).Int64()
					endTimeObj := time.Unix(ts, 0)

					containerRecommendation := &datahub_v1alpha1.ContainerRecommendation{
						Name: ser.Tags[string(recommendation_entity.ContainerName)],
						InitialLimitRecommendations:   []*datahub_v1alpha1.MetricData{},
						InitialRequestRecommendations: []*datahub_v1alpha1.MetricData{},
						LimitRecommendations:          []*datahub_v1alpha1.MetricData{},
						RequestRecommendations:        []*datahub_v1alpha1.MetricData{},
					}
					initialResourceLimitCPUData := []*datahub_v1alpha1.Sample{}
					initialResourceRequestCPUData := []*datahub_v1alpha1.Sample{}
					resourceLimitCPUData := []*datahub_v1alpha1.Sample{}
					resourceRequestCPUData := []*datahub_v1alpha1.Sample{}
					initialResourceLimitMemoryData := []*datahub_v1alpha1.Sample{}
					initialResourceRequestMemoryData := []*datahub_v1alpha1.Sample{}
					resourceLimitMemoryData := []*datahub_v1alpha1.Sample{}
					resourceRequestMemoryData := []*datahub_v1alpha1.Sample{}

					for columnIdx, column := range ser.Columns {
						if val[columnIdx] == nil {
							continue
						}

						if column == string(recommendation_entity.ContainerInitialResourceLimitCPU) {
							colVal := val[columnIdx].(json.Number).String()
							sampleObj := utils.GetSampleInstance(&timeObj, &endTimeObj, colVal)
							initialResourceLimitCPUData = append(initialResourceLimitCPUData, sampleObj)
						} else if column == string(recommendation_entity.ContainerInitialResourceRequestCPU) {
							colVal := val[columnIdx].(json.Number).String()
							sampleObj := utils.GetSampleInstance(&timeObj, &endTimeObj, colVal)
							initialResourceRequestCPUData = append(initialResourceRequestCPUData, sampleObj)
						} else if column == string(recommendation_entity.ContainerResourceLimitCPU) {
							colVal := val[columnIdx].(json.Number).String()
							sampleObj := utils.GetSampleInstance(&timeObj, &endTimeObj, colVal)
							resourceLimitCPUData = append(resourceLimitCPUData, sampleObj)
						} else if column == string(recommendation_entity.ContainerResourceRequestCPU) {
							colVal := val[columnIdx].(json.Number).String()
							sampleObj := utils.GetSampleInstance(&timeObj, &endTimeObj, colVal)
							resourceRequestCPUData = append(resourceRequestCPUData, sampleObj)
						} else if column == string(recommendation_entity.ContainerInitialResourceLimitMemory) {
							colVal := val[columnIdx].(json.Number).String()
							sampleObj := utils.GetSampleInstance(&timeObj, &endTimeObj, colVal)
							initialResourceLimitMemoryData = append(initialResourceLimitMemoryData, sampleObj)
						} else if column == string(recommendation_entity.ContainerInitialResourceRequestMemory) {
							colVal := val[columnIdx].(json.Number).String()
							sampleObj := utils.GetSampleInstance(&timeObj, &endTimeObj, colVal)
							initialResourceRequestMemoryData = append(initialResourceRequestMemoryData, sampleObj)
						} else if column == string(recommendation_entity.ContainerResourceLimitMemory) {
							colVal := val[columnIdx].(json.Number).String()
							sampleObj := utils.GetSampleInstance(&timeObj, &endTimeObj, colVal)
							resourceLimitMemoryData = append(resourceLimitMemoryData, sampleObj)
						} else if column == string(recommendation_entity.ContainerResourceRequestMemory) {
							colVal := val[columnIdx].(json.Number).String()
							sampleObj := utils.GetSampleInstance(&timeObj, &endTimeObj, colVal)
							resourceRequestMemoryData = append(resourceRequestMemoryData, sampleObj)
						} else if column == string(recommendation_entity.ContainerStartTime) {
							startTime, _ = val[columnIdx].(json.Number).Int64()
						} else if column == string(recommendation_entity.ContainerEndTime) {
							endTime, _ = val[columnIdx].(json.Number).Int64()
						} else if column == string(recommendation_entity.ContainerTopControllerName) {
							topControllerName = val[columnIdx].(string)
						} else if column == string(recommendation_entity.ContainerTopControllerKind) {
							topControllerKind = enumconv.KindEnum[val[columnIdx].(string)]
						}
					}
					if len(initialResourceLimitCPUData) > 0 {
						containerRecommendation.InitialLimitRecommendations = append(containerRecommendation.InitialLimitRecommendations,
							&datahub_v1alpha1.MetricData{
								MetricType: datahub_v1alpha1.MetricType_CPU_USAGE_SECONDS_PERCENTAGE,
								Data:       initialResourceLimitCPUData,
							})
					}
					if len(initialResourceLimitMemoryData) > 0 {
						containerRecommendation.InitialLimitRecommendations = append(containerRecommendation.InitialLimitRecommendations,
							&datahub_v1alpha1.MetricData{
								MetricType: datahub_v1alpha1.MetricType_MEMORY_USAGE_BYTES,
								Data:       initialResourceLimitMemoryData,
							})
					}
					if len(initialResourceRequestCPUData) > 0 {
						containerRecommendation.InitialRequestRecommendations = append(containerRecommendation.InitialRequestRecommendations,
							&datahub_v1alpha1.MetricData{
								MetricType: datahub_v1alpha1.MetricType_CPU_USAGE_SECONDS_PERCENTAGE,
								Data:       initialResourceRequestCPUData,
							})
					}
					if len(initialResourceRequestMemoryData) > 0 {
						containerRecommendation.InitialRequestRecommendations = append(containerRecommendation.InitialRequestRecommendations,
							&datahub_v1alpha1.MetricData{
								MetricType: datahub_v1alpha1.MetricType_MEMORY_USAGE_BYTES,
								Data:       initialResourceRequestMemoryData,
							})
					}
					if len(resourceLimitCPUData) > 0 {
						containerRecommendation.LimitRecommendations = append(containerRecommendation.LimitRecommendations,
							&datahub_v1alpha1.MetricData{
								MetricType: datahub_v1alpha1.MetricType_CPU_USAGE_SECONDS_PERCENTAGE,
								Data:       resourceLimitCPUData,
							})
					}
					if len(resourceLimitMemoryData) > 0 {
						containerRecommendation.LimitRecommendations = append(containerRecommendation.LimitRecommendations,
							&datahub_v1alpha1.MetricData{
								MetricType: datahub_v1alpha1.MetricType_MEMORY_USAGE_BYTES,
								Data:       resourceLimitMemoryData,
							})
					}
					if len(resourceRequestCPUData) > 0 {
						containerRecommendation.RequestRecommendations = append(containerRecommendation.RequestRecommendations,
							&datahub_v1alpha1.MetricData{
								MetricType: datahub_v1alpha1.MetricType_CPU_USAGE_SECONDS_PERCENTAGE,
								Data:       resourceRequestCPUData,
							})
					}
					if len(resourceRequestMemoryData) > 0 {
						containerRecommendation.RequestRecommendations = append(containerRecommendation.RequestRecommendations,
							&datahub_v1alpha1.MetricData{
								MetricType: datahub_v1alpha1.MetricType_MEMORY_USAGE_BYTES,
								Data:       resourceRequestMemoryData,
							})
					}

					foundPodRec := false
					for podRecommendationIdx, podRecommendation := range podRecommendations {
						if podRecommendation.GetStartTime() != nil && startTime != 0 && podRecommendation.GetStartTime().GetSeconds() == startTime &&
							podRecommendation.GetEndTime() != nil && endTime != 0 && podRecommendation.GetEndTime().GetSeconds() == endTime &&
							podRecommendation.GetNamespacedName().GetNamespace() == podNS && podRecommendation.GetNamespacedName().GetName() == podName {
							foundPodRec = true
							podRecommendations[podRecommendationIdx].ContainerRecommendations = append(podRecommendations[podRecommendationIdx].ContainerRecommendations, containerRecommendation)
							if startTime != 0 {
								podRecommendations[podRecommendationIdx].StartTime = &timestamp.Timestamp{
									Seconds: startTime,
								}
							}
							if endTime != 0 {
								podRecommendations[podRecommendationIdx].EndTime = &timestamp.Timestamp{
									Seconds: endTime,
								}
							}
						}
					}
					if !foundPodRec {
						podRec := &datahub_v1alpha1.PodRecommendation{
							NamespacedName: &datahub_v1alpha1.NamespacedName{
								Namespace: podNS,
								Name:      podName,
							},
							ContainerRecommendations: []*datahub_v1alpha1.ContainerRecommendation{
								containerRecommendation,
							},
							TopController: &datahub_v1alpha1.TopController{
								NamespacedName: &datahub_v1alpha1.NamespacedName{
									Namespace: podNS,
									Name:      topControllerName,
								},
								Kind: topControllerKind,
							},
						}
						if startTime != 0 {
							podRec.StartTime = &timestamp.Timestamp{
								Seconds: startTime,
							}
						}
						if endTime != 0 {
							podRec.EndTime = &timestamp.Timestamp{
								Seconds: endTime,
							}
						}
						podRecommendations = append(podRecommendations, podRec)
					}
				}
			}
		}
		return podRecommendations, nil
	} else {
		return podRecommendations, err
	}
}

func (c *ContainerRepository) combineClause(strList []string) string {
	ret := ""
	whereFlag := false

	for _, value := range strList {
		if value != "" && whereFlag == false {
			ret = fmt.Sprintf("WHERE %s", value)
			whereFlag = true
		} else if value != "" {
			ret += fmt.Sprintf(" AND %s", value)
		}
	}

	return ret
}

func (c *ContainerRepository) buildNameClause(in *datahub_v1alpha1.ListPodRecommendationsRequest) string {
	ret := ""
	namespace := in.GetNamespacedName().GetNamespace()
	if namespace == "" {
		return ret
	}

	ret = fmt.Sprintf(" \"namespace\"='%s'", namespace)
	return ret
}

func (c *ContainerRepository) buildKindClause(in *datahub_v1alpha1.ListPodRecommendationsRequest) string {
	ret := ""
	col := ""

	name := in.GetNamespacedName().GetName()
	if name == "" {
		return ret
	}

	kind := in.GetKind()

	switch kind {
	case datahub_v1alpha1.Kind_POD:
		col = string(recommendation_entity.ContainerPodName)
	case datahub_v1alpha1.Kind_DEPLOYMENT:
		col = string(recommendation_entity.ContainerTopControllerName)
	case datahub_v1alpha1.Kind_DEPLOYMENTCONFIG:
		col = string(recommendation_entity.ContainerTopControllerName)
	default:
		return ""
	}

	ret = fmt.Sprintf(" \"%s\"='%s'", col, name)
	return ret
}

func (c *ContainerRepository) buildApplyTimeClause(in *datahub_v1alpha1.ListPodRecommendationsRequest) string {
	ret := ""

	applyTime := in.GetQueryCondition().GetTimeRange().GetApplyTime().GetSeconds()
	if applyTime > 0 {
		ret = fmt.Sprintf(" \"end_time\">=%d AND \"start_time\"<=%d", applyTime, applyTime)
	}

	return ret
}
