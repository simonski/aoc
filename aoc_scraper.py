import os
import time
import requests
from bs4 import BeautifulSoup

# thank sam altman's baby for this script 
# pip install requests beautifulsoup4
# In Chrome: Application → Cookies → https://adventofcode.com
# export AOC_SESSION="xx"
# ./aoc_scraper.py

#
# ========================
# USER CONFIGURATION
# ========================

YEARS = [2015, 2016, 2017, 2018, 2019, 2020, 2021, 2022, 2023, 2024]  # Add/remove years here
SESSION_COOKIE = os.getenv("AOC_SESSION")  # Set via environment variable
OUTPUT_DIR = "aoc_downloads"
THROTTLE_SECONDS = 1  # Wait between requests

# ========================
# AOC Scraper
# ========================

HEADERS = {
    "User-Agent": "Mozilla/5.0 (compatible; aoc-downloader/2.0)",
    "Cookie": f"session={SESSION_COOKIE}"
}

def fetch_html(year: int, day: int) -> str:
    url = f"https://adventofcode.com/{year}/day/{day}"
    response = requests.get(url, headers=HEADERS)
    response.raise_for_status()
    return response.text

def fetch_input(year: int, day: int) -> str:
    url = f"https://adventofcode.com/{year}/day/{day}/input"
    response = requests.get(url, headers=HEADERS)
    response.raise_for_status()
    return response.text.strip()

def extract_markdown_from_html(html: str) -> str:
    soup = BeautifulSoup(html, "html.parser")
    articles = soup.find_all("article")
    markdown_sections = []
    for i, article in enumerate(articles):
        header = f"### Part {i + 1}"
        content = article.get_text()
        markdown_sections.append(f"{header}\n\n{content}")
    return "\n\n---\n\n".join(markdown_sections)

def save_to_file(path: str, content: str):
    os.makedirs(os.path.dirname(path), exist_ok=True)
    with open(path, "w", encoding="utf-8") as f:
        f.write(content)

def already_downloaded(path: str) -> bool:
    return os.path.isfile(path) and os.path.getsize(path) > 0

def process_day(year: int, day: int):
    base_path = os.path.join(OUTPUT_DIR, str(year), f"day{day:02}")
    problem_path = os.path.join(base_path, "problem.md")
    input_path = os.path.join(base_path, "input.txt")

    if already_downloaded(problem_path):
        print(f"[{year} Day {day:02}] Already downloaded, skipping.")
        return

    try:
        print(f"[{year} Day {day:02}] Fetching problem and input...")
        html = fetch_html(year, day)
        markdown = extract_markdown_from_html(html)
        puzzle_input = fetch_input(year, day)

        save_to_file(problem_path, markdown)
        save_to_file(input_path, puzzle_input)
    except requests.HTTPError as e:
        print(f"[{year} Day {day:02}] HTTP error: {e}")
    except Exception as e:
        print(f"[{year} Day {day:02}] Unexpected error: {e}")
    finally:
        time.sleep(THROTTLE_SECONDS)

def main():
    if not SESSION_COOKIE:
        print("⚠️  Please set your AOC_SESSION environment variable with your session cookie.")
        return

    for year in YEARS:
        for day in range(1, 26):
            process_day(year, day)

if __name__ == "__main__":
    main()
