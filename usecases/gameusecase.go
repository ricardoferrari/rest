package gameusecase

import (
	"errors"
	"fmt"
	"sync"
)

const GAME_ALREADY_EXISTS = "game already exists"
const GAME_NOT_FOUND = "game not found"

// Game represents a game entity
type Game struct {
	ID    string
	Title string
	Genre string
}

// GameUseCase represents the game use case
type GameUseCase struct {
	mu    sync.Mutex
	games map[string]Game
}

// NewGameUseCase creates a new GameUseCase
func NewGameUseCase() *GameUseCase {
	return &GameUseCase{
		games: make(map[string]Game),
	}
}

// CreateGame creates a new game
func (uc *GameUseCase) CreateGame(game Game) error {
	uc.mu.Lock()
	defer uc.mu.Unlock()

	if _, exists := uc.games[game.ID]; exists {
		return errors.New(GAME_ALREADY_EXISTS)
	}

	uc.games[game.ID] = game
	return nil
}

// GetGame retrieves a game by ID
func (uc *GameUseCase) GetGame(id string) (Game, error) {
	uc.mu.Lock()
	defer uc.mu.Unlock()

	fmt.Println("GetGame", id)

	game, exists := uc.games[id]
	if !exists {
		return Game{}, errors.New(GAME_NOT_FOUND)
	}

	return game, nil
}

// UpdateGame updates an existing game
func (uc *GameUseCase) UpdateGame(game Game) error {
	uc.mu.Lock()
	defer uc.mu.Unlock()

	if _, exists := uc.games[game.ID]; !exists {
		return errors.New(GAME_NOT_FOUND)
	}

	uc.games[game.ID] = game
	return nil
}

// DeleteGame deletes a game by ID
func (uc *GameUseCase) DeleteGame(id string) error {
	uc.mu.Lock()
	defer uc.mu.Unlock()

	if _, exists := uc.games[id]; !exists {
		return errors.New(GAME_NOT_FOUND)
	}

	delete(uc.games, id)
	return nil
}

// ListGames lists all games
func (uc *GameUseCase) ListGames() []Game {
	uc.mu.Lock()
	defer uc.mu.Unlock()

	games := make([]Game, 0, len(uc.games))
	for _, game := range uc.games {
		games = append(games, game)
	}

	return games
}
