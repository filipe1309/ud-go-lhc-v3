## Create a new section

# get section number from parameter
SEC=$1

# if no section number is provided, get the last section number
if [ -z "$SEC" ]; then
  SEC=$(ls -d sec-* | sort -n | tail -n 1 | sed 's/[^0-9]*//g')
  SEC=$((SEC + 1))
  echo -e "%{F#5f5f5f}No section number provided. Using $SEC as the next section number.%{F-}"
fi

echo "📝 Creating section $SEC"

# create a new section directory
mkdir sec-$SEC
touch sec-$SEC/main.go

# create a new section from template
cp scripts/main sec-$SEC/main.go

# replace {{ SEC }} with section number
sed -i '' "s/{{ SEC }}/$SEC/g" sec-$SEC/main.go
