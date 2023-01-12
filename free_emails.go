package gofreeemail

type GoFreeEmails interface {
	IsFreeEmail(domain string) bool
	AddDomainToList(domain string)
	RemoveDomainFromList(domain string)
}

type goFreeEmails struct {
	freeEmails map[string]struct{}
}

func NewGoFreeEmails() GoFreeEmails {
	return &goFreeEmails{
		freeEmails: freeEmails,
	}
}

func (g *goFreeEmails) IsFreeEmail(domain string) bool {
	_, ok := g.freeEmails[domain]
	return ok
}

func (g *goFreeEmails) AddDomainToList(domain string) {
	g.freeEmails[domain] = struct{}{}
}

func (g *goFreeEmails) RemoveDomainFromList(domain string) {
	delete(g.freeEmails, domain)
}
