# AIReview

A tool that uses AI to review code and provide feedback about your code. This is project used free model from [openrouter.ai
](https://openrouter.ai/)

> [!NOTE]
> Edit AI model you can in file `internal/api/promt.json`.

## Installation
To install AIReview:
```bash
git clone https://github.com/TheDEKKs/AIReview.git
cd AIReview
go mod tidy
go build -o aireview ./cmd/main.go
```

## Flags
| Flag | Type | Default | Description |
|------|------|---------|-------------|
| `cb` | string | — | Current Branch |
| `cp` | bool | `false` | Custom Prompt |
| `mb` | string | `main` | Main Branch |
| `o` | string | `a.md` | Out File |
| `sp` | string | — | Supplementation Prompt |

## Usage
```bash
./aireview -cb <current_branch> -mb <main_branch> -o <output_file> -sp "<supplementation_prompt>"
```

You can also use custom prompt from `config.json` and field `CustomPromt` by setting `-cp` flag to `true`:

```bash
./aireview -cb <current_branch> -mb <main_branch> -o <output_file> -cp=true -sp "<supplementation_prompt"
```





