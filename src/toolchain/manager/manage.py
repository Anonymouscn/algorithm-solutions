import argparse
import requests
import json
import os
import re
import sys
import shutil
import tempfile
import subprocess
import textwrap
import base64
import urllib.parse
import urllib.request
from bs4 import BeautifulSoup
from typing import Any, Dict, List
from wcwidth import wcwidth, wcswidth
from pathlib import Path
from urllib.parse import urlparse
from pathlib import Path


Align = str  # "left" | "right" | "center"


lang_support = ['java', 'go', 'python', 'rust']


leetcode_image_pattern = r"\[\s*(https?://[^\]\s]+(?:\s+[^\]\s]+)*)\s*\]"
leetcode_strong_pattern = r'^\s*示例\s*\d*[:：]|输入\s*\d*[:：]|输出\s*\d*[:：]|解释\s*\d*[:：]|提示\s*\d*[:：]'


leetcode_level_html = """
<p style="color: $color; display: inline-flex; item-align: center; justify-content: center; padding: 2px 10px; border: 2px solid #666; border-radius: 20px">$text</p>
"""
leetcode_level_color = {
    'EASY': '#1cb8b8',
    'MEDIUM': '#ffb800',
    'HARD': '#f8615c',
}
leetcode_level_translate_cn = {
    'EASY': '简单',
    'MEDIUM': '中等',
    'HARD': '困难',
}


headers = {
    "User-Agent": "Mozilla/5.0 (Windows NT 10.0; Win64; x64) AppleWebKit/537.36 (KHTML, like Gecko) Chrome/54.0.2840.99 Safari/537.36",
}

query_url = "https://leetcode.cn/graphql/"
problem_desc_url = "https://leetcode.cn/problems/$problem_name"

# 查询问题数据
query_payload = {
  "query": "\n    query problemsetQuestionListV2($filters: QuestionFilterInput, $limit: Int, $searchKeyword: String, $skip: Int, $sortBy: QuestionSortByInput, $categorySlug: String) {\n  problemsetQuestionListV2(\n    filters: $filters\n    limit: $limit\n    searchKeyword: $searchKeyword\n    skip: $skip\n    sortBy: $sortBy\n    categorySlug: $categorySlug\n  ) {\n    questions {\n      id\n      titleSlug\n      title\n      translatedTitle\n      questionFrontendId\n      paidOnly\n      difficulty\n      topicTags {\n        name\n        slug\n        nameTranslated\n      }\n      status\n      isInMyFavorites\n      frequency\n      acRate\n      contestPoint\n    }\n    totalLength\n    finishedLength\n    hasMore\n  }\n}\n    ",
  "variables": {
    "skip": 0,
    "limit": 100,
    "categorySlug": "all-code-essentials",
    "filters": {
      "filterCombineType": "ALL",
      "statusFilter": {
        "questionStatuses": [],
        "operator": "IS"
      },
      "difficultyFilter": {
        "difficulties": [],
        "operator": "IS"
      },
      "languageFilter": {
        "languageSlugs": [],
        "operator": "IS"
      },
      "topicFilter": {
        "topicSlugs": [],
        "operator": "IS"
      },
      "acceptanceFilter": {},
      "frequencyFilter": {},
      "frontendIdFilter": {},
      "lastSubmittedFilter": {},
      "publishedFilter": {},
      "companyFilter": {
        "companySlugs": [],
        "operator": "IS"
      },
      "positionFilter": {
        "positionSlugs": [],
        "operator": "IS"
      },
      "contestPointFilter": {
        "contestPoints": [],
        "operator": "IS"
      },
      "premiumFilter": {
        "premiumStatus": [],
        "operator": "IS"
      }
    },
    "searchKeyword": "",
    "sortBy": {
      "sortField": "CUSTOM",
      "sortOrder": "ASCENDING"
    }
  },
  "operationName": "problemsetQuestionListV2"
}

def get_script_abs_path() -> str:
    return os.path.abspath(__file__)

def get_script_abs_folder_path() -> str:
    return os.path.dirname(get_script_abs_path())

def create_category():
    print()

def create_problem_set(name:str):
    print()

# 创建
def create(args):
    match args.create:
        case 'folder': 
            create_folder(args.name)
        case 'file':
            create_file(args.name)
        case 'source':
            create_source(args.name)
        case 'category':
            create_category(args.name)
        case 'problem':
            match args.source:
                case 'leetcode':
                    args.source = 'leetcode-cn'
            create_problem(source=args.source, name=args.name)

# 路径安全检查
def path_safe_check(path: str, panic: bool = True, prefix: str = '', suffix: str = '') -> bool:
    unsafe = ".." in path
    if unsafe and panic:
        print(f'panic: path {prefix+path+suffix} is unsafe')
        exit(1)
    return not unsafe

# 创建目录 (相对于项目根路径, 可递归)
def create_folder(name: str):
    if name is not None:
        path_safe_check(name)
        path = f'{get_script_abs_folder_path()}/../../{name}'
        if exist_path(path) is False:
            os.makedirs(path)
            print(f'{path} created')

# 创建文件
def create_file(path: str, prefix: str = f'{get_script_abs_folder_path()}/../../', suffix: str = '', content: str = None):
    if path is not None:
        path_safe_check(path)
        path = prefix+path+suffix
        if exist_path(path) is False:
            with open(path, 'w', encoding='utf-8') as f:
                if content is not None:
                    f.write(content)
            print(f'{path} created')

# 创建问题源
def create_source(name: str):
    if name is not None:
        s = name.split('/')
        if len(s) == 0:
            return
        name = s[0]
        path_safe_check(name)
        create_folder(f'problems/{name}')
        create_file(path=f'problems/{name}/.gitkeep')

# 创建分类
def create_category(name: str):
    if name is not None:
        path_safe_check(name)
        create_folder(f'category/{name}')
        create_file(path=f'category/{name}/.gitkeep')

# 创建问题
def create_problem(name: str, source: str, content: str = None):
    if name is not None:
        path_safe_check(name)
        for lang in lang_support:
            create_folder(f'problems/{source}/{name}/solution/{lang}')
            create_file(path=f'problems/{source}/{name}/solution/{lang}/.gitkeep')
        create_file(path=f'problems/{source}/{name}/PROBLEM.md', content=content)

def exist_path(path: str) -> bool:
    return Path(path).expanduser().resolve().exists()

def exist_source(name: str) -> bool:
    return exist_path(f'problems/{name}')

def exist_category(name: str) -> bool:
    return exist_path(f'category/{name}')

def exist_problem(name: str, source: str) -> bool:
    return exist_path(f'problems/{source}/{name}/PROBLEM.md')

# 拉取数据源
def fetch(args):
    match args.fetch:
        case 'leetcode':
            fetch_leetcode_cn(args)
        case 'codeforces':
            fetch_codeforces(args)

def fetch_leetcode_cn(args):
    print('fetch leetcode')

def fetch_codeforces(args):
    print('fetch codeforces')

# 获取
def get(args):
    match args.get:
        case 'problem':
            get_problem(args)

def get_problem(args):
    if args.source is None:
        print('error: source arguments (-s, --source) cannot be empty')
        exit(1)
    match args.source:
        case 'leetcode':
            get_problems_in_leetcode_cn(args)
        case 'codeforces':
            get_problems_in_codeforces(args)
        case _:
            print(f'error: illegal source {args.source}')
            exit(1)


def _disp_width(s: str) -> int:
    # wcswidth 对不可见字符可能返回 -1，这里兜底
    w = wcswidth(s)
    return w if w >= 0 else sum(max(wcwidth(ch), 0) for ch in s)

def _wrap_by_disp_width(value: Any, width: int) -> List[str]:
    s = "" if value is None else str(value)
    if width <= 0:
        return [s]
    lines, cur, cur_w = [], [], 0
    for ch in s:
        ch_w = wcwidth(ch)
        if ch_w < 0:
            ch_w = 0
        # 超过列宽则换行
        if cur and cur_w + ch_w > width:
            lines.append("".join(cur))
            cur, cur_w = [ch], ch_w
        else:
            cur.append(ch)
            cur_w += ch_w
    if cur:
        lines.append("".join(cur))
    return lines if lines else [""]

def _pad_cell(s: str, width: int, align: Align) -> str:
    pad = width - _disp_width(s)
    if pad <= 0:
        return s
    if align == "right":
        return " " * pad + s
    if align == "center":
        left = pad // 2
        right = pad - left
        return " " * left + s + " " * right
    return s + " " * pad  # left

def print_table(
    rows: List[Any],
    columns: List[Dict[str, Any]],
    *,
    show_header: bool = True,
    header_sep: bool = True,
    col_sep: str = " ",
) -> None:
    # 规范化列配置
    cols = []
    for c in columns:
        if "width" not in c:
            raise ValueError("Each column must have a 'width'.")
        if ("key" not in c) and ("getter" not in c):
            raise ValueError("Each column must have either 'key' or 'getter'.")
        cols.append({
            "header": c.get("header", ""),
            "width": int(c["width"]),
            "align": c.get("align", "left"),
            "key": c.get("key"),
            "getter": c.get("getter"),
            "first_line_only": bool(c.get("first_line_only", False)),
        })

    # 表头 & 分隔线长度按“列宽+分隔符”计算（不再用 len(header_line)）
    if show_header:
        header_line = col_sep.join(_pad_cell(col["header"], col["width"], col["align"]) for col in cols)
        print(header_line)
        if header_sep:
            total_w = sum(col["width"] for col in cols) + (len(col_sep) * (len(cols) - 1))
            print("-" * total_w)

    for row in rows:
        wrapped_cols: List[List[str]] = []
        for col in cols:
            if col["getter"] is not None:
                val = col["getter"](row)
            else:
                if isinstance(row, dict):
                    val = row.get(col["key"])
                else:
                    val = getattr(row, col["key"])
            wrapped_cols.append(_wrap_by_disp_width(val, col["width"]))

        max_lines = max(len(lines) for lines in wrapped_cols)
        for i in range(len(wrapped_cols)):
            wrapped_cols[i] += [""] * (max_lines - len(wrapped_cols[i]))

        for line_idx in range(max_lines):
            parts = []
            for col, lines in zip(cols, wrapped_cols):
                cell_text = lines[line_idx]
                if col["first_line_only"] and line_idx > 0:
                    cell_text = ""
                parts.append(_pad_cell(cell_text, col["width"], col["align"]))
            print(col_sep.join(parts))


def get_problems_in_leetcode_cn(args, visable: bool = True):
    lock_page = False
    if args.filter is not None: 
        conditions = args.filter.split(';')
        for condition in conditions:
            cmd = []
            if '~' in condition:
                cmd = condition.split('~')
            if '=' in condition:
                cmd = condition.split('=')
            if len(cmd) != 2:
                print(f'Error: filter condition is illegal: {condition}')
                exit(1)
            match cmd[0].lower():
                case 'id':
                    lock_page = True
                    query_payload['variables']['limit'] = 1
                    query_payload['variables']['skip'] = query_payload['variables']['limit'] * (int(cmd[1]) - 1)
                case 'title': # 题目名称需要用户登录后 (cookie 通过鉴权) 才能调用 graphql 接口筛选
                    query_payload['variables']['searchKeyword'] = cmd[1]
                case 'difficulty':
                    difficulties = cmd[1].split(',')
                    available_options = ['EASY', 'MEDIUM', 'HARD']
                    for idx, difficulty in enumerate(difficulties):
                        difficulty = difficulty.upper()
                        if difficulty not in available_options:
                            print(f'Error: filter difficulty level is unsupported: {condition} (only in: easy, medium, hard)')
                            exit(1)
                        difficulties[idx] = difficulty
                    query_payload['variables']['filters']['difficultyFilter']['difficulties'] = difficulties
                case _:
                    print(f'Error: filter condition is unsupported: {condition}')
                    exit(1)
    if not lock_page:
        if args.size is not None:
            size = int(args.size)
            if size > 0:
                query_payload['variables']['limit'] = size
        if args.page is not None:
            page = int(args.page)
            if page > 1:
                query_payload['variables']['skip'] = query_payload['variables']['limit'] * (page - 1)
    response = requests.get(query_url, json=query_payload, headers=headers)
    data = response.json()
    if data is None:
        print('Error: no response form leetcode-cn')
        exit(1)
    # 接口错误输出
    if 'errors' in data and data['errors'] is not None:
            msg = ''
            for error in data['errors']:
                if len(msg) > 0:
                    msg += ';'
                msg += error['message']
            print(f'Error form leetcode-cn: {msg}')
            exit(1)
    data = data["data"]
    if data is None:
        print('Warning: no problem found form leetcode-cn')
        exit(0)
    problems_info = data["problemsetQuestionListV2"]
    # print("问题总数: ", problems_info["totalLength"])

    problem_list = problems_info["questions"]

    columns = [
        {"header": "ID",       "key": "questionFrontendId", "width": 6,  "align": "left",  "first_line_only": True},
        # {"header": "Title",    "key": "title",              "width": 40, "align": "left"},
        # {"header": "Slug",     "key": "titleSlug",          "width": 35, "align": "left"},
        {"header": "Title", "key": "translatedTitle",    "width": 32, "align": "left"},
        {"header": "AC Rate",       "getter": lambda p: f"{p['acRate']*100:.1f}%", "width": 12, "align": "left", "first_line_only": True},
        {"header": "Difficulty",   "key": "difficulty",             "width": 12, "align": "left",  "first_line_only": True},
        {"header": "Status",   "key": "status",             "width": 12, "align": "left",  "first_line_only": True},
    ]

    if visable:
        print_table(problem_list, columns, show_header=True, header_sep=True, col_sep=" ")

    if len(problem_list) == 1:
        get_problem_details_in_leetcode_cn(
            args=args, 
            slug=problem_list[0]['titleSlug'], 
            id=problem_list[0]['questionFrontendId'], 
            name=problem_list[0]['translatedTitle'],
            difficulty=problem_list[0]['difficulty'],
            paid_only=problem_list[0]['paidOnly'],
        )


def get_problem_details_in_leetcode_cn(args, slug: str, id: str, name: str, difficulty: str, paid_only: bool, visable: bool = True) -> str:
    if slug is None:
        print("error: cannot locate to problem")
        exit(1)
    url = problem_desc_url.replace('$problem_name', slug)
    response = requests.get(url)
    soup = BeautifulSoup(response.content, 'html.parser')
    meta_desc = soup.find('meta', attrs={'name': 'description'})
    raw_content = meta_desc['content']
    leetcode_problem_terminal_format(raw_content)

    rows = raw_content.split('\n')
    meta = rows[0].split('-')
    rows[0] = meta[1].strip()
    if not args.readonly:
        save_problem_details_in_leetcode_cn_as_markdown(
            args=args, 
            slug=slug, 
            id=id, 
            title=name, 
            difficulty=difficulty,
            url=url, 
            content_rows=rows,
        )
    return raw_content


def save_problem_details_in_leetcode_cn_as_markdown(args, slug: str, id: str, title: str, difficulty: str, url: str, content_rows: List):
    folder_name = f"{id}. {title}"
    target_path = f"problems/leetcode-cn/{folder_name}"
    difficulty = difficulty.upper()
    content_rows = [
        f"# {folder_name}", 
        leetcode_level_html.
            replace('$color', leetcode_level_color[difficulty]).
            replace('$text', leetcode_level_translate_cn[difficulty]),
    ] + content_rows + [
        "\r\n", 
        "## 参考引用", 
        f"[题目链接 - LeetCode(CN) - {folder_name}]({url})",
    ]
    rows = markdown_compile(content_rows)
    create_problem(name=folder_name, source='leetcode-cn', content='\r\n'.join(rows))
    print(f"problem saved in {target_path}")


def markdown_compile(rows: List[str]) -> List[str]:
    quote_state = False
    for idx, row in enumerate(rows):
        # 图片匹配
        new_row, matched = replace_and_get_old(row, leetcode_image_pattern)
        if matched:
            url = matched[0]
            rows[idx] = f"![{url}]({url})" # markdown 图片
            row = rows[idx]
        # 加粗匹配
        match_info = re.match(leetcode_strong_pattern, row)
        if match_info is not None:
            if match_info.start() > 0:
                row = f"{row[:match_info.start()]} **{match_info.group(0)}** {row[match_info.end():]}"
            else:
                row = f"**{match_info.group(0)}** {row[match_info.end():]}"
        # 引用匹配
        if "输入：" in row:
            quote_state = True
        if row.strip().replace('\n', '').replace('\r', '') == '':
            quote_state = False
        if quote_state:
            row = f"> {row}" # markdown 引用
    return rows

def leetcode_problem_terminal_format(content: str):
    rows = content.split('\n')
    meta = rows[0].split('-')
    rows[0] = meta[1].strip()
    print(meta[0].strip())
    for row in rows:
        # rend photo from url
        row, matched = replace_and_get_old(row, leetcode_image_pattern)
        for m in matched:
            if len(m) > 2:
                print()
                iterm2_imgcat_url(m, width='15%')
                print()
                print(f"图[{m}]")
                row = None
        if row is not None:
            print(row)


def replace_and_get_old(text: str, pattern: str, ns: str = ''):
    olds = re.findall(pattern, text)          # 旧字符串（包含[]）
    new_text = re.sub(pattern, ns, text)      # 替换为空
    return new_text, olds


def iterm2_imgcat_bytes(data: bytes, name: str = "image", width=None, height=None, preserve_aspect=True):
    b64 = base64.b64encode(data).decode("ascii")
    name_b64 = base64.b64encode(name.encode("utf-8")).decode("ascii")

    params = [
        f"name={name_b64}",
        f"size={len(data)}",
        "inline=1",
    ]
    if width is not None:
        params.append(f"width={width}")   # "40", "400px", "50%"
    if height is not None:
        params.append(f"height={height}")
    params.append(f"preserveAspectRatio={'1' if preserve_aspect else '0'}")

    osc = f"\033]1337;File={';'.join(params)}:{b64}\a"
    print(osc, end="")

def iterm2_imgcat_url(url: str, width=None, height=None, timeout=10):
    req = urllib.request.Request(
        url,
        headers=headers
    )
    with urllib.request.urlopen(req, timeout=timeout) as r:
        data = r.read()
    # 尝试从 URL 推断文件名
    path = urllib.parse.urlparse(url).path
    name = os.path.basename(path) or "image"
    iterm2_imgcat_bytes(data, name=name, width=width, height=height)


def get_problems_in_codeforces(args):
    print('codeforces')


def main():
    parser = argparse.ArgumentParser(description="Algorithm solution manager")
    # 算法项目管理
    parser.add_argument("-v", "--version", action="version", version="v1.0.0")
    parser.add_argument("-n", "--create", action="store", help="create categories, problem sets, or problems")
    parser.add_argument("--name", action="store", help="name of item")
    parser.add_argument("-g", "--get", action="store", help="get problems, categories, .etc information")
    parser.add_argument("--fetch", action="store", help="fetch problems, categories, .etc information")
    parser.add_argument("--today", action="store_true", default=False, help="get a daily problem (such as LeetCode, etc.)")
    parser.add_argument("--random", action="store_true", default=False, help="get a daily problem")
    parser.add_argument("-s", "--source", action="store", help="source of problem")
    parser.add_argument("--page", action="store", help="page of list item")
    parser.add_argument("--size", action="store", help="size of list item")
    parser.add_argument("--filter", action="store", help="filter condition of item")
    parser.add_argument("--update", action="store", help="update problems")
    parser.add_argument("-ro", "--readonly", action="store_true", default=False, help="read-only, do not save the document")
    # 账户管理 todo: 待开发
    parser.add_argument("--login", action="store", help="login platform/source")
    parser.add_argument("--account", action="store", help="")
    # 数据管理 todo: 待开发
    # 数据分析 todo: 待开发
    # AI 辅助: 待开发
    parser.set_defaults(func=lambda args: parser.print_help(args))
    args = parser.parse_args()

    if args.create is not None:
        args.func = create
    if args.get is not None:
        args.func = get
    if args.fetch is not None:
        args.func = fetch
    
    args.func(args)

if __name__ == "__main__":
    main()