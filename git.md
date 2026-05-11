# Git 速查手册(给我自己用)

> 用 HelloGo 这个仓库顺带练 git。重点是建立"知道在做什么"的肌肉记忆,而不是背命令。

---

## 心智模型:四个区域

```
工作区(working tree)  →  暂存区(staging / index)  →  本地仓库(.git)  →  远程仓库(GitHub)
        git add                  git commit                  git push
```

- **工作区**:你正在编辑的文件
- **暂存区**:`git add` 之后、`commit` 之前的中间状态。可以理解为"下一次 commit 会包含哪些改动"的购物车
- **本地仓库**:`.git/` 目录里的所有 commit 历史
- **远程仓库**:GitHub 上的副本

`git status` 会同时告诉你这四个区域之间的差异,**遇事不决先 `git status`**。

---

## 日常流程(90% 的时间在用这套)

```sh
git status                  # 看现在啥情况
git diff                    # 看具体改了什么(未暂存的)
git diff --staged           # 看暂存区里准备 commit 的东西
git add <file>              # 把改动放进暂存区
git add .                   # 把当前目录所有改动放进暂存区
git commit -m "信息"         # 提交到本地仓库
git push                    # 推到远程
```

写 commit message 的简单规则:
- 用一句话说清"做了什么",别写"修改了一些东西"
- 中文英文都行,统一就好
- 例子:`add 06-multiple-results` / `fix typo in 02-imports` / `update README with run instructions`

---

## 拉取远程更新

```sh
git pull                    # = git fetch + git merge,会产生 merge commit
git pull --rebase           # = git fetch + git rebase,保持线性历史(推荐单人仓库)
git fetch                   # 只下载远程的更新,不合并,可以先看再决定
```

**单人仓库默认 `git pull --rebase`**,历史更干净。可以一次性配置:

```sh
git config --global pull.rebase true
```

---

## 分支(branch)

练习项目可以一直在 master 上写,但分支是核心技能,趁早练。

```sh
git branch                  # 看所有本地分支
git branch <name>           # 新建分支(不切过去)
git checkout <name>         # 切到某个分支
git checkout -b <name>      # 新建并切过去(常用)
git switch <name>           # 新版命令,跟 checkout 等价
git switch -c <name>        # 等价 checkout -b

git merge <name>            # 把 name 分支合并到当前分支
git branch -d <name>        # 删除已合并的分支
git branch -D <name>        # 强删(没合并也删)
```

---

## 撤销 / 后悔药

按"后悔强度"从轻到重:

```sh
# 1. 改了文件还没 add,想恢复成上次 commit 的样子
git restore <file>          # 新版命令
git checkout -- <file>      # 老写法,等价

# 2. 已经 add 了想撤回暂存区(从购物车拿出来,文件改动还在)
git restore --staged <file>
git reset HEAD <file>       # 老写法,等价

# 3. 刚刚 commit 完想改 message,或者发现漏 add 了文件
git add <漏的文件>
git commit --amend          # 修改最近一次 commit(还没 push 时安全)

# 4. 想撤销一个已经 commit 的改动,但保留它做的事在工作区
git reset --soft HEAD~1     # 撤回最近 1 个 commit,改动回到暂存区
git reset HEAD~1            # 撤回最近 1 个 commit,改动回到工作区
git reset --hard HEAD~1     # ⚠️ 撤回最近 1 个 commit,改动直接扔掉,不可恢复
```

**`--hard` 是核武器,用之前先 `git status` 想清楚**。

---

## 远程相关

```sh
git remote -v               # 看当前关联的远程
git remote add origin <url> # 添加远程(初次)
git remote set-url origin <url>  # 改远程地址
git push -u origin master   # 第一次推,顺便绑定 upstream
git push                    # 之后直接这样
```

---

## 踩坑速查(实战记录)

> 这一节是我自己踩过的坑。每条都按 **症状 → 原因 → 修法** 的格式,以后碰到直接 Ctrl+F 搜报错关键字。

### 1. `'master' does not appear to be a git repository`(选项写错了)

```
git push -origin master
fatal: 'master' does not appear to be a git repository
```

**原因**:`-origin` 多了个连字符,被 git 当成选项,git 不认识就当远程地址解析失败。
**修法**:
```sh
git push origin master           # 没连字符
git push                         # 已经 -u 过的话直接这样
```

### 2. `Updates were rejected because the remote contains work that you do not have locally`

**原因**:远程比本地超前。**最常见**:在 GitHub 网页建仓库时勾选了 "Add a README" / "Add .gitignore" / "Add license",GitHub 替你生成了一个 commit,你本地是另起炉灶,两条历史分叉。
**修法**:
```sh
git pull --rebase origin master
# 如果两条历史完全没有共同祖先:
git pull --rebase --allow-unrelated-histories origin master
# 解决冲突 → git add <冲突文件> → git rebase --continue
git push
```
远程那个 commit 没价值且没人协作时,也可以直接 `git push --force-with-lease` 覆盖掉。
**预防**:新仓库**不要在 GitHub 网页勾选生成文件**,本地写完直接 push。

### 3. `LF will be replaced by CRLF`(Windows 换行警告)

**原因**:Windows 用 CRLF,Linux/Mac 用 LF。Git 默认 `core.autocrlf=true`,checkout 转 CRLF、commit 转 LF,仓库里永远是 LF。
**修法**:**无视它,不影响功能**。强迫症想关掉,仓库根加 `.gitattributes`:
```
* text=auto eol=lf
```

### 4. `cannot pull with rebase: You have unstaged changes`

**原因**:rebase 要重放 commit,工作区/暂存区有未提交改动会撞车,git 直接拒绝。
**修法**(二选一):
```sh
# A. 先 commit
git add . && git commit -m "..."
# B. 临时藏起来
git stash
git pull --rebase
git stash pop
```

### 5. commit 内容跟 message 对不上(`git add .` 误伤)

**症状**:commit 里多了你没想加的文件,因为 `git add .` 把之前残留在暂存区的东西也一起带走了。
**修法**(commit 完立刻发现,还没 push):
```sh
git reset --soft HEAD~1                    # 撤回 commit,改动留在暂存区
git restore --staged <不想要的文件>          # 从暂存区踢出去
git commit -m "..."                        # 重新 commit
```
**预防**:`commit` 之前永远 `git status` 看一眼**这次到底在提交什么**。

### 6. `.claude/` 这类工具目录被跟踪

**原因**:工具(Claude Code、IDE、依赖管理)会在仓库里生成本地状态目录,没 `.gitignore` 就全跟踪了。
**修法**:
1. 写 `.gitignore`(见下面 .gitignore 那节)
2. 已经被跟踪过的,光加 ignore 没用,要先从暂存区拿掉:
```sh
git rm -r --cached .claude/
git commit -m "stop tracking .claude/"
```

### 7. rebase 冲突,README 被当成 binary file

```
warning: Cannot merge binary files: README.md
CONFLICT (content): Merge conflict in README.md
```

**原因**:文件里有空字节(通常是 UTF-16 with BOM 编码),git 不当文本处理,**连 `<<<<<<<` 冲突标记都不给你插**。
**修法**:
```sh
# 直接告诉 git 用哪一边的版本
git checkout --theirs README.md          # rebase 时:用"正在重放的 commit"版本
git checkout --ours README.md            # rebase 时:用"目标分支"版本(注意!跟 merge 反着!)
git add README.md
git rebase --continue
```
**根治**:把文件转成 UTF-8 no BOM。PowerShell 里:
```powershell
$c = Get-Content xxx.md -Raw
[System.IO.File]::WriteAllText("$PWD\xxx.md", $c, (New-Object System.Text.UTF8Encoding $false))
```

### 8. `fatal: You are not currently on a branch`(rebase 中途的 detached HEAD)

**症状**:rebase 暂停时跑 `git push` 报这个。
**原因**:rebase 暂停期间 HEAD 是 detached 的,不在任何分支上,push 没有目标。
**修法**:**先把 rebase 走完或退出**:
```sh
git status                       # 看冲突在哪
# 解决后:
git add <冲突文件>
git rebase --continue            # 继续走
# 实在不想搞:
git rebase --abort               # 撤销整个 rebase,回到 pull 之前
```

### 9. `--ours` / `--theirs` 在 rebase 和 merge 里是反的

| | `--ours` 含义 | `--theirs` 含义 |
|---|---|---|
| merge 时 | 当前分支 | 被合并进来的分支 |
| **rebase 时** | **rebase 到的目标分支** | **正在重放的你的 commit** |

记不住就 `git rebase --abort` 改用 merge,或者反复试 `--ours` `--theirs` 各试一次看 diff。

### 10. `--force` vs `--force-with-lease`

- `--force`:无脑覆盖。**别人在你 fetch 之后 push 的 commit 会被你抹掉**,在协作仓库里是事故。
- `--force-with-lease`:只有当远程还停在你上次见过的位置时才覆盖。**永远优先用这条**。

### 11. commit 作者是 `你的名字`(全局 config 没改)

**原因**:`git config --global user.name` 当初设的就是占位符字符串。
**修法**(注意用英文双引号):
```sh
git config --global user.name "RoseKhlifa"
git config --global user.email "xxx@users.noreply.github.com"   # GitHub noreply 邮箱保护隐私
```
**已有 commit 也要改**(只在还没 push 或确定可以 force-push 时):
```sh
# 重写某个分叉点之后的所有 commit
git rebase <分叉点 hash> --exec "git commit --amend --no-edit --reset-author"

# 包括根 commit 一起改
git rebase --root --exec "git commit --amend --no-edit --reset-author"

git push --force-with-lease
```

### 12. shell 命令里的中文引号 `“...”` 不是引号

```sh
git config --global user.name “RoseKhlifa”          # 中文引号(U+201C / U+201D)
```

**原因**:`“` `”` 是 Unicode 标点,不是 shell 引号语法。zsh 把它当成普通字符,结果 `user.name` 实际存的是 `"RoseKhlifa"`(包括引号本身)。
**修法**:**永远用英文双引号 `"`**(Enter 旁边那个)。预防:写命令前切到英文输入法,或者关掉输入法的"智能引号"。

### 13. `remote: This repository moved`(仓库被重命名)

**症状**:push 时弹这条提示,push 仍然成功(走重定向)。
**原因**:你在 GitHub 上改了仓库名(包括大小写改动也算),本地 origin 还指着旧 URL。
**修法**:
```sh
git remote set-url origin https://github.com/<user>/<新名字>.git
git remote -v          # 确认
```

---

## .gitignore

不想被 git 跟踪的文件写进根目录的 `.gitignore`:

```
# Go
*.exe
*.test
*.out

# IDE
.vscode/
.idea/

# 系统
.DS_Store
Thumbs.db
```

已经被 git 跟踪过的文件,光加 .gitignore 不够,要先:

```sh
git rm --cached <file>
```

把它从暂存区里拿掉。

---

## rebase vs merge(简版)

- **merge**:把两条历史"拼"在一起,产生一个 merge commit。安全,历史不动。
- **rebase**:把一条历史"挪"到另一条之后,假装你一直在最新的代码上开发。历史更干净,但**改写了 commit 的 hash**。

**规则:别对已经 push 出去的、别人在用的 commit 做 rebase**。自己一个人的仓库随便 rebase。

---

## 看历史

```sh
git log                              # 完整历史
git log --oneline                    # 每个 commit 一行
git log --oneline --graph --all      # 带分支图(好看)
git log -p <file>                    # 看某个文件每次 commit 改了什么
git show <commit-hash>               # 看某次 commit 的具体内容
git blame <file>                     # 看每一行是谁、哪次 commit 改的
```

---

## 高频但容易忘的小操作

```sh
# 把工作区改动暂存起来,先做别的事
git stash
git stash pop                       # 把暂存的拿回来
git stash list                      # 看有几个 stash

# cherry-pick:把某个 commit 单独搬到当前分支
git cherry-pick <hash>

# 改最近一次 commit 的 message
git commit --amend

# 看远程有哪些分支
git branch -r
```

---

## 推荐配置(一次设置好)

```sh
git config --global user.name "你的名字"
git config --global user.email "your@email.com"
git config --global init.defaultBranch main   # 新仓库默认分支用 main
git config --global pull.rebase true          # pull 默认 rebase
git config --global core.editor "code --wait" # 用 VSCode 编辑 commit message
```

---

## 练习清单(配合 HelloGo 用)

- [ ] 每写完一节,单独 commit,message 写清楚是哪一节
- [ ] 主动用一次 `git commit --amend` 改 message
- [ ] 开一个分支写 06-multiple-results,合并回 master
- [ ] 故意制造一次冲突(两个分支改同一行),练解决
- [ ] 故意 `git reset --hard` 一次(确认没东西丢之后),感受核武器
- [ ] 写一份 `.gitignore`,把编译产物挡掉
- [ ] 用 `git stash` 暂存改动去做别的事再回来
- [ ] 用 `git cherry-pick` 把一个 commit 搬到另一个分支