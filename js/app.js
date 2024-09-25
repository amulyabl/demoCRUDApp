document.addEventListener('DOMContentLoaded', async () => {
    const userForm = document.getElementById('user-form');
    const userIdInput = document.getElementById('user-id');
    const usernameInput = document.getElementById('username');
    const emailInput = document.getElementById('email');
    const userList = document.getElementById('user-list');

    // Load users
    const loadUsers = async () => {
        const users = await fetchUsers();
        userList.innerHTML = '';
        users.forEach(user => {
            const li = document.createElement('li');
            li.textContent = `${user.username} (${user.email})`;
            li.appendChild(createEditButton(user));
            li.appendChild(createDeleteButton(user.id));
            userList.appendChild(li);
        });
    };

    const createEditButton = (user) => {
        const button = document.createElement('button');
        button.textContent = 'Edit';
        button.onclick = () => {
            userIdInput.value = user.id;
            usernameInput.value = user.username;
            emailInput.value = user.email;
        };
        return button;
    };

    const createDeleteButton = (id) => {
        const button = document.createElement('button');
        button.textContent = 'Delete';
        button.onclick = async () => {
            await deleteUser(id);
            loadUsers();
        };
        return button;
    };

    userForm.onsubmit = async (e) => {
        e.preventDefault();
        const user = {
            id: userIdInput.value,
            username: usernameInput.value,
            email: emailInput.value,
        };

        if (user.id) {
            await updateUser(user);
        } else {
            await createUser(user);
        }

        userIdInput.value = '';
        usernameInput.value = '';
        emailInput.value = '';
        loadUsers();
    };

    await loadUsers();
});
