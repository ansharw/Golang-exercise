package controllers

// type MockContext struct {
// 	mock.Mock
// }

// func (m *MockContext) Param(key string) string {
// 	args := m.Called(key)
// 	return args.String(0)
// }

// func (m *MockContext) MustGet(key interface{}) interface{} {
// 	args := m.Called(key)
// 	return args.Get(0)
// }

// func (m *MockContext) AbortWithStatusJSON(code int, obj interface{}) {
// 	m.Called(code, obj)
// }

// func (m *MockContext) JSON(code int, obj interface{}) {
// 	m.Called(code, obj)
// }

// func TestGetProduct(t *testing.T) {
// 	// create a mock gin.Context instance
// 	c := &MockContext{}

// 	// set up mock database
// 	db := database.NewMockDB()

// 	// set up expected result
// 	// expectedProduct := &models.Product{ID: 1, Name: "Product 1"}
// 	expectedProduct := &models.Product{
// 		GormModel: models.GormModel{
// 			ID: 1,
// 		},
// 		Title:       "test",
// 		Description: "test desc",
// 		UserID:      2,
// 	}

// 	// set up mock database query
// 	db.On("Where", "id = ? AND user_id = ?", 1, 1).Return(db)
// 	db.On("First", mock.AnythingOfType("*models.Product")).Return(nil).Run(func(args mock.Arguments) {
// 		product := args.Get(0).(*models.Product)
// 		product.ID = expectedProduct.GormModel.ID
// 		product.Title = expectedProduct.Title
// 		product.Description = expectedProduct.Description
// 	})

// 	// set up mock gin.Context properties
// 	c.On("MustGet", "userData").Return(map[string]interface{}{"id": float64(1)})
// 	c.On("Param", "productId").Return(strconv.Itoa(int(expectedProduct.ID)))

// 	// call the handler function with the mock context and database
// 	// GetProduct(&gin.Context{})

// 	// assert the response
// 	expectedBody := `{"id":1,"name":"test","description":"test descd","user_id":2}`
// 	c.AssertCalled(t, "JSON", http.StatusOK, expectedProduct, expectedBody)
// }
