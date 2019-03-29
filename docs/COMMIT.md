# commit 规范指南 #

## 简介 ##

本规范翻译自[the AngularJS commit conventions](https://gist.github.com/stephenparish/9941e89d80e2bc58a153).

## 基础 ##

commit message 应该清晰明了, 说明本次提交的目的. commit message 具有以下作用:

1. 提供更多的历史信息, 方便快速浏览

```
git log <last tag> HEAD --pretty=format:%s
```

2. 用于过滤某些 commit, 便于快速查找信息

```
git log <last tag> HEAD --grep feature
```

3. 直接从 commit message 生成 Change Log
4. 指示当前 commit 的作用, 方便跟踪工程历史

## 格式 ##

commit message 应该具有统一的格式, 常用的 [Angular 规范](https://docs.google.com/document/d/1QrDFcIiPjSLDn3EL15IJygNPiHORgU1_OOAqWjiDU5Y/edit#heading=h.greljkmo14y0) 如下:

每个 commit message 都包括三个部分: Header, body 和 footer.

```
<type>(<scope>): <subject>
<BLANK LINE>
<body>
<BLANK LINE>
<footer>
```

其中 Header 部分是必须的, body 和 footer 可以省略. 每一行都不应该超过 72 个字符, 以避免自动换行影响美观.

### Header ###

Header 部分只有一行, 包括三个字段: type(必需), scope(可选) 和 subject(必需).

#### type ####

用于说明 commit 的类别, 只允许使用下面 7 个标识:

- feat: 新功能(feature)
- fix: 修补bug
- docs: 文档
- style: 格式, 不影响代码运行的变动
- refactor: 重构(既不是新增功能, 也不是修改bug的代码变动)
- test: 增加测试
- chore: 构建过程或辅助工具的变动

#### scope ####

用户说明 commit 的影响范围, 比如 model, view, controller 等. 如果修改了多个 scope 可以使用 * 代替.

#### subject ####

commit 目的的简短描述, 不超过 50 个字符. 需注意以下内容:

- 以动词开头, 使用第一人称现在时. 比如应该是 change 而不是 changed 或 changes
- 第一个字母小写
- 结尾不加句号

### Body ###

是对本次 commit 的详细描述, 可以分成多行, 例如:

```
More detailed explanatory text, if necessary.  Wrap it to 
about 72 characters or so. 

Further paragraphs come after blank lines.

- Bullet points are okay, too
- Use a hanging indent
```

需注意以下内容:

- 使用第一人称现在时
- 第2行应该是空行
- 应该说明代码变动的动机, 以及与以前行为的对比

### Footer ###

只用于以下两种情况:

#### 不兼容变动 ####

如果当前代码与上一个版本不兼容, 则 Footer 部分以 BREAKING CHANGE 开头, 之后是对变动的描述, 以及变动理由和迁移方法:

```
BREAKING CHANGE: isolate scope bindings definition has changed.

    To migrate the code follow the example below:

    Before:

    scope: {
      myAttr: 'attribute',
    }

    After:

    scope: {
      myAttr: '@',
    }

    The removed `inject` wasn't generaly useful for directives so there should be no code using it.
```

#### 关闭 Issue ####

如果当前 commit 针对某个 issue, 那么可以在 Footer 部分进行关闭:

```
Closes #234
```

### Revert ###

如果当前 commit 用于撤销之前的 commit, 则必须以 revert: 开头, 后接被撤销的 commit 的 Header:

```
revert: feat(pencil): add 'graphiteWidth' option

This reverts commit 667ecc1654a317a13331b17617d973392f415f02.
```

Body 部分是固定的格式, 其中的 hash 是被撤销的 commit 的 SHA.