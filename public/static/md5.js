/*!
 * md5 编码解码
 * https://github.com/logotype/es-crypto.git
 */

window.Md5 = class {

	static hexArray = ['0', '1', '2', '3', '4', '5', '6', '7', '8', '9', 'a', 'b', 'c', 'd', 'e', 'f'];

	static hash(string) {
		return Md5.hex(Md5.md51(string));
	}

	static md5cycle(x, k) {
		let a = x[0],
			b = x[1],
			c = x[2],
			d = x[3];

		a = Md5.ff(a, b, c, d, k[0], 7, -680876936);
		d = Md5.ff(d, a, b, c, k[1], 12, -389564586);
		c = Md5.ff(c, d, a, b, k[2], 17, 606105819);
		b = Md5.ff(b, c, d, a, k[3], 22, -1044525330);
		a = Md5.ff(a, b, c, d, k[4], 7, -176418897);
		d = Md5.ff(d, a, b, c, k[5], 12, 1200080426);
		c = Md5.ff(c, d, a, b, k[6], 17, -1473231341);
		b = Md5.ff(b, c, d, a, k[7], 22, -45705983);
		a = Md5.ff(a, b, c, d, k[8], 7, 1770035416);
		d = Md5.ff(d, a, b, c, k[9], 12, -1958414417);
		c = Md5.ff(c, d, a, b, k[10], 17, -42063);
		b = Md5.ff(b, c, d, a, k[11], 22, -1990404162);
		a = Md5.ff(a, b, c, d, k[12], 7, 1804603682);
		d = Md5.ff(d, a, b, c, k[13], 12, -40341101);
		c = Md5.ff(c, d, a, b, k[14], 17, -1502002290);
		b = Md5.ff(b, c, d, a, k[15], 22, 1236535329);

		a = Md5.gg(a, b, c, d, k[1], 5, -165796510);
		d = Md5.gg(d, a, b, c, k[6], 9, -1069501632);
		c = Md5.gg(c, d, a, b, k[11], 14, 643717713);
		b = Md5.gg(b, c, d, a, k[0], 20, -373897302);
		a = Md5.gg(a, b, c, d, k[5], 5, -701558691);
		d = Md5.gg(d, a, b, c, k[10], 9, 38016083);
		c = Md5.gg(c, d, a, b, k[15], 14, -660478335);
		b = Md5.gg(b, c, d, a, k[4], 20, -405537848);
		a = Md5.gg(a, b, c, d, k[9], 5, 568446438);
		d = Md5.gg(d, a, b, c, k[14], 9, -1019803690);
		c = Md5.gg(c, d, a, b, k[3], 14, -187363961);
		b = Md5.gg(b, c, d, a, k[8], 20, 1163531501);
		a = Md5.gg(a, b, c, d, k[13], 5, -1444681467);
		d = Md5.gg(d, a, b, c, k[2], 9, -51403784);
		c = Md5.gg(c, d, a, b, k[7], 14, 1735328473);
		b = Md5.gg(b, c, d, a, k[12], 20, -1926607734);

		a = Md5.hh(a, b, c, d, k[5], 4, -378558);
		d = Md5.hh(d, a, b, c, k[8], 11, -2022574463);
		c = Md5.hh(c, d, a, b, k[11], 16, 1839030562);
		b = Md5.hh(b, c, d, a, k[14], 23, -35309556);
		a = Md5.hh(a, b, c, d, k[1], 4, -1530992060);
		d = Md5.hh(d, a, b, c, k[4], 11, 1272893353);
		c = Md5.hh(c, d, a, b, k[7], 16, -155497632);
		b = Md5.hh(b, c, d, a, k[10], 23, -1094730640);
		a = Md5.hh(a, b, c, d, k[13], 4, 681279174);
		d = Md5.hh(d, a, b, c, k[0], 11, -358537222);
		c = Md5.hh(c, d, a, b, k[3], 16, -722521979);
		b = Md5.hh(b, c, d, a, k[6], 23, 76029189);
		a = Md5.hh(a, b, c, d, k[9], 4, -640364487);
		d = Md5.hh(d, a, b, c, k[12], 11, -421815835);
		c = Md5.hh(c, d, a, b, k[15], 16, 530742520);
		b = Md5.hh(b, c, d, a, k[2], 23, -995338651);

		a = Md5.ii(a, b, c, d, k[0], 6, -198630844);
		d = Md5.ii(d, a, b, c, k[7], 10, 1126891415);
		c = Md5.ii(c, d, a, b, k[14], 15, -1416354905);
		b = Md5.ii(b, c, d, a, k[5], 21, -57434055);
		a = Md5.ii(a, b, c, d, k[12], 6, 1700485571);
		d = Md5.ii(d, a, b, c, k[3], 10, -1894986606);
		c = Md5.ii(c, d, a, b, k[10], 15, -1051523);
		b = Md5.ii(b, c, d, a, k[1], 21, -2054922799);
		a = Md5.ii(a, b, c, d, k[8], 6, 1873313359);
		d = Md5.ii(d, a, b, c, k[15], 10, -30611744);
		c = Md5.ii(c, d, a, b, k[6], 15, -1560198380);
		b = Md5.ii(b, c, d, a, k[13], 21, 1309151649);
		a = Md5.ii(a, b, c, d, k[4], 6, -145523070);
		d = Md5.ii(d, a, b, c, k[11], 10, -1120210379);
		c = Md5.ii(c, d, a, b, k[2], 15, 718787259);
		b = Md5.ii(b, c, d, a, k[9], 21, -343485551);

		x[0] = a + x[0] & 0xFFFFFFFF;
		x[1] = b + x[1] & 0xFFFFFFFF;
		x[2] = c + x[2] & 0xFFFFFFFF;
		x[3] = d + x[3] & 0xFFFFFFFF;
	}

	static cmn(q, a, b, x, s, t) {
		a = (a + q & 0xFFFFFFFF) + (x + t & 0xFFFFFFFF) & 0xFFFFFFFF;
		return (a << s | a >>> 32 - s) + b & 0xFFFFFFFF;
	}

	static ff(a, b, c, d, x, s, t) {
		return Md5.cmn(b & c | ~b & d, a, b, x, s, t);
	}

	static gg(a, b, c, d, x, s, t) {
		return Md5.cmn(b & d | c & ~d, a, b, x, s, t);
	}

	static hh(a, b, c, d, x, s, t) {
		return Md5.cmn(b ^ c ^ d, a, b, x, s, t);
	}

	static ii(a, b, c, d, x, s, t) {
		return Md5.cmn(c ^ (b | ~d), a, b, x, s, t);
	}

	static md51(s) {
		const n = s.length;
		const state = [1732584193, -271733879, -1732584194, 271733878];
		const tail = [0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0];

		let o = 0,
			i = 64;

		for (i; i <= n; i += 64) {
			Md5.md5cycle(state, Md5.md5blk(s.substring(i - 64, i)));
		}
		s = s.substring(i - 64);
		i = 0;
		o = s.length;
		for (i; i < o; i++) {
			tail[i >> 2] |= s.charCodeAt(i) << (i % 4 << 3);
		}
		tail[i >> 2] |= 0x80 << (i % 4 << 3);
		if (i > 55) {
			Md5.md5cycle(state, tail);
			for (i = 0; i < 16; i++) {
				tail[i] = 0;
			}
		}
		tail[14] = n * 8;
		Md5.md5cycle(state, tail);
		return state;
	}

	static md5blk(s) {
		const md5blks = [];
		let i = 0;

		for (i; i < 64; i += 4) {
			md5blks[i >> 2] = s.charCodeAt(i) + (s.charCodeAt(i + 1) << 8) + (s.charCodeAt(i + 2) << 16) + (s
				.charCodeAt(i + 3) << 24);
		}
		return md5blks;
	}

	static rhex(n) {
		let s = '';
		s += Md5.hexArray[n >> 4 & 0x0F] + Md5.hexArray[n >> 0 & 0x0F];
		s += Md5.hexArray[n >> 12 & 0x0F] + Md5.hexArray[n >> 8 & 0x0F];
		s += Md5.hexArray[n >> 20 & 0x0F] + Md5.hexArray[n >> 16 & 0x0F];
		s += Md5.hexArray[n >> 28 & 0x0F] + Md5.hexArray[n >> 24 & 0x0F];
		return s;
	}

	static hex(x) {
		const length = x.length;
		let i = 0;

		for (i; i < length; i++) {
			x[i] = Md5.rhex(x[i]);
		}
		return x.join('');
	}
}