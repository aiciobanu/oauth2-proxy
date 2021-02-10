package main

import (
	"html/template"
	"path"
	"strings"

	"github.com/oauth2-proxy/oauth2-proxy/v7/pkg/logger"
)

func loadTemplates(dir string) *template.Template {
	if dir == "" {
		return getTemplates()
	}
	logger.Printf("using custom template directory %q", dir)
	funcMap := template.FuncMap{
		"ToUpper": strings.ToUpper,
		"ToLower": strings.ToLower,
	}
	t, err := template.New("").Funcs(funcMap).ParseFiles(path.Join(dir, "sign_in.html"), path.Join(dir, "error.html"))
	if err != nil {
		logger.Fatalf("failed parsing template %s", err)
	}
	return t
}

func getTemplates() *template.Template {
	t, err := template.New("foo").Parse(`{{define "sign_in.html"}}
	<!DOCTYPE html>
	<html lang="en" charset="utf-8">
	  <head>
		<title>Boomerang - Log In</title>
		<meta name="viewport" content="width=device-width, initial-scale=1, maximum-scale=1, user-scalable=no" />
		<link rel="stylesheet" type="text/css" href="https://unpkg.com/carbon-components/css/carbon-components.min.css" />
		<link
		  href="data:image/x-icon;base64,iVBORw0KGgoAAAANSUhEUgAAAOQAAADkCAYAAACIV4iNAAAAAXNSR0IArs4c6QAAJydJREFUeAHtXQl8VcW5nznn3HuzIaBQrKJoBbUi+qz0CSS5SX0uFbUuVavWamtVtK/a1kpVSOIxCbhjxaoP970KT6tVW6u2chNAUKgPFOv2nkABUZA9JHc5Z943QX8NyV3OMtu9mfv7Qe6dmW+Z/5zvzPbNNwjpj0ZAI6AR0AhoBDQCGgGNgEZAI6AR0AhoBLwhgL0V06V6I1DWktjfTeLDXUwOxogcQAgajhHeC2E0iCA0CMpbiKAkQiSFMe7+C2lJQvBnGJNVkLYSEfw+Jsbb11rVH9k2dnvL0L/7HwLaID22edRecLDrpCZC8e8ghCcgQnb3SOqlWAcY6GIw5NcwMl49dUzN4jlnYccLoS5TWghog/TYnlZjogGMsMVj8ZDF8CboRZ8xTfOJzutqEmCsYKv60x8Q0Abpo5WthrbJCLk3+yAJXxSj1QYyHoya5l3b7erPwzPUHFRGQBukz9aJNCR+QRD5rU+y8MUx6oI56qMGMWckW2s+CM9Qc1ARAW2QAVol0pT4T+KiO2HBRjx+GMM6EnooapU3dNj/vi6A+ppEYQTEP1AKg+FHtUhj26WEkLulGCVVFKPt8N9New6O3Lb6ygmdfnTXZdVFQBtkiLaJNLVdTFwyS5pRgu6w4PN/poku7rLr/haiKppUEQS0QYZsCOgpL4Ke8l6ZRglmCWNncn9VVdXkTdeM3RKySppcIgLaIBmAD3PKn8Kc8j65RtldkTUmMi9Ntta+yKBamoUEBLRBMgJdIaOkw9gny83I5VvtCRsZVU+zEYSANkiGQKswp+xRnTWWYVzQ1Rz/a480/VVxBLRBMm4g6CknwfD1HgWGr1AzmFtifPuIoV+f8vEVo8CfVn9UR0AbJIcWgp7yMlh9vUsNowSzRHgZtqwfpuzqdzlUV7NkiIDBkJdm9SUC6eb4PeB++nNVAAHPosNcJ/1Wt5cRHDdRRS+tR18EdOP0xYRZCvSUlxPXncmMIQtGGL9QWWFcsOXa2k0s2GkebBHQBskWzz7crIbEr2HoemufDIkJGKMVIP6sdEv9WxLV0KKzIKCHrFlAYZmUaa27DfYhrmXJMywvOEy9H0F4Hjg1KDOsDlunUqHXPaSgloSjW41wdKtZkDjPYqC3nDPALLtooz1uq2ciXZAbAtoguUHbl3G0MdHiEtLQN0duCmyNfIRNfEbKji+Tq4mWroesAp+BVEtdo2HgmwSK9CQKfHFHuRl3QbSh7QxPBLoQNwS0QXKDNjvjVHPdNdjAd2TPlZpa6SJ3TrQp0WrbRD8XkppCD1klAR9pTMyCnukSSeLzi8Xoxd3Msh/qeWV+mHjkaoPkgaoHnmCMONrU9jD8Pd9DceFFYLHnfcM0T0natR8KF96PBeqhiaTGh4UUcuqY+IXg1va0JBXyioWtkYOdjPNmrGnuCXkL6kymCGiDZAqnP2Y09mq1FT8P9imf90cprPRAx0Uv6P1KYXgjbZDisM4qaa6NMweZQ86CzJezFpCfaBLi3glz3tv1Yg//xtBzSP4Ye5IwfMaC8nUbU3+CwvWeCGQUgp78a2bluWvtsTtkiO8PMrVBKtTKQ+3lVZudDa/BQs9RCqm1iyow930rWhY7uWPquM92ydA/mCCgDZIJjOyYDLyhffCOHc7rsKhyODuubDnBQ7MSW8ZE8Ox5jy1nzU3PIRV7BuixqJgZOQ62HZSNTg4XjYygnj1ldnu9YvAVvTq6h1S0CcvthcMzTlc7PZmhqIpUrSTcO3J2qjX+nMI6FpVqShpk9wLH1vS3sYO/hYh7CASG2Q8TMhTezEMhHsUecO8iXNWGUxAyJgkV6IKHdgXMbd6H9PcNy/jAjJClO6bGPy2qlsiiLPRA38g4bjvcurVXlmxVkhxsmBenm2sfUkWhYtZDOYOky+tgeD+DhzAaBlgw0PfAeF8Fr8zZGbtuQRheMmmjDYlvwmUebfCyGSJTj4KyMZ6caalT6iB2QZ0VLKCcQQ6YvmiPrs6uV2Gl8QhWeIE3zCdwIPeh8orY3dumHPUFK76i+MBL6gjA43WQN1CUzCBy6EkW6jwfhFbT7ERAOYOkag2y3x7U4Wx9BR7Cb7NtKNwJiyUPRcpwa7ENaS17Xg1ynL/AyKGCLSZsuQG+9586pu5SfQN0MFyVNEhalcE3Lh64ffv2l2H4Oi5Y1fJSdSBk3LTn7tatxXRzVMxu/67juM+HHc7nRYZFJsbPHD688pwlk8amWbDrTzyUNUjaCENumjdgyzYHjJJM4NEodChrGvjiYoruTQ8Rw7nFpwAPkwcmzHiCVw91CVxuj04x49kPGCltkBT/bqPc7vwZhq/VvNoDDPOBSmu3qzbbR2zmJYMlX5hTXggry/erEog5Z90g5CQY5RnaKHMi1CdDeYOkGne7lGXW/xmGrzV9asAsAX9qGmRSsrn+BWYsOTKCExi/BKfv2zmKYMMa45f2G7rX9/VVBt7gLAqDpFX50s/zZZ49JZUD2yX37r7boCs/m3w4zDPV/kC4Ddt1yXVqa9mt3cvDrRGnrbD37yoCXaWqWDQGSVESaJQfwXDwh8UQSBiGr3fAS+oKqU+RJ+H4leHWvqdoo8wPVlH5sq63R28fWGWeAHM+rhv98ICPguHxAog6PlX1M4Cp5vgvYauBLvIo/iHHrc6sfIF6YSmuqFT1isogKVIbrq7ZNsCKnQBDy0VckSPIgl6ytdVpe7lqettQrrJCMAccyIHm0AuAxWsh2IgiPWbdpvScI2ctjogSWGxyis4gKcA0GlpVZeXxMN5ezB1wQo7t6iT/YzXOreUuK6AAuoo5aIB1OhjnkoAsxJERcuKyf3Y8pvrIQxwgu0oqSoOkVdh0zdgtFZUmHFPCb+9aJQ6/qHM3Qa9bjYlrYDir5LybjhxipjURhq8fc0CAKUvYV/7BtEziv5gyLRFmSj5cfrDt9n3d0fU3egeiH7rAZTF6blCVdT41gMA8OBLuPCHizIcXyJ4cxTBhbWB8K0Rzn8yEWYkwKXqDpO1A53jJTjIXeq9DRLQL9MrvGWYEYpZOULI3itjz/404mQS42O0mAo9wMnAD3BA2LRyP0qEu2iFrzybYPiW+HhzGjwFDge0K/h9q+I6Teov6lvKX5l9C2q7+HwvCvgJl0j+1aArSGmlK/KdoqarKKwmDpODS0xuWGTua+qcKAZugQeDo/RJcM6fkkKuruf51OA4Fq68YdnDU/hAX3RltbDtbbS3FaFcSQ9aeUJXZb+znZFJtMKfcp2c6z+/0yFG1WXcZjbHKU04Q3rA63ATzyeuD0AqmSSLLOiZj18wTLFcpcSVnkBTdWENilLPzlL3IhY3XBlRVnUFXf5VqYVAm0jD3Segmz1FNryz6fGEiPD7ZWidk6pFFvvSkkhmy9kSSNigsuhwLaSKjAxyzvWP7/LLW+SN66qLC972tEXCHCFqogi4FdNgDwpX8ia6cFyhXstklaZC0tVJ29bsYm8fDV2E9FhyJGp3pSi+KNLYfqdITQ/1Ho+Vlp8Ki1yqV9MqmC2A4EkK4PD9y5kexbPmlnlayBkkbLt1SuwRh40T4KvLkxjA4FvV6WVPbf6j08NBI49BLngSLPErun/bEClaxq1euX/Owqk4YPXVl/b2kDZKClWmJz7cM4xT4KnALgAzIuO5LcP/j91k3WBh+sAn/jomMc+GMmRuGjwha6CnPjjW1t4qQpZKMkjdICjYN0WEa6EwICylyFTTmEjIbjkcpdUtysrX2RUzwVSo9hLl0cYk7JdqYODdXfimml+Qqa66Gija1nwON/Dh4sIh9EWE0NdNSPz2XXjLSYeX1Xlh5vViGbF8yMd4BjTWO9u6+6Iq0sNgHUzJIqeba38Mb6DLhahA0rXs/ULjg3AJHfG3vy2GRpxhOh1TAnvKzNAph7tqUTk6/MkjabOmWunthDiXeuwY25+lpEVUeHRrjxjQNelGssFXooHWnK6/bOjoe6w+LPP3OIOlDQUPew0kD8QsGhNwAUQh+HfTBZE3XZdf+n2EYP2XNlws/Qk6ONLVN5cJbIab9ag7ZG3dYcLkT3ro/753O+ze42v0C4vXM5C3HK3+YT86E+eTlXstLKwerwybGE5PN8b9I04Gz4H7ZQ36F6VQz/gtwRn/6q9+i/sIQ7A4I43ipKHmF5BxoDb0K3sz8oy8UUqRQPizGOYQ8Sf2VCxUt1vx+bZC2jd0DrSHnQ+O9JroBoWe+G44dKTFcpCFAzBguivkkrJDv7mSSvz9zNlE7cnvAB6pfGyTFjD6Mg62hp4lfcSRw5SW6F4bN9IUg/dPVWPeJgfCF0hXxoAAMr8f94Z1ESc4n+/Ucsmfbfxl1YD70XKN6pnP/DvMiAxvn0S0Z7rI8CCiaOK/g5IFNozptx9/0UK2iKdLve8ivWopGHTDN6HHg6yn25mWYF7mu81i0ae6ZX+ki8++B5pDJxTGfhDCdDnl8L3ux0tfz+W1LbZA9EOuyx6/AFj4J9il39EgW8dV0XfxEmZ04WoSwfDLoEN5A1nngZtiVr5wKeXQ0sz7TMUMFXVjpoA2yF5IwBPq7gcl54kNfkEjGIc9E7TYhgbp6VXuXn8nWmg8QwQ27JCr6A7x4JsWa2uiJnpL4aIPM0oyp5vo/QPK1WbL4JkGcHpJx/1Q5beEwvoIKc2+w4rfzvrKhsBbeSjiu+4DK0eW91WJnKW2QOdCC0IQ3wVzqoRzZ3JJhBXFEqrPrRdlzI7olZFjGT2Ck0MmtsuwYD0vucO9gx04eJ22QebA/bJ+qSbAdkshThEsWGOXYz52OJ2WH20/atR9CBRu5VJIxU8DsHBXm4GGrpQ0yD4JLJo1Nl5sRuDNDQnh+Qk6Z7ibsPOoJyTrtsPhvYei6TIiwkEIch9w12l4eDclGKrk2yALwb7UnbMRm5DQJK68IVl4bYk1zTy6gItfsOWdhh2D8M/GLXP6rBY4WB3+Q+UIZ533/NUBIG6QH1HYGzDKEO6HDdXjYcdFjMXvBSA9qcitCw6DArXePcBPAkjEmDSpG/vNaRW2QHpFKN9c+BEM3GQ/lQDeT+sOwW5ZWelSVS7FYufEb6CU3cWHOkikhFZlkRpmTNH6rpg3SB2JDrcqfwXxyuQ8SJkVhweLQjVs23ceEWUAm1JMJFriaApKLJSPke7GG9pPECmUjTRukDxzX2mN3QIAo6uLW4YOMSVG6iginQy5gwiwgk8OGV8zCCP1vQHKhZC52ZxTjiRBtkD4fk1Rr3T/AGVzKWUZ6KU1ZS2J/nyozK05XnbGBi+KUBXWre+6d9h8xq7wgRtogAwCdaok/Lmc+SQY4KfSYzDd/8vr4bPFH1QI0EiUhpOnIWYsjAamlkGmDDAh7RaXxKyD9LCB5YDJ481fDWcCrAzMISQjGSCA4ljLBuvJVB/xc91+2ZseP85VRLQ+mBPoTFIFoU+IHrkueCkofnA6n4VTKOOoIH5xHOEqrYe7rwKE+HBf+1PACWQVHykbRUyz8pYWXoHvIEBimmuuehmNKL4ZgEZCURFDGnSXTtc40jBsDKi+UDEYU+37orL9IqNAQwrRBhgCPkkZi3V4s20Ky8U0Oq65jpzkJKYtLVFka+Q22gJb6VlwCAXjwTCmW27S0QYZ8QDob6v4Jw6IpIdkEIgejnCbzqBYsbN0cSHHxRHuv3PCpErGLClVdG2QhhDzkTzVr7wajfMNDUbZF4PxkqqvrVrZMvXObYNIVV7TCO4XEkoRMkijds2htkJ6hyl2Qnh0Et+DLZFzzBsOx86zGtrrc2vHLmWvjDCJGUbipwVzyyIjd9i1+aLDhrA2SDY70ctilENfxCUbsfLHBxJW2wFJmVTwGL6KiWMFEGXKJL2AlFNYGyRB0syxCD/MKvBh2p/IwlxwXa0hMZFgVz6y22WM3wA78Hz0TSCwIOJ0r20m/UPW1QRZCyEd+V0P1SphL3uODhFlRF6NmZsx8MjIxetAniaTiZMDGrVvOkSTck1htkJ5g8l4oYpq3yBjC0TkSXKFOr24X/rnWqKOX36wRLjiQQLWHrdogAzVqbqIdds1acH96PHcJfjmEuNfz456bc3dALANLqXNurbLnwIvr29GGxDez58pP1QbJoQ0g5MctMkJewIrr4ZadiHOoUkGWLiHPFyykSAEX4VMVUaWPGtog+0ASPiFlT3gfXOr+HJ6Tfw7YQVL22xrMukXwElrvX2PxFBi53xMv1ZtEbZDecPJdyiDoAd9EDAhgSPb9AdMX7cGAlS8WdNgKQ/U/+SKSVJggfJRMD6d81dYGmQ+dEHlj9ql8ARZ3Pg/BIihprLOz84KgxGHowJVOgqN9EI0JTnelvhOEkjeNNkhOCHefrkdyFjqgp7qYU7Xysq2yoq/AUD2Tt5AimbAAJmWuXaj62iALIRQm3zTmhCEPSkvjk4KjwKig9EHpNtrjtmKC3g1KL5QOI22QQgFXQNhUVP2mpGErcjE+XhIEb0mS60ssIfiQITfNG+CLSEBh3UNyBHnnQgeRstBBkCvHIA2jKAySBqHe0olGc2z+QKy1QQaCzTsRHE+SYpCgYb2UAE+ELPaOjuSSrjtGsgZ9xGuD7AMJ2wTTKHuDLUeP3AiqWrp6+ziPpZkVqzbj78DCjvK3L3dXmJBDmFWcESNtkIyAzMWm0x63Gh7QdbnyeaaDo7vwB24unJGE7Y9PeNaLFW+ISjecFS9WfLRBskIyLx8sZV6FETkgr1qcMuGY02pOrJmyhRfH15kyZMBMGyQDEAuxMDBaVqgMj3yX4G/w4FuIJxzULgqDhHpogyzUmKWYDz3GP2XUCwxDSg+JjeLoIcHNcKCMdsknU/eQ+dBhlGcgLKXHgBfBPoyq4IsN7PFJqa8vJWlhjCO+aTgTaIPkDDBlD0NHOQ8oTCIFVK+PCPDWgbAexfAhyl1/rg1SwHNjxozNAsT0FUGQ0zdRQIoBh8CK4UOQ7iGLoZ1Y64gzpqwHVIqjNxzEklVff02HcdIfAf/SuofkjzHCkZSUBxROfUgxSGQ4UurruykJEX4FRCEdtUEWQohBPkZRlwGbICwkGWRx9JDg1qgNMshTVew0mbQr6VQBFn5/ZXdbQYCd4mgzrA2yOBqKrZYucoew5eiRGyYfeizJtBhsewxmypATM3hrfMqJdWC2esgaGDrvhIabGeq9NLuShBhyDBIjKfX1ixzMsVf5peFdXhskb4SBv0uMYQLE9BFhGEiKQcKDXhQGCT25Nsg+T01/SMDkYCnVNMwPZMglLpIzRPdZWQOTlT5JuBfXPSR3iOFsOkGHChCzqwgINrUX2uvjXRMF/SqSISsxrfcEIeJZjDZIz1CFKUiEh4oA97WFK+z9ZR0UFh5gy3frwAvrQDRIG6Rv4IqcoDsgL0HDRVcDggG/JlomlWfbxKBR72TI9iMT5rkfLLdHK3evpe4h/bRigLLpZLI+AFl4Egu9Gp6Jfw43oDe+AYP0cv+UoinwUtESvcjTBukFpRBlYIHjOyHIg5FivLUGxd8MRhyOirgZ4cPzQBpjtCAQHWcibZCcAYZ4OkfzFpGF/+s0tk2WdO5JLnGLwiBhjt3GHYwAArRBBgDNK0m0cd5oOJUufIFD7h0bWHikO6/t8a9yeFOyOa5khHVtkP9qJfbfsHMme6aFOOLOAWZ0dqFSPPLPnE1M2ONRMkT/LvXFaB5E5FPS31Yb5C4txfYHrDYKN0hYPXyW3rHBtibeuD333vwjoKRycWp6aw+BFF7unabKb22QnFoiYrd9C4arwuOimpbxMKcqFWRLMq74BayCWvUtYJnWi31T1UjRBsmrHTLupbxY5+ILc8d/XoNq/pYrn3s6JvXcZYQUAEPVdzrtmlUh2XAj1wbJAdrd7YW7wQTlXA6s87KE8IuP0gt+8hbilDnUXl4FIbWU7yFh6qhs70ibRhskhwd0m9P1Y2BbyYF1HpY4bRqx/8pTgGvWFnfDiUXhEECsJ7gCEZK5NsiQAPYm38/+pAwWc37TO533bwhH8Vj3PSK8BeXg70pYwMqhSs5kGK6+lWqpWZ6zgAIZ2iAZN8KazKpJwHJvxmzzs8PYNYh5c/5C/HKH3bK0EoarE/lJYMQZo0cYceLGRhskQ2iHz1hQTjC5hiFLj6zIs8nWGilnH6mCm7ZtPkn54SrGqXIj8nuPgEorZkmTXIKC121M/wyqtafoqmFk3ihaZk95rosu7Plbze/kj1vtCRvV1O1fWuke8l9YhPrWPWxD6OpQTIIQY/yXdEvtkiCkLGhiLfPphT7HsuDFk4dJTOWHq7T+2iAZPQVfbN18OQzbxMaSgbkjRob4l0APzEjKuQTqjXskqfj1s/FWjbLeOT0B0wbZE42A37sPISMi3DDACh6G3lHaub6RMz+KQYhL5YerBsaPyTr94veR0gbpF7Es5VNdyRmwyjgoSxbPpI5IGW7gKaAQ71Ub1p4B9VY8oBXsz5rmnYXqokq+NsiQLRFrmnss+KwK98qB2cZNO6bGpQX6hTpjOHwtfL/Vb3PBKOJJlV3letdHG2RvRHz8pk4AsCF+tw8SNkUxWv01q+I2NsyCcSlrnHciQeSwYNSiqDDBFpa2Pxukltogg6D2Jc1aZ9VU8MoZGYJFIFKYE1211h67IxAxIyIXu1MYseLHBpPnU3b8PX4C2HPWBhkQ06i94GDw4hY/ZMPoxVRz3dMB1WZCVma318OQdTwTZhyZyN6fDVI1bZABUOsOdZhJ3Qen4wVfiY23RcyyywKozJQk47hNTBnyYTYXVqAX8WHNj6s2yADYtmbar4bjVTUBSEORYANfLdOBnCofs9u/Cy+i74SqiABi08I3CBDDXIQ2SJ+QRhrbj4SN8Ot9koUuDquF81LX10o7XkUrQFdWXceR6qbnBUg41bEkade94qWsamW0Qfpokb3sxRUIOXCejkR8kIUvilGXgayLZAdmijW1/xAWsQ4PXyG+HEzTkODgz6ZO2iB94Lje6bgNHsiDfJAwKQrucZNlnuaglaBeOQS5LUwqxJMJxi912bVSrlFgUS1tkB5RjDW0nwRDNuFxchDGL6Rb4r/zqCa3YivWr7kaXkb7cRPAgjFcoGMQNJkFK1k8tEF6QL5iWtvXHew+4KEo4yL403Kz8kLGTH2zizUkRoGLnPL7jhBp9b5Ua90/fFdQIQJtkAUagwb/TXW6T8KKxtcKFGWcjYll4B9ts8duYMzYNzsHkbuBKOabUCQB3GcSqzCuEymShyxtkAVQff7dNrqiWl+gGPNsuI785q7m+F+ZM/bJMNqYoH66x/gkE1+coOnbp8TXixfMVqI2yDx4xprajofT8MKHanSLY8zelY15VBOSNch+e5BLyAwhwkIIAbxWDrf2vSMEC2VItUHmaIry1gV7O8R9XPjhW4zXRq2KM5dMGpvOoZqw5I7MVrrnOEyYwICCMDaukXhbdECts5Npg8yCS71NrExX6inhZ/0gEBNscZzeYf/7uixqCU2K2PPGgTcSRANQ/AMhTFIt8acU19Kzetogs0C1wGmfJsU1DuOfq+B/SV9IyMmAV5DyoTk6rJg1KUsTFm2SNsheTQcHjk+GeZPwvSzwwrk33Ry/r5c6Un7Oy7T/qhg8cuAuk6ldDdUrpYDESag2yB7AlrUk9ncIelR0z0AXcQ40h0CQLPmfqD3/UFQEHjnwAls01YoXTWgOry2rDfJLpKhrmJNCc0THxoEH68NyK3rKcnt0ymuj8SrX7R6XydC7L9Tec0Q4jU3rIlkXC/HCn/LVgZK/RHfl52tuh3kjnOQQ+cHrjag1cWujGgF8V25YC3Nn1cNyQDQhg9yYsqvfFdlSomTpHhKQppvfYIyCD/7iTmyZ30s2Vv+vqMbOJ6fMThwNQauuzFdGhTy4VOj9fYfsPU0FXXjo0O8NcmcoDjKLB7g5edLLcRA6L23XLMxZRmDGwBvaB2cc8ojouXOAKjrExD/9+IpRyQC0RUHSrw2Snm8kmRSdN1aJbC244/5X4AT9rEiZ+WTt6HDvAQyG5yujQp5h4NaMXbdABV146dCvDXK9s/0uGKrCqqLAD0bXpVvqZwqUmFdUtLHtPJg3/iBvIQUyYfFr/imHxtU/jxkSK1hx75+fSFP7T4jrPCiy9uDiNQPONv5apMx8sspa54/IdKXpVQQD85VTIG+LVRY5vNT2HLPh2i97SFjEGUNc965sgPBKg8WI+1UyRuqN4yS7tzhUN0ZkYOPS/mCM9NnrdwY55KZ5A2CY+t+wgFHOy/h68wWPkqenmnVKuXh1uwcSUt1bV9V+w4vs4VLyVS2Eb7/bh9yyPfMAuIUdWAgYZvkYP3vY8Mof2ZMwxFVW4wMRACY6EtwD/dYejPHjQebQy4v+kKOPivcrg4w0zr0CjPFMH/iEKgoP1JxqM37u3Ek4E4oRQ+Ly1sQ+6S7yKLBUfP0Ap5FpnrPeHr2dYfWVZ9Vvhqzdx4kIvlVUi4AxPnXqmLpzVLqXcOexMkSPKu0hCoegcjAmV8E+7eKg9MVK1y8McoC9eAjKOLNh3igknirMGZ8AYzxvzlnYUenBWOAmpsMWxwSVdMqmC7zMHldpayibjrzSSt4gbZsYnZmOJ+BB3IcXiD35wsP0MJxCOF81Y6RhLF0XX9VTVxW/w37j28MGR9U/GM0JvJKfQ0532q+DnvE4Tvjtwhbu3rgtdX18MjxUsJCrzqfcnrdv2nEegch5is8b0RemGT199ZXjO9VBT6wmJd1DwmHjE+CwsaBgUfiadHPdVaoZ45GzFkcymczTYIy7i320fEtzLMs8u8sev8I3ZQkRlGwPWWa/sV/GST0uwGHaASO8JN1SJ9Trx+szuGx1x83QXY/zWl5aOYymFPMVAKxwK8kesvuwsZP6b/69Au40DOP7qhpjtCFxOlx/8EtWDwsvPnR7KNNSX1RXj/PCoiQNcuX6Nb+DB5HvYWOM1sH99fWp5vjzvBonDN9Yy/wDXESU7LV71guMcenuuw3+Sc+0/vy95Awy0pT4KWz+X8SzUWGI+k7EtI5K2/E3ecoJyns/+5MyN5WZA/Sq+6musWLREz+bfHhH0LqWGl1JGSS9TJW4hK/TOMZ/HlhlVnfaNatUfRjWOKvugBHCEarqt1MvvA1j88TOhglr1NZTrHYlY5ADpi/aAxHnGYCPW4Am6Bl/e9qY+Mkbrq7ZJraZvEuLNiXgUlWi+j6eY1rGWRCDlh790p8eCKi+L9VD1dxf6eZ/q9P2MiziHJu7VIgcjHfAm+viVEvdkyG4cCeFRZxvwrzxLRBUyV1YCAGwX3spbBGJDZsSQl+RpCXRQ053E628jBEWHT4GkMapbozd4UgwPVamtjHC2cabtTHmNvGiN8ho09zTwCWMz53yGL1YaQ78NhjjO7khVCMHrlu/B4aqh6ihTXYt6PZGsrmWT1tlF1l0qUXtGNAdMS6TfoTD5n8SHMSvhi2Nmap53mR7wiKNiQvBGM/PlqdKGuDYvre57/nFgKdMzIp2Drm7vXC3bU7XItjiOJglgPDAvIdNfE7Kji9jyZcXLxqOxCVoEbyUhEVA8FsXeLn9vaqq8uhN14zd4pe2v5UvyiEr9AZ4q5N8lK0xYgLGePewwZGxxWKMQ+3lVXCKBfYbFTZGCGwcq8Df1cbo7dVSlENW2PxvhDiip3irYuFSMExYiCz8C7rRv7pwcWVKbHY2zIKX0kHKKNRLEcB1pWWWHbt9yrj1vbL0zxwIFF0PSXtHuNT0c6jPy/AvuIcHRhugR3wSDLEu3Vo/XlWvmxzthmDeeAlgcW6ufOnp4FpoWNFjOu1xxfSOUwA26SoEV4CGpJjvtI2B8DBHwbDtIBi+jYT5ygHwoO4GaRWQVoEwduBNvQ561HVwGvBTkLYYDPrVVHPN34t1gSFiz/834qTfgDqVBUePJyXeZGBUVwyr0zxRCMK7aBd1glS2FGi+XMxaAkPVkYrWpwNb1jGq3FuiKEY51Sq6IWvOmvSTDFjMekhlY4RDxidpYwz+MGqDDI6dcMpIQ9uvwCPpdOGCvQnsgPn4RDhkPNdbcV0qGwJ6yJoNFQXTrMa54xHBCZgXC4mc5wsCjLYj0zohY9fM80WnC/dBQPeQfSBRL6E7jCVCs5U0RoS3IRMfr42RzXOjDZINjty4fBXGUsn7GzHeCrdAH1fqdzZya9wsjK0saTpJIQRanUQDqCMkjKXPam+B7aPjYQEH3Pb0hxUCeg7JCkkOfCCM5bEOwfScp2ojmS9ga+O7/THUP4dm3oWlNshd4FDnR/elOEnydxiqDlFHK9AEo9UGwcfBlez/UEqvElFGtTdvicAarhrdwY27yGzVjBF8Fj+A4F7V2hjDtW8+aj2HzIeOpDwIbjxDteDGMJRaXGZWnbDNHrtBEiz9QqwesirWzNHGtrNd4v5eKbUw/utgc8ip/e2uRhltoIesMlDPITNqtx0Cxnh/jmw5yRCn5yBzyERtjGLg1wYpBueCUroPGzsuDWOpTMQ4OA3zuwaz7gfL7dGpghXQBZggoOeQTGAMz2Szs/5BthEQQukEFwihK+HOkpl2KDaa2C8C2iD9IsahPHUaJ8Q9kwNr/yzBL9XExtnJ5vhL/ok1RVgE9KJOWARD0oPTeC2w+Btscch/OcIeI0ZwfEpHFA/ZqsHJ9RwyOHahKSumtX0dmND9RunGSCPDRekFQtoYQ7drGAbaIMOgF4KWhh9Jd7rUGPcMwYYNKcbPDrUqa3fYNWvZMNRcgiIg/c0cVPFip5vvtt0Km/81UuuBMVwFghvSzbU3Fmt8Ian4cRCu55AcQC3EUpHN/y9MC5+btOteKaSvzheHgDZIcVh3S4ra8w91M+mF8EPafiP0hm+bZvT0Lnv8CsHV1+IKIKDnkAUAYpk9+MbFA4mTeRZ4yjNGhB+BOzYmaGNk2bLseOkekh2WeTnRAM+Rprbn4WzjyXkLcsvEnbDZ/0vY7L+XmwjNODQCelEnNITeGHx5/YEUY4QtjWXY6r5A6D1v2upSshDQPaQA5GMNiYkQPv0FGSf/Yb44c8TQvX7z8RWjkgKqqkWEREAbZEgAC5HHWuYf4KTSi2G/cVChsmzz8XrTwD/RLnBsUeXNTS/qcER42C1LK91k+jnhxghHpmJW+WHaGDk2LifWeg7JCVjKduPWzQ/C5v+hHEXswhqGp/8wTeMKiB7+WmaXHP2jWBDQPSSnlqKHjaFnZHq7cy5VYd4xz0DGaVPN+KHUGHOV0+nqI6DnkJzbqNyet2/GdY+HBZ16+FcNPeaI8CLhtmdElsJNmc8YBD2jg06FR1QVDtogBbdE1fS2oZkUGeM4aCTGZKRL8EhMwEgxqQS/0gq445LeaVkO1wYYkP4ZqPcpWN862LpYBweYV5uWsaQcVS3cbB+xWbDqWpxGQCOgEdAIaAQ0AhoBjYBGQCOgEdAIaAQ0Agoh8P+Zv29rvcQ9xwAAAABJRU5ErkJggg=="
		  rel="icon"
		  type="image/x-icon"
		/>
		<style>
		  body {
			display: flex;
			flex-direction: column;
			height: 100vh;
			background: white;
			font-family: "IBM Plex Sans", "Helvetica Neue", Helvetica, Arial, sans-serif;
			line-height: 1.5;
		  }
	
		  nav {
			display: flex;
			align-items: center;
			position: fixed;
			top: 0;
			left: 0;
			height: 3rem;
			width: 100vw;
			z-index: 1;
			background-color: #1b2834;
		  }
	
		  .logo {
			height: 2rem;
			width: 8.125rem;
			margin-left: 1.375rem;
		  }
	
		  .signin {
			display: flex;
			justify-content: center;
			align-items: center;
			position: absolute;
			top: 0;
			right: 0;
			left: 0;
			bottom: 0;
			height: 100vh;
			width: 100vw;
			padding: 1rem;
			background-color: rgba(23, 23, 23, 0.5);
		  }
	
		  .form {
			position: relative;
			display: flex;
			flex-direction: column;
			height: 100%;
			max-height: 24rem !important;
			max-width: 36rem;
			width: 100%;
			background: rgb(243, 243, 243);
			border-radius: 0.25rem;
		  }
	
		  .form__body {
			display: flex;
			flex-direction: column;
			padding: 1rem;
		  }
	
		  .errorNotification {
			bottom: 4rem;
			padding-bottom: 0 !important;
			position: absolute;
			width: 95%;
		  }
	
		  .hidden {
			display: none;
		  }
	
		  .bx--text-input {
			background: white;
			border-radius: 0.25rem;
			border-bottom: none;
		  }
	
		  .bx--text-input:focus,
		  .bx--text-input:active {
			outline-color: rgb(0, 114, 195);
		  }
	
		  .bx--modal-footer {
			background: transparent;
		  }
	
		  .bx--btn {
			border-bottom-right-radius: 0.25rem;
			border-bottom-left-radius: 0.25rem;
		  }
	
		  .bx--btn--primary {
			background-color: rgb(0, 114, 195);
		  }
	
		  .bx--btn--primary:hover {
			background-color: rgb(0, 114, 195);
		  }
		  .bx--btn--primary:focus,
		  .bx--btn--primary:active {
			background-color: rgb(0, 58, 109);
			border-color: rgb(0, 114, 195);
		  }
	
		  .svg-container {
			display: flex;
			flex-direction: column;
			justify-content: flex-end;
			height: 100%;
		  }
		</style>
	  </head>
	  <body>
		<nav>
		  <svg
			width="885px"
			height="190px"
			viewBox="0 0 885 190"
			version="1.1"
			xmlns="http://www.w3.org/2000/svg"
			alt="Boomerang Logo"
			class="logo"
		  >
			<title>Boomerang Logo</title>
			<g stroke="none" stroke-width="1" fill="none" fill-rule="evenodd">
			  <path
				d="M42.2694007,148.748317 C42.2694007,148.748317 48.6879071,140.757344 56.7594995,129.988419 C64.3688641,119.838478 74.6780371,108.0419 73.8380746,93.6039256 C73.0090465,79.3613146 60.3847579,66.2121423 51.7455695,57.0509807 C42.3369953,47.0754715 36,41.0012556 36,41.0012556 C36,41.0012556 45.6043998,37.5385339 55.1660559,43.4921426 C69.5994945,52.4778756 93.9146697,74.5888588 94.956422,92.4865649 C96.1492682,112.990782 72.5786279,137.690334 58.0129821,147.676807 C51.4423877,152.181136 42.2694007,148.748317 42.2694007,148.748317 Z M21.4630736,133.112607 C21.4630736,133.112607 26.0315742,127.446311 31.7778442,119.810351 C37.1947805,112.612651 44.5334993,104.247594 43.93608,94.0101521 C43.3456892,83.9111083 34.359298,74.5862911 28.2083894,68.0906101 C21.5112688,61.0172679 17,56.7098803 17,56.7098803 C17,56.7098803 23.8376898,54.2548198 30.6442536,58.4759594 C40.9188615,64.8482855 58.2279562,80.527377 58.968957,93.2178736 C59.8183969,107.757689 43.039449,125.272058 32.6704589,132.352421 C27.9935192,135.546607 21.4630736,133.112607 21.4630736,133.112607 Z"
				id="icon"
				fill="#F2F4F8"
			  ></path>
			  <path
				d="M124.36,146 L124.36,42.4 L142.28,42.4 L142.28,84.96 L142.84,84.96 C145.5,76.56 153.34,71.24 163,71.24 C181.48,71.24 191.7,84.96 191.7,109.32 C191.7,133.82 181.48,147.68 163,147.68 C153.34,147.68 145.64,142.08 142.84,133.82 L142.28,133.82 L142.28,146 L124.36,146 Z M157.26,132.84 C166.5,132.84 172.94,126.12 172.94,115.76 L172.94,103.16 C172.94,92.8 166.5,85.94 157.26,85.94 C148.72,85.94 142.28,90.56 142.28,97.7 L142.28,120.94 C142.28,128.5 148.72,132.84 157.26,132.84 Z M236.34,147.68 C215.48,147.68 202.46,132.84 202.46,109.32 C202.46,85.94 215.48,71.24 236.34,71.24 C257.34,71.24 270.36,85.94 270.36,109.32 C270.36,132.84 257.34,147.68 236.34,147.68 Z M236.34,133.26 C245.72,133.26 251.6,127.24 251.6,116.18 L251.6,102.6 C251.6,91.68 245.72,85.66 236.34,85.66 C227.1,85.66 221.22,91.68 221.22,102.6 L221.22,116.18 C221.22,127.24 227.1,133.26 236.34,133.26 Z M314.16,147.68 C293.3,147.68 280.28,132.84 280.28,109.32 C280.28,85.94 293.3,71.24 314.16,71.24 C335.16,71.24 348.18,85.94 348.18,109.32 C348.18,132.84 335.16,147.68 314.16,147.68 Z M314.16,133.26 C323.54,133.26 329.42,127.24 329.42,116.18 L329.42,102.6 C329.42,91.68 323.54,85.66 314.16,85.66 C304.92,85.66 299.04,91.68 299.04,102.6 L299.04,116.18 C299.04,127.24 304.92,133.26 314.16,133.26 Z M380.92,146 L363,146 L363,72.92 L380.92,72.92 L380.92,85.1 L381.62,85.1 C384.42,77.4 390.3,71.24 401.22,71.24 C411.16,71.24 418.86,76.14 422.36,85.94 L422.64,85.94 C425.3,77.82 433.42,71.24 445.04,71.24 C459.32,71.24 467.44,81.6 467.44,99.8 L467.44,146 L449.52,146 L449.52,101.62 C449.52,91.12 445.74,85.94 437.76,85.94 C430.76,85.94 424.18,89.86 424.18,97.7 L424.18,146 L406.26,146 L406.26,101.62 C406.26,91.12 402.48,85.94 394.5,85.94 C387.64,85.94 380.92,89.86 380.92,97.7 L380.92,146 Z M515.86,147.68 C494.3,147.68 481.42,132.7 481.42,109.32 C481.42,86.22 493.88,71.24 515.3,71.24 C538.12,71.24 548.62,88.04 548.62,108.06 L548.62,113.94 L500.04,113.94 L500.04,115.76 C500.04,126.26 506.06,133.4 517.82,133.4 C526.64,133.4 531.96,129.2 536.44,123.46 L546.1,134.24 C540.08,142.5 529.44,147.68 515.86,147.68 Z M515.58,84.68 C506.2,84.68 500.04,91.68 500.04,101.76 L500.04,102.88 L530,102.88 L530,101.62 C530,91.54 524.82,84.68 515.58,84.68 Z M581.36,146 L563.44,146 L563.44,72.92 L581.36,72.92 L581.36,88.04 L582.06,88.04 C583.88,80.34 589.76,72.92 601.38,72.92 L605.3,72.92 L605.3,89.86 L599.7,89.86 C587.66,89.86 581.36,93.22 581.36,101.2 L581.36,146 Z M633.98,147.68 C619.14,147.68 610.74,139.14 610.74,126.12 C610.74,111.14 622.08,103.72 641.96,103.72 L654.42,103.72 L654.42,98.4 C654.42,90.28 650.36,85.52 640.84,85.52 C632.44,85.52 627.68,89.72 624.18,94.76 L613.54,85.24 C618.86,76.84 627.4,71.24 642.1,71.24 C661.84,71.24 672.34,80.62 672.34,97.28 L672.34,131.72 L679.62,131.72 L679.62,146 L669.68,146 C661.98,146 657.22,140.96 656.1,133.4 L655.26,133.4 C652.88,142.78 644.9,147.68 633.98,147.68 Z M640,134.66 C647.98,134.66 654.42,131.02 654.42,124.16 L654.42,114.5 L642.94,114.5 C633.56,114.5 628.94,117.72 628.94,123.46 L628.94,125.84 C628.94,131.72 633,134.66 640,134.66 Z M711.1,146 L693.18,146 L693.18,72.92 L711.1,72.92 L711.1,85.1 L711.8,85.1 C714.74,77.26 720.9,71.24 732.24,71.24 C747.22,71.24 755.62,81.6 755.62,99.8 L755.62,146 L737.7,146 L737.7,101.62 C737.7,91.26 734.2,85.94 725.66,85.94 C718.24,85.94 711.1,89.86 711.1,97.7 L711.1,146 Z M840.44,152.16 C840.44,167.42 830.08,175.68 802.08,175.68 C777.44,175.68 768.06,169.24 768.06,158.46 C768.06,150.34 772.96,145.72 781.08,144.46 L781.08,142.92 C775.06,141.38 771.7,136.62 771.7,130.74 C771.7,123.18 778.14,119.4 785.14,117.86 L785.14,117.3 C776.32,113.24 771.84,105.68 771.84,96.02 C771.84,81.18 782.34,71.24 802.08,71.24 C806.56,71.24 811.18,71.8 814.82,73.06 L814.82,70.68 C814.82,63.54 818.46,60.04 825.18,60.04 L836.52,60.04 L836.52,73.76 L820.98,73.76 L820.98,75.72 C828.54,79.92 832.32,87.06 832.32,96.02 C832.32,110.72 821.96,120.52 802.08,120.52 C797.88,120.52 794.1,120.1 790.88,119.26 C788.08,120.52 785.7,122.62 785.7,125.84 C785.7,129.48 788.78,131.3 795.22,131.3 L814.82,131.3 C832.74,131.3 840.44,139 840.44,152.16 Z M823.64,154.26 C823.64,149.64 820.28,146.84 811.18,146.84 L786.26,146.84 C783.6,148.8 782.48,151.46 782.48,154.4 C782.48,160 786.68,163.36 797.88,163.36 L806.84,163.36 C818.46,163.36 823.64,160.42 823.64,154.26 Z M802.08,108.34 C810.62,108.34 815.24,104.42 815.24,97.14 L815.24,94.76 C815.24,87.34 810.62,83.56 802.08,83.56 C793.54,83.56 788.92,87.34 788.92,94.76 L788.92,97.14 C788.92,104.42 793.54,108.34 802.08,108.34 Z"
				id="boomerang"
				fill="#F2F4F8"
				fill-rule="nonzero"
			  ></path>
			</g>
		  </svg>
		</nav>
		<main class="signin">
		  <form class="form" method="POST" action="{{.FormAction}}">
			<div class="bx--modal-header">
			  <p class="bx--modal-header__heading bx--type-beta">G’day! Log in to Boomerang</p>
			</div>
			<div class="form__body">
			  <input id="rd" type="hidden" name="rd" value="{{.Redirect}}" />
			  <label class="bx--label" for="username">Username</label
			  ><input
				class="bx--text-input bx--text__input"
				type="text"
				name="username"
				id="username"
				size="10"
				autocomplete="off"
			  /><br />
			  <label class="bx--label" for="password">Password</label
			  ><input
				class="bx--text-input bx--text__input"
				type="password"
				name="password"
				id="password"
				size="10"
				autocomplete="off"
			  /><br />
			  {{if .Error}}
			  <div
				id="errorNotification"
				role="alert"
				kind="error"
				class="errorNotification bx--inline-notification bx--inline-notification--error"
			  >
				<div class="bx--inline-notification__details">
				  <svg
					focusable="false"
					preserveAspectRatio="xMidYMid meet"
					xmlns="http://www.w3.org/2000/svg"
					width="20"
					height="20"
					viewBox="0 0 20 20"
					aria-hidden="true"
					class="bx--inline-notification__icon"
					style="will-change: transform;"
				  >
					<path d="M10,1c-5,0-9,4-9,9s4,9,9,9s9-4,9-9S15,1,10,1z M13.5,14.5l-8-8l1-1l8,8L13.5,14.5z"></path>
					<path d="M13.5,14.5l-8-8l1-1l8,8L13.5,14.5z" data-icon-path="inner-path" opacity="0"></path>
					<title>closes notification</title>
				  </svg>
				  <div class="bx--inline-notification__text-wrapper">
					<p class="bx--inline-notification__title">Login Failed</p>
					<div class="bx--inline-notification__subtitle">We can't find that username and password</div>
				  </div>
				</div>
				<button
				  type="button"
				  aria-label="closes notification"
				  title="closes notification"
				  class="bx--inline-notification__close-button"
				>
				  <svg
					focusable="false"
					preserveAspectRatio="xMidYMid meet"
					aria-label="close notification"
					xmlns="http://www.w3.org/2000/svg"
					width="20"
					height="20"
					viewBox="0 0 32 32"
					role="img"
					class="bx--inline-notification__close-icon"
					style="will-change: transform;"
				  >
					<path
					  d="M24 9.4L22.6 8 16 14.6 9.4 8 8 9.4 14.6 16 8 22.6 9.4 24 16 17.4 22.6 24 24 22.6 17.4 16 24 9.4z"
					></path>
				  </svg>
				</button>
			  </div>
			  {{end}}
			</div>
			<div class="bx--modal-footer">
			  <button class="bx--btn bx--btn--field bx--btn--primary" type="submit" class="btn">Log In</button>
			</div>
		  </form>
		</main>
		<div class="svg-container">
		  <svg xmlns="http://www.w3.org/2000/svg" viewBox="0 0 932.39 241.45" alt="Sydney skyline">
			<title>Sydney Skyline</title>
			<defs>
			  <style>
				.cls-1 {
				  fill: #f2f4f8;
				}
				.cls-2 {
				  fill: #a2a9b0;
				}
				.cls-3 {
				  fill: #009d9a;
				}
				.cls-4 {
				  fill: #3ddbd9;
				}
				.cls-5 {
				  fill: #878d96;
				}
				.cls-6 {
				  fill: #c1c7cd;
				}
				.cls-7 {
				  fill: #dde1e6;
				}
				.cls-8 {
				  fill: #121619;
				}
			  </style>
			</defs>
			<title>Skyline-color</title>
			<rect class="cls-1" x="163.07" y="97.51" width="9.36" height="100.5" />
			<polygon
			  class="cls-1"
			  points="170.54 87.27 164.97 87.27 160.86 87.27 160.86 88.78 174.64 88.78 174.64 87.27 170.54 87.27"
			/>
			<rect class="cls-1" x="165.97" y="75.55" width="3.57" height="9.72" />
			<rect class="cls-2" x="171.54" y="78.67" width="1.4" height="6.6" />
			<rect class="cls-2" x="162.57" y="78.67" width="1.39" height="6.6" />
			<polygon
			  class="cls-2"
			  points="175.64 90.78 159.86 90.78 157.92 90.78 157.92 95.51 177.58 95.51 177.58 90.78 175.64 90.78"
			/>
			<rect class="cls-3" x="94.97" y="140.44" width="2.57" height="57.57" />
			<rect class="cls-4" x="99.54" y="140.44" width="4.63" height="57.57" />
			<path class="cls-4" d="M74.88,168.22h6.83a1,1,0,0,1,1,1V198H93V140.44H74.88Z" />
			<path class="cls-5" d="M269.59,198V127.4a1,1,0,0,0-1-1h-2.73V74.25H273V198Z" />
			<rect class="cls-6" x="265.86" y="60.5" width="7.16" height="11.75" />
			<rect class="cls-6" x="256.07" y="60.5" width="7.8" height="11.75" />
			<rect class="cls-6" x="275.02" y="60.5" width="7.8" height="11.75" />
			<rect class="cls-7" x="275.02" y="74.25" width="15.69" height="123.76" />
			<rect class="cls-7" x="248.18" y="74.25" width="15.69" height="52.15" />
			<path
			  class="cls-4"
			  d="M233.08,102.22l-.39-.08v1.79a15.44,15.44,0,0,1,1.59.21,10.15,10.15,0,0,0,2.27.21v-1.78A12.59,12.59,0,0,1,233.08,102.22Z"
			/>
			<path
			  class="cls-4"
			  d="M375.49,119.68l-.13,0v1.75a11.41,11.41,0,0,1,1.26.21,5.28,5.28,0,0,0,1.6.2V120A8.32,8.32,0,0,1,375.49,119.68Z"
			/>
			<path
			  class="cls-4"
			  d="M364.82,119.68l-.13,0v1.75a11,11,0,0,1,1.26.21,5.32,5.32,0,0,0,1.6.2V120A8.32,8.32,0,0,1,364.82,119.68Z"
			/>
			<rect class="cls-6" x="460.44" y="97.92" width="24.28" height="5.61" />
			<path class="cls-6" d="M363.56,129.06l.13,0,.13,0h10.41l.13,0,.13,0h5.15v5.36h-20v-5.36Z" />
			<rect class="cls-1" x="56.83" y="170.22" width="23.87" height="27.79" />
			<rect class="cls-1" x="456.37" y="105.54" width="39.85" height="92.47" />
			<path
			  class="cls-1"
			  d="M391.59,164.53a.5.5,0,0,1-.13,0h-3.87V198h17.25V164.5H391.72A.5.5,0,0,1,391.59,164.53Z"
			/>
			<path class="cls-1" d="M347.66,164.53l-.13,0H308.16V198h49.45V164.5h-9.82A.5.5,0,0,1,347.66,164.53Z" />
			<rect class="cls-1" x="232.69" y="128.4" width="34.9" height="69.6" />
			<polygon class="cls-5" points="551.78 194.42 897.86 194.42 901.47 198.01 546.97 198.01 551.78 194.42" />
			<polygon
			  class="cls-6"
			  points="895.85 192.42 554.46 192.42 558.74 189.22 620.36 189.22 620.37 189.22 620.37 189.22 624.49 189.22 624.49 189.22 624.5 189.22 631.57 189.22 631.58 189.22 631.58 189.22 767.98 189.22 767.99 189.22 767.99 189.22 868.61 189.22 868.61 189.22 868.62 189.22 892.63 189.22 895.85 192.42"
			/>
			<path
			  class="cls-6"
			  d="M583.7,130.1c-17.87,17.41-23.26,50.65-24.15,57.12h18c1.21-5.2,8.94-35.56,30.9-67.75C603.77,120,591.9,122.11,583.7,130.1Z"
			/>
			<path
			  class="cls-6"
			  d="M630.56,114.13c-21.72,25-32.9,65.54-34.84,73.09h7.58A265,265,0,0,1,611.65,162c5.05-13,13.33-31.13,24.33-44.16Z"
			/>
			<path
			  class="cls-6"
			  d="M741.7,112.22a111.94,111.94,0,0,1,7.54,14c1.82-3,6-9.39,10.37-13.31l-12.89-7.57C746.23,105.65,744.8,106.88,741.7,112.22Z"
			/>
			<path class="cls-6" d="M802.87,120.3c-16,17.33-30.76,59-33.49,66.92h1.28c2.64-5.11,23.44-44.34,49.9-63.14Z" />
			<path
			  class="cls-6"
			  d="M625.9,187.22h5a264.64,264.64,0,0,1,14.42-26.6c8.09-13.05,20.59-30.27,34.68-39.4l-7.59-5.91C645.32,135.24,628.79,179.09,625.9,187.22Z"
			/>
			<path
			  class="cls-3"
			  d="M631,112l6.36,4.32a93.51,93.51,0,0,1,43.19-25.66c-1.26-4.39-6.66-22.91-10.11-30.73-38.6,35.56-59.38,68.45-70,89.84-9.8,19.72-13,33.53-13.83,37.47h7.12c1.58-6.25,13.06-49.23,36-75.08A1,1,0,0,1,631,112Z"
			/>
			<path
			  class="cls-3"
			  d="M777.25,58.15c-27.11,13.13-41,43.48-43.84,50.18,1.69,0,3.42,0,5.2.11a1,1,0,0,1,.75.38s.43.56,1.11,1.54c3.6-6.07,5.22-7,6.33-7.06a1,1,0,0,1,.56.14l13.93,8.18a142.51,142.51,0,0,1,14.94-9Z"
			/>
			<path
			  class="cls-3"
			  d="M823.89,121.73v-24c-36.6,8.07-51.88,51.76-62,80.87.16,2.8.26,5.68.28,8.63h5.11c1-2.81,4.81-13.94,10.29-26.71,8.5-19.82,16.66-33.95,24.23-42a1,1,0,0,1,.94-.29l20,4.28C823.12,122.26,823.5,122,823.89,121.73Z"
			/>
			<path
			  class="cls-3"
			  d="M671.88,113.24a1,1,0,0,1,1.2,0l8.79,6.85a122.44,122.44,0,0,1,19-7.55c1.33-1.7,2.67-3.38,4-5-1.48-2.07-6.54-9-10.73-13-47.18,20.7-64,62.46-72.47,92.68h2.09C626,181,643.09,133.89,671.88,113.24Z"
			/>
			<path
			  class="cls-3"
			  d="M655.45,187.22h19.77c2.84-4.33,25.25-36.59,70-47.2A101.39,101.39,0,0,0,735.1,118.2C685.78,131,659.88,178.42,655.45,187.22Z"
			/>
			<path
			  class="cls-3"
			  d="M819.24,187.22h10.59c2.75-5.46,25.55-49.27,57.84-67.3l1.13-3.84C881.66,117.7,850.42,128,819.24,187.22Z"
			/>
			<path
			  class="cls-1"
			  d="M682.23,122.17C657.54,137.28,637,179,633.14,187.22h20.08c2.95-6.07,29.51-57.9,82.08-71.13a1,1,0,0,1,1.08.42c16.72,25.62,19.47,64.15,19.82,70.71h4c-.29-45.54-19.57-73.35-22.1-76.8C713.44,109.52,697.91,114.47,682.23,122.17Z"
			/>
			<path
			  class="cls-1"
			  d="M853.34,113.47a119,119,0,0,0-22.5,6.84h0a32.88,32.88,0,0,0-7.18,4.06l0,0,0,0c-25.74,17.13-46.53,55-50.66,62.8H817C853,118.14,888.63,114,890.13,113.81a1,1,0,0,1,.86.36,1,1,0,0,1,.19.92L870,187.22h4.21c12-43.15,20.32-71,22.1-76.9C883.47,109.2,868.29,110.31,853.34,113.47Z"
			/>
			<path
			  class="cls-1"
			  d="M824.7,95.51a1,1,0,0,1,.83.2,1,1,0,0,1,.36.78v24c1.07-.59,2.25-1.19,3.56-1.78V89.52a137.71,137.71,0,0,0-51.87,14.69h0a140.67,140.67,0,0,0-15.65,9.4c-5.24,4.1-10.59,13-11.7,14.9a147.31,147.31,0,0,1,11.27,45.1C771.44,145.31,788,102.73,824.7,95.51Z"
			/>
			<path
			  class="cls-1"
			  d="M731.24,108.33c2-5,16.72-39.11,46.63-52.68a1,1,0,0,1,1,.07,1,1,0,0,1,.45.86l-1,45.07,1.08-.53c.56-23,.89-42.61,1-48C752.63,63.29,727,83,704.13,111.67A110.62,110.62,0,0,1,731.24,108.33Z"
			/>
			<path
			  class="cls-1"
			  d="M694,92.44a1,1,0,0,1,1.07.18c4.21,3.84,9.29,10.7,11.19,13.34.59-.72,1.19-1.43,1.79-2.13-5.75-9-11.84-13.41-13.23-14.34A124.17,124.17,0,0,0,682,92.28s0,0,0,0h0c-17.62,4.85-31.75,13.28-43.79,26-19.39,22.33-30.78,61.47-32.82,68.94h14.23c3.63-13.07,9.78-32.17,21.38-50A118.18,118.18,0,0,1,694,92.44Z"
			/>
			<path
			  class="cls-1"
			  d="M670.06,57.51a1,1,0,0,1,1.57.3c3.51,7.14,9.46,27.58,10.81,32.29q1.53-.41,3.09-.78C679.36,71.11,673.64,57.07,672,53.19c-15.5,9.81-35.91,31.9-60.71,65.69-22.13,31.73-30.18,62.06-31.68,68.34h4.88C586,179.63,599.24,122.31,670.06,57.51Z"
			/>
			<rect class="cls-4" x="115.06" y="121.61" width="22.27" height="76.39" />
			<path class="cls-4" d="M305,198h1.12V163.5a1,1,0,0,1,1-1h2V79.17H305Z" />
			<rect class="cls-4" x="298.93" y="79.17" width="4.11" height="118.84" />
			<rect class="cls-4" x="292.71" y="79.17" width="4.23" height="118.84" />
			<path class="cls-4" d="M323.38,147.31l.24-.16a1,1,0,0,1,.54-.15h3.33V79.17h-4.11Z" />
			<path class="cls-4" d="M311.16,162.5h2V152.43a1,1,0,0,1,1-1h1.1V79.17h-4.11Z" />
			<polygon class="cls-4" points="317.27 151.22 321.38 148.59 321.38 79.17 317.27 79.17 317.27 151.22" />
			<rect class="cls-4" x="323.38" y="68.03" width="4.11" height="9.14" />
			<rect class="cls-4" x="317.27" y="68.03" width="4.11" height="9.14" />
			<rect class="cls-4" x="311.16" y="68.03" width="4.11" height="9.14" />
			<rect class="cls-4" x="305.04" y="68.03" width="4.11" height="9.14" />
			<rect class="cls-4" x="298.93" y="68.03" width="4.11" height="9.14" />
			<path class="cls-4" d="M329.49,147h1.38a1,1,0,0,1,.54.15l2.62,1.68V79.17h-4.54Z" />
			<rect class="cls-4" x="329.49" y="68.03" width="4.54" height="9.14" />
			<rect class="cls-4" x="292.71" y="68.03" width="4.23" height="9.14" />
			<path class="cls-4" d="M181.79,97.51h-7.36V198h2.66V117.08a1,1,0,0,1,.64-.94c1.42-.55,2.77-1.14,4.06-1.75Z" />
			<polygon
			  class="cls-4"
			  points="153.71 97.51 153.71 198.01 161.07 198.01 161.07 97.51 156.92 97.51 153.71 97.51"
			/>
			<rect class="cls-4" x="444.42" y="159.69" width="9.96" height="38.31" />
			<rect class="cls-4" x="438.7" y="159.69" width="3.72" height="38.31" />
			<rect class="cls-4" x="432.97" y="159.69" width="3.72" height="38.31" />
			<rect class="cls-4" x="427.25" y="159.69" width="3.72" height="38.31" />
			<rect class="cls-4" x="419.36" y="159.69" width="5.89" height="38.31" />
			<rect class="cls-4" x="348.66" y="142.94" width="41.93" height="4.05" />
			<rect class="cls-4" x="348.66" y="149" width="41.93" height="4.05" />
			<polygon
			  class="cls-4"
			  points="358.61 136.42 348.66 136.42 348.66 140.94 390.59 140.94 390.59 136.42 380.64 136.42 358.61 136.42"
			/>
			<path class="cls-4" d="M348.66,162.5h9.95a1,1,0,0,1,1,1V198h26V163.5a1,1,0,0,1,1-1h4V155H348.66Z" />
			<path
			  class="cls-4"
			  d="M532.19,128.89V198h11.53a1.06,1.06,0,0,1,.17-.2l13.58-10.13a147.75,147.75,0,0,1,4.68-21.76c3.34-11.29,9.54-26.91,20.16-37.26.22-.21.44-.42.67-.62A182.47,182.47,0,0,0,532.19,128.89Z"
			/>
			<path
			  class="cls-4"
			  d="M570.53,125V110.49a93,93,0,0,0-12.27-.75,81.57,81.57,0,0,0-12,.86v14.61A194.66,194.66,0,0,1,570.53,125Z"
			/>
			<path class="cls-3" d="M572.53,125.07c3.78.24,7.74.6,11.81,1.12V114.3c-1.15-.65-4.77-2.39-11.81-3.52Z" />
			<path class="cls-3" d="M532.19,126.86c1.94-.34,6.18-1,12.1-1.49V110.92a53.62,53.62,0,0,0-12.1,3.46Z" />
			<rect class="cls-3" x="139.33" y="148.76" width="8.67" height="49.25" />
			<rect class="cls-3" x="43.37" y="152.47" width="29.51" height="4.76" />
			<path class="cls-3" d="M43.37,198H54.83V169.22a1,1,0,0,1,1-1h17v-9H43.37Z" />
			<path class="cls-3" d="M218,198h12.69V127.4a1,1,0,0,1,1-1h14.49V113.62H218Z" />
			<path
			  class="cls-3"
			  d="M410.24,119.38c-9.35,0-10.87,1.87-11,2.13v41h6.63a1,1,0,0,1,1,1V198h10.52V158.69a1,1,0,0,1,1-1h2.92V121.57C421,121.22,419.37,119.38,410.24,119.38Z"
			/>
			<path class="cls-3" d="M593.57,112.22l-.06,9.1a51.85,51.85,0,0,1,16.46-4c1.24-1.69,2.47-3.34,3.69-5Z" />
			<rect class="cls-3" x="516" y="124.5" width="14.19" height="73.51" />
			<path
			  class="cls-3"
			  d="M337.25,153.27,330.57,149h-6.11l-6.68,4.27a1,1,0,0,1-.54.16h-2.07v9.07h24.69v-9.07h-2.07A1,1,0,0,1,337.25,153.27Z"
			/>
			<path class="cls-3" d="M183.79,113.4c.59-.3,1.16-.61,1.72-.92V100.81h-1.72Z" />
			<rect class="cls-3" x="150" y="100.81" width="1.71" height="97.2" />
			<path
			  class="cls-7"
			  d="M894.83,210.18c-2,2-3.71,3.7-7.28,3.7s-5.3-1.71-7.28-3.7-4.29-4.3-8.7-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.7-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.27,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.69,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.69,2.29-8.7,4.3-3.69,3.7-7.27,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.69,2.29-8.7,4.3-3.69,3.7-7.27,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.27,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.69,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.69,2.29-8.7,4.3-3.69,3.7-7.27,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.69,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.69,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.69,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.27,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.69,3.7-7.27,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.69,3.7-7.27,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.27,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.27,3.7S84,212.17,82,210.18s-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.27,3.7S52,212.17,50,210.18s-4.28-4.3-8.69-4.3a9.86,9.86,0,0,0-5,1.25V200H908.5v7.13a9.87,9.87,0,0,0-5-1.26C899.11,205.88,896.84,208.17,894.83,210.18Z"
			/>
			<path
			  class="cls-8"
			  d="M402.17,128.75a57.55,57.55,0,0,1,16.13.14l.17,0a1,1,0,0,0,1-.83,1,1,0,0,0-.81-1.16,58.86,58.86,0,0,0-16.77-.15,1,1,0,0,0-.84,1.14A1,1,0,0,0,402.17,128.75Z"
			/>
			<path
			  class="cls-8"
			  d="M418.36,133.58h.11a1,1,0,0,0,.1-2,100.47,100.47,0,0,0-16.64-.15,1,1,0,1,0,.18,2A99.53,99.53,0,0,1,418.36,133.58Z"
			/>
			<path
			  class="cls-8"
			  d="M134.62,123.51H117.77a1,1,0,0,0-1,1v8a1,1,0,0,0,1,1h16.85a1,1,0,0,0,1-1v-8A1,1,0,0,0,134.62,123.51Zm-1,8H118.77v-6h14.85Z"
			/>
			<path
			  class="cls-8"
			  d="M134.62,135H117.77a1,1,0,0,0-1,1v8a1,1,0,0,0,1,1h16.85a1,1,0,0,0,1-1v-8A1,1,0,0,0,134.62,135Zm-1,8H118.77v-6h14.85Z"
			/>
			<path
			  class="cls-8"
			  d="M134.62,146.43H117.77a1,1,0,0,0-1,1v8a1,1,0,0,0,1,1h16.85a1,1,0,0,0,1-1v-8A1,1,0,0,0,134.62,146.43Zm-1,8H118.77v-6h14.85Z"
			/>
			<path
			  class="cls-8"
			  d="M134.62,157.89H117.77a1,1,0,0,0-1,1v8a1,1,0,0,0,1,1h16.85a1,1,0,0,0,1-1v-8A1,1,0,0,0,134.62,157.89Zm-1,8H118.77v-6h14.85Z"
			/>
			<path
			  class="cls-8"
			  d="M134.62,169.35H117.77a1,1,0,0,0-1,1v8a1,1,0,0,0,1,1h16.85a1,1,0,0,0,1-1v-8A1,1,0,0,0,134.62,169.35Zm-1,8H118.77v-6h14.85Z"
			/>
			<path class="cls-8" d="M60.17,176.45a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,60.17,176.45Z" />
			<path class="cls-8" d="M64.47,176.45a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,64.47,176.45Z" />
			<path class="cls-8" d="M68.77,176.45a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,68.77,176.45Z" />
			<path class="cls-8" d="M73.07,176.45a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,73.07,176.45Z" />
			<path class="cls-8" d="M77.37,176.45a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,77.37,176.45Z" />
			<path class="cls-8" d="M60.17,182.39a1,1,0,0,0,1-1v-2.21a1,1,0,0,0-2,0v2.21A1,1,0,0,0,60.17,182.39Z" />
			<path class="cls-8" d="M64.47,182.39a1,1,0,0,0,1-1v-2.21a1,1,0,0,0-2,0v2.21A1,1,0,0,0,64.47,182.39Z" />
			<path class="cls-8" d="M68.77,182.39a1,1,0,0,0,1-1v-2.21a1,1,0,0,0-2,0v2.21A1,1,0,0,0,68.77,182.39Z" />
			<path class="cls-8" d="M73.07,182.39a1,1,0,0,0,1-1v-2.21a1,1,0,0,0-2,0v2.21A1,1,0,0,0,73.07,182.39Z" />
			<path class="cls-8" d="M77.37,182.39a1,1,0,0,0,1-1v-2.21a1,1,0,0,0-2,0v2.21A1,1,0,0,0,77.37,182.39Z" />
			<path class="cls-8" d="M60.17,188.32a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,60.17,188.32Z" />
			<path class="cls-8" d="M64.47,188.32a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,64.47,188.32Z" />
			<path class="cls-8" d="M68.77,188.32a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,68.77,188.32Z" />
			<path class="cls-8" d="M73.07,188.32a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,73.07,188.32Z" />
			<path class="cls-8" d="M77.37,188.32a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,77.37,188.32Z" />
			<path class="cls-8" d="M60.17,194.26a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,60.17,194.26Z" />
			<path class="cls-8" d="M64.47,194.26a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,64.47,194.26Z" />
			<path class="cls-8" d="M68.77,194.26a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,68.77,194.26Z" />
			<path class="cls-8" d="M73.07,194.26a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,73.07,194.26Z" />
			<path class="cls-8" d="M77.37,194.26a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,77.37,194.26Z" />
			<path class="cls-8" d="M236.48,131.16a1,1,0,0,0-1,1v4.73a1,1,0,0,0,2,0v-4.73A1,1,0,0,0,236.48,131.16Z" />
			<path class="cls-8" d="M240.78,131.16a1,1,0,0,0-1,1v4.73a1,1,0,0,0,2,0v-4.73A1,1,0,0,0,240.78,131.16Z" />
			<path class="cls-8" d="M245.08,131.16a1,1,0,0,0-1,1v4.73a1,1,0,0,0,2,0v-4.73A1,1,0,0,0,245.08,131.16Z" />
			<path class="cls-8" d="M249.37,131.16a1,1,0,0,0-1,1v4.73a1,1,0,0,0,2,0v-4.73A1,1,0,0,0,249.37,131.16Z" />
			<path class="cls-8" d="M253.67,131.16a1,1,0,0,0-1,1v4.73a1,1,0,0,0,2,0v-4.73A1,1,0,0,0,253.67,131.16Z" />
			<path class="cls-8" d="M236.48,140.9a1,1,0,0,0-1,1v4.73a1,1,0,0,0,2,0V141.9A1,1,0,0,0,236.48,140.9Z" />
			<path class="cls-8" d="M240.78,140.9a1,1,0,0,0-1,1v4.73a1,1,0,0,0,2,0V141.9A1,1,0,0,0,240.78,140.9Z" />
			<path class="cls-8" d="M245.08,140.9a1,1,0,0,0-1,1v4.73a1,1,0,0,0,2,0V141.9A1,1,0,0,0,245.08,140.9Z" />
			<path class="cls-8" d="M249.37,140.9a1,1,0,0,0-1,1v4.73a1,1,0,0,0,2,0V141.9A1,1,0,0,0,249.37,140.9Z" />
			<path class="cls-8" d="M253.67,140.9a1,1,0,0,0-1,1v4.73a1,1,0,0,0,2,0V141.9A1,1,0,0,0,253.67,140.9Z" />
			<path class="cls-8" d="M236.48,150.65a1,1,0,0,0-1,1v4.72a1,1,0,0,0,2,0v-4.72A1,1,0,0,0,236.48,150.65Z" />
			<path class="cls-8" d="M240.78,150.65a1,1,0,0,0-1,1v4.72a1,1,0,0,0,2,0v-4.72A1,1,0,0,0,240.78,150.65Z" />
			<path class="cls-8" d="M245.08,150.65a1,1,0,0,0-1,1v4.72a1,1,0,0,0,2,0v-4.72A1,1,0,0,0,245.08,150.65Z" />
			<path class="cls-8" d="M249.37,150.65a1,1,0,0,0-1,1v4.72a1,1,0,0,0,2,0v-4.72A1,1,0,0,0,249.37,150.65Z" />
			<path class="cls-8" d="M253.67,150.65a1,1,0,0,0-1,1v4.72a1,1,0,0,0,2,0v-4.72A1,1,0,0,0,253.67,150.65Z" />
			<path class="cls-8" d="M236.48,160.39a1,1,0,0,0-1,1v4.72a1,1,0,0,0,2,0v-4.72A1,1,0,0,0,236.48,160.39Z" />
			<path class="cls-8" d="M240.78,160.39a1,1,0,0,0-1,1v4.72a1,1,0,0,0,2,0v-4.72A1,1,0,0,0,240.78,160.39Z" />
			<path class="cls-8" d="M245.08,160.39a1,1,0,0,0-1,1v4.72a1,1,0,0,0,2,0v-4.72A1,1,0,0,0,245.08,160.39Z" />
			<path class="cls-8" d="M249.37,160.39a1,1,0,0,0-1,1v4.72a1,1,0,0,0,2,0v-4.72A1,1,0,0,0,249.37,160.39Z" />
			<path class="cls-8" d="M253.67,160.39a1,1,0,0,0-1,1v4.72a1,1,0,0,0,2,0v-4.72A1,1,0,0,0,253.67,160.39Z" />
			<path class="cls-8" d="M390.27,170.28a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,390.27,170.28Z" />
			<path class="cls-8" d="M394.23,170.28a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,394.23,170.28Z" />
			<path class="cls-8" d="M398.2,170.28a1,1,0,0,0,1-1v-2.2a1,1,0,1,0-2,0v2.2A1,1,0,0,0,398.2,170.28Z" />
			<path class="cls-8" d="M402.17,170.28a1,1,0,0,0,1-1v-2.2a1,1,0,1,0-2,0v2.2A1,1,0,0,0,402.17,170.28Z" />
			<path class="cls-8" d="M312.37,167.28a1,1,0,0,0-1,1v1.2a1,1,0,1,0,2,0v-1.2A1,1,0,0,0,312.37,167.28Z" />
			<path class="cls-8" d="M318.17,167.28a1,1,0,0,0-1,1v1.2a1,1,0,0,0,2,0v-1.2A1,1,0,0,0,318.17,167.28Z" />
			<path class="cls-8" d="M324,167.28a1,1,0,0,0-1,1v1.2a1,1,0,1,0,2,0v-1.2A1,1,0,0,0,324,167.28Z" />
			<path class="cls-8" d="M329.77,167.28a1,1,0,0,0-1,1v1.2a1,1,0,1,0,2,0v-1.2A1,1,0,0,0,329.77,167.28Z" />
			<path class="cls-8" d="M335.57,167.28a1,1,0,0,0-1,1v1.2a1,1,0,0,0,2,0v-1.2A1,1,0,0,0,335.57,167.28Z" />
			<path class="cls-8" d="M341.37,167.28a1,1,0,0,0-1,1v1.2a1,1,0,1,0,2,0v-1.2A1,1,0,0,0,341.37,167.28Z" />
			<path class="cls-8" d="M347.17,167.28a1,1,0,0,0-1,1v1.2a1,1,0,0,0,2,0v-1.2A1,1,0,0,0,347.17,167.28Z" />
			<path class="cls-8" d="M353,167.28a1,1,0,0,0-1,1v1.2a1,1,0,1,0,2,0v-1.2A1,1,0,0,0,353,167.28Z" />
			<path class="cls-8" d="M312.37,173.66a1,1,0,0,0-1,1v1.2a1,1,0,0,0,2,0v-1.2A1,1,0,0,0,312.37,173.66Z" />
			<path class="cls-8" d="M318.17,173.66a1,1,0,0,0-1,1v1.2a1,1,0,0,0,2,0v-1.2A1,1,0,0,0,318.17,173.66Z" />
			<path class="cls-8" d="M324,173.66a1,1,0,0,0-1,1v1.2a1,1,0,0,0,2,0v-1.2A1,1,0,0,0,324,173.66Z" />
			<path class="cls-8" d="M329.77,173.66a1,1,0,0,0-1,1v1.2a1,1,0,0,0,2,0v-1.2A1,1,0,0,0,329.77,173.66Z" />
			<path class="cls-8" d="M335.57,173.66a1,1,0,0,0-1,1v1.2a1,1,0,0,0,2,0v-1.2A1,1,0,0,0,335.57,173.66Z" />
			<path class="cls-8" d="M341.37,173.66a1,1,0,0,0-1,1v1.2a1,1,0,0,0,2,0v-1.2A1,1,0,0,0,341.37,173.66Z" />
			<path class="cls-8" d="M347.17,173.66a1,1,0,0,0-1,1v1.2a1,1,0,0,0,2,0v-1.2A1,1,0,0,0,347.17,173.66Z" />
			<path class="cls-8" d="M353,173.66a1,1,0,0,0-1,1v1.2a1,1,0,0,0,2,0v-1.2A1,1,0,0,0,353,173.66Z" />
			<path class="cls-8" d="M312.37,180a1,1,0,0,0-1,1v1.2a1,1,0,0,0,2,0V181A1,1,0,0,0,312.37,180Z" />
			<path class="cls-8" d="M318.17,180a1,1,0,0,0-1,1v1.2a1,1,0,0,0,2,0V181A1,1,0,0,0,318.17,180Z" />
			<path class="cls-8" d="M324,180a1,1,0,0,0-1,1v1.2a1,1,0,0,0,2,0V181A1,1,0,0,0,324,180Z" />
			<path class="cls-8" d="M329.77,180a1,1,0,0,0-1,1v1.2a1,1,0,0,0,2,0V181A1,1,0,0,0,329.77,180Z" />
			<path class="cls-8" d="M335.57,180a1,1,0,0,0-1,1v1.2a1,1,0,0,0,2,0V181A1,1,0,0,0,335.57,180Z" />
			<path class="cls-8" d="M341.37,180a1,1,0,0,0-1,1v1.2a1,1,0,0,0,2,0V181A1,1,0,0,0,341.37,180Z" />
			<path class="cls-8" d="M347.17,180a1,1,0,0,0-1,1v1.2a1,1,0,1,0,2,0V181A1,1,0,0,0,347.17,180Z" />
			<path class="cls-8" d="M353,180a1,1,0,0,0-1,1v1.2a1,1,0,0,0,2,0V181A1,1,0,0,0,353,180Z" />
			<path class="cls-8" d="M390.27,176.06a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,390.27,176.06Z" />
			<path class="cls-8" d="M394.23,176.06a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,394.23,176.06Z" />
			<path class="cls-8" d="M398.2,176.06a1,1,0,0,0,1-1v-2.2a1,1,0,1,0-2,0v2.2A1,1,0,0,0,398.2,176.06Z" />
			<path class="cls-8" d="M402.17,176.06a1,1,0,0,0,1-1v-2.2a1,1,0,1,0-2,0v2.2A1,1,0,0,0,402.17,176.06Z" />
			<path class="cls-8" d="M390.27,181.84a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,390.27,181.84Z" />
			<path class="cls-8" d="M394.23,181.84a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,394.23,181.84Z" />
			<path class="cls-8" d="M398.2,181.84a1,1,0,0,0,1-1v-2.2a1,1,0,1,0-2,0v2.2A1,1,0,0,0,398.2,181.84Z" />
			<path class="cls-8" d="M402.17,181.84a1,1,0,0,0,1-1v-2.2a1,1,0,1,0-2,0v2.2A1,1,0,0,0,402.17,181.84Z" />
			<path class="cls-8" d="M390.27,187.63a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,390.27,187.63Z" />
			<path class="cls-8" d="M394.23,187.63a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,394.23,187.63Z" />
			<path class="cls-8" d="M398.2,187.63a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,398.2,187.63Z" />
			<path class="cls-8" d="M402.17,187.63a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,402.17,187.63Z" />
			<path class="cls-8" d="M390.27,193.41a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,390.27,193.41Z" />
			<path class="cls-8" d="M394.23,193.41a1,1,0,0,0,1-1v-2.2a1,1,0,0,0-2,0v2.2A1,1,0,0,0,394.23,193.41Z" />
			<path class="cls-8" d="M398.2,193.41a1,1,0,0,0,1-1v-2.2a1,1,0,1,0-2,0v2.2A1,1,0,0,0,398.2,193.41Z" />
			<path class="cls-8" d="M402.17,193.41a1,1,0,0,0,1-1v-2.2a1,1,0,1,0-2,0v2.2A1,1,0,0,0,402.17,193.41Z" />
			<path
			  class="cls-8"
			  d="M919.5,213.88c-3.58,0-5.3-1.71-7.28-3.7-.54-.54-1.11-1.1-1.72-1.63V199a1,1,0,0,0-1-1h-5.23a.86.86,0,0,0-.07-.1l-5.22-5.2h0l-5.23-5.2a1.05,1.05,0,0,0-.71-.29H876.23c13.39-48.25,22.21-77.2,22.3-77.49a1,1,0,0,0-.13-.85,1,1,0,0,0-.73-.44c-23.55-2.34-51,3.16-66.22,9.45V88.44a1,1,0,0,0-.32-.74,1,1,0,0,0-.76-.26,140.14,140.14,0,0,0-49,12.69c.63-26.5,1-48.17,1-48.39a1,1,0,0,0-1.34-1c-25.82,9.15-49.87,26.44-71.62,51.47C702.69,92,695.81,87.76,695.51,87.58a1,1,0,0,0-.68-.13q-3.76.63-7.33,1.43c-7.31-21.6-14-37.39-14.09-37.55a1,1,0,0,0-.62-.56,1,1,0,0,0-.82.1c-14.69,8.92-33.79,28.92-56.82,59.49l-22.57-.14a1,1,0,0,0-1,1l-.07,11a37.77,37.77,0,0,0-5.17,3.07V113.74a1,1,0,0,0-.42-.81c-.29-.21-7.45-5.19-27.66-5.19-17.43,0-27.13,4.9-27.53,5.11a1,1,0,0,0-.54.89v8.76H515a1,1,0,0,0-1,1V198h-7.28V101.54a1,1,0,0,0-1.33-1l-8.34,3H486.72V96.92a1,1,0,0,0-1-1H459.44a1,1,0,0,0-1,1v6.62h-3.07a1,1,0,0,0-1,1v53.15H423.28V121.38s0-.07,0-.11c-.11-.92-1.37-3.89-13-3.89-12.09,0-13,3.07-13,4V162.5h-4.62V135.42a1,1,0,0,0-1-1h-9.95v-6.36a1,1,0,0,0-1-1h-5.28v-3.65a7.5,7.5,0,0,1,.83.16,8.89,8.89,0,0,0,1.83.25,4.78,4.78,0,0,0,1.52-.25,1,1,0,0,0,.68-.95v-3.88a1,1,0,0,0-1.47-.89,4.16,4.16,0,0,1-2.77-.11,5.73,5.73,0,0,0-1.59-.25h-.13a.45.45,0,0,0-.11,0l-.23.07-.13.09a.75.75,0,0,0-.17.14.67.67,0,0,0-.09.14.64.64,0,0,0-.1.18.65.65,0,0,0,0,.19.8.8,0,0,0,0,.15v8.58h-8.67v-3.65a7.09,7.09,0,0,1,.83.16,8.89,8.89,0,0,0,1.83.25,4.73,4.73,0,0,0,1.52-.25,1,1,0,0,0,.68-.95v-3.88a1,1,0,0,0-1.47-.89,4.16,4.16,0,0,1-2.77-.11,5.73,5.73,0,0,0-1.59-.25h-.13a.45.45,0,0,0-.11,0l-.23.07-.13.09a.75.75,0,0,0-.17.14.67.67,0,0,0-.09.14.64.64,0,0,0-.1.18,1.29,1.29,0,0,0,0,.19.8.8,0,0,0,0,.15v8.58h-4.08a1,1,0,0,0-1,1v6.36h-9.95a1,1,0,0,0-1,1V162.5h-4.8V152.43a1,1,0,0,0-1-1h-2.77L336,150.11V67a1,1,0,0,0-1-1H291.71a1,1,0,0,0-1,1v5.22h-1.52V54.76a1,1,0,0,0-1-1H275v-5a1,1,0,1,0-2,0v5h-2.58V37.48a1,1,0,0,0-2,0V53.76h-2.58v-5a1,1,0,1,0-2,0v5H250.55a1,1,0,0,0-1,1V72.25h-2.37a1,1,0,0,0-1,1v38.37H232.69v-5.69q.59.06,1.26.18a14,14,0,0,0,2.27.25,5.12,5.12,0,0,0,1.65-.25,1,1,0,0,0,.68-1v-3.88a1,1,0,0,0-1.46-.89c-.81.42-2.5.09-3.62-.13a8.08,8.08,0,0,0-1.74-.23h-.13a.35.35,0,0,0-.11,0l-.23.07a.6.6,0,0,0-.13.09.75.75,0,0,0-.17.14.71.71,0,0,0-.1.14c0,.06-.07.11-.1.18a1.08,1.08,0,0,0,0,.19.83.83,0,0,0,0,.15v10.6H217a1,1,0,0,0-1,1V198h-5.84V87.94a1,1,0,0,0-.85-1,1,1,0,0,0-1.1.68,40.49,40.49,0,0,1-6.15,10.74,52.27,52.27,0,0,1-7.44,7.78v-29a1,1,0,0,0-2,0v30.58a59.18,59.18,0,0,1-5.11,3.54V99.81a1,1,0,0,0-1-1h-2.72v-2.3a1,1,0,0,0-1-1h-3.21V89.78a1,1,0,0,0-1-1h-1.94V86.27a1,1,0,0,0-1-1h-.71v-7.6a1,1,0,0,0-1-1h-2.39V74.55a1,1,0,0,0-1-1h-1.79v-12a1,1,0,0,0-2,0v12H165a1,1,0,0,0-1,1v2.12h-2.4a1,1,0,0,0-1,1v7.6h-.71a1,1,0,0,0-1,1v2.51h-1.94a1,1,0,0,0-1,1v5.73h-3.21a1,1,0,0,0-1,1v2.3H149a1,1,0,0,0-1,1v46.95h-8.67V120.61a1,1,0,0,0-1-1h-.59V60.89a1,1,0,1,0-2,0v11H116.65v-11a1,1,0,0,0-2,0v58.72h-.59a1,1,0,0,0-1,1V198h-6.89V139.44a1,1,0,0,0-1-1H73.88a1,1,0,0,0-1,1v11H42.37a1,1,0,0,0-1,1V198h-6a1,1,0,0,0-1,1v9.53c-.61.54-1.18,1.09-1.72,1.64-2,2-3.7,3.7-7.28,3.7a1,1,0,0,0,0,2c4.41,0,6.69-2.28,8.7-4.29s3.69-3.71,7.27-3.71,5.29,1.72,7.27,3.71,4.29,4.29,8.7,4.29S64,213.6,66,211.59s3.69-3.71,7.27-3.71,5.29,1.72,7.27,3.71,4.29,4.29,8.7,4.29,6.68-2.28,8.69-4.29,3.69-3.71,7.27-3.71,5.29,1.72,7.28,3.71,4.28,4.29,8.69,4.29,6.68-2.28,8.69-4.29,3.7-3.71,7.27-3.71,5.29,1.72,7.27,3.71,4.29,4.29,8.69,4.29,6.69-2.28,8.69-4.29,3.7-3.71,7.28-3.71,5.29,1.72,7.27,3.71,4.29,4.29,8.69,4.29,6.68-2.28,8.69-4.29,3.7-3.71,7.27-3.71,5.29,1.72,7.27,3.71,4.29,4.29,8.69,4.29,6.68-2.28,8.69-4.29,3.7-3.71,7.28-3.71,5.29,1.72,7.27,3.71,4.29,4.29,8.69,4.29,6.68-2.28,8.69-4.29,3.7-3.71,7.27-3.71,5.29,1.72,7.28,3.71,4.28,4.29,8.69,4.29,6.68-2.28,8.69-4.29,3.69-3.71,7.27-3.71,5.29,1.72,7.27,3.71,4.29,4.29,8.69,4.29,6.68-2.28,8.69-4.29,3.7-3.71,7.28-3.71,5.29,1.72,7.27,3.71,4.29,4.29,8.69,4.29,6.68-2.28,8.69-4.29,3.7-3.71,7.28-3.71,5.29,1.72,7.27,3.71,4.28,4.29,8.69,4.29,6.68-2.28,8.69-4.29,3.7-3.71,7.28-3.71,5.28,1.72,7.27,3.71,4.28,4.29,8.69,4.29,6.68-2.28,8.69-4.29,3.7-3.71,7.28-3.71,5.29,1.72,7.27,3.71,4.28,4.29,8.69,4.29,6.68-2.28,8.69-4.29,3.7-3.71,7.28-3.71,5.29,1.72,7.27,3.71,4.29,4.29,8.69,4.29,6.69-2.28,8.69-4.29,3.7-3.71,7.28-3.71,5.29,1.72,7.28,3.71,4.28,4.29,8.69,4.29,6.68-2.28,8.69-4.29,3.7-3.71,7.27-3.71,5.29,1.72,7.28,3.71,4.28,4.29,8.69,4.29,6.68-2.28,8.69-4.29,3.7-3.71,7.28-3.71,5.29,1.72,7.27,3.71,4.29,4.29,8.69,4.29,6.69-2.28,8.69-4.29,3.7-3.71,7.28-3.71,5.29,1.72,7.28,3.71,4.28,4.29,8.69,4.29,6.68-2.28,8.69-4.29,3.7-3.71,7.28-3.71,5.29,1.72,7.27,3.71,4.29,4.29,8.69,4.29,6.69-2.28,8.69-4.29,3.7-3.71,7.28-3.71,5.29,1.72,7.28,3.71,4.28,4.29,8.69,4.29,6.68-2.28,8.69-4.29,3.7-3.71,7.28-3.71,5.29,1.72,7.27,3.71,4.29,4.29,8.69,4.29,6.69-2.28,8.7-4.29,3.7-3.71,7.27-3.71,5.29,1.72,7.28,3.71,4.28,4.29,8.69,4.29,6.68-2.28,8.69-4.29,3.7-3.71,7.27-3.71,5.3,1.72,7.28,3.71,4.29,4.29,8.69,4.29,6.69-2.28,8.7-4.29,3.7-3.71,7.27-3.71,5.29,1.72,7.28,3.71,4.28,4.29,8.69,4.29,6.68-2.28,8.69-4.29,3.7-3.71,7.28-3.71,5.29,1.72,7.28,3.71,4.28,4.29,8.69,4.29,6.68-2.28,8.69-4.29,3.7-3.71,7.28-3.71,5.29,1.72,7.28,3.71,4.28,4.29,8.69,4.29,6.68-2.28,8.69-4.29,3.7-3.71,7.28-3.71,5.3,1.72,7.28,3.71,4.29,4.29,8.7,4.29,6.68-2.28,8.69-4.29,3.7-3.71,7.28-3.71,5.3,1.72,7.28,3.71,4.29,4.29,8.7,4.29a1,1,0,0,0,0-2ZM375.36,119.65l.13,0a8.32,8.32,0,0,0,2.73.36v1.77a5.28,5.28,0,0,1-1.6-.2,11.41,11.41,0,0,0-1.26-.21Zm-10.67,0,.13,0a8.32,8.32,0,0,0,2.73.36v1.77a5.32,5.32,0,0,1-1.6-.2,11,11,0,0,0-1.26-.21Zm-132-17.51.39.08a12.59,12.59,0,0,0,3.47.35v1.78a10.15,10.15,0,0,1-2.27-.21,15.44,15.44,0,0,0-1.59-.21ZM901.47,198H547l4.81-3.59H897.86Zm-8.84-8.79,3.21,3.2H554.46l4.28-3.2H892.63ZM631,112a1,1,0,0,0-1.31.16c-22.93,25.85-34.41,68.83-36,75.08h-7.12c.79-3.94,4-17.75,13.83-37.47,10.65-21.39,31.43-54.28,70-89.84,3.45,7.82,8.85,26.34,10.11,30.73a93.51,93.51,0,0,0-43.19,25.66Zm5,5.83c-11,13-19.28,31.12-24.33,44.16a265,265,0,0,0-8.35,25.25h-7.58c1.94-7.55,13.12-48.1,34.84-73.09Zm111.38-14.37a1,1,0,0,0-.56-.14c-1.11.07-2.73,1-6.33,7.06-.68-1-1.08-1.5-1.11-1.54a1,1,0,0,0-.75-.38c-1.78-.07-3.51-.1-5.2-.11,2.8-6.7,16.73-37,43.84-50.18l-1,44.52a142.51,142.51,0,0,0-14.94,9ZM759.61,113c-4.41,3.92-8.55,10.32-10.37,13.31a111.94,111.94,0,0,0-7.54-14c3.1-5.34,4.53-6.57,5-6.84Zm-21.55-2.53c2.53,3.45,21.81,31.26,22.1,76.8h-4c-.25-4.68-1.73-25.68-8.64-46.64a.22.22,0,0,1,0-.08c0-.06-.06-.11-.08-.17a103.69,103.69,0,0,0-11.1-23.82,1,1,0,0,0-1.08-.42c-52.57,13.23-79.13,65.06-82.08,71.13H633.14c3.86-8.2,24.4-49.94,49.09-65C697.91,114.47,713.44,109.52,738.06,110.42Zm-82.61,76.8c4.43-8.8,30.33-56.22,79.65-69A101.39,101.39,0,0,1,745.26,140c-44.79,10.61-67.2,42.87-70,47.2Zm90.45-45.3a192.93,192.93,0,0,1,8.3,45.3H677.64C682.41,180.32,704.39,151.69,745.9,141.92Zm-115,45.3h-5c2.89-8.13,19.42-52,46.54-71.91l7.59,5.91c-14.09,9.13-26.59,26.35-34.68,39.4A264.64,264.64,0,0,0,630.93,187.22Zm50.94-67.1-8.79-6.85a1,1,0,0,0-1.2,0c-28.79,20.65-45.92,67.73-48.1,74h-2.09c8.47-30.22,25.29-72,72.47-92.68,4.19,4,9.25,10.94,10.73,13-1.35,1.64-2.69,3.32-4,5A122.44,122.44,0,0,0,681.87,120.12Zm120.86-1.89a1,1,0,0,0-.94.29c-7.57,8-15.73,22.17-24.23,42-5.48,12.77-9.34,23.9-10.29,26.71h-5.11c0-2.95-.12-5.83-.28-8.63,10.13-29.11,25.41-72.8,62-80.87v24c-.39.27-.77.53-1.11.78Zm17.83,5.85c-26.46,18.8-47.26,58-49.9,63.14h-1.28c2.73-7.93,17.54-49.59,33.49-66.92Zm67.11-4.16c-32.29,18-55.09,61.84-57.84,67.3H819.24c31.18-59.22,62.42-69.52,69.56-71.14Zm-.81,2.78-19,64.52H832.08C836.17,179.26,857.35,140.24,886.86,122.7Zm-56-2.39h0a119,119,0,0,1,22.5-6.84c14.95-3.16,30.13-4.27,42.92-3.15-1.78,5.9-10.13,33.75-22.1,76.9H870l19.53-66.35h0l1.7-5.78a1,1,0,0,0-.19-.92,1,1,0,0,0-.86-.36c-1.5.14-37.09,4.33-73.15,73.41H772.92c4.13-7.85,24.92-45.67,50.66-62.8l0,0,0,0A32.88,32.88,0,0,1,830.84,120.31Zm-1.39-30.79v29.23c-1.31.59-2.49,1.19-3.56,1.78v-24a1,1,0,0,0-.36-.78,1,1,0,0,0-.83-.2c-36.67,7.22-53.26,49.8-63.2,78.1a147.31,147.31,0,0,0-11.27-45.1c1.11-1.9,6.46-10.8,11.7-14.9a140.67,140.67,0,0,1,15.65-9.4h0A137.71,137.71,0,0,1,829.45,89.52ZM780.3,53.15c-.08,5.36-.41,24.93-1,48l-1.08.53,1-45.07a1,1,0,0,0-.45-.86,1,1,0,0,0-1-.07c-29.91,13.57-44.6,47.67-46.63,52.68a110.62,110.62,0,0,0-27.11,3.34C727,83,752.63,63.29,780.3,53.15ZM694.78,89.49c1.39.93,7.48,5.32,13.23,14.34-.6.7-1.2,1.41-1.79,2.13-1.9-2.64-7-9.5-11.19-13.34a1,1,0,0,0-1.07-.18,118.18,118.18,0,0,0-53,44.77c-11.6,17.84-17.75,36.94-21.38,50H605.38c2-7.47,13.43-46.61,32.82-68.94,12-12.71,26.17-21.14,43.79-26h0s0,0,0,0A124.17,124.17,0,0,1,694.78,89.49ZM672,53.19c1.62,3.88,7.34,17.92,13.51,36.13q-1.56.38-3.09.78c-1.35-4.71-7.3-25.15-10.81-32.29a1,1,0,0,0-1.57-.3c-70.82,64.8-84.07,122.12-85.55,129.71h-4.88c1.5-6.28,9.55-36.61,31.68-68.34C636.11,85.09,656.52,63,672,53.19Zm-78.45,59,20.09.13c-1.22,1.63-2.45,3.28-3.69,5a51.85,51.85,0,0,0-16.46,4Zm14.91,7.25c-22,32.19-29.69,62.55-30.9,67.75h-18c.89-6.47,6.28-39.71,24.15-57.12C591.9,122.11,603.77,120,608.48,119.47Zm-36-8.69c7,1.13,10.66,2.87,11.81,3.52v11.89c-4.07-.52-8-.88-11.81-1.12Zm-26.24-.18a81.57,81.57,0,0,1,12-.86,93,93,0,0,1,12.27.75V125a194.66,194.66,0,0,0-24.24.25Zm-14.1,3.78a53.62,53.62,0,0,1,12.1-3.46v14.45c-5.92.5-10.16,1.15-12.1,1.49Zm0,14.51A182.47,182.47,0,0,1,583,128c-.23.2-.45.41-.67.62-10.62,10.35-16.82,26-20.16,37.26a147.75,147.75,0,0,0-4.68,21.76l-13.58,10.13a1.06,1.06,0,0,0-.17.2H532.19ZM516,124.5h14.19V198H516Zm-17.78-19.26,6.5-2.29V198h-6.5Zm-37.78-7.32h24.28v5.62H460.44Zm-4.07,7.62h39.85V198H456.37Zm-11.95,54.15h9.95V198h-9.95Zm-5.72,0h3.72V198H438.7Zm-5.73,0h3.73V198H433Zm-5.72,0H431V198h-3.72Zm-2,0V198h-5.89V159.69Zm-26-38.18c.16-.26,1.68-2.13,11-2.13,9.13,0,10.8,1.84,11,2.19v36.12h-2.92a1,1,0,0,0-1,1V198H406.84V163.5a1,1,0,0,0-1-1h-6.63Zm-7.62,43a.5.5,0,0,0,.13,0h13.12V198H387.59V164.5h3.87A.5.5,0,0,0,391.59,164.53Zm-32-35.47h3.95l.13,0,.13,0h10.41l.13,0,.13,0h5.15v5.36h-20Zm-10.95,7.36h41.93v4.52H348.66Zm0,6.52h41.93V147H348.66Zm0,6.06h41.93v4H348.66Zm0,6h41.93v7.46h-4a1,1,0,0,0-1,1V198h-26V163.5a1,1,0,0,0-1-1h-9.95Zm-1.13,9.46.13,0a.5.5,0,0,0,.13,0h9.82V198H308.16V164.5ZM329.49,68H334v9.14h-4.54Zm0,11.14H334v69.66l-2.62-1.68a1,1,0,0,0-.54-.15h-1.38ZM323.38,68h4.11v9.14h-4.11Zm0,11.14h4.11V147h-3.33a1,1,0,0,0-.54.15l-.24.16ZM317.27,68h4.11v9.14h-4.11Zm0,11.14h4.11v69.42l-4.11,2.63Zm.51,74.1,6.68-4.27h6.11l6.68,4.27a1,1,0,0,0,.54.16h2.07v9.07H315.17v-9.07h2.07A1,1,0,0,0,317.78,153.27ZM311.16,68h4.11v9.14h-4.11Zm0,11.14h4.11v72.26h-1.1a1,1,0,0,0-1,1V162.5h-2ZM305,68h4.12v9.14H305Zm0,11.14h4.12V162.5h-2a1,1,0,0,0-1,1V198H305ZM298.93,68H303v9.14h-4.11Zm0,11.14H303V198h-4.11ZM292.71,68h4.22v9.14h-4.22Zm0,11.14h4.22V198h-4.22ZM275,55.76h12.17V72.25h-2.37V59.5a1,1,0,0,0-1-1H275Zm0,4.74h7.8V72.25H275Zm0,13.75h15.69V198H275Zm-4.58-18.49H273V58.5h-2.58Zm-4.58,0h2.58V58.5h-2.58Zm0,4.74H273V72.25h-7.16Zm0,13.75H273V198h-3.43V127.4a1,1,0,0,0-1-1h-2.73ZM251.55,55.76h12.31V58.5h-8.79a1,1,0,0,0-1,1V72.25h-2.52Zm12.31,4.74V72.25h-7.79V60.5ZM248.18,74.25h15.68V126.4H248.18Zm19.41,54.15V198h-34.9V128.4ZM218,113.62h28.18V126.4H231.69a1,1,0,0,0-1,1V198H218Zm-34.21-12.81h1.72v11.67c-.56.31-1.13.62-1.72.92ZM183.13,116a.84.84,0,0,0,.14-.06c1.26-.61,2.46-1.24,3.61-1.9,0,0,0,0,.06,0a53.64,53.64,0,0,0,21.22-21.31v22h-9.35a1,1,0,0,0,0,2h9.35v4.07h-9.35a1,1,0,0,0,0,2h9.35v4.08h-9.35a1,1,0,0,0,0,2h9.35v4.07h-9.35a1,1,0,0,0,0,2h9.35v4.08h-9.35a1,1,0,0,0,0,2h9.35V145h-9.35a1,1,0,0,0,0,2h9.35v51H179.09V117.76C180.5,117.19,181.84,116.58,183.13,116Zm-1.34-18.45v16.88c-1.29.61-2.64,1.2-4.06,1.75a1,1,0,0,0-.64.94V198h-2.66V97.51h7.36ZM171.54,78.67h1.39v6.6h-1.39ZM166,75.55h3.57v9.72H166Zm-3.4,3.12H164v6.6h-1.4Zm-1.71,8.6h13.78v1.51H160.86Zm-2.94,3.51h19.66v4.73H157.92Zm14.51,6.73V198h-9.36V97.51Zm-18.72,0h7.36V198h-7.36Zm-3.71,3.3h1.71V198H150Zm-10.67,47.95H148V198h-8.67Zm-3.59-42.56v7h-7.1Zm-9.55,6.57-8-7.9h16.08Zm2.45-9.9,7.1-7v7Zm-2.45-.4-8-7.9h16.08Zm2.45-9.9,7.1-7v7Zm-2.45-.4-8-7.9h16.08Zm2.45-9.9,7.1-7v7Zm5.65-8.36-8.1,8-8.09-8Zm-17.64,1.38,7.1,7h-7.1Zm0,10.3,7.1,7h-7.1Zm0,10.31,7.1,7h-7.1Zm0,10.3,7.1,7h-7.1Zm0,9h19.09v4.44H116.65Zm-1.59,6.44h22.27V198H115.06ZM99.54,140.44h4.63V198H99.54Zm-4.57,0h2.57V198H95Zm-20.09,0H93V198H82.71V169.22a1,1,0,0,0-1-1H74.88Zm5.83,29.78V198H56.83V170.22ZM43.37,152.47H72.88v4.76H43.37Zm0,6.76H72.88v9h-17a1,1,0,0,0-1,1V198H43.37ZM908.5,207.14a9.87,9.87,0,0,0-5-1.26c-4.41,0-6.68,2.29-8.69,4.3s-3.71,3.7-7.28,3.7-5.3-1.71-7.28-3.7-4.29-4.3-8.7-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.7-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.27,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.69,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.69,2.29-8.7,4.3-3.69,3.7-7.27,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.69,2.29-8.7,4.3-3.69,3.7-7.27,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.27,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.69,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.69,2.29-8.7,4.3-3.69,3.7-7.27,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.69,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.69,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.69,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.27,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.69,3.7-7.27,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.69,3.7-7.27,3.7-5.29-1.71-7.27-3.7-4.29-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.28,3.7-5.29-1.71-7.27-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.27,3.7-5.29-1.71-7.28-3.7-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.27,3.7S84,212.17,82,210.18s-4.28-4.3-8.69-4.3-6.68,2.29-8.69,4.3-3.7,3.7-7.27,3.7S52,212.17,50,210.18s-4.28-4.3-8.69-4.3a9.86,9.86,0,0,0-5,1.25V200H908.5Z"
			/>
			<rect class="cls-7" x="118.77" y="125.51" width="14.85" height="6.02" />
			<rect class="cls-7" x="118.77" y="136.97" width="14.85" height="6.02" />
			<rect class="cls-7" x="118.77" y="148.43" width="14.85" height="6.02" />
			<rect class="cls-7" x="118.77" y="159.89" width="14.85" height="6.02" />
			<rect class="cls-7" x="118.77" y="171.35" width="14.85" height="6.02" />
			<path
			  class="cls-4"
			  d="M677.64,187.22H754.2a192.93,192.93,0,0,0-8.3-45.3C704.39,151.69,682.41,180.32,677.64,187.22Z"
			/>
			<path class="cls-4" d="M832.08,187.22h35.78l19-64.52C857.35,140.24,836.17,179.26,832.08,187.22Z" />
			<polygon class="cls-5" points="498.22 105.25 498.22 198.01 504.72 198.01 504.72 102.95 498.22 105.25" />
			<path
			  class="cls-6"
			  d="M186.94,114s0,0-.06,0c-1.15.66-2.35,1.29-3.61,1.9a.84.84,0,0,1-.14.06c-1.29.62-2.63,1.23-4,1.8V198h29.07V147h-9.35a1,1,0,0,1,0-2h9.35v-4.07h-9.35a1,1,0,0,1,0-2h9.35v-4.08h-9.35a1,1,0,0,1,0-2h9.35v-4.07h-9.35a1,1,0,0,1,0-2h9.35v-4.08h-9.35a1,1,0,0,1,0-2h9.35v-4.07h-9.35a1,1,0,0,1,0-2h9.35v-22A53.64,53.64,0,0,1,186.94,114Z"
			/>
		  </svg>
		</div>
	
		<script>
		  //IFEE to format redirect input for form submission
		  if (window.location.hash) {
			(function() {
			  var input = document.getElementById("rd");
			  inputs.value += window.location.hash;
			})();
		  }
	
		  //IFEE to hide error notification if they start typing again
		  (function() {
			function hideErrorNotification() {
			  var errorNotification = document.getElementById("errorNotification");
			  if (errorNotification) {
				errorNotification.style.display = "none";
			  }
			}
	
			var usernameInput = document.getElementById("username");
			var passwordInput = document.getElementById("password");
			usernameInput.addEventListener("keydown", hideErrorNotification);
			passwordInput.addEventListener("keydown", hideErrorNotification);
	
			var closeNotifcationButton = document.getElementsByClassName("bx--inline-notification__close-button")[0];
			closeNotifcationButton.addEventListener("click", hideErrorNotification);
		  })();
		</script>
	  </body>
	</html>
	{{end}}`)
	if err != nil {
		logger.Fatalf("failed parsing template %s", err)
	}

	t, err = t.Parse(`{{define "error.html"}}
	<!DOCTYPE html>
    <html lang="en">
      <head>
        <meta charset="utf-8" />
        <meta name="viewport" content="width=device-width, initial-scale=1, shrink-to-fit=no" />
        <meta name="theme-color" content="#1b2834" />
        <title>Boomerang - 403</title>
        <link href="https://fonts.googleapis.com/css?family=IBM+Plex+Sans:300,400,600&display=swap" rel="stylesheet" />
        <link
          href="data:image/x-icon;base64,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"
          rel="icon"
          type="image/x-icon"
        />
        <style type="text/css">
          :root {
            --red: #d91e28;
            --blue: #0072c3;
            --gray: #50565b;
          }
          h1,
          p {
            margin: 0;
          }
          body {
            background: #f2f4f8;
            font-family: "IBM Plex Sans", sans-serif;
            font-weight: 600;
            box-sizing: border-box;
          }
          .container {
            align-items: center;
            display: flex;
            flex-direction: column;
            margin: 180px auto;
            padding: 0 2rem;
          }
          .title {
            color: var(--red);
            font-size: 32px;
            font-weight: 600;
            line-height: 42px;
            margin: 0;
            text-align: center;
          }
          .header {
            color: var(--gray);
            font-size: 24px;
            font-weight: 600;
            line-height: 32px;
            margin-top: 30px;
            text-align: center;
          }
          .message {
            color: var(--gray);
            font-size: 16px;
            font-weight: 300;
            height: 20px;
            margin-top: 9px;
            text-align: center;
            text-decoration: none;
          }
          .graphic {
            margin-top: 50px;
            max-width: 520px;
            width: 100%;
          }
        </style>
      </head>
      <body>
        <main class="container">
          <h1 class="title">403 - Access Forbidden</h1>
          <p class="header">You’ve found yourself in deep water.</p>
          <p class="message">You shouldn’t be here - contact the local authorities if you disagree.</p>
          <svg
            id="Layer_1"
            class="graphic"
            data-name="Layer 1"
            xmlns="http://www.w3.org/2000/svg"
            viewBox="0 0 570.04 265.53"
          >
            <defs>
              <style>
                .cls-1 {
                  fill: #f2f4f8;
                }
                .cls-2 {
                  fill: #3ddbd9;
                }
                .cls-3 {
                  fill: #373d42;
                }
                .cls-4 {
                  fill: none;
                  stroke: #343a3f;
                  stroke-linecap: round;
                  stroke-linejoin: round;
                  stroke-width: 5px;
                }
              </style>
            </defs>
            <title>great-reef-serpent-error-403</title>
            <path
              class="cls-1"
              d="M487.83,131.73c-.8-5.82-.84-13.66,2.94-14.94,4.38-1.48,6.5,4.78,7.14,12.34s-1,13.9-4.89,14.16C489.63,143.53,488.62,137.55,487.83,131.73Z"
            />
            <path
              class="cls-2"
              d="M469,27.1c42.11-15.32,68.11,27.2,52.19,54.27a40.13,40.13,0,0,1-3,4.41c-14.9,19-43.06,18.74-53.73,4.2l-8,.6a4,4,0,0,0-3.69,4.29s.46,12,1.66,27.27c-12.4,66.49-37.93,101.28-37.93,101.28H366S393,191.87,416,132.56C435.33,82.78,433.65,40,469,27.1Zm40.36,37.23c2.46-.7,5.62-6.25.57-8.38-1.87-.78-3.28.83-3.84,2.79C505.16,62,506.84,65.06,509.37,64.33Zm-27.27,1c2.45-.7,5.61-6.25.57-8.38-1.87-.78-3.29.83-3.85,2.79C477.88,63,479.56,66.06,482.1,65.33Z"
            />
            <path
              class="cls-2"
              d="M163.42,223.42H125.84s-2.64-26.89-21.75-26.89c-21.36,0-23.54,26.89-23.54,26.89h-37S52.36,163,106.5,163C161.29,163,163.42,223.42,163.42,223.42Z"
            />
            <path
              class="cls-3"
              d="M482.67,57c5,2.13,1.88,7.68-.57,8.38s-4.22-2.32-3.28-5.59C479.38,57.78,480.8,56.17,482.67,57Z"
            />
            <path
              class="cls-1"
              d="M497.91,129.13c-.64-7.56-2.76-13.82-7.14-12.34-3.78,1.28-3.74,9.12-2.94,14.94s1.8,11.8,5.19,11.56C496.93,143,498.56,136.7,497.91,129.13Zm25.71-43.58a4,4,0,0,1,4.28,3.69s-3.68,45.8,3.22,69.46l0,.31-54,5.34s-1.31,5-2.27,8.78c-.68,2.72-2.23,6.19-5.53,6.43l-.2,0c-3.85.29-6-2.58-7.5-6.46-3.59-12.69-5.87-33.86-7.21-51-1.2-15.24-1.66-27.27-1.66-27.27a4,4,0,0,1,3.69-4.29l8-.6c10.67,14.54,38.83,14.8,53.73-4.2l.21.16Z"
            />
            <path
              class="cls-3"
              d="M509.94,56c5.05,2.13,1.89,7.68-.57,8.38s-4.21-2.32-3.27-5.59C506.66,56.78,508.07,55.17,509.94,56Z"
            />
            <path
              class="cls-1"
              d="M477.16,164.35l54-5.34,13-1.28a22,22,0,0,1-.27,9.78c-1.22,3.8-2.26,6-6.53,6.43l-68,5.62c3.3-.24,4.85-3.71,5.53-6.43C475.85,169.31,477.16,164.35,477.16,164.35Z"
            />
            <path
              class="cls-2"
              d="M237.43,223.42H196.9s0-88.21,69.2-88.21c65.38,0,69.41,88.21,69.41,88.21H294.6S295.14,175,267.12,175C236.75,175,237.43,223.42,237.43,223.42Z"
            />
            <path
              class="cls-4"
              d="M80.55,223.42s2.18-26.89,23.54-26.89c19.11,0,21.75,26.89,21.75,26.89h37.58S161.29,163,106.5,163c-54.14,0-62.92,60.47-62.92,60.47Z"
            />
            <path
              class="cls-4"
              d="M196.9,223.42h40.53S236.75,175,267.12,175c28,0,27.48,48.44,27.48,48.44h40.91s-4-88.21-69.41-88.21C196.9,135.21,196.9,223.42,196.9,223.42Z"
            />
            <path
              class="cls-4"
              d="M454.45,122.12v0c-12.4,66.49-37.93,101.28-37.93,101.28H366S393,191.87,416,132.56C435.33,82.78,433.65,40,469,27.1c42.11-15.32,68.11,27.2,52.19,54.27a40.13,40.13,0,0,1-3,4.41c-14.9,19-43.06,18.74-53.73,4.2a21.41,21.41,0,0,1-3.78-8.61"
            />
            <path class="cls-4" d="M518.41,85.94l5.21-.39a4,4,0,0,1,4.28,3.69s-3.68,45.8,3.22,69.46" />
            <path
              class="cls-4"
              d="M469.36,179.56c3.3-.24,4.85-3.71,5.53-6.43,1-3.82,2.27-8.78,2.27-8.78l54-5.34,13-1.28a22,22,0,0,1-.27,9.78c-1.22,3.8-2.26,6-6.53,6.43l-68,5.62-.2,0c-3.85.29-6-2.58-7.5-6.46-3.59-12.69-5.87-33.86-7.21-51-1.2-15.24-1.66-27.27-1.66-27.27a4,4,0,0,1,3.69-4.29l8-.6"
            />
            <path class="cls-4" d="M394.85,233.66c50,1,61.72-10.24,39.33-14.8" />
            <path class="cls-4" d="M380.84,233.9s-40.53-1.22-28.54-13.21" />
            <path class="cls-4" d="M279.8,220.69s-13.6,13.36,22.83,13.21" />
            <path class="cls-4" d="M248.4,220.69s20,10.43-12,13.21" />
            <path class="cls-4" d="M186.24,218.5c-16,4.92-3.47,15.4,35.33,15.4" />
            <path class="cls-4" d="M35.33,216.76c-15,1.74-13.28,16,8.25,17.14" />
            <path class="cls-4" d="M71.36,233.9s37.77,1,25.17-7.7" />
            <path class="cls-4" d="M116.11,224.59c-1.75,12.93,50.5,11,54,5.48" />
            <path class="cls-4" d="M458,223.42c8.75,1.42,17.3,13.29-15.53,18.25" />
            <polyline class="cls-4" points="468.73 118.66 467.3 130.26 476.55 129.48 481.63 129.05" />
            <polyline class="cls-4" points="475.84 121.06 476.55 129.48 477.78 143.95" />
            <path
              class="cls-4"
              d="M490.77,116.79c-3.78,1.28-3.74,9.12-2.94,14.94s1.8,11.8,5.19,11.56c3.91-.26,5.54-6.59,4.89-14.16S495.15,115.31,490.77,116.79Z"
            />
            <path
              class="cls-4"
              d="M503.72,118.7c-.2-2.37,7.59-5.32,10.6-.13,3.76,6.5-4.61,9.9-4.61,9.9s10.58.67,5.95,9.49c-2.66,5.07-8.16,3.71-10.33-.26"
            />
          </svg>
        </main>
      </body>
    </html>
	{{end}}`)
	if err != nil {
		logger.Fatalf("failed parsing template %s", err)
	}
	return t
}
