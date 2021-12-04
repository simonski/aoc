YEAR=$1 
DAY=$2

sed -e "s/yyyy_dd/${YEAR}_${DAY}/g" -e "s/package app/package aoc${YEAR}/g" -e "s/Day XX/Day ${DAY}/g" -e "s/yyyyDdd/${YEAR}D${DAY}/g" app/aocYYYY_DD.go > app/aoc${YEAR}/aoc${YEAR}_${DAY}.go
sed -e "s/yyyy_dd/${YEAR}_${DAY}/g" -e "s/package app/package aoc${YEAR}/g" -e "s/Day XX/Day ${DAY}/g" -e "s/yyyyDdd/${YEAR}D${DAY}/g" app/aocYYYY_DD_data.go > app/aoc${YEAR}/aoc${YEAR}_${DAY}_data.go
sed -e "s/yyyy_dd/${YEAR}_${DAY}/g" -e "s/package app/package aoc${YEAR}/g" -e "s/Day XX/Day ${DAY}/g" -e "s/yyyyDdd/${YEAR}D${DAY}/g" app/aocYYYY_DD_test.go > app/aoc${YEAR}/aoc${YEAR}_${DAY}_test.go

echo "ok!"
