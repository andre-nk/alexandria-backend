package activity

type Service interface {
	CreateActivity(activity Activity) (Activity, error)
	GetActivityByAffiliateID(id string) ([]Activity, error)
	GetActivityByID(id string) (Activity, error)
	MarkActivityAsRead(id string) error
}

type service struct {
	repository Repository
}

func NewService(repository Repository) *service {
	return &service{repository}
}

func (service *service) CreateActivity(activity Activity) (Activity, error) {
	newActivity, err := service.repository.CreateActivity(activity)
	if err != nil {
		return newActivity, err
	}

	return newActivity, nil
}

func (service *service) GetActivityByAffiliateID(id string) ([]Activity, error) {
	activities, err := service.repository.GetActivityByAffiliateID(id)
	if err != nil {
		return activities, err
	}

	return activities, nil
}

func (service *service) GetActivityByID(id string) (Activity, error) {
	activity, err := service.repository.GetActivityByID(id)
	if err != nil {
		return activity, err
	}

	return activity, nil
}

func (service *service) MarkActivityAsRead(id string) error {
	err := service.repository.MarkActivityAsRead(id)
	if err != nil {
		return err
	}

	return nil
}
