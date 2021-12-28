package persistence

import (
	"errors"
	f "fmt"

	uuid "github.com/google/uuid"
	primitive "go.mongodb.org/mongo-driver/bson/primitive"

	uuidHelpers "react-go-chat/lib/helpers"
)

var _ = primitive.A{}

const (
	USER_PREFIX      string = "user"
	MESSAGE_PREFIX   string = "msg"
	ROOM_PREFIX      string = "room"
	GROUP_PREFIX     string = "group"
	REACTION_PREFIX  string = "reaction"
	PREFIX_ERROR_MSG string = "INVALID_PREFIX"
	UUID_ERROR_MSG   string = "INVALID_ID"
)

type User struct {
	ID        Identifier `json:"_id,omitempty" bson:"_id,omitempty" primitive:"_id,omitempty"`
	FirstName string     `json:"first_name,omitempty" bson:"first_name,omitempty" primitive:"first_name,omitempty"`
	LastName  string     `json:"last_name,omitempty" bson:"last_name,omitempty" primitive:"last_name,omitempty"`
	Age       int        `json:"age,omitempty" bson:"age,omitempty" primitive:"age,omitempty"`
	EMailId   string     `json:"email_id,omitempty" bson:"email_id,omitempty" primitive:"email_id,omitempty"`
}

func (u *User) String() string {
	return f.Sprintf(
		"id: %s, first_name: %s, last_name: %s, Age: %d, e-Mail-id %s",
		u.ID, u.FirstName, u.LastName, u.Age, u.EMailId,
	)
}

func NewUser(
	firstName string,
	lastName string,
	age int,
	email_id string,
) (user *User) {
	identifier := Identifier{
		Prefix: USER_PREFIX,
		ID:     uuid.New().String(),
	}
	user = &User{
		ID:        identifier,
		FirstName: firstName,
		LastName:  lastName,
		Age:       age,
		EMailId:   email_id,
	}
	return
}

type Message struct {
	ID          Identifier   `json:"_id,omitempty" bson:"_id,omitempty" primitive:"_id,omitempty"`
	Sender      Identifier   `json:"sender,omitempty" bson:"sender,omitempty" primitive:"sender,omitempty"`
	Data        string       `json:"data,omitempty" bson:"data,omitempty" primitive:"data,omitempty"`
	RepliedTo   Identifier   `json:"replied_to,omitempty" bson:"replied_to,omitempty" primitive:"replied_to,omitempty"`
	ChatRoom    Identifier   `json:"chat_room,omitempty" bson:"chat_room,omitempty" primitive:"chat_room,omitempty"`
	GroupInRoom Identifier   `json:"group,omitempty" bson:"group,omitempty" primitive:"group,omitempty"`
	Reactions   []Identifier `json:"reactions,omitempty" bson:"reactions,omitempty" primitive:"reactions,omitempty"`
}

func NewMessage(
	identifier Identifier,
	sender Identifier,
	data string,
	repliedTo Identifier,
	chatRoom Identifier,
	groupInRoom Identifier,
	reactions []Identifier,
) (message *Message, err error) {

	if identifier.ID == "" {

		identifier = Identifier{
			Prefix: MESSAGE_PREFIX,
			ID:     uuid.New().String(),
		}
	} else {
		if identifier.Prefix != MESSAGE_PREFIX {
			return nil, errors.New(
				PREFIX_ERROR_MSG,
			)
		} else if valid := uuidHelpers.IsValidUUID(identifier.ID); !valid {
			return nil, errors.New(UUID_ERROR_MSG)
		}
	}

	message = &Message{
		ID:          identifier,
		Sender:      sender,
		Data:        data,
		RepliedTo:   repliedTo,
		ChatRoom:    chatRoom,
		GroupInRoom: groupInRoom,
		Reactions:   reactions,
	}
	return
}

func (m *Message) ReactToMessage(
	reaction_id Identifier,
) (err error) {
	if reaction_id.Prefix != REACTION_PREFIX {
		return errors.New(PREFIX_ERROR_MSG)
	}

	return
}

type ChatRoom struct {
	ID        Identifier   `json:"_id,omitempty" bson:"_id,omitempty" primitive:"_id,omitempty"`
	Name      string       `json:"name,omitempty" bson:"name,omitempty" primitive:"name,omitempty"`
	Admin     Identifier   `json:"admin,omitempty" bson:"admin,omitempty" primitive:"admin,omitempty"`
	CreatedBy Identifier   `json:"created_by,omitempty" bson:"created_by,omitempty" primitive:"created_by,omitempty"`
	Users     []Identifier `json:"users,omitempty" bson:"users,omitempty" primitive:"users,omitempty"`
	Groups    []Identifier `json:"groups,omitempty" bson:"groups,omitempty" primitive:"groups,omitempty"`
}

func NewChatRoom(
	name string,
	admin Identifier,
	created_by Identifier,
	users []Identifier,
) (
	chatRoom *ChatRoom, err error,
) {
	if (admin.Prefix != USER_PREFIX) || (created_by.Prefix != USER_PREFIX) {
		return nil, errors.New(PREFIX_ERROR_MSG)
	}

	return
}

type GroupInRoom struct {
	ID        Identifier   `json:"_id,omitempty" bson:"_id,omitempty" primitive:"_id,omitempty"`
	Name      string       `json:"name,omitempty" bson:"name,omitempty" primitive:"name,omitempty"`
	Admin     Identifier   `json:"admin,omitempty" bson:"admin,omitempty" primitive:"admin,omitempty"`
	CreatedBy Identifier   `json:"created_by,omitempty" bson:"created_by,omitempty" primitive:"created_by,omitempty"`
	Users     []Identifier `json:"users,omitempty" bson:"users,omitempty" primitive:"users,omitempty"`
}

type Reaction struct {
	ID   Identifier `json:"_id,omitempty" bson:"_id,omitempty" primitive:"_id,omitempty"`
	Name string     `json:"name,omitempty" bson:"name,omitempty" primitive:"name,omitempty"`
}

type Identifier struct {
	Prefix string `json:"prefix,omitempty" bson:"prefix,omitempty" primitive:"prefix,omitempty" default:"user"`
	ID     string `json:"_id,omitempty" bson:"_id,omitempty" primitive:"_id,omitempty"`
}
