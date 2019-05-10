document.querySelector('.download-form').addEventListener('submit', e => {
	e.preventDefault();

	const u = new URL(e.target.querySelector('#url').value);
	location.pathname = u.pathname.slice(u.pathname.indexOf('/p/'))
		+ (u.pathname.endsWith('/') ? '' : '/');
});