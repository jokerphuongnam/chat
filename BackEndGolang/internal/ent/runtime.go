// Code generated by ent, DO NOT EDIT.

package ent

import (
	"chat-backend/internal/ent/authorize"
	"chat-backend/internal/ent/member"
	"chat-backend/internal/ent/message"
	"chat-backend/internal/ent/room"
	"chat-backend/internal/ent/roominfo"
	"chat-backend/internal/ent/schema"
	"chat-backend/internal/ent/user"
	"chat-backend/internal/ent/usernamepassword"

	"github.com/google/uuid"
)

// The init function reads all schema descriptors with runtime code
// (default values, validators, hooks and policies) and stitches it
// to their package variables.
func init() {
	authorizeFields := schema.Authorize{}.Fields()
	_ = authorizeFields
	// authorizeDescToken is the schema descriptor for token field.
	authorizeDescToken := authorizeFields[1].Descriptor()
	// authorize.TokenValidator is a validator for the "token" field. It is called by the builders before save.
	authorize.TokenValidator = authorizeDescToken.Validators[0].(func(string) error)
	// authorizeDescID is the schema descriptor for id field.
	authorizeDescID := authorizeFields[0].Descriptor()
	// authorize.DefaultID holds the default value on creation for the id field.
	authorize.DefaultID = authorizeDescID.Default.(func() uuid.UUID)
	memberFields := schema.Member{}.Fields()
	_ = memberFields
	// memberDescNickName is the schema descriptor for nick_name field.
	memberDescNickName := memberFields[3].Descriptor()
	// member.NickNameValidator is a validator for the "nick_name" field. It is called by the builders before save.
	member.NickNameValidator = func() func(string) error {
		validators := memberDescNickName.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
		}
		return func(nick_name string) error {
			for _, fn := range fns {
				if err := fn(nick_name); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	messageFields := schema.Message{}.Fields()
	_ = messageFields
	// messageDescID is the schema descriptor for id field.
	messageDescID := messageFields[0].Descriptor()
	// message.DefaultID holds the default value on creation for the id field.
	message.DefaultID = messageDescID.Default.(func() uuid.UUID)
	roomFields := schema.Room{}.Fields()
	_ = roomFields
	// roomDescColor is the schema descriptor for color field.
	roomDescColor := roomFields[1].Descriptor()
	// room.ColorValidator is a validator for the "color" field. It is called by the builders before save.
	room.ColorValidator = roomDescColor.Validators[0].(func(string) error)
	// roomDescID is the schema descriptor for id field.
	roomDescID := roomFields[0].Descriptor()
	// room.DefaultID holds the default value on creation for the id field.
	room.DefaultID = roomDescID.Default.(func() uuid.UUID)
	roominfoFields := schema.RoomInfo{}.Fields()
	_ = roominfoFields
	// roominfoDescID is the schema descriptor for id field.
	roominfoDescID := roominfoFields[0].Descriptor()
	// roominfo.DefaultID holds the default value on creation for the id field.
	roominfo.DefaultID = roominfoDescID.Default.(func() uuid.UUID)
	userFields := schema.User{}.Fields()
	_ = userFields
	// userDescName is the schema descriptor for name field.
	userDescName := userFields[1].Descriptor()
	// user.NameValidator is a validator for the "name" field. It is called by the builders before save.
	user.NameValidator = userDescName.Validators[0].(func(string) error)
	// userDescID is the schema descriptor for id field.
	userDescID := userFields[0].Descriptor()
	// user.DefaultID holds the default value on creation for the id field.
	user.DefaultID = userDescID.Default.(func() uuid.UUID)
	usernamepasswordFields := schema.UsernamePassword{}.Fields()
	_ = usernamepasswordFields
	// usernamepasswordDescUsername is the schema descriptor for username field.
	usernamepasswordDescUsername := usernamepasswordFields[1].Descriptor()
	// usernamepassword.UsernameValidator is a validator for the "username" field. It is called by the builders before save.
	usernamepassword.UsernameValidator = func() func(string) error {
		validators := usernamepasswordDescUsername.Validators
		fns := [...]func(string) error{
			validators[0].(func(string) error),
			validators[1].(func(string) error),
			validators[2].(func(string) error),
		}
		return func(username string) error {
			for _, fn := range fns {
				if err := fn(username); err != nil {
					return err
				}
			}
			return nil
		}
	}()
	// usernamepasswordDescID is the schema descriptor for id field.
	usernamepasswordDescID := usernamepasswordFields[0].Descriptor()
	// usernamepassword.DefaultID holds the default value on creation for the id field.
	usernamepassword.DefaultID = usernamepasswordDescID.Default.(func() uuid.UUID)
}
