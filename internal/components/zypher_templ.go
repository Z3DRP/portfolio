// Code generated by templ - DO NOT EDIT.

// templ: version: v0.2.778
package components

//lint:file-ignore SA4006 This context is only used if a nested component is present.

import "github.com/a-h/templ"
import templruntime "github.com/a-h/templ/runtime"

import "github.com/a-h/templ"

func ZypherTool() templ.Component {
	return templruntime.GeneratedTemplate(func(templ_7745c5c3_Input templruntime.GeneratedComponentInput) (templ_7745c5c3_Err error) {
		templ_7745c5c3_W, ctx := templ_7745c5c3_Input.Writer, templ_7745c5c3_Input.Context
		if templ_7745c5c3_CtxErr := ctx.Err(); templ_7745c5c3_CtxErr != nil {
			return templ_7745c5c3_CtxErr
		}
		templ_7745c5c3_Buffer, templ_7745c5c3_IsBuffer := templruntime.GetBuffer(templ_7745c5c3_W)
		if !templ_7745c5c3_IsBuffer {
			defer func() {
				templ_7745c5c3_BufErr := templruntime.ReleaseBuffer(templ_7745c5c3_Buffer)
				if templ_7745c5c3_Err == nil {
					templ_7745c5c3_Err = templ_7745c5c3_BufErr
				}
			}()
		}
		ctx = templ.InitializeContext(ctx)
		templ_7745c5c3_Var1 := templ.GetChildren(ctx)
		if templ_7745c5c3_Var1 == nil {
			templ_7745c5c3_Var1 = templ.NopComponent
		}
		ctx = templ.ClearChildren(ctx)
		_, templ_7745c5c3_Err = templ_7745c5c3_Buffer.WriteString("<section id=\"descrips\"><h2>Zypher Encryption Tool</h2><p id=\"zyphash-description\">A tool I made to try to mitigate hash colisions and make it harder for any plain text I hash to be cracked.<br>The Zyphash algorithm is customizable which allows for fine grain control. The idea is a plain text value<br>is first fed through a Zypher then fed through a hash, the iteration count for both of these can be set allowing for<br>multiple rounds of cyphering and hashing. It is a one way encryption process so to validate any passwords, etc <br>the value will have to be fed through the Zyphash algorithm with the same parameters in order for the values to match. Enter a value into the Zyphash section and play around with the settings to see how it works. The Zyphash algorithm lets you customize the following settings:<ul><li>Restrict Hash Shift flag (if set to true will not shift hex digits out of the A-F 0-9 range)</li><li>Hash Iteraction count (the number of times the hashing algorithm should be applied)</li><li>Cypher Shift count (which is the number a character should shift by, if set to 3 A would shift to D)</li><li>Alternate flag (tells the algorithm to shift odd number characters to shift by the negative of shift count)</li><li>Cypher Iteration count (the number of times the plaintext value should have the zypher algorithm appllied to it)</li></ul></p><p id=\"zypher-description\">Another tool I made, but this one is used by the Zyphash alogrithm, the Zypher Cypher encrypts plaintext by<br>feeding it through a customizable cypher.<br>Enter in a value into one of the plain text fields in the forms below and play around with the settings and see how it works.<br>The Zypher Cypher algorithm, lets you customize it in the follwoing settings:<ul><li>Shift count (which is the number a character should shift by, if set to 3 A would shift to D)</li><li>Alternate flag (tells the algorithm to shift odd number characters to shift by the negative of shift count) </li><li>Iteration count (the number of times the plaintext value should have the zypher algorithm appllied to it)</li></ul></p></section><section id=\"demos\"><article id=\"zyphash\"><form hx-get=\"/zyphash\" target=\"zhash\"><label for=\"zyphash-usrval\">PlainText</label> <input id=\"zyphash-usrval\" type=\"text\"> <label for=\"restrict-hash-shift\">Restrict Hash Shifts</label> <input id=\"restrict-hash-shift\" type=\"checkbox\"> <label for=\"hash-iter\">Hash Iteration Count</label> <input id=\"hash-iter\" type=\"number\" step=\"1\"> <button type=\"submit\">Calculate</button></form><div id=\"zhash\" hidden></div></article><article id=\"zypher\"><form hx-get=\"/zypher\" target=\"zyph\"><label for=\"zypher-usrval\">PlainText</label> <input id=\"zypher-usrval\" type=\"text\"> <label for=\"shift-count\">Shift By</label> <input id=\"shift-count\" type=\"number\" step=\"1\"> <label for=\"iteration-count\">Iteration Count</label> <input id=\"iteration-count\" type=\"number\" step=\"1\"> <button type=\"submit\">Calculate</button></form><div id=\"zyph\" hidden></div></article></section><section id=\"video-demo\"></section>")
		if templ_7745c5c3_Err != nil {
			return templ_7745c5c3_Err
		}
		return templ_7745c5c3_Err
	})
}

var _ = templruntime.GeneratedTemplate
