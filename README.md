CLI that fixes broken srt subtitles
	You need a gemini API key, you can get for free at: https://aistudio.google.com/app/apikey 

INSTALLATION:

To clone the repo:

    git clone https://github.com/M4rcus6/srt-fixer.git
    
    cd srt-fixer

Build executable: 
    go build


SETUP:

	Before using the tool, please set the google API KEY as an env variable:

	For Linux/macOS:
		export GEMINI_API_KEY="YOUR_API_KEY

	For Windows:
		$env:GEMINI_API_KEY="YOUR_API_KEY"

USAGE:

    #On Linux/macOS
    ./srt-fixer fix path/to/your/subtitles.srt

    #On Windows
    ./srt-fixer.exe fix path/to/your/subtitles.srt

    
