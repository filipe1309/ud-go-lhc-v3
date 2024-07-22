## Create a new section

clear

BOLD="\033[1m"
DARKGRAY="\033[1;30m"
NC="\033[0m"

# get section number from parameter
SEC=$1

# get description from parameter
DESC=$2

# if no section number is provided, get the last section number
if [ -z "$SEC" ]; then
  # SEC=$(ls -d sec-* | tail -n 1 | sed 's/[^0-9]{2}//g')
  SEC=$(ls -d sec-* | sort -n | tail -n 1 | cut -c 5- | cut -d '-' -f 1) # first 2 numbers of the last section
  SEC=$((SEC + 1))
  echo -e "${DARKGRAY}${BOLD}No section number provided. Using $SEC${NC}"
fi

# if no description is provided, ask for it
if [ -z "$DESC" ]; then
  read -p "Description: " DESC
fi

echo -e "${BOLD}📝 Creating section $SEC - $DESC${NC}"

# convert DESC to lowercase and replace spaces with dashes
DESC_SNAKE_CASE=$(echo $DESC | tr '[:upper:]' '[:lower:]' | tr ' ' '-')

# create a new section directory
mkdir sec-$SEC-$DESC_SNAKE_CASE
touch sec-$SEC-$DESC_SNAKE_CASE/main.go

# create a new section from template
cp scripts/main sec-$SEC-$DESC_SNAKE_CASE/main.go

# replace {{ SEC }} with section number
sed -i '' "s/{{ SEC }}/$SEC/g" sec-$SEC-$DESC_SNAKE_CASE/main.go

# replace {{ DESC }} with description
sed -i '' "s/{{ DESC }}/$DESC/g" sec-$SEC-$DESC_SNAKE_CASE/main.go

echo -e "${BOLD}✅ Section $SEC-$DESC_SNAKE_CASE created${NC}"
