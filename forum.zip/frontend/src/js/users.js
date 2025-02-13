let users = [];

async function fetchData(params) {
    let apikey = await getApikey()
        users = await fetch('/aut/users', {Headers})
}
async function getApikey() {
    let JWTheaders = new Headers()
    ///ENDPOINT KHASNA NCHOFO CHNO FIHA .
    const JWTrequest = new Request('/jwt',{
        method: 'GET',
        headers: JWTheaders
    })
}