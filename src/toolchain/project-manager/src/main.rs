use clap::{Parser, ArgAction};
use std::path::PathBuf;

/// CLI 参数定义
#[derive(Parser, Debug)]
#[command(
    name = "pm",
    about = "A tool to manage this project",
    version = env!("CARGO_PKG_VERSION"),
    disable_version_flag = true // 关掉 clap 默认的 -V/--version，避免自定义冲突
)]

struct Cli {
    /// 输出版本号并退出（自定义为 -v / --version）
    #[arg(short = 'v', long = "version", action = ArgAction::SetTrue)]
    show_version: bool,

    #[arg(long)]
    pattern: Option<String>,

    #[arg(long)]
    path: Option<PathBuf>,
}

/// 解析命令行参数的方法（可复用到任何 Parser）
fn parse_cli<T: Parser>() -> T {
    T::parse()
}

/// 统一处理“打印版本并退出”的逻辑
fn handle_common_flags(cli: &Cli) {
    if cli.show_version {
        println!("{}", env!("CARGO_PKG_VERSION"));
        std::process::exit(0);
    }
}

fn main() {
    let cli: Cli = parse_cli();
    handle_common_flags(&cli);

    println!("parsed args: {:#?}", cli);
}
