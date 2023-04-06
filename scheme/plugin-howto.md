# Endorsement Plugin

"Endorsement Plugin" is an umbrella term for the custom code that handles
extraction and normalisation of reference values and trust anchors for a
specific "Attestation Scheme."

The values are extracted from a reference integrity manifest (RIM), e.g., CoRIM,
and normalised into a key-value representation ready to be stored in the
appropriate KVStore.

When designing an Endorsement Plugin, the first thing to decide is the shape of
your KVStore records.  Specifically, the following two things must be defined
for both reference values and trust anchors:

1. The attributes of each record type
1. The lookup keys for locating the records

These design choices are dictated not only by the Endorsement format, but also
by the Evidence format, i.e., they depend on the Attestation Scheme as a whole.
In particular, the lookup keys should be structured so that they can be
assembled entirely from data available in Evidence.

Each KVStore _key_ is expressed as a custom URI (see
[KVStore#uri-format](../kvstore/README.md#uri-format) for the details).

Each KVStore _value_ is a JSON object with some fixed keys and a customisable
payload.  The common structure is as follows:

```json
{
  "scheme": <attestation scheme>
  "type": <record type>
  "attributes": {
    <any defined by the combo scheme/type>
  }
}
```

## A Fictitious TPM-based Attestation Scheme

It may be easier to look at a concrete use case to see how these abstract
principles are put into practice.  To this end we will use an imaginary
TPM-based Attestation Scheme called `MY_TPM`, and assume that CoRIM is the RIM
format.  More specifically, the assumption is that trust anchors and reference
values are conveyed as CoMID attestation verification key and reference value
triples respectively.

In `MY_TPM` trust anchors are per-device raw public keys, whilst reference
values are digests associated with specific PCRs that must be the same for all
devices in a "class" (e.g., the deployed fleet).

### Trust Anchor Design

Trust anchors are unique to a given device - note that this is an arbitrary
choice we made for our `MY_TPM` example, it is not generally true.

The natural way this is expressed in CoMID is by means of an "instance
environment" made of an "instance" unique identifier and, optionally, a "class"
identifier.

Each record will contain exactly one key.

#### Normalised Record Layout

The `scheme` is `MY_TPM` and `type` is `VERIFICATION_KEY`.

The `attributes` bag contains the raw public key (`my-tpm.pkey`) alongside any
identification metadata needed to synthesise the lookup key for this record.  In
this case, `my-tpm.instance-id` and `my-tpm.class-id`.  Note that the types of
the attributes depend on their native CoMID representation.

```json
{
  "scheme": "MY_TPM",
  "type": "VERIFICATION_KEY",
  "attributes": {
    "my-tpm.class-id": "<e.g., UUID>",
    "my-tpm.instance-id": "<e.g., UUID>",
    "my-tpm.pkey": "<Base64 URL-safe encoded SPKI>"
  }
}
```

#### Lookup Keys

The lookup key for trust anchor follows the KVStore URI convention (see above),
using the device instance identifier as path:

```
MY_TPM://<tenant-id>/<instance-id>
```

### Reference Values Design

Reference Values are shared by all devices of the same class.  Again, while
common this is not necessarily true for all Attestation Schemes.

The natural way this is expressed in CoMID is by means of a "class environment",
i.e., one with only a "class" identifier.

Each record will contain exactly one reference value.  Note that, for algorithm
agility reasons, the same PCR may be associated with more than one reference
value.  Each of these will have have their own record.

#### Normalised Record Layout

The `scheme` is `MY_TPM` and `type` is `REFERENCE_VALUE`.

The `attributes` bag contains the PCR index (`my-tpm.pcr`) and the associated
measurement comprising the hash algorithm (`my-tpm.alg-id`) and the actual value
(`my-tpm.digest`).

Any identification metadata needed to synthesise the lookup key for this record
are stashed in the `attributes`.  In this case, `my-tpm.class-id` that, for a
matching record, is normally the same as the one found in the trust anchor case.

Again, the exact types of the attributes depend on their native CoMID
representation.

```json
{
  "scheme": "MY_TPM",
  "type": "REFERENCE_VALUE",
  "attributes": {
    "my-tpm.pcr": "<PCR index>",
    "my-tpm.alg-id": "<hash algorithm identifier>",
    "my-tpm.digest": "<hash value>",
    "my-tpm.class-id": "<must match the one in the trust anchor record>"
  }
}
```

#### Lookup Keys

The lookup key for reference values follows the KVStore URI convention (see
above) using the device class identifier as path:

```
MY_TPM://<tenant-id>/<class-id>
```