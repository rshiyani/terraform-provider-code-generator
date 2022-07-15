import sys, time, threading, os

def format() :
    os.chdir(os.getcwd()+"/output")
    os.system("gofmt -w .")

def loadingAnimation(process) :
    while process.is_alive() :
        chars = ["🙂","🙂 🙂","🙂 🙂 🙂","🙂 🙂 🙂 🙂","🙂 🙂 🙂 🙂 🙂"] 
        for char in chars:
            sys.stdout.write('\r'+'Formatting '+char)
            time.sleep(.5)
            sys.stdout.flush()
        sys.stdout.write('\n'+'Done ')

def format_all_files():       
    loading_process = threading.Thread(target=format)
    loading_process.start()

    loadingAnimation(loading_process)
    loading_process.join()