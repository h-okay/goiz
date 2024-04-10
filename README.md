<p align="center">
  <img src="https://github.com/h-okay/quill/assets/86803100/028a9a4d-c1a6-4172-ab62-d1a63ba0268f" width="100" />
    <h1 align="center" style="border:none; margin:0;">GOIZ</h1>
    <p align="center">Engaging quizes in your terminal.</p>
</p>

## 📍 Overview

`Goiz` delivers an interactive quiz experience, challenging players to answer questions within set time constraints.

## 🚀 Getting Started

**Requirements**

Ensure you have the following dependencies installed on your system:

- **Go**: `version 1.22.1`

### ⚙️ Installation

```sh
git clone https://github.com/h-okay/goiz
cd goiz
go build -o goiz
./goiz --help
```

For pre-built binaries check [releases](https://github.com/h-okay/goiz/releases).

### 🤖 Running goiz

| Paramteter | Type   | Description                                                                                  | Default      |
| ---------- | ------ | -------------------------------------------------------------------------------------------- | ------------ |
| amount     | int    | amount of questions to be loaded from the csv file, if invalid all file contents are loaded. | 5            |
| csv        | string | a csv file in the format of 'question,answer'                                                | problems.csv |
| time       | int    | time limit per question, in seconds                                                          | 3            |

```sh
./goiz -csv=my.csv -amount=3 -time=5
```

---

## 📄 License

This project is protected under the [MIT](https://choosealicense.com/licenses/mit/) License. For more details, refer to the [LICENCE](LICENCE) file.
