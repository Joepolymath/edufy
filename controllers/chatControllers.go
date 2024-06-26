package controllers

// import (
// 	"Learnium/adapters"
// 	"Learnium/database"
// 	"Learnium/logger"
// 	"Learnium/models"
// 	"Learnium/serializers"
// 	"context"
// 	"encoding/json"
// 	"github.com/gofiber/contrib/websocket"
// 	"github.com/google/uuid"
// 	"go.uber.org/zap"
// 	"gorm.io/gorm"
// 	"strconv"
// 	"strings"
// 	"time"
// )

// func sendErrorToClient(c *websocket.Conn, db *gorm.DB, errorMessage string, conversationUserID uuid.UUID) {

// 	if conversationUserID != uuid.Nil {
// 		err := db.Model(&models.ConversationUser{}).Where("id =?", conversationUserID).Update("is_online", false).Error
// 		if err != nil {
// 			if err != gorm.ErrRecordNotFound {
// 				logger.Error(context.Background(), "Error updating conversation user to offline", zap.Error(err))
// 			}
// 		}
// 	}

// 	// Create a JSON message with an error field

// 	errorMsg := map[string]string{"error": errorMessage}
// 	// Send the error message as JSON to the client
// 	if err := c.WriteJSON(errorMsg); err != nil {
// 		logger.Error(context.Background(), "Error writing json when closing socket")
// 	}
// 	err := c.WriteMessage(websocket.CloseMessage, websocket.FormatCloseMessage(websocket.CloseNormalClosure, ""))
// 	if err != nil {
// 		logger.Error(context.Background(), "error closing with message", zap.Error(err))
// 	} // Perform websocket close handshake
// 	err = c.Close()
// 	if err != nil {
// 		logger.Error(context.Background(), "error closing connection", zap.Error(err))
// 	} // Close the connection after sending the error message
// }

// type ConversationUserListMessageType struct {
// 	Type   string `json:"type"`
// 	Page   int    `json:"page"`
// 	Limit  int    `json:"limit"`
// 	Search string `json:"search"`
// 	// Add other fields specific to your messages
// }

// func NewDefaultConversationUserListMessageType() ConversationUserListMessageType {
// 	return ConversationUserListMessageType{
// 		Type:   "read_all",
// 		Page:   1,
// 		Limit:  10,
// 		Search: "default_search",
// 	}
// }

// func ConversationUserListWebSocket(c *websocket.Conn, hub adapters.HubInterface) {
// 	var school models.School
// 	var conversationUser models.ConversationUser
// 	var studentConversationUser []models.ConversationUser
// 	var staffConversationUser []models.ConversationUser
// 	var serializer serializers.StaffAndStudentListSerializer
// 	var studentTotal int64
// 	var staffTotal int64
// 	var loggedInConversationUserID uuid.UUID

// 	var (
// 		err error
// 		_   int
// 		msg []byte
// 	)

// 	// Create a context with a timeout of one hour
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Hour)
// 	defer cancel() // Make sure to call cancel to release resources when done

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	// Route identifier for this WebSocket
// 	route := "user-list"
// 	// Register the WebSocket connection with the Hub
// 	hub.Register(c, route)
// 	defer func() {
// 		// Unregister the WebSocket connection when the function exits
// 		hub.Unregister(c, route)
// 	}()

// 	var receivedMessage ConversationUserListMessageType
// 	receivedMessage = NewDefaultConversationUserListMessageType()

// 	schoolCode := c.Query("school_code")
// 	user := c.Locals("user").(models.User)

// 	page, err := strconv.Atoi(c.Query("page", "1")) // default to page 1 if not provided
// 	if err != nil {
// 		sendErrorToClient(c, db, "Invalid page number", loggedInConversationUserID)
// 		return
// 	}
// 	limit, err := strconv.Atoi(c.Query("limit", "10")) // default to 10 items per page if not provided
// 	if err != nil {
// 		sendErrorToClient(c, db, "Invalid page number", loggedInConversationUserID)
// 		return
// 	}

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		// Handle the error and send an error message to the client
// 		sendErrorToClient(c, db, "Failed to retrieve school data", loggedInConversationUserID)
// 		return
// 	}

// 	// Check if it's a school admin
// 	isAdmin := school.IsSchoolAdminOrOwner(ctx, db, schoolCode, user.ID)
// 	if !isAdmin {
// 		// Handle the case where the user is not an admin and send an error message
// 		sendErrorToClient(c, db, "Unauthorized", loggedInConversationUserID)
// 		return
// 	}

// 	conversationUser, err = conversationUser.GetOrCreateConversationUser(ctx, db, user.ID, school.ID, *school.SchoolCode)
// 	if err != nil {
// 		sendErrorToClient(c, db, "conversation user does not exist", loggedInConversationUserID)
// 		return
// 	}
// 	// set the logged in user id
// 	loggedInConversationUserID = conversationUser.ID

// 	// Continue handling WebSocket messages
// 	for {
// 		if _, msg, err = c.ReadMessage(); err != nil {
// 			logger.Error(ctx, "error reading message:", zap.Error(err))
// 			break
// 		}
// 		logger.Info(ctx, "Received message:", zap.Any("msg", string(msg)))

// 		if err := json.Unmarshal(msg, &receivedMessage); err != nil {
// 			logger.Error(ctx, "error parsing message:", zap.Error(err))
// 			sendErrorToClient(c, db, err.Error(), loggedInConversationUserID)
// 			break
// 		}

// 		if receivedMessage.Search == "" {
// 			// Get all the student conversations matching the search criteria
// 			err = db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).
// 				Model(&conversationUser).
// 				Where("school_id = ?  AND user_type = ? ",
// 					school.ID, "STUDENT").Preload("User").
// 				Find(&studentConversationUser).Error
// 			if err != nil {
// 				logger.Error(ctx, "error filtering student", zap.Error(err))
// 				// Handle the error and send an error message to the client
// 				sendErrorToClient(c, db, "Failed to retrieve student conversations", loggedInConversationUserID)
// 				break
// 			}

// 			// Get all the staff conversations matching the search criteria
// 			err = db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).
// 				Model(&conversationUser).
// 				Where("school_id = ? AND user_type = ? ",
// 					school.ID, "STAFF").Preload("User").
// 				Find(&staffConversationUser).Error
// 			if err != nil {
// 				logger.Error(ctx, "error filtering student", zap.Error(err))
// 				// Handle the error and send an error message to the client
// 				sendErrorToClient(c, db, "Failed to retrieve staff conversations", loggedInConversationUserID)
// 				break
// 			}

// 		} else {
// 			receivedMessage.Search = "%" + receivedMessage.Search + "%"
// 			// Get all the student conversations matching the search criteria
// 			err = db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).
// 				Model(&conversationUser).
// 				Where("school_id = ?  AND user_type = ? AND (user_id IN (SELECT id FROM users WHERE (first_name ILIKE  ? OR last_name ILIKE  ?)))",
// 					school.ID, "STUDENT", receivedMessage.Search, receivedMessage.Search).Preload("User").
// 				Find(&studentConversationUser).Error
// 			if err != nil {
// 				logger.Error(ctx, "error filtering student", zap.Error(err))
// 				// Handle the error and send an error message to the client
// 				sendErrorToClient(c, db, "Failed to retrieve student conversations", loggedInConversationUserID)
// 				break
// 			}

// 			// Get all the staff conversations matching the search criteria
// 			err = db.Scopes(adapters.NewPaginateAdapter(limit, page).PaginatedResult).
// 				Model(&conversationUser).
// 				Where("school_id = ? AND user_type = ? AND (user_id IN (SELECT id FROM users WHERE (first_name ILIKE ? OR last_name ILIKE ?)))",
// 					school.ID, "STAFF", receivedMessage.Search, receivedMessage.Search).Preload("User").
// 				Find(&staffConversationUser).Error
// 			if err != nil {
// 				logger.Error(ctx, "error filtering student", zap.Error(err))
// 				// Handle the error and send an error message to the client
// 				sendErrorToClient(c, db, "Failed to retrieve staff conversations", loggedInConversationUserID)
// 				break
// 			}

// 		}

// 		err = db.Model(&conversationUser).Where("school_id = ?  AND user_type = ?", school.ID, "STUDENT").Count(&studentTotal).Error
// 		err = db.Model(&conversationUser).Where("school_id = ?  AND user_type = ?", school.ID, "STAFF").Count(&staffTotal).Error
// 		staffSerializer := serializers.ConversationUserListSerializer{
// 			Total: staffTotal,
// 			Data:  staffConversationUser,
// 		}

// 		studentSerializer := serializers.ConversationUserListSerializer{
// 			Total: studentTotal,
// 			Data:  studentConversationUser,
// 		}

// 		serializer = serializers.StaffAndStudentListSerializer{
// 			Staff:   staffSerializer,
// 			Student: studentSerializer,
// 		}

// 		// Send the conversationData as JSON to the client
// 		if err := c.WriteJSON(serializer); err != nil {
// 			// Handle the error and send an error message to the client
// 			sendErrorToClient(c, db, "Failed to send conversation data", loggedInConversationUserID)
// 			break
// 		}
// 	}
// }

// type ConversationMessageUserMessageType struct {
// 	Type    string `json:"type"`
// 	Page    int    `json:"page"`
// 	Message string `json:"message"`
// 	FileUrl string `json:"file_url"`
// 	Limit   int    `json:"limit"`
// }

// func NewDefaultConversationMessageUserMessageType() ConversationMessageUserMessageType {
// 	return ConversationMessageUserMessageType{
// 		Type:    "read_all",
// 		Page:    1,
// 		Message: "",
// 		FileUrl: "",
// 		Limit:   10,
// 	}
// }

// func ConversationMessageUserWebsocket(c *websocket.Conn, hub adapters.HubInterface) {
// 	var school models.School
// 	var conversation models.Conversation
// 	var conversationUser1 models.ConversationUser
// 	var conversationUser2 models.ConversationUser
// 	var messages []models.Message
// 	var message models.Message
// 	var serializer serializers.ConversationMessagesListSerializer
// 	var loggedInConversationUserID uuid.UUID
// 	var (
// 		err                      error
// 		mt                       int
// 		msg                      []byte
// 		total                    int64
// 		conversationUserID1Found bool
// 		conversationUserID2      uuid.UUID
// 	)

// 	// Create a context with a timeout of one hour
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Hour)
// 	defer cancel() // Make sure to call cancel to release resources when done

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	var receivedMessage ConversationMessageUserMessageType
// 	receivedMessage = NewDefaultConversationMessageUserMessageType()

// 	schoolCode := c.Query("school_code")
// 	user := c.Locals("user").(models.User)
// 	conversationName := c.Params("name")

// 	// Split the string into individual UUIDs using underscores as the separator
// 	uuidStrings := strings.Split(conversationName, "__")
// 	if len(uuidStrings) != 2 {
// 		sendErrorToClient(c, db, "Invalid conversation name", loggedInConversationUserID)
// 	}

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		// Handle the error and send an error message to the client
// 		sendErrorToClient(c, db, "Failed to retrieve school data", loggedInConversationUserID)
// 		return
// 	}

// 	conversationUser1, err = conversationUser2.GetOrCreateConversationUser(ctx, db, user.ID, school.ID, schoolCode)
// 	if err != nil {
// 		sendErrorToClient(c, db, "error fetching current user conversation user", loggedInConversationUserID)
// 		return
// 	}
// 	loggedInConversationUserID = conversationUser1.ID

// 	// Check if the user's UUID matches either of the conversation participants
// 	for _, uuidString := range uuidStrings {
// 		uuidValue, err := uuid.Parse(uuidString)
// 		if uuidValue == conversationUser1.ID {
// 			conversationUserID1Found = true
// 		} else {
// 			conversationUserID2, err = uuid.Parse(uuidString)
// 			if err != nil {
// 				logger.Error(ctx, "invalid uuid passed")
// 				sendErrorToClient(c, db, "error parsing conversation user id", loggedInConversationUserID)
// 			}
// 		}
// 	}

// 	if conversationUserID1Found == false {
// 		sendErrorToClient(c, db, "conversation user id not found", loggedInConversationUserID)
// 		return
// 	}

// 	conversationUser2, err = conversationUser2.GetConversationUserWithID(ctx, db, conversationUserID2)
// 	if err != nil {
// 		sendErrorToClient(c, db, "error fetching other user conversation user", loggedInConversationUserID)
// 		return
// 	}

// 	conversation, err = conversation.GetOrConnectConversation(ctx, db, school.ID, conversationUser1.ID, conversationUser2.ID)
// 	if err != nil {
// 		logger.Error(ctx, "unable to  get or create conversation between both users", zap.Error(err))
// 		sendErrorToClient(c, db, "unable to  get or create conversation between both users", loggedInConversationUserID)
// 		return
// 	}

// 	if conversation.Name == nil {
// 		logger.Error(ctx, "the conversation name is nil", zap.Error(err))
// 	}
// 	route := *conversation.Name
// 	// Register the WebSocket connection with the Hub
// 	hub.Register(c, route)
// 	defer func() {
// 		// Unregister the WebSocket connection when the function exits
// 		hub.Unregister(c, route)
// 	}()

// 	// continue websocket
// 	for {
// 		if mt, msg, err = c.ReadMessage(); err != nil {
// 			logger.Error(ctx, "error reading message:", zap.Error(err))
// 			sendErrorToClient(c, db, "error reading message", loggedInConversationUserID)
// 			break
// 		}
// 		logger.Info(ctx, "Received message:", zap.Any("msg", string(msg)))

// 		if err := json.Unmarshal(msg, &receivedMessage); err != nil {
// 			logger.Error(ctx, "error parsing message:", zap.Error(err))
// 			sendErrorToClient(c, db, err.Error(), loggedInConversationUserID)
// 			break
// 		}

// 		switch receivedMessage.Type {
// 		case "read_all":
// 			err := db.Model(&models.Message{}).
// 				Where(
// 					"conversation_id = ? AND school_id =? AND read =? AND from_user_id =?",
// 					conversation.ID,
// 					school.ID,
// 					false,
// 					conversationUser1.ID).Update("read", true).Error
// 			if err != nil {
// 				sendErrorToClient(c, db, "error reading all messages", loggedInConversationUserID)
// 				break
// 			}
// 			err = c.WriteMessage(mt, []byte(`{"message": "successfully read message"}`))
// 			if err != nil {
// 				sendErrorToClient(c, db, "error reading all messages", loggedInConversationUserID)
// 				break
// 			}
// 		case "send_message":
// 			read := false
// 			message = models.Message{
// 				SchoolID:       &school.ID,
// 				ConversationID: &conversation.ID,
// 				FromUserID:     &conversationUser1.ID,
// 				ToUserID:       &conversationUser2.ID,
// 				Content:        &receivedMessage.Message,
// 				File:           &receivedMessage.FileUrl,
// 				Read:           &read,
// 			}
// 			err = db.Model(&models.Message{}).
// 				Create(&message).
// 				Preload("FromUser").
// 				Preload("ToUser").First(&message).Error
// 			if err != nil {
// 				logger.Error(ctx, "error creating message", zap.Error(err))
// 				sendErrorToClient(c, db, "error sending message ", loggedInConversationUserID)
// 				break
// 			}
// 			// update the conversation message
// 			err := db.Model(&models.Conversation{}).Where("id = ?", conversation.ID).Update("last_message", message.Content).Error
// 			if err != nil {
// 				logger.Error(ctx, "error updating conversation", zap.Error(err))
// 				sendErrorToClient(c, db, "error updating conversation", loggedInConversationUserID)
// 				break
// 			}

// 			// Convert the message to JSON
// 			messageJSON, err := json.Marshal(message)
// 			if err != nil {
// 				logger.Error(ctx, "error converting message to JSON", zap.Error(err))
// 				sendErrorToClient(c, db, "error sending message", loggedInConversationUserID)
// 				break
// 			}

// 			// Broadcast the message to all connected clients
// 			hub.Broadcast(c, route, messageJSON)

// 			// broadcast to the user notification
// 			hub.Broadcast(c, conversationUser2.ID.String(), messageJSON)

// 			err = c.WriteMessage(mt, []byte(`{"message": "successfully sent message"}`))
// 			if err != nil {
// 				sendErrorToClient(c, db, "error sending message", loggedInConversationUserID)
// 				break
// 			}
// 		case "get_messages":
// 			err := db.Model(&models.Message{}).
// 				Where("conversation_id = ?", conversation.ID).Count(&total).Error
// 			if err != nil {
// 				sendErrorToClient(c, db, err.Error(), loggedInConversationUserID)
// 			}
// 			err = db.Scopes(adapters.NewPaginateAdapter(receivedMessage.Limit, receivedMessage.Page).PaginatedResult).Model(&models.Message{}).
// 				Where("conversation_id = ? AND school_id =? ", conversation.ID, school.ID).Find(&messages).Error
// 			if err != nil {
// 				sendErrorToClient(c, db, "error getting messages", loggedInConversationUserID)
// 				break
// 			}
// 			err = db.Model(&models.Message{}).
// 				Where("conversation_id = ? AND school_id =? ", conversation.ID, school.ID).Count(&total).Error
// 			if err != nil {
// 				sendErrorToClient(c, db, "error getting all messages", loggedInConversationUserID)
// 			}
// 			serializer = serializers.ConversationMessagesListSerializer{
// 				Total:        total,
// 				Conversation: conversation,
// 				FromUser:     conversationUser1,
// 				ToUser:       conversationUser2,
// 				Messages:     messages,
// 			}

// 			// Send the conversationData as JSON to the client
// 			if err := c.WriteJSON(serializer); err != nil {
// 				// Handle the error and send an error message to the client
// 				sendErrorToClient(c, db, "Failed to send conversation data", loggedInConversationUserID)
// 				break

// 			}

// 		}
// 	}
// }

// type ConversationMessageUserMessageNotificationType struct {
// 	Type string `json:"type"`
// }

// func NewDefaultConversationMessageUserMessageNotificationType() ConversationMessageUserMessageNotificationType {
// 	return ConversationMessageUserMessageNotificationType{
// 		Type: "read_all",
// 	}
// }

// func ConversationUserMessageNotificationWebsocket(c *websocket.Conn, hub adapters.HubInterface) {
// 	var school models.School
// 	var conversationUser models.ConversationUser
// 	var messages []models.Message
// 	var loggedInConversationUserID uuid.UUID
// 	var (
// 		err error
// 		_   int
// 		msg []byte
// 	)

// 	// Create a context with a timeout of one hour
// 	ctx, cancel := context.WithTimeout(context.Background(), 1*time.Hour)
// 	defer cancel() // Make sure to call cancel to release resources when done

// 	// Open the database connection
// 	db := database.DBConnection()
// 	defer database.CloseDB(db)

// 	var receivedMessage ConversationMessageUserMessageNotificationType
// 	receivedMessage = NewDefaultConversationMessageUserMessageNotificationType()

// 	schoolCode := c.Query("school_code")
// 	user := c.Locals("user").(models.User)

// 	school, err = school.RetrieveBySchoolCode(ctx, db, schoolCode)
// 	if err != nil {
// 		// Handle the error and send an error message to the client
// 		sendErrorToClient(c, db, "Failed to retrieve school data", loggedInConversationUserID)
// 		return
// 	}

// 	conversationUser, err = conversationUser.GetOrCreateConversationUser(ctx, db, user.ID, school.ID, schoolCode)
// 	if err != nil {
// 		sendErrorToClient(c, db, "error fetching current user conversation user", loggedInConversationUserID)
// 		return
// 	}

// 	loggedInConversationUserID = conversationUser.ID
// 	route := loggedInConversationUserID.String()
// 	// Register the WebSocket connection with the Hub
// 	hub.Register(c, route)
// 	defer func() {
// 		// Unregister the WebSocket connection when the function exits
// 		hub.Unregister(c, route)
// 	}()

// 	// continue websocket
// 	for {
// 		if _, msg, err = c.ReadMessage(); err != nil {
// 			logger.Error(ctx, "error reading message:", zap.Error(err))
// 			sendErrorToClient(c, db, "error reading message", loggedInConversationUserID)
// 			break
// 		}
// 		logger.Info(ctx, "Received message:", zap.Any("msg", string(msg)))

// 		if err := json.Unmarshal(msg, &receivedMessage); err != nil {
// 			logger.Error(ctx, "error parsing message:", zap.Error(err))
// 			sendErrorToClient(c, db, err.Error(), loggedInConversationUserID)
// 			break
// 		}

// 		switch receivedMessage.Type {
// 		case "clear_all":
// 			messages = []models.Message{}
// 			if err != nil {
// 				sendErrorToClient(c, db, "error reading all messages", loggedInConversationUserID)
// 				break
// 			}
// 		case "notification":

// 		}
// 		// Send the conversationData as JSON to the client
// 		if err := c.WriteJSON(messages); err != nil {
// 			// Handle the error and send an error message to the client
// 			sendErrorToClient(c, db, "Failed to send conversation data", loggedInConversationUserID)
// 			break

// 		}
// 	}
// }
