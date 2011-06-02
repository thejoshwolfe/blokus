package andyai

// TODO: move this to blokus package that andyai imports
type BlokusState struct {
    WhoseTurn int // which player the game is waiting on
    [][]int Board // 2d grid of colored cells
    [4][]int PiecesLeft // for each player, list of pieces left
}

func TakeTurn(BlokusState state) {
    // returns move, heuristic score for each player
    func getBestMove(BlokusState state) {
        list_of_possible_moves := generateMoveList(state, state.WhoseTurn)
        best_score := nil
        best_move := nil
        for move in list_of_possible_moves {
            test_board = doMove(test_board, move)
            while not gameOver(test_board) {
                enemy_move = getBestMove(test_board)
                test_board = doMove(test_board, enemy_move)
            }
            board_score = boardScore(test_board)
            if best_score == nil || board_score > best_score {
                best_score = board_score
                best_move = move
            }
        }
    }
    best_move, scores = getBestMove(state)
    return best_move
}
