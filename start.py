#!/usr/bin/env python3
import sys, os

USAGE="""
./start.py init         - creates the empty programs to be used
"""

def get_title(filename:str, day:str)->str:
    iday = int(day)
    with open(filename, 'r') as f:
        for line in f.readlines():
            line = line.strip()
            start = f"--- Day {iday}:"
            if line.startswith(start):
                splits = line.split("---")
                s = splits[1][8:]
                return s.strip()
    return None

def get_filename_or_die(filename:str)->str:
    if not os.path.isfile(filename):
        print(f"{filename} does not exist.")
        sys.exit(1)
    else:
        return filename

def token_switch(source_filename:str, target_filename:str, package:str, year:str, day:str, title:str, data:str=""):

    with open(source_filename, 'r') as f_in:
        with open(target_filename, 'w') as f_out:
            for line in f_in.readlines():
                line = line.replace("TOKEN_TITLE", title)
                line = line.replace("TOKEN_DAY", day)
                line = line.replace("TOKEN_YEAR", year)
                line = line.replace("TOKEN_PACKAGE", package)
                line = line.replace("TOKEN_START", package)
                line = line.replace("TOKEN_DATA", data)
                f_out.write(line)

def copy_file(source_filename:str, target_filename:str):
    with open(source_filename, 'r') as f_in:
        with open(target_filename, 'w') as f_out:
            f_out.write(f_in.read())

        
def run_day(year:str, day:str)->bool:
    root_dir = f"app"
    target_dir = f"{root_dir}/y{year}/d{day}"
    source_dir = f"aoc_downloads/{year}/day{day}"
    if not os.path.exists(source_dir):
        print(f"Source dir {source_dir} does not exist.")
        return False
    
    if os.path.exists(target_dir):
        print(f"Target dir {target_dir} exists.")
        return False


    problem_filename = get_filename_or_die(f"{source_dir}/problem.md")
    title = get_title(problem_filename, day)

    if not os.path.exists(target_dir):
        os.makedirs(target_dir)

    copy_file(f"{source_dir}/problem.md", f"{target_dir}/problem.md")
    token_switch("app/template/template_blog.md", f"{target_dir}/blog.md", f"d{day}", year, day, title)
    token_switch("app/template/template_program_test.go", f"{target_dir}/program_test.go", f"d{day}", year, day, title)
    token_switch("app/template/template_program.go", f"{target_dir}/program.go", f"d{day}", year, day, title)

    data = open(f"{source_dir}/input.txt").read()
    token_switch("app/template/template_data.go", f"{target_dir}/data.go", f"d{day}", year, day, title, data)

    package_dir = f"{root_dir}/y{year}"
    token_switch("app/template/template_app.go.template", f"{package_dir}/app.go", f"y{year}", year, day, title)

    print(f"year={year}")
    print(f"day={day}")
    print(f"target_dir={target_dir}")

    return True

def main():
    if len(sys.argv) == 1:
        print(USAGE)
        sys.exit(1)
                 
    command = sys.argv[1]
    if command == "init":
        for year in range(2015, 2025):
            for day in range(1, 26):
                if day < 10:
                    day = f"0{day}"
                source_dir = f"aoc_downloads/{year}/day{day}"
                if os.path.exists(source_dir):
                    run_day(str(year), str(day))
                else:
                    print(f"{source_dir} does not exist.")
    else:
        print(USAGE)    

if __name__ == "__main__":
    main()

