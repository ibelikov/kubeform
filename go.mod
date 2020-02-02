module kubeform.dev/kubeform

go 1.12

require (
	github.com/appscode/go v0.0.0-20191119085241-0887d8ec2ecc
	github.com/dave/jennifer v1.4.0
	github.com/go-openapi/spec v0.19.6
	github.com/gobuffalo/flect v0.2.0
	github.com/golang/glog v0.0.0-20160126235308-23def4e6c14b
	github.com/hashicorp/terraform v0.12.20
	github.com/hashicorp/terraform-config-inspect v0.0.0-20191212124732-c6ae6269b9d7
	github.com/terraform-providers/terraform-provider-aws v2.32.0+incompatible
	github.com/terraform-providers/terraform-provider-azurerm v1.42.0
	github.com/terraform-providers/terraform-provider-digitalocean v1.13.0
	github.com/terraform-providers/terraform-provider-google v2.17.0+incompatible
	github.com/terraform-providers/terraform-provider-linode v1.9.1
	k8s.io/api v0.16.4
	k8s.io/apimachinery v0.16.4
	k8s.io/client-go v11.0.0+incompatible
	k8s.io/kube-openapi v0.0.0-20200130172213-cdac1c71ff9f
	kmodules.xyz/client-go v0.0.0-20200201171629-70cdbdd3321b
)

replace (
	github.com/keybase/go-crypto => github.com/keybase/go-crypto v0.0.0-20200123153347-de78d2cb44f4
	k8s.io/client-go => k8s.io/client-go v0.16.4
)
