YEAR=$1 
DAY=$2
sed -e "s/package app/package aoc${YEAR}/g" -e "s/Day XX/Day ${DAY}/g" -e "s/XXDXX/${YEAR}D${DAY}/g" app/aoc20XX_XX.go > app/aoc${YEAR}/aoc20${YEAR}_${DAY}.go
sed -e "s/package app/package aoc${YEAR}/g" -e "s/Day XX/Day ${DAY}/g" -e "s/XXDXX/${YEAR}D${DAY}/g" app/aoc20XX_XX_data.go > app/aoc${YEAR}/aoc20${YEAR}_${DAY}_data.go
sed -e "s/package app/package aoc${YEAR}/g" -e "s/Day XX/Day ${DAY}/g" -e "s/XXDXX/${YEAR}D${DAY}/g" app/aoc20XX_XX_test.go > app/aoc${YEAR}/aoc20${YEAR}_${DAY}_test.go

echo "ok!"
