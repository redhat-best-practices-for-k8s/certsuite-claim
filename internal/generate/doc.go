// Copyright (C) 2020-2026 Red Hat, Inc.
//
// This program is free software; you can redistribute it and/or modify it under the terms of the GNU General Public
// License as published by the Free Software Foundation; either version 2 of the License, or (at your option) any later
// version.
//
// This program is distributed in the hope that it will be useful, but WITHOUT ANY WARRANTY; without even the implied
// warranty of MERCHANTABILITY or FITNESS FOR A PARTICULAR PURPOSE. See the GNU General Public License for more details.
//
// You should have received a copy of the GNU General Public License along with this program; if not, write to the Free
// Software Foundation, Inc., 51 Franklin Street, Fifth Floor, Boston, MA 02110-1301, USA.

/*
Package generate is responsible for generating GoLang representations of the redhat-best-practices-for-k8s claim defined by
`claim-schema.json`.  A modified version of `https://github.com/sebrandon1/generate` is utilized to emit GoLang structs
and their corresponding `MarshallJSON` and `UnmarshallJSON` interface implementations.  A modified version was
utilized since the original version (`https://github.com/sebrandon1/generate`) blindly remaps JSON "object" types to GoLang
structs.  In this case, we have no struct definition for some facets of our claim, such as "LshwOutput" and
"JunitResults". This is due to the fact that others define those JSON structures, and we are not in the business of
redefinition.  This is a rare case in which we want to allow GoLang to represent arbitrary JSON data.  If deeper
meaning is needed for such properties at a later point, GoLang provides a reflection interface which can aid in
providing more granular type information.
*/
package generate
