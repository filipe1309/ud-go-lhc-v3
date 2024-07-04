## Create a new section

# get section number from parameter
SEC=$1

# create a new section directory
mkdir sec-$SEC
touch sec-$SEC/main.go

# create a new section from template
cp scripts/main sec-$SEC/main.go

# replace {{ SEC }} with section number
sed -i '' "s/{{ SEC }}/$SEC/g" sec-$SEC/main.go
