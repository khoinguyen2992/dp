package builder

// We want to create a new account in our system
// Before saving account to database, we have to put enough information into it
// We have to generate its password and referral code.
// However, customers want their own format of these information
// So we need a way which helps to construct complex account object and is flexible enough to customize

// Account is the complex object we want to create
type Account struct {
	Password     string
	ReferralCode string
}

// Process defines methods/steps to create an account
type Process interface {
	InitAccount()
	GenPassword()
	GenReferralCode()
	GetAccount() *Account
}

// Director defines how process is used
type Director struct {
	process Process
}

// SetProcess sets process to director
func (d *Director) SetProcess(process Process) {
	d.process = process
}

// Build calls process's methods to create an account
func (d *Director) Build() {
	d.process.InitAccount()
	d.process.GenPassword()
	d.process.GenReferralCode()
}

// CitibankProcess is process to create an account of Citibank
type CitibankProcess struct {
	account *Account
}

// InitAccount ...
func (p *CitibankProcess) InitAccount() {
	p.account = &Account{}
}

// GenPassword ...
func (p *CitibankProcess) GenPassword() {
	p.account.Password = "citi"
}

// GenReferralCode ...
func (p *CitibankProcess) GenReferralCode() {
	p.account.ReferralCode = "CITI"
}

// GetAccount ...
func (p *CitibankProcess) GetAccount() *Account {
	return p.account
}

// HSBCProcess is process to create an account of HSBC
type HSBCProcess struct {
	account *Account
}

// InitAccount ...
func (p *HSBCProcess) InitAccount() {
	p.account = &Account{}
}

// GenPassword ...
func (p *HSBCProcess) GenPassword() {
	p.account.Password = "hsbc"
}

// GenReferralCode ...
func (p *HSBCProcess) GenReferralCode() {
	p.account.ReferralCode = "HSBC"
}

// GetAccount ...
func (p HSBCProcess) GetAccount() *Account {
	return p.account
}
