package data

import (
	"errors"
	"github.com/google/uuid"
	"github.com/tnyidea/react-admin-dataprovider/go/types"
)

func (d *DB) Count() (int, error) {
	users, err := d.FindAllUsers()
	if err != nil {
		return 0, err
	}
	return len(users), nil
}

func (d *DB) CreateUser(user types.User) error {

	tx := d.memDB.Txn(true)
	defer tx.Abort()

	if user.UUID == "" {
		user.UUID = uuid.NewString()
	}

	err := tx.Insert("user", user)
	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (d *DB) FindAllUsers() ([]types.User, error) {
	tx := d.memDB.Txn(false)
	defer tx.Abort()

	result, err := tx.Get("user", "id")
	if err != nil {
		return nil, err
	}

	var users []types.User
	for v := result.Next(); v != nil; v = result.Next() {
		user := v.(types.User)
		users = append(users, user)
	}

	tx.Commit()

	return users, nil
}

func (d *DB) FindUserByUUID(uuid string) (types.User, error) {
	tx := d.memDB.Txn(false)
	defer tx.Abort()

	v, err := tx.First("user", "id", uuid)
	if err != nil {
		return types.User{}, err
	}

	tx.Commit()

	if v == nil {
		return types.User{}, errors.New("error: UUID " + uuid + " not found")
	}

	return v.(types.User), nil
}

func (d *DB) UpdateUser(v types.User) error {
	return d.CreateUser(v)
}

func (d *DB) DeleteAllUsers() error {
	tx := d.memDB.Txn(true)
	defer tx.Abort()

	_, err := tx.DeleteAll("user", "id")
	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}

func (d *DB) DeleteUserByUUID(uuid string) error {
	tx := d.memDB.Txn(true)
	defer tx.Abort()

	err := tx.Delete("user", types.User{
		UUID: uuid,
	})
	if err != nil {
		return err
	}

	tx.Commit()

	return nil
}
