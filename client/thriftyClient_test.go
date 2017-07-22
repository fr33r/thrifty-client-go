package client

import (
	"errors"
	"testify/assert"
	"testify/suite"
	"testing"
	"thrifty-client-go/gen-go/thrifty"
	"thrifty-client-go/mocks"
)

type ThriftClientTestSuite struct {
	suite.Suite

	//system under test (SUT).
	sut                     *ThriftyClient
	personServiceClientMock *mocks.PersonServiceClient
}

func TestSuiteSetup(t *testing.T) {
	suite.Run(t, new(ThriftClientTestSuite))
}

func (suite *ThriftClientTestSuite) SetupSuite() {}

func (suite *ThriftClientTestSuite) SetupTest() {
	suite.personServiceClientMock = &mocks.PersonServiceClient{}
	suite.sut = &ThriftyClient{Client: suite.personServiceClientMock} //mock dependencies.
}

func (suite *ThriftClientTestSuite) BeforeTest(suiteName, testName string) {}

func (suite *ThriftClientTestSuite) TestGet_OutcomeIs_PersonRetrieved() {

	//arrange.
	const id int32 = 1
	const GET string = "Get"
	expectedPerson := thrifty.NewPerson()
	expectedPerson.GivenName = "Jon"
	expectedPerson.Surname = "Freer"
	expectedPerson.Age = 25
	var expectedError error = nil

	suite.personServiceClientMock.On(GET, id).Return(expectedPerson, expectedError)

	//action.
	actualPerson, actualError := suite.sut.Get(id)

	//assert.
	assert.Equal(suite.T(), expectedError, actualError)
	assert.Equal(suite.T(), expectedPerson.GivenName, actualPerson.GivenName)
	assert.Equal(suite.T(), expectedPerson.Surname, actualPerson.Surname)
	assert.Equal(suite.T(), expectedPerson.Age, actualPerson.Age)
}

func (suite *ThriftClientTestSuite) TestGet_OutcomeIs_Error() {

	//arrange.
	const id int32 = 1
	const GET string = "Get"
	var expectedPerson *thrifty.Person = nil
	expectedError := errors.New("Something happened!")

	suite.personServiceClientMock.On(GET, id).Return(expectedPerson, expectedError)

	//action.
	actualPerson, actualErr := suite.sut.Get(id)

	//assert.
	assert.Equal(suite.T(), expectedError, actualErr)
	assert.Equal(suite.T(), expectedPerson, actualPerson)
}

func (suite *ThriftClientTestSuite) AfterTest(suiteName, testName string) {}

func (suite *ThriftClientTestSuite) TearDownTest() {
	suite.sut = nil
	suite.personServiceClientMock = nil
}

func (suite *ThriftClientTestSuite) TearDownSuite() {}
