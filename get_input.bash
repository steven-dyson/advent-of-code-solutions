BLUE=$(tput setaf 4)
BOLD=$(tput bold)
RESET=$(tput sgr0)

echo ""
printf "=== Advent of Code Input Downloader ===\n\n"

echo "${BOLD}${BLUE}[AoC Helper] What is your session ID?${RESET}"
echo ""
read -r session_token

echo ""
echo "${BOLD}${BLUE}[AoC Helper] What year are you working on?${RESET}"
echo ""
read -r year

echo ""
echo "${BOLD}${BLUE}[AoC Helper] What day are you working on?${RESET}"
echo ""
read -r day

# Directory
DIR="$year/day$day"
mkdir -p "$DIR"

wget --header="Cookie: session=$session_token" \
	-O "$DIR/input.txt" \
	"https://adventofcode.com/$year/day/$day/input"

echo "${BOLD}${BLUE}[AoC Helper] You are all set!${RESET}"
