# Move to Gogs

## Usage

### How to obtain API token

Login into your [Gogs](https://gogs.io/) instance. Go to **Your Settings** under your User avatar menu. Select **Applications** and click on **Generate New Token**.

### Create organization

```
move2gogs --token-file file-with-token --create-org --org someorg
```

### Mirror local repository

Without project name specified tool will use git directory name by default.
In following example it will be `repo-xyz`.

```
move2gogs --token-file file-with-token --repo /path/to/repo-xyz
```

Or you can specify project name with `--project name` as in following example.

```
move2gogs --token-file file-with-token --repo /path/to/repo-xyz --project foobar
```
