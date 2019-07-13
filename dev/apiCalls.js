// this is a script to test the API as I"m building

const { spawn } = require('child_process')

const shell = (comm) => spawn(comm, {shell: true, stdio: "inherit"})

const commandLine = process.argv[2]

const cutty = '{"name":"Cutty","password":"yadayadayada","email":"cutty@example.com"}'

const Elaine = '{"name":"Elaine","password":"yadayadayada","email":"elaine@example.com"}'

const Login = `curl http://localhost:3000/login -X POST -d '${cutty}'`
const LoginElaine = `curl http://localhost:3000/login -X POST -d '${Elaine}'`

const SignUp = `curl http://localhost:3000/signup -X POST -d '${cutty}'`

const SetDB = "dropdb test && createdb test && cd dev && psql -U jonathancannon -d test -a -f sql.sql"

const trans = `curl -H "Authorization: Bearer eyJhbGciOiJIUzI1NiIsInR5cCI6IkpXVCJ9.eyJlbWFpbCI6ImN1dHR5QGV4YW1wbGUuY29tIiwiaXNzIjoiYnJva2VyIn0.sFCN6BlhfHFkGwSWi1aTPpLOBTU9UnS_dTQ1pRPeg2I" http://localhost:3000/trans -d '${cutty}'`

if(commandLine === "elaine"){

	shell(LoginElaine)
}

if(commandLine === 'login'){

	console.log(Login)
	shell(Login)
} 
else if(commandLine === "signup"){

	console.log(SignUp)

	shell(SignUp)
}
else if(commandLine === "setup"){

	shell(SetDB)
}
else if(commandLine === "trans"){

	shell(trans)
}
else if(commandLine === "buy"){

	const buyObj = "('0', 'BUY', 'FB', '1', '100')"
	
	const buyScript = `curl http://localhost:3000/trans -X POST -d '${buyObj}'`
	
	shell(buyScript)
}

// INSERT INTO users (name, email, password) VALUES ('Jerry', 'jerry@example.com', 'yadayada'), ('George', 'george@example.com', 'imbackbaby'), ('Elaine', 'elaine@example.com', 'yaadayadayaada'), ('Kramer', 'kramer@example.com', 'ugotitmadeintheshade');
// INSERT INTO transactions (userID, TYPE, SYMBOL, QUANTITY, PRICE) VALUES ('0', 'BUY', 'FB', '1', '100');
