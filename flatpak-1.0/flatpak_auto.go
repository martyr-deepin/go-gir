package flatpak

/*
#cgo pkg-config: flatpak
#include <flatpak.h>
#include <stdlib.h>
#include <strings.h>
static void ProgressCallbackWrapper(const char* status, guint progress, gboolean estimating, gpointer user_data);
static void ProgressCallbackWrapper(const char* status, guint progress, gboolean estimating, gpointer user_data) {
    GClosure* closure = user_data;
    GValue params[3];
    bzero(params, 3*sizeof(GValue));
    g_value_init(&params[0], G_TYPE_STRING);
    g_value_set_string(&params[0], status);
    g_value_init(&params[1], G_TYPE_UINT);
    g_value_set_uint(&params[1], progress);
    g_value_init(&params[2], G_TYPE_BOOLEAN);
    g_value_set_boolean(&params[2], estimating);
    g_closure_invoke(closure, NULL, 3, params, NULL);
}
static FlatpakInstalledRef* _flatpak_installation_install(FlatpakInstallation* self, const char* remote_name, FlatpakRefKind kind, const char* name, const char* arch, const char* branch, GClosure* progress_data_for_progress, GCancellable* cancellable, GError **error) {
    return flatpak_installation_install(self, remote_name, kind, name, arch, branch, ProgressCallbackWrapper, progress_data_for_progress, cancellable, error);
}
static FlatpakInstalledRef* _flatpak_installation_install_bundle(FlatpakInstallation* self, GFile* file, GClosure* progress_data_for_progress, GCancellable* cancellable, GError **error) {
    return flatpak_installation_install_bundle(self, file, ProgressCallbackWrapper, progress_data_for_progress, cancellable, error);
}
static gboolean _flatpak_installation_uninstall(FlatpakInstallation* self, FlatpakRefKind kind, const char* name, const char* arch, const char* branch, GClosure* progress_data_for_progress, GCancellable* cancellable, GError **error) {
    return flatpak_installation_uninstall(self, kind, name, arch, branch, ProgressCallbackWrapper, progress_data_for_progress, cancellable, error);
}
static FlatpakInstalledRef* _flatpak_installation_update(FlatpakInstallation* self, FlatpakUpdateFlags flags, FlatpakRefKind kind, const char* name, const char* arch, const char* branch, GClosure* progress_data_for_progress, GCancellable* cancellable, GError **error) {
    return flatpak_installation_update(self, flags, kind, name, arch, branch, ProgressCallbackWrapper, progress_data_for_progress, cancellable, error);
}
*/
import "C"
import "github.com/electricface/go-auto-gir/gio-2.0"
import "github.com/electricface/go-auto-gir/glib-2.0"
import "github.com/electricface/go-auto-gir/gobject-2.0"
import "github.com/electricface/go-auto-gir/util"
import "unsafe"

type ProgressCallback func(status string, progress uint, estimating bool)

// Object Installation
type Installation struct {
	gobject.Object
}

func (v Installation) native() *C.FlatpakInstallation {
	return (*C.FlatpakInstallation)(v.Ptr)
}
func wrapInstallation(p *C.FlatpakInstallation) (v Installation) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapInstallation(p unsafe.Pointer) (v Installation) {
	v.Ptr = p
	return
}
func (v Installation) GetType() gobject.Type {
	return gobject.Type(C.flatpak_installation_get_type())
}
func (v Installation) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapInstallation(unsafe.Pointer(ptr)), nil
	}
}

// InstallationNewForPath is a wrapper around flatpak_installation_new_for_path().
func InstallationNewForPath(path gio.File, user bool, cancellable gio.Cancellable) (Installation, error) {
	var err glib.Error
	ret0 := C.flatpak_installation_new_for_path((*C.GFile)(path.Ptr), C.gboolean(util.Bool2Int(user)) /*go:.util*/, (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return Installation{}, err.GoValue()
	}
	return wrapInstallation(ret0), nil
}

// InstallationNewSystem is a wrapper around flatpak_installation_new_system().
func InstallationNewSystem(cancellable gio.Cancellable) (Installation, error) {
	var err glib.Error
	ret0 := C.flatpak_installation_new_system((*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return Installation{}, err.GoValue()
	}
	return wrapInstallation(ret0), nil
}

// InstallationNewSystemWithId is a wrapper around flatpak_installation_new_system_with_id().
func InstallationNewSystemWithId(id string, cancellable gio.Cancellable) (Installation, error) {
	id0 := C.CString(id)
	var err glib.Error
	ret0 := C.flatpak_installation_new_system_with_id(id0, (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(id0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return Installation{}, err.GoValue()
	}
	return wrapInstallation(ret0), nil
}

// InstallationNewUser is a wrapper around flatpak_installation_new_user().
func InstallationNewUser(cancellable gio.Cancellable) (Installation, error) {
	var err glib.Error
	ret0 := C.flatpak_installation_new_user((*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return Installation{}, err.GoValue()
	}
	return wrapInstallation(ret0), nil
}

// CreateMonitor is a wrapper around flatpak_installation_create_monitor().
func (self Installation) CreateMonitor(cancellable gio.Cancellable) (gio.FileMonitor, error) {
	var err glib.Error
	ret0 := C.flatpak_installation_create_monitor(self.native(), (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return gio.FileMonitor{}, err.GoValue()
	}
	return gio.WrapFileMonitor(unsafe.Pointer(ret0)) /*gir:Gio*/, nil
}

// DropCaches is a wrapper around flatpak_installation_drop_caches().
func (self Installation) DropCaches(cancellable gio.Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.flatpak_installation_drop_caches(self.native(), (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// FetchRemoteMetadataSync is a wrapper around flatpak_installation_fetch_remote_metadata_sync().
func (self Installation) FetchRemoteMetadataSync(remote_name string, ref Ref, cancellable gio.Cancellable) (glib.Bytes, error) {
	remote_name0 := C.CString(remote_name)
	var err glib.Error
	ret0 := C.flatpak_installation_fetch_remote_metadata_sync(self.native(), remote_name0, ref.native(), (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(remote_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return glib.Bytes{}, err.GoValue()
	}
	return glib.WrapBytes(unsafe.Pointer(ret0)) /*gir:GLib*/, nil
}

// FetchRemoteRefSync is a wrapper around flatpak_installation_fetch_remote_ref_sync().
func (self Installation) FetchRemoteRefSync(remote_name string, kind RefKind, name string, arch string, branch string, cancellable gio.Cancellable) (RemoteRef, error) {
	remote_name0 := C.CString(remote_name)
	name0 := C.CString(name)
	arch0 := C.CString(arch)
	branch0 := C.CString(branch)
	var err glib.Error
	ret0 := C.flatpak_installation_fetch_remote_ref_sync(self.native(), remote_name0, C.FlatpakRefKind(kind), name0, arch0, branch0, (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(remote_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(name0))        /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(arch0))        /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(branch0))      /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return RemoteRef{}, err.GoValue()
	}
	return wrapRemoteRef(ret0), nil
}

// FetchRemoteSizeSync is a wrapper around flatpak_installation_fetch_remote_size_sync().
func (self Installation) FetchRemoteSizeSync(remote_name string, ref Ref, cancellable gio.Cancellable) (bool, uint64, uint64, error) {
	remote_name0 := C.CString(remote_name)
	var download_size0 C.guint64
	var installed_size0 C.guint64
	var err glib.Error
	ret0 := C.flatpak_installation_fetch_remote_size_sync(self.native(), remote_name0, ref.native(), &download_size0, &installed_size0, (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(remote_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, 0, 0, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, uint64(download_size0), uint64(installed_size0), nil
}

// GetCurrentInstalledApp is a wrapper around flatpak_installation_get_current_installed_app().
func (self Installation) GetCurrentInstalledApp(name string, cancellable gio.Cancellable) (InstalledRef, error) {
	name0 := C.CString(name)
	var err glib.Error
	ret0 := C.flatpak_installation_get_current_installed_app(self.native(), name0, (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return InstalledRef{}, err.GoValue()
	}
	return wrapInstalledRef(ret0), nil
}

// GetDisplayName is a wrapper around flatpak_installation_get_display_name().
func (self Installation) GetDisplayName() string {
	ret0 := C.flatpak_installation_get_display_name(self.native())
	ret := C.GoString(ret0)
	return ret
}

// GetId is a wrapper around flatpak_installation_get_id().
func (self Installation) GetId() string {
	ret0 := C.flatpak_installation_get_id(self.native())
	ret := C.GoString(ret0)
	return ret
}

// GetInstalledRef is a wrapper around flatpak_installation_get_installed_ref().
func (self Installation) GetInstalledRef(kind RefKind, name string, arch string, branch string, cancellable gio.Cancellable) (InstalledRef, error) {
	name0 := C.CString(name)
	arch0 := C.CString(arch)
	branch0 := C.CString(branch)
	var err glib.Error
	ret0 := C.flatpak_installation_get_installed_ref(self.native(), C.FlatpakRefKind(kind), name0, arch0, branch0, (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(name0))   /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(arch0))   /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(branch0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return InstalledRef{}, err.GoValue()
	}
	return wrapInstalledRef(ret0), nil
}

// GetIsUser is a wrapper around flatpak_installation_get_is_user().
func (self Installation) GetIsUser() bool {
	ret0 := C.flatpak_installation_get_is_user(self.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPath is a wrapper around flatpak_installation_get_path().
func (self Installation) GetPath() gio.File {
	ret0 := C.flatpak_installation_get_path(self.native())
	return gio.WrapFile(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// GetPriority is a wrapper around flatpak_installation_get_priority().
func (self Installation) GetPriority() int {
	ret0 := C.flatpak_installation_get_priority(self.native())
	return int(ret0)
}

// GetRemoteByName is a wrapper around flatpak_installation_get_remote_by_name().
func (self Installation) GetRemoteByName(name string, cancellable gio.Cancellable) (Remote, error) {
	name0 := (*C.gchar)(C.CString(name))
	var err glib.Error
	ret0 := C.flatpak_installation_get_remote_by_name(self.native(), name0, (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return Remote{}, err.GoValue()
	}
	return wrapRemote(ret0), nil
}

// GetStorageType is a wrapper around flatpak_installation_get_storage_type().
func (self Installation) GetStorageType() StorageType {
	ret0 := C.flatpak_installation_get_storage_type(self.native())
	return StorageType(ret0)
}

// Install is a wrapper around flatpak_installation_install().
func (self Installation) Install(remote_name string, kind RefKind, name string, arch string, branch string, progress ProgressCallback, cancellable gio.Cancellable) (InstalledRef, error) {
	remote_name0 := C.CString(remote_name)
	name0 := C.CString(name)
	arch0 := C.CString(arch)
	branch0 := C.CString(branch)
	progress0 := (*C.GClosure)(gobject.ClosureNew(progress).Ptr) /*gir:GObject*/
	var err glib.Error
	ret0 := C._flatpak_installation_install(self.native(), remote_name0, C.FlatpakRefKind(kind), name0, arch0, branch0, progress0, (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(remote_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(name0))        /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(arch0))        /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(branch0))      /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return InstalledRef{}, err.GoValue()
	}
	return wrapInstalledRef(ret0), nil
}

// InstallBundle is a wrapper around flatpak_installation_install_bundle().
func (self Installation) InstallBundle(file gio.File, progress ProgressCallback, cancellable gio.Cancellable) (InstalledRef, error) {
	progress0 := (*C.GClosure)(gobject.ClosureNew(progress).Ptr) /*gir:GObject*/
	var err glib.Error
	ret0 := C._flatpak_installation_install_bundle(self.native(), (*C.GFile)(file.Ptr), progress0, (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return InstalledRef{}, err.GoValue()
	}
	return wrapInstalledRef(ret0), nil
}

// InstallRefFile is a wrapper around flatpak_installation_install_ref_file().
func (self Installation) InstallRefFile(ref_file_data glib.Bytes, cancellable gio.Cancellable) (RemoteRef, error) {
	var err glib.Error
	ret0 := C.flatpak_installation_install_ref_file(self.native(), (*C.GBytes)(ref_file_data.Ptr), (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return RemoteRef{}, err.GoValue()
	}
	return wrapRemoteRef(ret0), nil
}

// Launch is a wrapper around flatpak_installation_launch().
func (self Installation) Launch(name string, arch string, branch string, commit string, cancellable gio.Cancellable) (bool, error) {
	name0 := C.CString(name)
	arch0 := C.CString(arch)
	branch0 := C.CString(branch)
	commit0 := C.CString(commit)
	var err glib.Error
	ret0 := C.flatpak_installation_launch(self.native(), name0, arch0, branch0, commit0, (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(name0))   /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(arch0))   /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(branch0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(commit0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// LoadAppOverrides is a wrapper around flatpak_installation_load_app_overrides().
func (self Installation) LoadAppOverrides(app_id string, cancellable gio.Cancellable) (string, error) {
	app_id0 := C.CString(app_id)
	var err glib.Error
	ret0 := C.flatpak_installation_load_app_overrides(self.native(), app_id0, (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(app_id0)) /*ch:<stdlib.h>*/
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return "", err.GoValue()
	}
	return ret, nil
}

// ModifyRemote is a wrapper around flatpak_installation_modify_remote().
func (self Installation) ModifyRemote(remote Remote, cancellable gio.Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.flatpak_installation_modify_remote(self.native(), remote.native(), (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// RemoveRemote is a wrapper around flatpak_installation_remove_remote().
func (self Installation) RemoveRemote(name string, cancellable gio.Cancellable) (bool, error) {
	name0 := C.CString(name)
	var err glib.Error
	ret0 := C.flatpak_installation_remove_remote(self.native(), name0, (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Uninstall is a wrapper around flatpak_installation_uninstall().
func (self Installation) Uninstall(kind RefKind, name string, arch string, branch string, progress ProgressCallback, cancellable gio.Cancellable) (bool, error) {
	name0 := C.CString(name)
	arch0 := C.CString(arch)
	branch0 := C.CString(branch)
	progress0 := (*C.GClosure)(gobject.ClosureNew(progress).Ptr) /*gir:GObject*/
	var err glib.Error
	ret0 := C._flatpak_installation_uninstall(self.native(), C.FlatpakRefKind(kind), name0, arch0, branch0, progress0, (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(name0))   /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(arch0))   /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(branch0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Update is a wrapper around flatpak_installation_update().
func (self Installation) Update(flags UpdateFlags, kind RefKind, name string, arch string, branch string, progress ProgressCallback, cancellable gio.Cancellable) (InstalledRef, error) {
	name0 := C.CString(name)
	arch0 := C.CString(arch)
	branch0 := C.CString(branch)
	progress0 := (*C.GClosure)(gobject.ClosureNew(progress).Ptr) /*gir:GObject*/
	var err glib.Error
	ret0 := C._flatpak_installation_update(self.native(), C.FlatpakUpdateFlags(flags), C.FlatpakRefKind(kind), name0, arch0, branch0, progress0, (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(name0))   /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(arch0))   /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(branch0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return InstalledRef{}, err.GoValue()
	}
	return wrapInstalledRef(ret0), nil
}

// UpdateRemoteSync is a wrapper around flatpak_installation_update_remote_sync().
func (self Installation) UpdateRemoteSync(name string, cancellable gio.Cancellable) (bool, error) {
	name0 := C.CString(name)
	var err glib.Error
	ret0 := C.flatpak_installation_update_remote_sync(self.native(), name0, (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Object Ref
type Ref struct {
	gobject.Object
}

func (v Ref) native() *C.FlatpakRef {
	return (*C.FlatpakRef)(v.Ptr)
}
func wrapRef(p *C.FlatpakRef) (v Ref) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRef(p unsafe.Pointer) (v Ref) {
	v.Ptr = p
	return
}
func (v Ref) GetType() gobject.Type {
	return gobject.Type(C.flatpak_ref_get_type())
}
func (v Ref) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRef(unsafe.Pointer(ptr)), nil
	}
}

// FormatRef is a wrapper around flatpak_ref_format_ref().
func (self Ref) FormatRef() string {
	ret0 := C.flatpak_ref_format_ref(self.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetArch is a wrapper around flatpak_ref_get_arch().
func (self Ref) GetArch() string {
	ret0 := C.flatpak_ref_get_arch(self.native())
	ret := C.GoString(ret0)
	return ret
}

// GetBranch is a wrapper around flatpak_ref_get_branch().
func (self Ref) GetBranch() string {
	ret0 := C.flatpak_ref_get_branch(self.native())
	ret := C.GoString(ret0)
	return ret
}

// GetCommit is a wrapper around flatpak_ref_get_commit().
func (self Ref) GetCommit() string {
	ret0 := C.flatpak_ref_get_commit(self.native())
	ret := C.GoString(ret0)
	return ret
}

// GetKind is a wrapper around flatpak_ref_get_kind().
func (self Ref) GetKind() RefKind {
	ret0 := C.flatpak_ref_get_kind(self.native())
	return RefKind(ret0)
}

// GetName is a wrapper around flatpak_ref_get_name().
func (self Ref) GetName() string {
	ret0 := C.flatpak_ref_get_name(self.native())
	ret := C.GoString(ret0)
	return ret
}

// RefParse is a wrapper around flatpak_ref_parse().
func RefParse(ref string) (Ref, error) {
	ref0 := C.CString(ref)
	var err glib.Error
	ret0 := C.flatpak_ref_parse(ref0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(ref0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return Ref{}, err.GoValue()
	}
	return wrapRef(ret0), nil
}

// Object InstalledRef
type InstalledRef struct {
	Ref
}

func (v InstalledRef) native() *C.FlatpakInstalledRef {
	return (*C.FlatpakInstalledRef)(v.Ptr)
}
func wrapInstalledRef(p *C.FlatpakInstalledRef) (v InstalledRef) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapInstalledRef(p unsafe.Pointer) (v InstalledRef) {
	v.Ptr = p
	return
}
func (v InstalledRef) GetType() gobject.Type {
	return gobject.Type(C.flatpak_installed_ref_get_type())
}
func (v InstalledRef) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapInstalledRef(unsafe.Pointer(ptr)), nil
	}
}

// GetDeployDir is a wrapper around flatpak_installed_ref_get_deploy_dir().
func (self InstalledRef) GetDeployDir() string {
	ret0 := C.flatpak_installed_ref_get_deploy_dir(self.native())
	ret := C.GoString(ret0)
	return ret
}

// GetInstalledSize is a wrapper around flatpak_installed_ref_get_installed_size().
func (self InstalledRef) GetInstalledSize() uint64 {
	ret0 := C.flatpak_installed_ref_get_installed_size(self.native())
	return uint64(ret0)
}

// GetIsCurrent is a wrapper around flatpak_installed_ref_get_is_current().
func (self InstalledRef) GetIsCurrent() bool {
	ret0 := C.flatpak_installed_ref_get_is_current(self.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetLatestCommit is a wrapper around flatpak_installed_ref_get_latest_commit().
func (self InstalledRef) GetLatestCommit() string {
	ret0 := C.flatpak_installed_ref_get_latest_commit(self.native())
	ret := C.GoString(ret0)
	return ret
}

// GetOrigin is a wrapper around flatpak_installed_ref_get_origin().
func (self InstalledRef) GetOrigin() string {
	ret0 := C.flatpak_installed_ref_get_origin(self.native())
	ret := C.GoString(ret0)
	return ret
}

// LoadMetadata is a wrapper around flatpak_installed_ref_load_metadata().
func (self InstalledRef) LoadMetadata(cancellable gio.Cancellable) (glib.Bytes, error) {
	var err glib.Error
	ret0 := C.flatpak_installed_ref_load_metadata(self.native(), (*C.GCancellable)(cancellable.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return glib.Bytes{}, err.GoValue()
	}
	return glib.WrapBytes(unsafe.Pointer(ret0)) /*gir:GLib*/, nil
}

// Object RemoteRef
type RemoteRef struct {
	Ref
}

func (v RemoteRef) native() *C.FlatpakRemoteRef {
	return (*C.FlatpakRemoteRef)(v.Ptr)
}
func wrapRemoteRef(p *C.FlatpakRemoteRef) (v RemoteRef) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRemoteRef(p unsafe.Pointer) (v RemoteRef) {
	v.Ptr = p
	return
}
func (v RemoteRef) GetType() gobject.Type {
	return gobject.Type(C.flatpak_remote_ref_get_type())
}
func (v RemoteRef) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRemoteRef(unsafe.Pointer(ptr)), nil
	}
}

// GetRemoteName is a wrapper around flatpak_remote_ref_get_remote_name().
func (self RemoteRef) GetRemoteName() string {
	ret0 := C.flatpak_remote_ref_get_remote_name(self.native())
	ret := C.GoString(ret0)
	return ret
}

// Object BundleRef
type BundleRef struct {
	Ref
}

func (v BundleRef) native() *C.FlatpakBundleRef {
	return (*C.FlatpakBundleRef)(v.Ptr)
}
func wrapBundleRef(p *C.FlatpakBundleRef) (v BundleRef) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapBundleRef(p unsafe.Pointer) (v BundleRef) {
	v.Ptr = p
	return
}
func (v BundleRef) GetType() gobject.Type {
	return gobject.Type(C.flatpak_bundle_ref_get_type())
}
func (v BundleRef) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapBundleRef(unsafe.Pointer(ptr)), nil
	}
}

// BundleRefNew is a wrapper around flatpak_bundle_ref_new().
func BundleRefNew(file gio.File) (BundleRef, error) {
	var err glib.Error
	ret0 := C.flatpak_bundle_ref_new((*C.GFile)(file.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return BundleRef{}, err.GoValue()
	}
	return wrapBundleRef(ret0), nil
}

// GetAppstream is a wrapper around flatpak_bundle_ref_get_appstream().
func (self BundleRef) GetAppstream() glib.Bytes {
	ret0 := C.flatpak_bundle_ref_get_appstream(self.native())
	return glib.WrapBytes(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetFile is a wrapper around flatpak_bundle_ref_get_file().
func (self BundleRef) GetFile() gio.File {
	ret0 := C.flatpak_bundle_ref_get_file(self.native())
	return gio.WrapFile(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// GetIcon is a wrapper around flatpak_bundle_ref_get_icon().
func (self BundleRef) GetIcon(size int) glib.Bytes {
	ret0 := C.flatpak_bundle_ref_get_icon(self.native(), C.int(size))
	return glib.WrapBytes(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetInstalledSize is a wrapper around flatpak_bundle_ref_get_installed_size().
func (self BundleRef) GetInstalledSize() uint64 {
	ret0 := C.flatpak_bundle_ref_get_installed_size(self.native())
	return uint64(ret0)
}

// GetMetadata is a wrapper around flatpak_bundle_ref_get_metadata().
func (self BundleRef) GetMetadata() glib.Bytes {
	ret0 := C.flatpak_bundle_ref_get_metadata(self.native())
	return glib.WrapBytes(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetOrigin is a wrapper around flatpak_bundle_ref_get_origin().
func (self BundleRef) GetOrigin() string {
	ret0 := C.flatpak_bundle_ref_get_origin(self.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetRuntimeRepoUrl is a wrapper around flatpak_bundle_ref_get_runtime_repo_url().
func (self BundleRef) GetRuntimeRepoUrl() string {
	ret0 := C.flatpak_bundle_ref_get_runtime_repo_url(self.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// Object Remote
type Remote struct {
	gobject.Object
}

func (v Remote) native() *C.FlatpakRemote {
	return (*C.FlatpakRemote)(v.Ptr)
}
func wrapRemote(p *C.FlatpakRemote) (v Remote) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapRemote(p unsafe.Pointer) (v Remote) {
	v.Ptr = p
	return
}
func (v Remote) GetType() gobject.Type {
	return gobject.Type(C.flatpak_remote_get_type())
}
func (v Remote) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRemote(unsafe.Pointer(ptr)), nil
	}
}

// RemoteNew is a wrapper around flatpak_remote_new().
func RemoteNew(name string) Remote {
	name0 := C.CString(name)
	ret0 := C.flatpak_remote_new(name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapRemote(ret0)
}

// GetAppstreamDir is a wrapper around flatpak_remote_get_appstream_dir().
func (self Remote) GetAppstreamDir(arch string) gio.File {
	arch0 := C.CString(arch)
	ret0 := C.flatpak_remote_get_appstream_dir(self.native(), arch0)
	C.free(unsafe.Pointer(arch0))             /*ch:<stdlib.h>*/
	return gio.WrapFile(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// GetAppstreamTimestamp is a wrapper around flatpak_remote_get_appstream_timestamp().
func (self Remote) GetAppstreamTimestamp(arch string) gio.File {
	arch0 := C.CString(arch)
	ret0 := C.flatpak_remote_get_appstream_timestamp(self.native(), arch0)
	C.free(unsafe.Pointer(arch0))             /*ch:<stdlib.h>*/
	return gio.WrapFile(unsafe.Pointer(ret0)) /*gir:Gio*/
}

// GetDefaultBranch is a wrapper around flatpak_remote_get_default_branch().
func (self Remote) GetDefaultBranch() string {
	ret0 := C.flatpak_remote_get_default_branch(self.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetDisabled is a wrapper around flatpak_remote_get_disabled().
func (self Remote) GetDisabled() bool {
	ret0 := C.flatpak_remote_get_disabled(self.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetGpgVerify is a wrapper around flatpak_remote_get_gpg_verify().
func (self Remote) GetGpgVerify() bool {
	ret0 := C.flatpak_remote_get_gpg_verify(self.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetName is a wrapper around flatpak_remote_get_name().
func (self Remote) GetName() string {
	ret0 := C.flatpak_remote_get_name(self.native())
	ret := C.GoString(ret0)
	return ret
}

// GetNodeps is a wrapper around flatpak_remote_get_nodeps().
func (self Remote) GetNodeps() bool {
	ret0 := C.flatpak_remote_get_nodeps(self.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetNoenumerate is a wrapper around flatpak_remote_get_noenumerate().
func (self Remote) GetNoenumerate() bool {
	ret0 := C.flatpak_remote_get_noenumerate(self.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPrio is a wrapper around flatpak_remote_get_prio().
func (self Remote) GetPrio() int {
	ret0 := C.flatpak_remote_get_prio(self.native())
	return int(ret0)
}

// GetTitle is a wrapper around flatpak_remote_get_title().
func (self Remote) GetTitle() string {
	ret0 := C.flatpak_remote_get_title(self.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetUrl is a wrapper around flatpak_remote_get_url().
func (self Remote) GetUrl() string {
	ret0 := C.flatpak_remote_get_url(self.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// SetDefaultBranch is a wrapper around flatpak_remote_set_default_branch().
func (self Remote) SetDefaultBranch(default_branch string) {
	default_branch0 := C.CString(default_branch)
	C.flatpak_remote_set_default_branch(self.native(), default_branch0)
	C.free(unsafe.Pointer(default_branch0)) /*ch:<stdlib.h>*/
}

// SetDisabled is a wrapper around flatpak_remote_set_disabled().
func (self Remote) SetDisabled(disabled bool) {
	C.flatpak_remote_set_disabled(self.native(), C.gboolean(util.Bool2Int(disabled)) /*go:.util*/)
}

// SetGpgKey is a wrapper around flatpak_remote_set_gpg_key().
func (self Remote) SetGpgKey(gpg_key glib.Bytes) {
	C.flatpak_remote_set_gpg_key(self.native(), (*C.GBytes)(gpg_key.Ptr))
}

// SetGpgVerify is a wrapper around flatpak_remote_set_gpg_verify().
func (self Remote) SetGpgVerify(gpg_verify bool) {
	C.flatpak_remote_set_gpg_verify(self.native(), C.gboolean(util.Bool2Int(gpg_verify)) /*go:.util*/)
}

// SetNodeps is a wrapper around flatpak_remote_set_nodeps().
func (self Remote) SetNodeps(nodeps bool) {
	C.flatpak_remote_set_nodeps(self.native(), C.gboolean(util.Bool2Int(nodeps)) /*go:.util*/)
}

// SetNoenumerate is a wrapper around flatpak_remote_set_noenumerate().
func (self Remote) SetNoenumerate(noenumerate bool) {
	C.flatpak_remote_set_noenumerate(self.native(), C.gboolean(util.Bool2Int(noenumerate)) /*go:.util*/)
}

// SetPrio is a wrapper around flatpak_remote_set_prio().
func (self Remote) SetPrio(prio int) {
	C.flatpak_remote_set_prio(self.native(), C.int(prio))
}

// SetTitle is a wrapper around flatpak_remote_set_title().
func (self Remote) SetTitle(title string) {
	title0 := C.CString(title)
	C.flatpak_remote_set_title(self.native(), title0)
	C.free(unsafe.Pointer(title0)) /*ch:<stdlib.h>*/
}

// SetUrl is a wrapper around flatpak_remote_set_url().
func (self Remote) SetUrl(url string) {
	url0 := C.CString(url)
	C.flatpak_remote_set_url(self.native(), url0)
	C.free(unsafe.Pointer(url0)) /*ch:<stdlib.h>*/
}

type Error int

const (
	ErrorAlreadyInstalled Error = 0
	ErrorNotInstalled           = 1
)

type RefKind int

const (
	RefKindApp     RefKind = 0
	RefKindRuntime         = 1
)

type StorageType int

const (
	StorageTypeDefault  StorageType = 0
	StorageTypeHardDisk             = 1
	StorageTypeSdcard               = 2
	StorageTypeMmc                  = 3
)

type InstallFlags int

const (
	InstallFlagsNone           InstallFlags = 0
	InstallFlagsNoStaticDeltas              = 1
)

type UpdateFlags int

const (
	UpdateFlagsNone           UpdateFlags = 0
	UpdateFlagsNoDeploy                   = 1
	UpdateFlagsNoPull                     = 2
	UpdateFlagsNoStaticDeltas             = 4
)
