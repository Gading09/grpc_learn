package constant

const (
	ACTIVE            = "ACTIVE"
	NONACTIVE         = "NONACTIVE"
	EXPIRED           = "EXPIRED"
	BLOCK             = "BLOCK"
	DITOLAK_BY_SYSTEM = "Ditolak oleh Sistem"
	DITINDAKLANJUTI   = "Ditindaklanjuti"

	/* TODO flow status otp */
	LOGIN_MAGIC_PROCESS = 1
	LOGIN_MAGIC_SUCCESS = 2
	LOGIN_MAGIC_FAILED  = 3

	/* TODO flow status perorangan */
	PRESCREENING_PERORANGAN_PROCESS = 4 // Menunggu Prescreening
	PRESCREENING_PERORANGAN_SUCCESS = 5 // Lolos Prescreening
	PRESCREENING_PERORANGAN_FAILED  = 6 // Gagal Prescreening

	/* TODO flow status badan usaha */
	PRESCREENING_BADAN_USAHA_PROCESS = 7 // Menunggu Prescreening
	PRESCREENING_BADAN_USAHA_SUCCESS = 8 // Lolos Prescreening
	PRESCREENING_BADAN_USAHA_FAILED  = 9 // Gagal Prescreening

	/* TODO status apply */
	NEW             = 10 // Pengajuan Baru dari Webview
	CONTINUE        = 11 // Di Tindak Lanjuti dari Dashboard (Pengajuan dalam proses analisa)
	APPROVED_ADK    = 12 // Tanda Tangan Selesai
	APPROVED_KANPUS = 13 // Persetujuan Kuasa Debet (Menunggu konfirmasi kuasa debet)
	APPROVED_ADMIN  = 14
	BRISPOT_PROCESS = 15
	REJECTED        = 16 // Di Tolak dari Dashboard (Pengajuan gagal)
	APPROVE         = 17
	ACTIVED         = 18 // (Limit Aktif)
	NOT_ACTIVED     = 19
	SUSPEND         = 20
	ACCEPT_CREDIT   = 26 // Lihat Penawaran Kredit dari Dashboard (Pengajuan berhasil, lihat penawaran limit)
	ACCEPT_USER     = 27 // Terima Limit dari Webview (Menunggu proses kelengkapan dokumen tanda tangan Anda)
	REJECTED_USER   = 28 // Tolak Limit dari Webview (Anda menolak penawaran kredit)

	/* TODO status re-apply */
	NEW_REAPPLY             = 101
	CONTINUE_REAPPLY        = 111
	APPROVED_ADK_REAPPLY    = 121
	APPROVED_KANPUS_REAPPLY = 131
	APPROVED_REAPPLY        = 141
	REJECTED_REAPPLY        = 161

	/* TODO flow status lolos prescreening */
	INFORMASI_PEMOHON_SUCCESS   = 10
	INFORMASI_USAHA_SUCCESS     = 21
	PROFILE_KEUANGAN_SUCCESS    = 22
	INFORMASI_PENDUKUNG_SUCCESS = 23
	UPLOAD_DOKUMEN_SUCCESS      = 24
	TNC_PENGAJUAN_SUCCESS       = 25

	/* status untuk tipe reset pengajuan */
	RESET_SLIK                 = 1
	RESET_REJECTED_ADK_OR_USER = 2

	/* TODO gender comchain */
	MALE   = 1
	FEMALE = 2

	/* TODO marital status */
	NOT_MARRY = 1
	MARRY     = 2

	/* TODO pipeline status comchain */
	SUDAH_CAIR = "5"
	DITOLAK    = "6"

	/* response code */
	SUCCESS           = "00"
	FAILED_INTERNAL   = "01"
	FAILED_NOT_FOUND  = "02"
	FAILED_REQUIRED   = "03"
	FAILED_AUTHORIZED = "04"
	FAILED_EXIST      = "05"

	/* Status Apply Micro */
	AFTER_VERIFY_OTP = 1
	ACCEPT_TNC       = 2

	/* Notification Type */
	SMS    = "1"
	WA     = "2"
	EMAIL  = "3"
	PRODUK = "DELIMA OTP"

	/* Status OTP */
	LOCK_ACCOUNT_OTP   = "Anda telah mengirim ulang kode sebanyak 3x, akun akan dikunci selama 1x24jam"
	TRY_AGAIN_OTP      = "Anda Telah Mengirim OTP. Silakan Coba Lagi Nanti"
	TIME_OUT_OTP       = "Pengisian Kode OTP melebihi batas waktu"
	WRONG_INPUT_OTP    = "Kode OTP yang dimasukkan salah/expired"
	SUCCESS_VERIFY_OTP = "OTP Verified"

	/* Alert type otp */
	ALERT  = "01"
	INFO   = "02"
	DRAWER = "03"

	/* Status Token Micro */
	INVALID_REGISTERED_USER = "NIK/Nomor Handphone yang terdaftar tidak sesuai"
	ALREADY_REGISTERED_USER = "NIK/Nomor Handphone sudah terdaftar"

	/* GENDER APPLY BRISPOT */
	MAN   = "L"
	WOMAN = "P"

	// MARITAL STATUS BRISPOT
	SINGLE        = 1
	MARRIED       = 2
	WIDOWER_WIDOW = 3

	/* Status Apply Micro */
	ACCEPT_PLAFON   = 3  // TERIMA LIMIT
	REJECT_PLAFON   = 4  // TOLAK LIMIT
	PROFILE_DONE    = 5  // SUCCESS INFORMASI PEMOHON
	BUSINESS_DONE   = 6  // SUCCESS INFORMASI USAHA
	FINANCE_DONE    = 7  // SUCCESS INFORMASI KEUANGAN
	PARTNER_DONE    = 8  // SUCCESS INFORMASI PASANGAN
	ACTIVATION_DONE = 9  // SUCCESS LIMIT AKTIF BISA TRANSAKSI
	ZERO_DONE       = 10 // DITOLAK OLEH MANTRI
	FIVE_DONE       = 11 // DITOLAK PEMUTUS

	/* otp message */
	OTP_MESSAGE_SMS = "WASPADA PENIPUAN, JANGAN BERI KODE OTP INI KE SIAPAPUN BAHKAN PIHAK BANK"

	/* TUJUAN PENGAJUAN BRISPOT */
	KREDIT_INVESTASI   = "ki"
	KREDIT_MODAL_KERJA = "kmk"

	/* response code */
	SUCCESS_BRISPOT = "0000"

	/* STATUS APPLY BRISPOT */
	ZERO             = "0"
	ONE              = "1"
	ONE_HUNDRED      = "100"
	ONE_HUNDRED_FIVE = "105" // RITEL
	TWO_HUNDRED      = "200" // MIKRO
	TWO_HUNDRED_FIVE = "205"
	FOUR             = "4"
	FIVE             = "5"
	X0               = "X0"
	X1               = "X1"
	X2               = "X2"
	X3               = "X3"
	XP               = "XP"
	XR               = "XR"

	/* JOB APPLY BRISPOT */
	SWASTA  = "1"
	BUMN    = "2"
	PNS     = "3"
	LAINNYA = "4"

	/* status account */
	ACCOUNT_ACTIVE    = "1"
	ACCOUNT_NONACTIVE = "2"
	ACCOUNT_EXPIRED   = "3"
	ACCOUNT_SUSPEND   = "4"
)

const (
	DataPendukungPerpanjanganTableName = "data_pendukung_perpanjangan"
	DataProfilePerpanjanganTableName   = "data_profile_perpanjangan"
	DataUsahaPerpanjanganTableName     = "data_usaha_perpanjangan"
	DataKeuanganPerpanjanganTableName  = "data_keuangan_perpanjangan"
	DataPinjamanPerpanjanganTableName  = "data_pinjaman_perpanjangan"
	DataDokumenPerpanjanganTableName   = "data_dokumen_perpanjangan"
	DataProfileTableName               = "data_profile"
	DataPengajuanTableName             = "data_pengajuan"
	DataPengajuanPerpanjanganTableName = "data_pengajuan_perpanjangan"
	DataPengajuanPeroranganTableName   = "data_pengajuan_perorangan"
	DataPengajuanBadanUsahaTableName   = "data_pengajuan_badan_usaha"
	DataUserTableName                  = "data_user"
)
