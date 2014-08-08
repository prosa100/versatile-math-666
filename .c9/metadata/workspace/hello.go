{"filter":false,"title":"hello.go","tooltip":"/hello.go","undoManager":{"mark":62,"position":62,"stack":[[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":0,"column":0},"end":{"row":0,"column":13}},"text":"package hello"},{"action":"insertText","range":{"start":{"row":0,"column":13},"end":{"row":1,"column":0}},"text":"\n"},{"action":"insertLines","range":{"start":{"row":1,"column":0},"end":{"row":13,"column":0}},"lines":["","import (","    \"fmt\"","    \"net/http\"",")","","func init() {","    http.HandleFunc(\"/\", handler)","}","","func handler(w http.ResponseWriter, r *http.Request) {","    fmt.Fprint(w, \"Hello, world!\")"]},{"action":"insertText","range":{"start":{"row":13,"column":0},"end":{"row":13,"column":1}},"text":"}"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":1},"end":{"row":14,"column":0}},"text":"\n"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":14,"column":0},"end":{"row":15,"column":0}},"text":"\n"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":15,"column":0},"end":{"row":15,"column":1}},"text":"f"}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":15,"column":0},"end":{"row":15,"column":1}},"text":"f"},{"action":"insertText","range":{"start":{"row":15,"column":0},"end":{"row":15,"column":3}},"text":"fmt"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":15,"column":3},"end":{"row":15,"column":4}},"text":"."}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":15,"column":3},"end":{"row":15,"column":4}},"text":"."}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":15,"column":2},"end":{"row":15,"column":3}},"text":"t"}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":15,"column":1},"end":{"row":15,"column":2}},"text":"m"}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":15,"column":0},"end":{"row":15,"column":1}},"text":"f"}]}],[{"group":"doc","deltas":[{"action":"removeLines","range":{"start":{"row":14,"column":0},"end":{"row":15,"column":0}},"nl":"\n","lines":[""]}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":13,"column":1},"end":{"row":14,"column":0}},"text":"\n"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":12,"column":34},"end":{"row":13,"column":0}},"text":"\n"},{"action":"insertText","range":{"start":{"row":13,"column":0},"end":{"row":13,"column":4}},"text":"    "}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":4},"end":{"row":13,"column":5}},"text":"f"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":5},"end":{"row":13,"column":6}},"text":"m"}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":13,"column":4},"end":{"row":13,"column":6}},"text":"fm"},{"action":"insertText","range":{"start":{"row":13,"column":4},"end":{"row":13,"column":7}},"text":"fmt"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":7},"end":{"row":13,"column":8}},"text":"."}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":8},"end":{"row":13,"column":9}},"text":"F"}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":13,"column":8},"end":{"row":13,"column":9}},"text":"F"},{"action":"insertText","range":{"start":{"row":13,"column":8},"end":{"row":13,"column":14}},"text":"Fprint"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":14},"end":{"row":13,"column":15}},"text":"("}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":15},"end":{"row":13,"column":16}},"text":"\""}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":13,"column":15},"end":{"row":13,"column":16}},"text":"\""}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":13,"column":14},"end":{"row":13,"column":15}},"text":"("}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":13,"column":8},"end":{"row":13,"column":14}},"text":"Fprint"}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":13,"column":7},"end":{"row":13,"column":8}},"text":"."}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":13,"column":4},"end":{"row":13,"column":7}},"text":"fmt"}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":13,"column":0},"end":{"row":13,"column":4}},"text":"    "}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":12,"column":34},"end":{"row":13,"column":0}},"text":"\n"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":12,"column":34},"end":{"row":13,"column":0}},"text":"\n"},{"action":"insertText","range":{"start":{"row":13,"column":0},"end":{"row":13,"column":4}},"text":"    "}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":4},"end":{"row":13,"column":29}},"text":"for i := 0; i < 10; i++ {"},{"action":"insertText","range":{"start":{"row":13,"column":29},"end":{"row":14,"column":0}},"text":"\n"},{"action":"insertLines","range":{"start":{"row":14,"column":0},"end":{"row":15,"column":0}},"lines":["    sum += i"]},{"action":"insertText","range":{"start":{"row":15,"column":0},"end":{"row":15,"column":1}},"text":"}"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":14,"column":0},"end":{"row":14,"column":4}},"text":"    "},{"action":"insertText","range":{"start":{"row":15,"column":0},"end":{"row":15,"column":4}},"text":"    "}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":15,"column":5},"end":{"row":16,"column":0}},"text":"\n"},{"action":"insertText","range":{"start":{"row":16,"column":0},"end":{"row":16,"column":4}},"text":"    "}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":16,"column":4},"end":{"row":16,"column":34}},"text":"fmt.Fprint(w, \"Hello, world!\")"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":16,"column":33},"end":{"row":16,"column":34}},"text":"."}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":16,"column":33},"end":{"row":16,"column":34}},"text":"."}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":16,"column":33},"end":{"row":16,"column":34}},"text":","}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":16,"column":34},"end":{"row":16,"column":35}},"text":" "}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":16,"column":35},"end":{"row":16,"column":36}},"text":"w"}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":16,"column":35},"end":{"row":16,"column":36}},"text":"w"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":16,"column":35},"end":{"row":16,"column":36}},"text":"1"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":16,"column":36},"end":{"row":16,"column":37}},"text":"0"}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":16,"column":36},"end":{"row":16,"column":37}},"text":"0"}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":16,"column":35},"end":{"row":16,"column":36}},"text":"1"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":16,"column":35},"end":{"row":16,"column":36}},"text":"s"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":16,"column":36},"end":{"row":16,"column":37}},"text":"u"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":16,"column":37},"end":{"row":16,"column":38}},"text":"m"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":22},"end":{"row":13,"column":23}},"text":"0"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":23},"end":{"row":13,"column":24}},"text":"0"}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":13,"column":23},"end":{"row":13,"column":24}},"text":"0"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":23},"end":{"row":13,"column":24}},"text":"0"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":24},"end":{"row":13,"column":25}},"text":"0"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":25},"end":{"row":13,"column":26}},"text":"0"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":26},"end":{"row":13,"column":27}},"text":"0"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":12,"column":34},"end":{"row":13,"column":0}},"text":"\n"},{"action":"insertText","range":{"start":{"row":13,"column":0},"end":{"row":13,"column":4}},"text":"    "}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":4},"end":{"row":13,"column":5}},"text":"s"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":5},"end":{"row":13,"column":6}},"text":"e"}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":13,"column":5},"end":{"row":13,"column":6}},"text":"e"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":5},"end":{"row":13,"column":6}},"text":"u"}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":13,"column":4},"end":{"row":13,"column":6}},"text":"su"},{"action":"insertText","range":{"start":{"row":13,"column":4},"end":{"row":13,"column":7}},"text":"sum"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":7},"end":{"row":13,"column":8}},"text":"="}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":8},"end":{"row":13,"column":9}},"text":"0"}]}],[{"group":"doc","deltas":[{"action":"insertText","range":{"start":{"row":13,"column":9},"end":{"row":13,"column":10}},"text":";"}]}],[{"group":"doc","deltas":[{"action":"removeText","range":{"start":{"row":13,"column":9},"end":{"row":13,"column":10}},"text":";"}]}]]},"ace":{"folds":[],"scrolltop":18,"scrollleft":0,"selection":{"start":{"row":13,"column":9},"end":{"row":13,"column":9},"isBackwards":false},"options":{"guessTabSize":true,"useWrapMode":false,"wrapToView":true},"firstLineState":{"row":0,"state":"start","mode":"ace/mode/golang"}},"timestamp":1407533679373,"hash":"dee44fac5384203747789e3818a1c7c065299f7f"}