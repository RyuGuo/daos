# Container Management

DAOS containers are datasets managed by the users. A container is the unit of
snapshot and has a type. It can be a POSIX namespace, an HDF5 file or any other
new application-specific data model. The chapter explains how to manager
container, while the subsequent ones detail how to access a DAOS container from
applications.

## Container Creation/Destroy

Containers can be created and destroyed through the `daos(1)` utility.
provided to manage containers.

**To create and then query a container labeled `mycont` on a pool
(labeled `tank`):**
```bash
$ daos cont create tank --label mycont
  Container UUID : daefe12c-45d4-44f7-8e56-995d02549041
  Container Label: mycont
  Container Type : unknown
Successfully created container daefe12c-45d4-44f7-8e56-995d02549041

$ daos cont query tank mycont
  Container UUID             : daefe12c-45d4-44f7-8e56-995d02549041
  Container Label            : mycont
  Container Type             : unknown
  Pool UUID                  : 0d1fad71-5681-48d4-acdd-7bb2e786f12e
  Number of snapshots        : 0
  Latest Persistent Snapshot : 0
  Highest Aggregated Epoch   : 263546931609567249
  Container redundancy factor: 0
  Snapshot Epochs            :
```

The label can be up to 127 characters long and must only include
alphanumeric characters, colon (':'), period ('.') or underscore ('\_').

The container type (i.e., POSIX or HDF5) can be passed via the --type option.
As shown below, the pool UUID, container UUID, and container attributes can be
stored in the extended attributes of a POSIX file or directory for convenience.
Then subsequent invocations of the daos tools need to reference the path
to the POSIX file or directory.

```bash
$ daos cont create tank --path /tmp/mycontainer --type=POSIX --oclass=SX
  Container UUID : 30e5d364-62c9-4ddf-9284-1021359455f2
  Container Type : POSIX

Successfully created container 30e5d364-62c9-4ddf-9284-1021359455f2 type POSIX

$ daos cont query --path /tmp/mycontainer
  Container UUID             : 30e5d364-62c9-4ddf-9284-1021359455f2
  Container Type             : POSIX
  Pool UUID                  : 0d1fad71-5681-48d4-acdd-7bb2e786f12e
  Number of snapshots        : 0
  Latest Persistent Snapshot : 0
  Highest Aggregated Epoch   : 263548861715283973
  Container redundancy factor: 0
  Snapshot Epochs            :
  Object Class               : SX
  Chunk Size                 : 1.0 MiB
```

**To list all containers available in a pool:**

```bash
$ daos cont list tank
UUID                                 Label
----                                 -----
30e5d364-62c9-4ddf-9284-1021359455f2 container_label_not_set
daefe12c-45d4-44f7-8e56-995d02549041 mycont
```

**To destroy a container:**
```bash
$ daos cont destroy tank mycont
Successfully destroyed container mycont

$ daos cont destroy --path /tmp/mycontainer
Successfully destroyed container 30e5d364-62c9-4ddf-9284-1021359455f2
```

## Container Properties

Container properties are the main mechanism that one can use to control the
behavior of container. This includes the type of middleware, whether some
features like deduplication or checksum are enabled. Some properties are
immutable after creation creation, while some others can be dynamically
changed.

### Querying Properties

The user-level administration `daos` utility may be used to query a
container's properties. Refer to the manual page for full usage details.

```bash
$ daos cont get-prop tank mycont
# -OR- --path interface shown below
$ daos cont get-prop --path=/tmp/mycontainer
Properties for container mycont
Name                  Value
----                  -----
Highest Allocated OID 0
Checksum              off
Checksum Chunk Size   32 KiB
Compression           off
Deduplication         off
Dedupe Threshold      4.0 KiB
EC Cell Size          1.0 MiB
Encryption            off
Group                 jlombard@
Label                 mycont
Layout Type           unknown (0)
Layout Version        1
Max Snapshot          0
Owner                 jlombard@
Redundancy Factor     rf0
Redundancy Level      rank (1)
Server Checksumming   off
Health                HEALTHY
Access Control List   A::OWNER@:rwdtTaAo, A:G:GROUP@:rwtT
```

Additionally, a container's properties may be retrieved using the
libdaos API `daos_cont_query()` function. Refer to the file
src/include/daos\_cont.h Doxygen comments and the online documentation
available [here](https://daos-stack.github.io/html/).

### Changing Properties

By default, a container will inherit a set of default value for each property.
Those can be overridden at container creation time via the `--properties` option.

```bash
$ daos cont create tank --label mycont2 --properties cksum:sha1,dedup:hash,rf:1
  Container UUID : a6286ead-1952-4faa-bf87-00fc0f3785aa
  Container Label: mycont2
  Container Type : unknown
Successfully created container a6286ead-1952-4faa-bf87-00fc0f3785aa

$ daos cont query tank mycont2
Properties for container mycont2
Name                  Value
----                  -----
Highest Allocated OID 0
Checksum              sha1
Checksum Chunk Size   32 KiB
Compression           off
Deduplication         hash
Dedupe Threshold      4.0 KiB
EC Cell Size          1.0 MiB
Encryption            off
Group                 jlombard@
Label                 mycont2
Layout Type           unknown (0)
Layout Version        1
Max Snapshot          0
Owner                 jlombard@
Redundancy Factor     rf1
Redundancy Level      rank (1)
Server Checksumming   off
Health                HEALTHY
Access Control List   A::OWNER@:rwdtTaAo, A:G:GROUP@:rwtT
```

Mutable properties can be modified after container creation via the `set-prop`
option.

```bash
$ daos cont set-prop tank mycont2 --properties label:mycont3
Properties were successfully set
```

This effectively changed the container label.

```bash
$ daos cont get-prop tank mycont2
ERROR: daos: DER_NONEXIST(-1005): The specified entity does not exist

$ daos cont get-prop tank mycont3
Properties for container mycont3
Name                  Value
----                  -----
Highest Allocated OID 0
Checksum              sha1
Checksum Chunk Size   32 KiB
Compression           off
Deduplication         hash
Dedupe Threshold      4.0 KiB
EC Cell Size          1.0 MiB
Encryption            off
Group                 jlombard@
Label                 mycont3
Layout Type           unknown (0)
Layout Version        1
Max Snapshot          0
Owner                 jlombard@
Redundancy Factor     rf1
Redundancy Level      rank (1)
Server Checksumming   off
Health                HEALTHY
Access Control List   A::OWNER@:rwdtTaAo, A:G:GROUP@:rwtT
```

### Property Values

The table below summarizes the different container properties available.

| **Container Property**  | **Immutable**   | **Description** |
| ------------------------| --------------- | ----------------|
| label			  | No              | String associate with a containers. e.g., "Cat\_Pics" or "ResNet50\_training\_data".|
| owner                   | Yes             | User acting as the owner of the container.|
| group                   | Yes             | Group acting as the owner of the container|
| layout\_type            | Yes             | The container type (POSIX, HDF5, ...)|
| layout\_ver             | Yes             | A version of the layout that can be used by I/O middleware to handle interoperability.|
| rf                      | Yes             | The redundancy factor that drives the minimal data protection required for objects stored in the container. e.g., RF1 means no data protection, RF3 only allows 3-way replication or erasure code N+2.|
| rf\_lvl                 | Yes             | The fault domain level to use for data redundancy placement. This is used to determine object placement.|
| ec\_cell                | Yes             | Erasure code cell size for erasure-coded objects.|
| cksum                   | Yes             | Checksum off, or algorithm to use (adler32, crc[16|32|64] or sha[1|256|512]).|
| cksum\_size             | Yes             | Checksum chunk size.|
| srv\_cksum              | Yes             | Perform additional checksum verification on server (default: off).|
| max\_snapshot           | No              | Impose a upper limit on number of snapshots to retain (default: 0, no limitation).|
| acl                     | No              | Container access control list.|
| compression             | Yes             | Whether online compression is enabled (off, lz4, deflate[1-4])|
| dedup                   | Yes             | Inline deduplication off, hash based (hash) or using memory compare (memcmp)|
| dedup\_threshold        | Yes             | Minimum I/O size to consider for deduplication|
| encryption              | Yes             | Inline encryption off, or algorithm
to use (XTS[128|256], CBC[128|192|256] or GCM]129|256]|
| status                  | No              | Current state of the container|
| alloc\_oid               | No              | Maximum allocated object ID by container allocator|

Refer to the [Data Integrity](#Data-Integrity) and [Access Control Lists](#Access_Control_Lists)
sections for more details on the checksum and access-related properties.

Refer to the [Inline Deduplication](#Inline_Deduplication_\(Preview\)) section
for details about additional properties for that preview feature that are not
listed here.

While those properties are currently stored persistently with container
metadata, many of them are still under development. The ability to modify some
of these properties on an existing container will also be provided in a future
release.

## Data Integrity

DAOS allows to detect and fix (when data protection is enabled) silent data
corruptions. This is done by calculating checksums for both data and metadata
in the DAOS library on the client side and storing those checksums persistently
in SCM. The checksums will then be validated on access and on update/write as
well on the server side if server verify option is enabled.

Corrupted data will never be returned to the application. When a corruption is
detected, DAOS will try to read from a different replica, if any.  If the
original data cannot be recovered, then an error will be reported to the
application.

To enable and configure checksums, the following container properties are used
during container create.

- cksum (`DAOS_PROP_CO_CSUM`): the type of checksum algorithm to use.
  Supported values are adler32, crc[16|32|64] or sha[1|256|512]. By default,
  checksum is disabled for new containers.
- cksum\_size (`DAOS_PROP_CO_CSUM_CHUNK_SIZE`): defines the chunk size used for
  creating checksums of array types. (default is 32K).
- srv\_cksum (`DAOS_PROP_CO_CSUM_SERVER_VERIFY`): Because of the probable decrease to
  IOPS, in most cases, it is not desired to verify checksums on an object
  update on the server side. It is sufficient for the client to verify on
  a fetch because any data corruption, whether on the object update,
  storage, or fetch, will be caught. However, there is an advantage to
  knowing if corruption happens on an update. The update would fail
  right away, indicating to the client to retry the RPC or report an
  error to upper levels.

For instance, to create a new container with crc64 checksum enabled and
checksum verification on the server side, one can use the following command
line:

```bash
$ daos cont create tank --label mycont --properties cksum:crc64,srv_cksum:on
Successfully created container dfa09efd-4529-482c-b7cd-748c29ef7419

$ daos cont get-prop  tank mycont4 | grep cksum
Checksum              crc64
Checksum Chunk Size   32 KiB
Server Checksumming   on
```

!!! note
    Note that currently, once a container is created, its checksum configuration
    cannot be changed.

## Inline Deduplication (Preview)

Data deduplication (dedup) is a process that allows to eliminate duplicated
data copies in order to decrease capacity requirements. DAOS has some initial
support of inline dedup.

When dedup is enabled, each DAOS server maintains a per-pool table indexing
extents by their hash (i.e. checksum). Any new I/Os bigger than the
deduplication threshold will thus be looked up in this table to find out
whether an existing extent with the same signature has already been stored.
If an extent is found, then two options are provided:

- Transferring the data from the client to the server and doing a memory compare
  (i.e. memcmp) of the two extents to verify that they are indeed identical.
- Trusting the hash function and skipping the data transfer. To minimize issue
  with hash collision, a cryptographic hash function (i.e. SHA256) is used in
  this case. The benefit of this approarch is that the data to be written does
  not need to be transferred to the server. Data processing is thus greatly
  accelerated.

The inline dedup feature can be enabled on a per-container basis. To enable and
configure dedup, the following container properties are used:

- dedup (`DAOS_PROP_CO_DEDUP`): Type of dedup mechanism to use. Supported values
  are off (default), memcmp (memory compare) or hash (hash-based using SHA256).
- dedup\_threshold (`DAOS_PROP_CO_DEDUP_THRESHOLD`): defines the minimal I/O size
  to consider the I/O for dedup (default is 4K).

!!! warning
    Dedup is a feature preview in 2.0 and has some known
    limitations. Aggregation of deduplicated extents isn't supported and the
    checksum tree isn't persistent yet. This means that aggregation is disabled
    for a container with dedplication enabled and duplicated extents won't be
    matched after a server restart.
    NVMe isn't supported for dedup enabled container, so please make sure not
    using dedup on the pool with NVMe enabled.

## Compression & Encryption (unsupported)

The compression (`DAOS_PROP_CO_COMPRESS`) and encryption
(`DAOS_PROP_CO_ENCRYPT`) properties are reserved for configuring respectively
online compression and encryption.
These features are currently not on the roadmap.

## Snapshot & Rollback

The `daos` tool provides container {create/destroy}-snap and list-snaps
commands.

```bash
$ daos cont create-snap tank mycont
snapshot/epoch 262508437483290624 has been created

$ daos cont list-snaps tank mycont
Container's snapshots :
262508437483290624

$ daos cont destroy-snap tank mycont -e 262508437483290624
```

The max\_snapshot (`DAOS_PROP_CO_SNAPSHOT_MAX`) property is used to limit the
maximum number of snapshots to retain. When a new snapshot is taken, and the
threshold is reached, the oldest snapshot will be automatically deleted.

Rolling back the content of a container to a snapshot is planned for future
DAOS versions.

## User Attributes

Similar to POSIX extended attributes, users can attach some metadata to each
container through the `daos cont [set|get|list|del]-attr` commands or via the
`daos_cont_{list/get/set}_attr()` functions of the libdaos API.

```bash
$ daos cont set-attr tank mycont import_date "12/01/2021"

$ daos cont list-attr tank mycont
Attributes for container mycont:
Name
----
import_date

$ daos cont get-attr tank mycont import_date
Attributes for container mycont:
Name        Value
----        -----
import_date 12/01/2021

$ daos cont del-attr tank mycont import_date

$ daos cont list-attr tank mycont
Attributes for container mycont:
  No attributes found.
```

## Access Control Lists

Client user and group access for containers is controlled by
[Access Control Lists (ACLs)](https://daos-stack.github.io/overview/security/#access-control-lists).

Access-controlled container accesses include:

* Opening the container for access.

* Reading and writing data in the container.

  * Reading and writing objects.

  * Getting, setting, and listing user attributes.

  * Getting, setting, and listing snapshots.

* Deleting the container (if the pool does not grant the user permission).

* Getting and setting container properties.

* Getting and modifying the container ACL.

* Modifying the container's owner.


This is reflected in the set of supported
[container permissions](https://daos-stack.github.io/overview/security/#permissions).

### Pool vs. Container Permissions

In general, pool permissions are separate from container permissions, and access
to one does not guarantee access to the other. However, a user must have
permission to connect to a container's pool before they can access the
container in any way, regardless of their permissions on that container.
Once the user has connected to a pool, container access decisions are based on
the individual container ACL. A user need not have read/write access to a pool
in order to open a container with read/write access, for example.

There is one situation in which the pool can grant a container-level permission:
Container deletion. If a user has Delete permission on a pool, this grants them
the ability to delete *any* container in the pool, regardless of their
permissions on that container.

If the user does not have Delete permission on the pool, they will only be able
to delete containers for which they have been explicitly granted Delete
permission in the container's ACL.

### ACL at Container Creation

To create a container labeled mycont in a pool labeled tank with a custom ACL:

```bash
$ daos cont create <pool_label> --label <container_label> --acl-file=<path>
```

The ACL file format is detailed in the
[security overview](https://daos-stack.github.io/overview/security/#acl-file).

### Displaying ACL

To view a container's ACL:

```bash
$ daos cont get-acl <pool_label> <container_label>
```

The output is in the same string format used in the ACL file during creation,
with one ACE per line.

### Modifying a Container's ACL

For all of these commands using an ACL file, the ACL file must be in the format
noted above for container creation.

#### Overwriting the ACL

To replace a container's ACL with a new ACL:

```bash
$ daos cont overwrite-acl <pool_label> <container_label> --acl-file=<path>
```

#### Adding and Updating ACEs

To add or update multiple entries in an existing container ACL:

```bash
$ daos cont update-acl <pool_label> <container_label> --acl-file=<path>
```

To add or update a single entry in an existing container ACL:

```bash
$ daos cont update-acl <pool_label> <container_label> --entry <ACE>
```

If there is no existing entry for the principal in the ACL, the new entry is
added to the ACL. If there is already an entry for the principal, that entry
is replaced with the new one.

#### Removing an ACE

To delete an entry for a given principal in an existing container ACL:

```bash
$ daos cont delete-acl <pool_label> <container_label> --principal=<principal>
```

The `principal` argument refers to the
[principal](https://daos-stack.github.io/overview/security/#principal), or
identity, of the entry to be removed.

For the delete operation, the `principal` argument must be formatted as follows:

* Named user: `u:username@`
* Named group: `g:groupname@`
* Special principals:
  * `OWNER@`
  * `GROUP@`
  * `EVERYONE@`

The entry for that principal will be completely removed. This does not always
mean that the principal will have no access. Rather, their access to the
container will be decided based on the remaining ACL rules.

### Ownership

The ownership of the container corresponds to the special principals `OWNER@`
and `GROUP@` in the ACL. These values are a part of the container properties.
They may be set on container creation and changed later.

#### Privileges

The owner-user (`OWNER@`) has some implicit privileges on their container.
These permissions are silently included alongside any permissions that the
user was explicitly granted by entries in the ACL.

The owner-user will always have the following implicit capabilities:

* Open container
* Set ACL (A)
* Get ACL (a)

Because the owner's special permissions are implicit, they do not need to be
specified in the `OWNER@` entry. After
[determining](https://daos-stack.github.io/overview/security/#enforcement)
the user's privileges from the container ACL, DAOS checks whether the user
requesting access is the owner-user. If so, DAOS grants the owner's
implicit permissions to that user, in addition to any permissions granted by
the ACL.

In contrast, the owner-group (`GROUP@`) has no special permissions beyond those
explicitly granted by the `GROUP@` entry in the ACL.

#### Setting Ownership at Creation

The default owner user and group are the effective user and group of the user
creating the container. However, a specific user and/or group may be specified
at container creation time.

```bash
$ daos cont create <pool_label> <container_label> --user=<owner-user> --group=<owner-group>
```

The user and group names are case sensitive and must be formatted as
[DAOS ACL user/group principals](https://daos-stack.github.io/overview/security/#principal).

#### Changing Ownership

To change the owner user:

```bash
$ daos cont set-owner <pool_label> <container_label> --user=<owner-user>
```

To change the owner group:

```bash
$ daos cont set-owner <pool_label> <container_label> --group=<owner-group>
```

The user and group names are case sensitive and must be formatted as
[DAOS ACL user/group principals](https://daos-stack.github.io/overview/security/#principal).
