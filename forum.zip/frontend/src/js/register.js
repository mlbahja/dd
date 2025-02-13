document.querySelector('#signUpModal .btn_s').addEventListener('click', function handleRegister(a) {

    a.preventDefault();
    console.log('++++++++++')

    const username = document.querySelector("#signUpModal #username").value.trim()
    
    console.log(username,"++++++++++++++++++++++++++++++++++++\n");
    
    const email = document.querySelector("#email").value.trim()
    const password = document.querySelector("#signUpModal #password").value.trim()

    // const first_name = document.querySelector("")
    const confirmPassword = document.querySelector("#confirm-password").value.trim()


    if (username == "" || email == "" || password == "" || confirmPassword == "") {
        displayToast('var(--red)', "all fields are required!!")
        return
    } else if (username.length < 3 || username.length > 20) {
        displayToast('var(--red)', "username must be between  3 and 20 chars !!")
        return
    } else if (email.length < 5 || email.length > 40) {
        displayToast('var(--red)', "email must be between  3 and 40 chars !!")
        return
    } else if (password.length < 6 || password.length > 20) {
        displayToast('var(--red)', "Password must be between  6 and 20 chars !!")
        return
    }
    else if (password !== confirmPassword) {
        displayToast('var(--red)', "Password mismatch")
        return
    }

    fetch("/auth/register", {
        method: "POST",
        headers: { "Content-Type": "application/json" },
        body: JSON.stringify({ username,nickname, email, password }),
    })
        .then(res => {
            if (!res.ok) {
                return res.json().then(errorData => {
                    throw new Error(errorData.Message || 'Something went wrong, please try again');
                });
            }
            return res.json()
        })
        .then(() => {
            //display a popup
            displayToast('var(--green)', 'registered, please login')
            displayPopup("openLogin")
        })
        .catch(err => displayToast('var(--red)', err))

})