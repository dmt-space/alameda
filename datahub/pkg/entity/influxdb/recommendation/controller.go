package recommendation

type controllerTag string
type controllerField string

const (
	ControllerTime      controllerTag = "time"
	ControllerNamespace controllerTag = "namespace"
	ControllerName      controllerTag = "name"

	ControllerType            controllerField = "type"
	ControllerCurrentReplicas controllerField = "current_replicas"
	ControllerDesiredReplicas controllerField = "desired_replicas"
	ControllerCreateTime      controllerField = "create_time"
)

var (
	// ControllerTags is list of tags of alameda_controller_recommendation measurement
	ControllerTags = []controllerTag{
		ControllerTime,
		ControllerNamespace,
		ControllerName,
	}
	// ControllerFields is list of fields of alameda_controller_recommendation measurement
	ControllerField = []controllerField{
		ControllerCurrentReplicas,
		ControllerDesiredReplicas,
		ControllerCreateTime,
		ControllerType,
	}
)
