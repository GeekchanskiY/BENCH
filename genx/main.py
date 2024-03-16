import tkinter as tk

class App(tk.Tk):
    def __init__(self,
                screenName: str | None = None,
                baseName: str | None = None,
                className: str = "Tk",
                useTk: bool = True,
                sync: bool = False,
                use: str | None = None) -> None:
        super().__init__(screenName, baseName, className, useTk, sync, use)

def main():
    app = App()
    tk.Text().pack()
    app.mainloop()
    

if __name__ == '__main__':
    main()