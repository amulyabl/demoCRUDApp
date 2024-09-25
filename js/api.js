const API_URL = 'http://localhost:8090/user'; // Change this to your backend URL

async function fetchUsers() {
    const response = await fetch(API_URL);
    return response.json();
}

async function createUser(user) {
    const response = await fetch(API_URL, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(user)
    });
    return response.json();
}

async function updateUser(user) {
    const response = await fetch(`${API_URL}/${user.id}`, {
        method: 'PUT',
        headers: {
            'Content-Type': 'application/json'
        },
        body: JSON.stringify(user)
    });
    return response.json();
}

async function deleteUser(id) {
    await fetch(`${API_URL}/${id}`, {
        method: 'DELETE'
    });
}
