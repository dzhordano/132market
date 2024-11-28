package integration

import (
	"context"
	"dzhordano/132market/services/users/internal/application/command"
	"dzhordano/132market/services/users/internal/application/interfaces"
	"dzhordano/132market/services/users/internal/application/services"
	"dzhordano/132market/services/users/internal/domain/entities"
	"dzhordano/132market/services/users/internal/domain/repository"
	"dzhordano/132market/services/users/internal/infrastructure/db/postgres"
	"dzhordano/132market/services/users/pkg/migration/goose"
	"log"
	"log/slog"
	"os"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"github.com/joho/godotenv"
	"github.com/stretchr/testify/suite"
)

func TestSuite(t *testing.T) {
	if testing.Short() {
		t.Skip("skipping integration test")
	}

	suite.Run(t, new(BaseSuite))
}

type BaseSuite struct {
	suite.Suite

	pool *pgxpool.Pool
	svc  interfaces.UserService
	repo repository.UserRepository
}

func (s *BaseSuite) SetupSuite() {
	if err := godotenv.Load("./../../.env"); err != nil {
		log.Fatalf("Error loading .env file: %v", err)
	}

	dsn := os.Getenv("POSTGRES_TEST_DSN")
	if dsn == "" {
		log.Fatal("POSTGRES_TEST_DSN variable is not set")
	}

	s.pool = postgres.NewPool(dsn)
	if err := goose.Run(context.TODO(), "./../../migrations", dsn, "up"); err != nil {
		log.Fatalf("Failed to run migrations: %v", err)
	}

	s.repo = postgres.NewUserRepository(s.pool)
	s.svc = services.NewUserService(slog.New(slog.NewJSONHandler(os.Stdout, &slog.HandlerOptions{Level: slog.LevelDebug})), s.repo)

	s.seedData()
}

var (
	testUser1, testUser2, testUser3 *entities.User
)

func (s *BaseSuite) seedData() {
	u1, err := entities.NewUser("u1", "email1@mail.ru", "p@ssw0Rd")
	if err != nil {
		s.T().Fatalf("Failed to create user 1: %v", err)
	}
	u2, err := entities.NewUser("u2", "email2@mail.ru", "p@ssw0Rd")
	if err != nil {
		s.T().Fatalf("Failed to create user 2: %v", err)
	}
	u3, err := entities.NewUser("u3", "email3@mail.ru", "p@ssw0Rd")
	if err != nil {
		s.T().Fatalf("Failed to create user 3: %v", err)
	}

	if testUser1, err = s.repo.Save(context.TODO(), u1); err != nil {
		s.T().Fatalf("Failed to save user 1: %v", err)
	}
	if testUser2, err = s.repo.Save(context.TODO(), u2); err != nil {
		s.T().Fatalf("Failed to save user 2: %v", err)
	}
	if testUser3, err = s.repo.Save(context.TODO(), u3); err != nil {
		s.T().Fatalf("Failed to save user 3: %v", err)
	}
}

// todo данный тест меня на напрягах держит (какой-то слабый)

func (s *BaseSuite) TestCreateUser() {
	resp, err := s.svc.CreateUser(context.TODO(), &command.CreateUserCommand{
		Name:         "u4",
		Email:        "email4@mail.ru",
		PasswordHash: "p@ssw0Rd",
	})
	s.NoError(err)

	s.Equal("u4", resp.Result.Name)
	s.Equal("email4@mail.ru", resp.Result.Email)
}

func (s *BaseSuite) TestDeleteUser() {
	state := testUser1.State

	err := s.svc.DeleteUser(context.TODO(), testUser1.ID.String())
	s.NoError(err)

	user, _ := s.repo.FindById(context.TODO(), testUser1.ID)

	s.NotEqual(state, user.State)

	testUser1.State = entities.StateDeleted
}

func (s *BaseSuite) TestUpdateUser() {
	oldUser1 := testUser1

	updatedUser1, err := s.svc.UpdateUser(context.TODO(), &command.UpdateUserCommand{
		ID:           testUser1.ID.String(),
		Name:         "u1_NewName",
		Email:        "email1@mail.ru",
		PasswordHash: "p@ssw0Rd",
	})
	s.NoError(err)

	s.NotEqual(updatedUser1.Result.Name, oldUser1.Name)
}

func (s *BaseSuite) TestFindUserById() {
	resp, err := s.svc.FindUserById(context.TODO(), testUser1.ID.String())
	s.NoError(err)

	s.Equal(testUser1.ID, resp.Result.ID)
	s.Equal(testUser1.Name, resp.Result.Name)
	s.Equal(testUser1.Email, resp.Result.Email)
	s.Equal(testUser1.State.String(), resp.Result.State)
	s.Equal(testUser1.Status.String(), resp.Result.Status)
	s.Equal(testUser1.Roles[0].String(), resp.Result.Roles[0])
	s.Equal(testUser1.CreatedAt, resp.Result.CreatedAt)
	s.Equal(testUser1.LastSeenAt, resp.Result.LastSeenAt)
}

func (s *BaseSuite) TestFindAllUsersNoFilter() {
	users, err := s.svc.FindAllUsers(context.TODO(), 0, 3)
	s.NoError(err)

	s.Equal(3, len(users.Result))

	s.Equal(testUser1.ID, users.Result[0].ID)
	s.Equal(testUser2.ID, users.Result[1].ID)
	s.Equal(testUser3.ID, users.Result[2].ID)

	s.Equal(testUser1.Name, users.Result[0].Name)
	s.Equal(testUser2.Name, users.Result[1].Name)
	s.Equal(testUser3.Name, users.Result[2].Name)

	s.Equal(testUser1.Email, users.Result[0].Email)
	s.Equal(testUser2.Email, users.Result[1].Email)
	s.Equal(testUser3.Email, users.Result[2].Email)

	s.Equal(testUser1.State.String(), users.Result[0].State)
	s.Equal(testUser2.State.String(), users.Result[1].State)
	s.Equal(testUser3.State.String(), users.Result[2].State)

	s.Equal(testUser1.CreatedAt, users.Result[0].CreatedAt)
	s.Equal(testUser2.CreatedAt, users.Result[1].CreatedAt)
	s.Equal(testUser3.CreatedAt, users.Result[2].CreatedAt)

	s.Equal(testUser1.LastSeenAt, users.Result[0].LastSeenAt)
	s.Equal(testUser2.LastSeenAt, users.Result[1].LastSeenAt)
	s.Equal(testUser3.LastSeenAt, users.Result[2].LastSeenAt)

	s.Equal(testUser1.Roles[0].String(), users.Result[0].Roles[0])
	s.Equal(testUser2.Roles[0].String(), users.Result[1].Roles[0])
	s.Equal(testUser3.Roles[0].String(), users.Result[2].Roles[0])

	s.Equal(testUser1.Status.String(), users.Result[0].Status)
	s.Equal(testUser2.Status.String(), users.Result[1].Status)
	s.Equal(testUser3.Status.String(), users.Result[2].Status)
}

func (s *BaseSuite) TestFindAllUsersOffset2() {
	users, err := s.svc.FindAllUsers(context.TODO(), 2, 1)
	s.NoError(err)

	s.Equal(1, len(users.Result))

	s.Equal(testUser3.ID, users.Result[0].ID)

	s.Equal(testUser3.Name, users.Result[0].Name)

	s.Equal(testUser3.Email, users.Result[0].Email)

	s.Equal(testUser3.State.String(), users.Result[0].State)

	s.Equal(testUser3.CreatedAt, users.Result[0].CreatedAt)

	s.Equal(testUser3.LastSeenAt, users.Result[0].LastSeenAt)

	s.Equal(testUser3.Roles[0].String(), users.Result[0].Roles[0])

	s.Equal(testUser3.Status.String(), users.Result[0].Status)
}

func (s *BaseSuite) TestFindUserByCredentials() {
	resp, err := s.svc.FindUserByCredentials(context.TODO(), testUser1.Email, testUser1.PasswordHash)
	s.NoError(err)

	s.Equal(testUser1.ID, resp.Result.ID)
	s.Equal(testUser1.Name, resp.Result.Name)
	s.Equal(testUser1.Email, resp.Result.Email)
	s.Equal(testUser1.State.String(), resp.Result.State)
	s.Equal(testUser1.Status.String(), resp.Result.Status)
	s.Equal(testUser1.Roles[0].String(), resp.Result.Roles[0])
	s.Equal(testUser1.CreatedAt, resp.Result.CreatedAt)
	s.Equal(testUser1.LastSeenAt, resp.Result.LastSeenAt)
}
