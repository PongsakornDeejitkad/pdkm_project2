package repository

import (
	"log"
	"order-management/domain"
	"order-management/entity"

	"gorm.io/gorm"
)

type customerRepository struct {
	db *gorm.DB
}

func NewRepository(db *gorm.DB) domain.CustomerRepository {
	return &customerRepository{
		db: db,
	}
}

type adminRepository struct {
	db *gorm.DB
}

func (r *customerRepository) CreateCustomer(customer entity.Customer) error {
	if err := r.db.Create(&customer).Error; err != nil {
		log.Println("Create Customer error: ", err)
		return err
	}
	return nil
}

func (r *customerRepository) ListCustomers() ([]entity.Customer, error) {
	customers := []entity.Customer{}
	if err := r.db.Find(&customers).Error; err != nil {
		log.Println("List customers error: ", err)
		return nil, err
	}
	return customers, nil
}

func (r *customerRepository) GetCustomer(id int) (*entity.Customer, error) {
	customer := &entity.Customer{}
	if err := r.db.First(&customer, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Customer not found")
			return nil, err
		}
		log.Println("Error retrieving customer:", err)
		return nil, err
	}
	return customer, nil

}

func (r *customerRepository) DeleteCustomer(id int) error {
	existingCustomer := &entity.Customer{}
	if err := r.db.First(&existingCustomer, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Customer not found")
			return err
		}
		log.Println("Error retrieving Customer:", err)
		return err
	}

	if err := r.db.Model(&entity.Customer{}).Where("id = ?", id).Delete(&existingCustomer).Error; err != nil {
		log.Println("Error delete customer:", err)
		return err
	}
	return nil

}

func (r *customerRepository) UpdateCustomer(id int, customer entity.Customer) error {
	existingCustomer := &entity.Customer{}
	if err := r.db.First(&existingCustomer, id).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Customer not found")
			return err
		}
		log.Println("Error retrieving customer:", err)
		return err
	}
	if err := r.db.Model(&entity.Admin{}).Where("id = ?", id).Updates(&customer).Error; err != nil {
		log.Println("Error updating customer:", err)
		return err
	}
	return nil
}

func (r *customerRepository) GetCustomerByEmail(email string) (entity.Customer, error) {
	customer := entity.Customer{}

	if err := r.db.Where("email = ?", email).First(&customer).Error; err != nil {
		if err == gorm.ErrRecordNotFound {
			log.Println("Customer not found")
			return customer, err
		}

		log.Println("GetCustomerByEmail error:", err)
		return customer, err
	}


	return customer, nil
}
