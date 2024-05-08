package bikeshed_test

import (
	"log"
	"testing"

	"github.com/jcbhmr/go-bikeshed"
)

func init() {
	log.SetFlags(0)
}

const Template = `
<pre class='metadata'>
Title: Your Spec Title
Shortname: your-spec
Level: 1
Status: w3c/UD
Group: WGNAMEORWHATEVER
Repository: org/repo-name
URL: http://example.com/url-this-spec-will-live-at
Editor: Your Name, Your Company http://example.com/your-company, your-email@example.com, http://example.com/your-personal-website
Abstract: A short description of your spec, one or two sentences.
Complain About: accidental-2119 yes, missing-example-ids yes
Markup Shorthands: markdown yes, css no
</pre>

Introduction {#intro}
=====================

Introduction here.
`

func TestRun(t *testing.T) {
	html, err := bikeshed.Run(bikeshed.Parameters{
		Text: Template,
	})
	if err != nil {
		t.Fatal(err)
	}
	log.Printf("html=%#+v", html)
}
