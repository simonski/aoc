YEAR=$1 
DAY=$2

read -p "Year: " YEAR
read -p "Day : " DAY
read -p "Title : " TITLE

if [ "$YEAR" == "" ]
then
    echo "Error, usage $0 YYYY"
    exit 1
fi

if [ "$DAY" == "" ]
then
    echo "Error, usage $0 MM"
    exit 1
fi

if [ "$TITLE" == "" ]
then
    echo "Error, you need the title."
    exit 1
fi

mkdir -p app/aoc$YEAR/d${DAY}

sed -e "s/TOKEN_PACKAGE/d${DAY}/g" -e "s/TOKEN_TITLE/${TITLE}/g" -e "s/TOKEN_YEAR/${YEAR}/g" -e "s/TOKEN_DAY/${DAY}/g" app/template/template_data.go > app/aoc${YEAR}/d${DAY}/data.go

sed -e "s/TOKEN_PACKAGE/d${DAY}/g" -e "s/TOKEN_TITLE/${TITLE}/g" -e "s/TOKEN_YEAR/${YEAR}/g" -e "s/TOKEN_DAY/${DAY}/g" app/template/template_blog.md > app/aoc${YEAR}/d${DAY}/blog.md

sed -e "s/TOKEN_PACKAGE/d${DAY}/g" -e "s/TOKEN_TITLE/${TITLE}/g" -e "s/TOKEN_YEAR/${YEAR}/g" -e "s/TOKEN_DAY/${DAY}/g" app/template/template_program_test.go > app/aoc${YEAR}/d${DAY}/program_test.go

sed -e "s/TOKEN_PACKAGE/d${DAY}/g" -e "s/TOKEN_TITLE/${TITLE}/g" -e "s/TOKEN_YEAR/${YEAR}/g" -e "s/TOKEN_DAY/${DAY}/g" app/template/template_program.go > app/aoc${YEAR}/d${DAY}/program.go


# git checkout main
# git pull origin main
# git checkout -b feature/${YEAR}_${DAY}
# echo "ok!"
