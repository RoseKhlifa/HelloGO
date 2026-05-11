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

## 这次遇到的坑:remote 和本地历史不同步

症状:`git push` 被 rejected,提示 `Updates were rejected because the remote contains work that you do not have locally`。

```sh
git pull --rebase origin master                          # 多数情况这条够用
git pull --rebase --allow-unrelated-histories origin master  # 历史完全不相交时用这条
# 有冲突的话:打开冲突文件、改完、git add、git rebase --continue
git push
```

下次开新仓库的避坑姿势:**在 GitHub 网页创建仓库时,不要勾 "Add a README" / "Add .gitignore" / "Add license"**,自己本地写完再 push,历史就一条线。

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