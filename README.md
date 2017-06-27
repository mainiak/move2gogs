# Move to Gogs

## Usage

### How to obtain API token

Login into your [Gogs](https://gogs.io/) instance. Go to **Your Settings** under your User avatar menu. Select **Applications** and click on **Generate New Token**.

### Define proper server URI

Server URI/URL should contain protocol as well.
Using **HTTPS** whenever possible is strongly recommended.

```
move2gogs --server https://foo.example.com
```

### Mirror local repository

Without project name specified tool will use git directory name by default.
In following example it will be `repo-xyz`.

```
move2gogs --server server-uri --token-file file-with-token --repo /path/to/repo-xyz
```

Or you can specify project name as you can see in following example.

```
move2gogs --server server-uri --token-file file-with-token --repo /path/to/repo-xyz --project foobar
```

### Use organizations

You can create organization with following example:

```
move2gogs --server server-uri --token-file file-with-token --create-org --org someorg
```

And then use organization in another command:

```
move2gogs --server server-uri --token-file file-with-token --org someorg --repo /path/to/repo-xyz
```
