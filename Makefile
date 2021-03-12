NAME			=		n_puzzle

SRC_FILES		=		main.go\
						create_puzzle.go\
						generate.go\
						get_user_entry.go\
						header.go\
						helper.go\
						is_valid_taquin.go\
						manage_command.go\
						manage_load.go\
						manage_play.go\
						moves.go\
						solve.go\
						webapp.go\
						PriorityQueue.go\
						set_cmd.go



all: build

build:
	go get github.com/eiannone/keyboard
	go build -o $(NAME) $(SRC_FILES)

run:
	go get github.com/eiannone/keyboard
	go run $(SRC_FILES)

clean:
	rm -rf $(NAME)

fclean: clean
	rm -rf go.mod
	rm -rf go.sum

re: clean all