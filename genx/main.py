import tkinter as tk
from tkinter import Menu, Tk, filedialog, messagebox, Text, ttk

class App(Tk):
    def __init__(self,
                screenName: str | None = None,
                baseName: str | None = None,
                className: str = "Tk",
                useTk: bool = True,
                sync: bool = False,
                use: str | None = None) -> None:
        super().__init__(screenName, baseName, className, useTk, sync, use)
        self.wm_title = "GenX IDE"
        self.textbar = tk.Text(self, wrap="word")
        self.textbar.pack(expand=True, fill="both")
        self.textbar.pack()
        menubar = Menu(self)
        self.config(menu=menubar)
        file_menu = Menu(menubar, tearoff=0)
        menubar.add_cascade(label="File", menu=file_menu)
        file_menu.add_command(label="New", command=lambda: self.textbar.delete(1.0, tk.END))
        file_menu.add_command(label="Open", command=self.open_file)
        file_menu.add_command(label="Save", command=self.save_file)
        file_menu.add_separator()
        file_menu.add_command(label="Exit", command=self.exit_app)
        help_menu = Menu(menubar, tearoff=0)
        menubar.add_cascade(label="Help", menu=help_menu)
        help_menu.add_command(label="About", command=self.about_info)
        self.scrollbar = ttk.Scrollbar(self, orient="vertical", command=self.textbar.yview)
        self.scrollbar.pack(side="right", fill="y")
        self.textbar.configure(yscrollcommand=self.scrollbar.set)
        self.sidebar_frame = ttk.Frame(self, width=200)
        self.sidebar_frame.pack(side="left", fill="y")

        sidebar_label = ttk.Label(self.sidebar_frame, text="File List")
        sidebar_label.pack()

        self.bottom_bar_frame = ttk.Frame(self, height=100)
        self.bottom_bar_frame.pack(side="bottom", fill="x")

        bottom_bar_label = ttk.Label(self.bottom_bar_frame, text="Shell")
        bottom_bar_label.pack()
        sidebar_button = ttk.Button(self, text="Toggle Sidebar", command=self.toggle_sidebar)
        sidebar_button.pack(side="left")

        bottom_bar_button = ttk.Button(self, text="Toggle Bottom Bar", command=self.toggle_bottom_bar)
        bottom_bar_button.pack(side="right")
        
    def open_file(self):
        filepath = filedialog.askopenfilename(filetypes=[("Text files", "*.txt"), ("All files", "*.*")])
        if filepath:
            with open(filepath, "r") as f:
                self.textbar.delete(1.0, tk.END)
                self.textbar.insert(1.0, f.read())

    # Function to save a file
    def save_file(self):
        filepath = filedialog.asksaveasfilename(defaultextension=".txt", filetypes=[("Text files", "*.txt"), ("All files", "*.*")])
        if filepath:
            with open(filepath, "w") as f:
                f.write(self.textbar.get(1.0, tk.END))
    
    def exit_app(self):
        if messagebox.askokcancel("Quit", "Do you want to quit?"):
            self.destroy()
    
    def about_info(self):
        messagebox.showinfo("About", "GenX - Python IDE\nVersion 0.0.1\nCreated by GeekchanskiY")

    # Function to toggle sidebar visibility
    def toggle_sidebar(self):
        if self.sidebar_frame.winfo_ismapped():
            self.sidebar_frame.grid_remove()
        else:
            self.sidebar_frame.grid()

    # Function to toggle bottom bar visibility
    def toggle_bottom_bar(self):
        if self.bottom_bar_frame.winfo_ismapped():
            self.bottom_bar_frame.grid_remove()
        else:
            self.bottom_bar_frame.grid()
        

def main():
    app = App()
    app.mainloop()
    

if __name__ == '__main__':
    main()