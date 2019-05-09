document.querySelector('.download-form').addEventListener('submit', e => {
	e.preventDefault();

	const urlField = e.target.querySelector('#url');
	let u;
	try {
		u = new URL(urlField.value);
		if (u.host !== 'www.instagram.com' || !u.pathname.startsWith('/p/')) {
			urlField.setCustomValidity('Must be Instagram post URL');
			return;
		}

		urlField.setCustomValidity('');
	} catch (e) {
		urlField.setCustomValidity(e.toString());
		return;
	}

	location.pathname = u.pathname;
});