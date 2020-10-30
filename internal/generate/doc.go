// Package generate is responsible for generating GoLang representations of the test-network-function claim defined by
// `claim-schema.json`.  A modified version of `https://github.com/a-h/generate` is utilized to emit GoLang structs
// and their corresponding `MarshallJSON` and `UnmarshallJSON` interface implementations.  A modified version was
// utilized since the original version (`https://github.com/a-h/generate`) blindly remaps JSON "object" types to GoLang
// structs.  In this case, we have no struct definition for some facets of our claim, such as "LshwOutput" and
// "JunitResults". This is due to the fact that others define those JSON structures, and we are not in the business of
// redefinition.  This is a rare case in which we want to allow GoLang to represent arbitrary JSON data.  If deeper
// meaning is needed for such properties at a later point, GoLang provides a reflection interface which can aid in
// providing more granular type information.
package generate
