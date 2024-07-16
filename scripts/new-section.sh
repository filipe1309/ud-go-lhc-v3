## Create a new section

clear

BOLD="\033[1m"
DARKGRAY="\033[1;30m"
NC="\033[0m"

# get section number from parameter
SEC=$1

# if no section number is provided, get the last section number
if [ -z "$SEC" ]; then
  SEC=$(ls -d sec-* | sort -n | tail -n 1 | sed 's/[^0-9]*//g')
  SEC=$((SEC + 1))
  echo -e "${DARKGRAY}${BOLD}No section number provided. Using $SEC${NC}"
fi

echo -e "${BOLD}📝 Creating section $SEC${NC}"

# create a new section directory
mkdir sec-$SEC
touch sec-$SEC/main.go

# create a new section from template
cp scripts/main sec-$SEC/main.go

# replace {{ SEC }} with section number
sed -i '' "s/{{ SEC }}/$SEC/g" sec-$SEC/main.go

echo -e "${BOLD}✅ Section $SEC created${NC}"
