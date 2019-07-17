export const filled = (...fields) => {

    for (let i = 0; i< fields.length; i++){
        
        if(fields[i] === ""){
            return false 
        } 
    }

    return true
}