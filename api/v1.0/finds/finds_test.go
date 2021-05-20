package finds

import (
	"testing"

	"github.com/gin-gonic/gin"
)

func setupServer(app *gin.Engine) {
}

func TestfindsPOST(t *testing.T) {

	// // create mock
	// mc := gomock.NewController(t)
	// defer mc.Finish()

	// mockSvc := servicehandler.NewMockServiceHandler(mc)

	// type test struct {
	// 	name        string
	// 	user        user.User
	// 	requestBody request.BodyFlowsPOST
	// 	reqFlow     *flow.Flow
	// 	resFlow     *flow.Flow
	// }

	// tests := []test{
	// 	{
	// 		"normal",
	// 		user.User{
	// 			ID:         1,
	// 			Permission: user.PermissionAdmin,
	// 		},
	// 		request.BodyFlowsPOST{
	// 			Name:   "test name",
	// 			Detail: "test detail",
	// 			Actions: []action.Action{
	// 				{
	// 					Type: action.TypeAnswer,
	// 				},
	// 			},
	// 		},
	// 		&flow.Flow{
	// 			Name:    "test name",
	// 			Detail:  "test detail",
	// 			Persist: true,
	// 			Actions: []action.Action{
	// 				{
	// 					Type: action.TypeAnswer,
	// 				},
	// 			},
	// 		},
	// 		&flow.Flow{
	// 			ID:      uuid.FromStringOrNil("264b18d4-82fa-11eb-919b-9f55a7f6ace1"),
	// 			Name:    "test name",
	// 			Detail:  "test detail",
	// 			Persist: true,
	// 			Actions: []action.Action{
	// 				{
	// 					Type: action.TypeAnswer,
	// 				},
	// 			},
	// 		},
	// 	},
	// 	{
	// 		"webhook",
	// 		user.User{
	// 			ID:         1,
	// 			Permission: user.PermissionAdmin,
	// 		},
	// 		request.BodyFlowsPOST{
	// 			Name:       "test name",
	// 			Detail:     "test detail",
	// 			WebhookURI: "https://test.com/webhook",
	// 			Actions: []action.Action{
	// 				{
	// 					Type: action.TypeAnswer,
	// 				},
	// 			},
	// 		},
	// 		&flow.Flow{
	// 			Name:       "test name",
	// 			Detail:     "test detail",
	// 			Persist:    true,
	// 			WebhookURI: "https://test.com/webhook",
	// 			Actions: []action.Action{
	// 				{
	// 					Type: action.TypeAnswer,
	// 				},
	// 			},
	// 		},
	// 		&flow.Flow{
	// 			ID:         uuid.FromStringOrNil("7413ccc8-82fa-11eb-9826-3b9adf174ed9"),
	// 			Name:       "test name",
	// 			Detail:     "test detail",
	// 			Persist:    true,
	// 			WebhookURI: "https://test.com/webhook",
	// 			Actions: []action.Action{
	// 				{
	// 					Type: action.TypeAnswer,
	// 				},
	// 			},
	// 		},
	// 	},
	// }

	// for _, tt := range tests {
	// 	t.Run(tt.name, func(t *testing.T) {

	// 		w := httptest.NewRecorder()
	// 		_, r := gin.CreateTestContext(w)

	// 		r.Use(func(c *gin.Context) {
	// 			c.Set(common.OBJServiceHandler, mockSvc)
	// 			c.Set("user", tt.user)
	// 		})
	// 		setupServer(r)

	// 		// create body
	// 		body, err := json.Marshal(tt.requestBody)
	// 		if err != nil {
	// 			t.Errorf("Could not marshal the request. err: %v", err)
	// 		}

	// 		mockSvc.EXPECT().FlowCreate(&tt.user, tt.reqFlow).Return(tt.resFlow, nil)
	// 		req, _ := http.NewRequest("POST", "/v1.0/flows", bytes.NewBuffer(body))
	// 		req.Header.Set("Content-Type", "application/json")

	// 		r.ServeHTTP(w, req)
	// 		if w.Code != http.StatusOK {
	// 			t.Errorf("Wrong match. expect: %d, got: %d", http.StatusOK, w.Code)
	// 		}
	// 	})
	// }
}
