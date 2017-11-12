package gio

/*
#cgo pkg-config: gio-2.0 gio-unix-2.0
#include <gio/gdesktopappinfo.h>
#include <gio/gfiledescriptorbased.h>
#include <gio/gio.h>
#include <gio/gunixconnection.h>
#include <gio/gunixcredentialsmessage.h>
#include <gio/gunixfdlist.h>
#include <gio/gunixfdmessage.h>
#include <gio/gunixinputstream.h>
#include <gio/gunixmounts.h>
#include <gio/gunixoutputstream.h>
#include <gio/gunixsocketaddress.h>
#include <stdlib.h>
#include <strings.h>
static void AsyncReadyCallbackWrapper(GObject* source_object, GAsyncResult* res, gpointer user_data);
static void FileMeasureProgressCallbackWrapper(gboolean reporting, guint64 current_size, guint64 num_dirs, guint64 num_files, gpointer user_data);
static void FileProgressCallbackWrapper(goffset current_num_bytes, goffset total_num_bytes, gpointer user_data);
static void _g_app_info_launch_default_for_uri_async(const char* uri, GAppLaunchContext* launch_context, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_app_info_launch_default_for_uri_async(uri, launch_context, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_append_to_async(GFile* file, GFileCreateFlags flags, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_append_to_async(file, flags, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static gboolean _g_file_copy(GFile* source, GFile* destination, GFileCopyFlags flags, GCancellable* cancellable, GClosure* progress_callback_data_for_progress_callback, GError **error) {
    return g_file_copy(source, destination, flags, cancellable, FileProgressCallbackWrapper, progress_callback_data_for_progress_callback, error);
}
static void _g_file_copy_async(GFile* source, GFile* destination, GFileCopyFlags flags, int io_priority, GCancellable* cancellable, GClosure* progress_callback_data_for_progress_callback, GClosure* user_data_for_callback) {
    return g_file_copy_async(source, destination, flags, io_priority, cancellable, FileProgressCallbackWrapper, progress_callback_data_for_progress_callback, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_create_async(GFile* file, GFileCreateFlags flags, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_create_async(file, flags, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_create_readwrite_async(GFile* file, GFileCreateFlags flags, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_create_readwrite_async(file, flags, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_delete_async(GFile* file, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_delete_async(file, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_eject_mountable_with_operation(GFile* file, GMountUnmountFlags flags, GMountOperation* mount_operation, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_eject_mountable_with_operation(file, flags, mount_operation, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_enumerate_children_async(GFile* file, const char* attributes, GFileQueryInfoFlags flags, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_enumerate_children_async(file, attributes, flags, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_find_enclosing_mount_async(GFile* file, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_find_enclosing_mount_async(file, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_make_directory_async(GFile* file, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_make_directory_async(file, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static gboolean _g_file_measure_disk_usage(GFile* file, GFileMeasureFlags flags, GCancellable* cancellable, GClosure* progress_data_for_progress_callback, guint64* disk_usage, guint64* num_dirs, guint64* num_files, GError **error) {
    return g_file_measure_disk_usage(file, flags, cancellable, FileMeasureProgressCallbackWrapper, progress_data_for_progress_callback, disk_usage, num_dirs, num_files, error);
}
static void _g_file_measure_disk_usage_async(GFile* file, GFileMeasureFlags flags, gint io_priority, GCancellable* cancellable, GClosure* progress_data_for_progress_callback, GClosure* user_data_for_callback) {
    return g_file_measure_disk_usage_async(file, flags, io_priority, cancellable, FileMeasureProgressCallbackWrapper, progress_data_for_progress_callback, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_mount_enclosing_volume(GFile* location, GMountMountFlags flags, GMountOperation* mount_operation, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_mount_enclosing_volume(location, flags, mount_operation, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_mount_mountable(GFile* file, GMountMountFlags flags, GMountOperation* mount_operation, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_mount_mountable(file, flags, mount_operation, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static gboolean _g_file_move(GFile* source, GFile* destination, GFileCopyFlags flags, GCancellable* cancellable, GClosure* progress_callback_data_for_progress_callback, GError **error) {
    return g_file_move(source, destination, flags, cancellable, FileProgressCallbackWrapper, progress_callback_data_for_progress_callback, error);
}
static void _g_file_open_readwrite_async(GFile* file, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_open_readwrite_async(file, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_poll_mountable(GFile* file, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_poll_mountable(file, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_query_filesystem_info_async(GFile* file, const char* attributes, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_query_filesystem_info_async(file, attributes, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_query_info_async(GFile* file, const char* attributes, GFileQueryInfoFlags flags, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_query_info_async(file, attributes, flags, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_read_async(GFile* file, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_read_async(file, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_replace_async(GFile* file, const char* etag, gboolean make_backup, GFileCreateFlags flags, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_replace_async(file, etag, make_backup, flags, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_set_attributes_async(GFile* file, GFileInfo* info, GFileQueryInfoFlags flags, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_set_attributes_async(file, info, flags, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_set_display_name_async(GFile* file, const char* display_name, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_set_display_name_async(file, display_name, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_start_mountable(GFile* file, GDriveStartFlags flags, GMountOperation* start_operation, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_start_mountable(file, flags, start_operation, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_stop_mountable(GFile* file, GMountUnmountFlags flags, GMountOperation* mount_operation, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_stop_mountable(file, flags, mount_operation, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_trash_async(GFile* file, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_trash_async(file, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_unmount_mountable_with_operation(GFile* file, GMountUnmountFlags flags, GMountOperation* mount_operation, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_unmount_mountable_with_operation(file, flags, mount_operation, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_output_stream_close_async(GOutputStream* stream, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_output_stream_close_async(stream, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_output_stream_flush_async(GOutputStream* stream, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_output_stream_flush_async(stream, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_output_stream_splice_async(GOutputStream* stream, GInputStream* source, GOutputStreamSpliceFlags flags, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_output_stream_splice_async(stream, source, flags, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_output_stream_write_bytes_async(GOutputStream* stream, GBytes* bytes, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_output_stream_write_bytes_async(stream, bytes, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_output_stream_query_info_async(GFileOutputStream* stream, const char* attributes, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_output_stream_query_info_async(stream, attributes, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_io_stream_close_async(GIOStream* stream, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_io_stream_close_async(stream, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_io_stream_splice_async(GIOStream* stream1, GIOStream* stream2, GIOStreamSpliceFlags flags, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_io_stream_splice_async(stream1, stream2, flags, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_input_stream_close_async(GInputStream* stream, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_input_stream_close_async(stream, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_input_stream_read_bytes_async(GInputStream* stream, gsize count, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_input_stream_read_bytes_async(stream, count, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_input_stream_skip_async(GInputStream* stream, gsize count, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_input_stream_skip_async(stream, count, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_io_stream_query_info_async(GFileIOStream* stream, const char* attributes, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_io_stream_query_info_async(stream, attributes, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_enumerator_close_async(GFileEnumerator* enumerator, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_enumerator_close_async(enumerator, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_enumerator_next_files_async(GFileEnumerator* enumerator, int num_files, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_enumerator_next_files_async(enumerator, num_files, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_mount_eject_with_operation(GMount* mount, GMountUnmountFlags flags, GMountOperation* mount_operation, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_mount_eject_with_operation(mount, flags, mount_operation, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_mount_guess_content_type(GMount* mount, gboolean force_rescan, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_mount_guess_content_type(mount, force_rescan, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_mount_remount(GMount* mount, GMountMountFlags flags, GMountOperation* mount_operation, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_mount_remount(mount, flags, mount_operation, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_mount_unmount_with_operation(GMount* mount, GMountUnmountFlags flags, GMountOperation* mount_operation, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_mount_unmount_with_operation(mount, flags, mount_operation, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_drive_eject_with_operation(GDrive* drive, GMountUnmountFlags flags, GMountOperation* mount_operation, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_drive_eject_with_operation(drive, flags, mount_operation, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_drive_poll_for_media(GDrive* drive, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_drive_poll_for_media(drive, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_drive_start(GDrive* drive, GDriveStartFlags flags, GMountOperation* mount_operation, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_drive_start(drive, flags, mount_operation, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_drive_stop(GDrive* drive, GMountUnmountFlags flags, GMountOperation* mount_operation, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_drive_stop(drive, flags, mount_operation, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_volume_eject_with_operation(GVolume* volume, GMountUnmountFlags flags, GMountOperation* mount_operation, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_volume_eject_with_operation(volume, flags, mount_operation, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_volume_mount(GVolume* volume, GMountMountFlags flags, GMountOperation* mount_operation, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_volume_mount(volume, flags, mount_operation, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_file_input_stream_query_info_async(GFileInputStream* stream, const char* attributes, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_file_input_stream_query_info_async(stream, attributes, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_dbus_connection_call(GDBusConnection* connection, const gchar* bus_name, const gchar* object_path, const gchar* interface_name, const gchar* method_name, GVariant* parameters, const GVariantType* reply_type, GDBusCallFlags flags, gint timeout_msec, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_dbus_connection_call(connection, bus_name, object_path, interface_name, method_name, parameters, reply_type, flags, timeout_msec, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_dbus_connection_call_with_unix_fd_list(GDBusConnection* connection, const gchar* bus_name, const gchar* object_path, const gchar* interface_name, const gchar* method_name, GVariant* parameters, const GVariantType* reply_type, GDBusCallFlags flags, gint timeout_msec, GUnixFDList* fd_list, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_dbus_connection_call_with_unix_fd_list(connection, bus_name, object_path, interface_name, method_name, parameters, reply_type, flags, timeout_msec, fd_list, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_dbus_connection_close(GDBusConnection* connection, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_dbus_connection_close(connection, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_dbus_connection_flush(GDBusConnection* connection, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_dbus_connection_flush(connection, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_dbus_connection_new(GIOStream* stream, const gchar* guid, GDBusConnectionFlags flags, GDBusAuthObserver* observer, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_dbus_connection_new(stream, guid, flags, observer, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_dbus_connection_new_for_address(const gchar* address, GDBusConnectionFlags flags, GDBusAuthObserver* observer, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_dbus_connection_new_for_address(address, flags, observer, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_async_initable_init_async(GAsyncInitable* initable, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_async_initable_init_async(initable, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_socket_address_enumerator_next_async(GSocketAddressEnumerator* enumerator, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_socket_address_enumerator_next_async(enumerator, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_dtls_connection_close_async(GDtlsConnection* conn, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_dtls_connection_close_async(conn, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_dtls_connection_handshake_async(GDtlsConnection* conn, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_dtls_connection_handshake_async(conn, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_dtls_connection_shutdown_async(GDtlsConnection* conn, gboolean shutdown_read, gboolean shutdown_write, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_dtls_connection_shutdown_async(conn, shutdown_read, shutdown_write, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_loadable_icon_load_async(GLoadableIcon* icon, int size, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_loadable_icon_load_async(icon, size, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_network_monitor_can_reach_async(GNetworkMonitor* monitor, GSocketConnectable* connectable, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_network_monitor_can_reach_async(monitor, connectable, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_proxy_connect_async(GProxy* proxy, GIOStream* connection, GProxyAddress* proxy_address, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_proxy_connect_async(proxy, connection, proxy_address, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_proxy_resolver_lookup_async(GProxyResolver* resolver, const gchar* uri, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_proxy_resolver_lookup_async(resolver, uri, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_tls_database_lookup_certificate_for_handle_async(GTlsDatabase* self, const gchar* handle, GTlsInteraction* interaction, GTlsDatabaseLookupFlags flags, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_tls_database_lookup_certificate_for_handle_async(self, handle, interaction, flags, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_tls_database_lookup_certificate_issuer_async(GTlsDatabase* self, GTlsCertificate* certificate, GTlsInteraction* interaction, GTlsDatabaseLookupFlags flags, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_tls_database_lookup_certificate_issuer_async(self, certificate, interaction, flags, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_tls_database_verify_chain_async(GTlsDatabase* self, GTlsCertificate* chain, const gchar* purpose, GSocketConnectable* identity, GTlsInteraction* interaction, GTlsDatabaseVerifyFlags flags, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_tls_database_verify_chain_async(self, chain, purpose, identity, interaction, flags, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_tls_interaction_ask_password_async(GTlsInteraction* interaction, GTlsPassword* password, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_tls_interaction_ask_password_async(interaction, password, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_tls_interaction_request_certificate_async(GTlsInteraction* interaction, GTlsConnection* connection, GTlsCertificateRequestFlags flags, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_tls_interaction_request_certificate_async(interaction, connection, flags, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_tls_connection_handshake_async(GTlsConnection* conn, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_tls_connection_handshake_async(conn, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_buffered_input_stream_fill_async(GBufferedInputStream* stream, gssize count, int io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_buffered_input_stream_fill_async(stream, count, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_dbus_proxy_call(GDBusProxy* proxy, const gchar* method_name, GVariant* parameters, GDBusCallFlags flags, gint timeout_msec, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_dbus_proxy_call(proxy, method_name, parameters, flags, timeout_msec, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_dbus_proxy_call_with_unix_fd_list(GDBusProxy* proxy, const gchar* method_name, GVariant* parameters, GDBusCallFlags flags, gint timeout_msec, GUnixFDList* fd_list, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_dbus_proxy_call_with_unix_fd_list(proxy, method_name, parameters, flags, timeout_msec, fd_list, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_dbus_proxy_new(GDBusConnection* connection, GDBusProxyFlags flags, GDBusInterfaceInfo* info, const gchar* name, const gchar* object_path, const gchar* interface_name, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_dbus_proxy_new(connection, flags, info, name, object_path, interface_name, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_dbus_proxy_new_for_bus(GBusType bus_type, GDBusProxyFlags flags, GDBusInterfaceInfo* info, const gchar* name, const gchar* object_path, const gchar* interface_name, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_dbus_proxy_new_for_bus(bus_type, flags, info, name, object_path, interface_name, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_data_input_stream_read_line_async(GDataInputStream* stream, gint io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_data_input_stream_read_line_async(stream, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_data_input_stream_read_until_async(GDataInputStream* stream, const gchar* stop_chars, gint io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_data_input_stream_read_until_async(stream, stop_chars, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_data_input_stream_read_upto_async(GDataInputStream* stream, const gchar* stop_chars, gssize stop_chars_len, gint io_priority, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_data_input_stream_read_upto_async(stream, stop_chars, stop_chars_len, io_priority, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_permission_acquire_async(GPermission* permission, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_permission_acquire_async(permission, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_permission_release_async(GPermission* permission, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_permission_release_async(permission, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_resolver_lookup_by_address_async(GResolver* resolver, GInetAddress* address, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_resolver_lookup_by_address_async(resolver, address, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_resolver_lookup_by_name_async(GResolver* resolver, const gchar* hostname, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_resolver_lookup_by_name_async(resolver, hostname, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_resolver_lookup_records_async(GResolver* resolver, const gchar* rrname, GResolverRecordType record_type, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_resolver_lookup_records_async(resolver, rrname, record_type, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_resolver_lookup_service_async(GResolver* resolver, const gchar* service, const gchar* protocol, const gchar* domain, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_resolver_lookup_service_async(resolver, service, protocol, domain, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_socket_connection_connect_async(GSocketConnection* connection, GSocketAddress* address, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_socket_connection_connect_async(connection, address, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_socket_client_connect_async(GSocketClient* client, GSocketConnectable* connectable, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_socket_client_connect_async(client, connectable, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_socket_client_connect_to_host_async(GSocketClient* client, const gchar* host_and_port, guint16 default_port, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_socket_client_connect_to_host_async(client, host_and_port, default_port, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_socket_client_connect_to_service_async(GSocketClient* client, const gchar* domain, const gchar* service, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_socket_client_connect_to_service_async(client, domain, service, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_socket_client_connect_to_uri_async(GSocketClient* client, const gchar* uri, guint16 default_port, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_socket_client_connect_to_uri_async(client, uri, default_port, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_socket_listener_accept_async(GSocketListener* listener, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_socket_listener_accept_async(listener, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_socket_listener_accept_socket_async(GSocketListener* listener, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_socket_listener_accept_socket_async(listener, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_subprocess_communicate_async(GSubprocess* subprocess, GBytes* stdin_buf, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_subprocess_communicate_async(subprocess, stdin_buf, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_subprocess_communicate_utf8_async(GSubprocess* subprocess, const char* stdin_buf, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_subprocess_communicate_utf8_async(subprocess, stdin_buf, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_subprocess_wait_async(GSubprocess* subprocess, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_subprocess_wait_async(subprocess, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_subprocess_wait_check_async(GSubprocess* subprocess, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_subprocess_wait_check_async(subprocess, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static GTask* _g_task_new(gpointer source_object, GCancellable* cancellable, GClosure* callback_data_for_callback) {
    return g_task_new(source_object, cancellable, AsyncReadyCallbackWrapper, callback_data_for_callback);
}
static void _g_task_report_error(gpointer source_object, GClosure* callback_data_for_callback, gpointer source_tag, GError* error) {
    return g_task_report_error(source_object, AsyncReadyCallbackWrapper, callback_data_for_callback, source_tag, error);
}
static void _g_unix_connection_receive_credentials_async(GUnixConnection* connection, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_unix_connection_receive_credentials_async(connection, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void _g_unix_connection_send_credentials_async(GUnixConnection* connection, GCancellable* cancellable, GClosure* user_data_for_callback) {
    return g_unix_connection_send_credentials_async(connection, cancellable, AsyncReadyCallbackWrapper, user_data_for_callback);
}
static void AsyncReadyCallbackWrapper(GObject* source_object, GAsyncResult* res, gpointer user_data) {
    GClosure* closure = user_data;
    GValue params[2];
    bzero(params, 2*sizeof(GValue));
    g_value_init(&params[0], G_TYPE_OBJECT);
    g_value_set_object(&params[0], source_object);
    g_value_init(&params[1], g_async_result_get_type());
    g_value_set_object(&params[1], res);
    g_closure_invoke(closure, NULL, 2, params, NULL);
    g_closure_unref(closure);
}
static void FileMeasureProgressCallbackWrapper(gboolean reporting, guint64 current_size, guint64 num_dirs, guint64 num_files, gpointer user_data) {
    GClosure* closure = user_data;
    GValue params[4];
    bzero(params, 4*sizeof(GValue));
    g_value_init(&params[0], G_TYPE_BOOLEAN);
    g_value_set_boolean(&params[0], reporting);
    g_value_init(&params[1], G_TYPE_UINT64);
    g_value_set_uint64(&params[1], current_size);
    g_value_init(&params[2], G_TYPE_UINT64);
    g_value_set_uint64(&params[2], num_dirs);
    g_value_init(&params[3], G_TYPE_UINT64);
    g_value_set_uint64(&params[3], num_files);
    g_closure_invoke(closure, NULL, 4, params, NULL);
}
static void FileProgressCallbackWrapper(goffset current_num_bytes, goffset total_num_bytes, gpointer user_data) {
    GClosure* closure = user_data;
    GValue params[2];
    bzero(params, 2*sizeof(GValue));
    g_value_init(&params[0], G_TYPE_INT64);
    g_value_set_int64(&params[0], current_num_bytes);
    g_value_init(&params[1], G_TYPE_INT64);
    g_value_set_int64(&params[1], total_num_bytes);
    g_closure_invoke(closure, NULL, 2, params, NULL);
}
*/
import "C"
import "github.com/electricface/go-auto-gir/glib-2.0"
import "github.com/electricface/go-auto-gir/gobject-2.0"
import "github.com/electricface/go-auto-gir/util"
import "unsafe"

// Interface AppInfo
type AppInfo struct {
	Ptr unsafe.Pointer
}

func (v AppInfo) native() *C.GAppInfo {
	return (*C.GAppInfo)(v.Ptr)
}
func wrapAppInfo(p *C.GAppInfo) AppInfo {
	return AppInfo{unsafe.Pointer(p)}
}
func WrapAppInfo(p unsafe.Pointer) AppInfo {
	return AppInfo{p}
}
func (v AppInfo) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAppInfo(p unsafe.Pointer) interface{} {
	return WrapAppInfo(p)
}
func (v AppInfo) GetType() gobject.Type {
	return gobject.Type(C.g_app_info_get_type())
}
func (v AppInfo) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAppInfo(unsafe.Pointer(ptr)), nil
	}
}

// AddSupportsType is a wrapper around g_app_info_add_supports_type().
func (appinfo AppInfo) AddSupportsType(content_type string) (bool, error) {
	content_type0 := C.CString(content_type)
	var err glib.Error
	ret0 := C.g_app_info_add_supports_type(appinfo.native(), content_type0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(content_type0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// CanDelete is a wrapper around g_app_info_can_delete().
func (appinfo AppInfo) CanDelete() bool {
	ret0 := C.g_app_info_can_delete(appinfo.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// CanRemoveSupportsType is a wrapper around g_app_info_can_remove_supports_type().
func (appinfo AppInfo) CanRemoveSupportsType() bool {
	ret0 := C.g_app_info_can_remove_supports_type(appinfo.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Delete is a wrapper around g_app_info_delete().
func (appinfo AppInfo) Delete() bool {
	ret0 := C.g_app_info_delete(appinfo.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Dup is a wrapper around g_app_info_dup().
func (appinfo AppInfo) Dup() AppInfo {
	ret0 := C.g_app_info_dup(appinfo.native())
	return wrapAppInfo(ret0)
}

// Equal is a wrapper around g_app_info_equal().
func (appinfo1 AppInfo) Equal(appinfo2 AppInfo) bool {
	ret0 := C.g_app_info_equal(appinfo1.native(), appinfo2.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetCommandline is a wrapper around g_app_info_get_commandline().
func (appinfo AppInfo) GetCommandline() string {
	ret0 := C.g_app_info_get_commandline(appinfo.native())
	ret := C.GoString(ret0)
	return ret
}

// GetDescription is a wrapper around g_app_info_get_description().
func (appinfo AppInfo) GetDescription() string {
	ret0 := C.g_app_info_get_description(appinfo.native())
	ret := C.GoString(ret0)
	return ret
}

// GetDisplayName is a wrapper around g_app_info_get_display_name().
func (appinfo AppInfo) GetDisplayName() string {
	ret0 := C.g_app_info_get_display_name(appinfo.native())
	ret := C.GoString(ret0)
	return ret
}

// GetExecutable is a wrapper around g_app_info_get_executable().
func (appinfo AppInfo) GetExecutable() string {
	ret0 := C.g_app_info_get_executable(appinfo.native())
	ret := C.GoString(ret0)
	return ret
}

// GetIcon is a wrapper around g_app_info_get_icon().
func (appinfo AppInfo) GetIcon() Icon {
	ret0 := C.g_app_info_get_icon(appinfo.native())
	return wrapIcon(ret0)
}

// GetId is a wrapper around g_app_info_get_id().
func (appinfo AppInfo) GetId() string {
	ret0 := C.g_app_info_get_id(appinfo.native())
	ret := C.GoString(ret0)
	return ret
}

// GetName is a wrapper around g_app_info_get_name().
func (appinfo AppInfo) GetName() string {
	ret0 := C.g_app_info_get_name(appinfo.native())
	ret := C.GoString(ret0)
	return ret
}

// GetSupportedTypes is a wrapper around g_app_info_get_supported_types().
func (appinfo AppInfo) GetSupportedTypes() []string {
	ret0 := C.g_app_info_get_supported_types(appinfo.native())
	var ret0Slice []*C.char
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString(elem)
		ret[idx] = elemG
	}
	return ret
}

// Launch is a wrapper around g_app_info_launch().
func (appinfo AppInfo) Launch(files glib.List, launch_context AppLaunchContext) (bool, error) {
	var err glib.Error
	ret0 := C.g_app_info_launch(appinfo.native(), (*C.GList)(files.Ptr), launch_context.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// LaunchUris is a wrapper around g_app_info_launch_uris().
func (appinfo AppInfo) LaunchUris(uris glib.List, launch_context AppLaunchContext) (bool, error) {
	var err glib.Error
	ret0 := C.g_app_info_launch_uris(appinfo.native(), (*C.GList)(uris.Ptr), launch_context.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// RemoveSupportsType is a wrapper around g_app_info_remove_supports_type().
func (appinfo AppInfo) RemoveSupportsType(content_type string) (bool, error) {
	content_type0 := C.CString(content_type)
	var err glib.Error
	ret0 := C.g_app_info_remove_supports_type(appinfo.native(), content_type0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(content_type0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetAsDefaultForExtension is a wrapper around g_app_info_set_as_default_for_extension().
func (appinfo AppInfo) SetAsDefaultForExtension(extension string) (bool, error) {
	extension0 := C.CString(extension)
	var err glib.Error
	ret0 := C.g_app_info_set_as_default_for_extension(appinfo.native(), extension0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(extension0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetAsDefaultForType is a wrapper around g_app_info_set_as_default_for_type().
func (appinfo AppInfo) SetAsDefaultForType(content_type string) (bool, error) {
	content_type0 := C.CString(content_type)
	var err glib.Error
	ret0 := C.g_app_info_set_as_default_for_type(appinfo.native(), content_type0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(content_type0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetAsLastUsedForType is a wrapper around g_app_info_set_as_last_used_for_type().
func (appinfo AppInfo) SetAsLastUsedForType(content_type string) (bool, error) {
	content_type0 := C.CString(content_type)
	var err glib.Error
	ret0 := C.g_app_info_set_as_last_used_for_type(appinfo.native(), content_type0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(content_type0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// ShouldShow is a wrapper around g_app_info_should_show().
func (appinfo AppInfo) ShouldShow() bool {
	ret0 := C.g_app_info_should_show(appinfo.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SupportsFiles is a wrapper around g_app_info_supports_files().
func (appinfo AppInfo) SupportsFiles() bool {
	ret0 := C.g_app_info_supports_files(appinfo.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SupportsUris is a wrapper around g_app_info_supports_uris().
func (appinfo AppInfo) SupportsUris() bool {
	ret0 := C.g_app_info_supports_uris(appinfo.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// AppInfoCreateFromCommandline is a wrapper around g_app_info_create_from_commandline().
func AppInfoCreateFromCommandline(commandline string, application_name string, flags AppInfoCreateFlags) (AppInfo, error) {
	commandline0 := C.CString(commandline)
	application_name0 := C.CString(application_name)
	var err glib.Error
	ret0 := C.g_app_info_create_from_commandline(commandline0, application_name0, C.GAppInfoCreateFlags(flags), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(commandline0))      /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(application_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return AppInfo{}, err.GoValue()
	}
	return wrapAppInfo(ret0), nil
}

// AppInfoGetAll is a wrapper around g_app_info_get_all().
func AppInfoGetAll() glib.List {
	ret0 := C.g_app_info_get_all()
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapAppInfo(p) }) /*gir:GLib*/
}

// AppInfoGetAllForType is a wrapper around g_app_info_get_all_for_type().
func AppInfoGetAllForType(content_type string) glib.List {
	content_type0 := C.CString(content_type)
	ret0 := C.g_app_info_get_all_for_type(content_type0)
	C.free(unsafe.Pointer(content_type0)) /*ch:<stdlib.h>*/
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapAppInfo(p) }) /*gir:GLib*/
}

// AppInfoGetDefaultForType is a wrapper around g_app_info_get_default_for_type().
func AppInfoGetDefaultForType(content_type string, must_support_uris bool) AppInfo {
	content_type0 := C.CString(content_type)
	ret0 := C.g_app_info_get_default_for_type(content_type0, C.gboolean(util.Bool2Int(must_support_uris)) /*go:.util*/)
	C.free(unsafe.Pointer(content_type0)) /*ch:<stdlib.h>*/
	return wrapAppInfo(ret0)
}

// AppInfoGetDefaultForUriScheme is a wrapper around g_app_info_get_default_for_uri_scheme().
func AppInfoGetDefaultForUriScheme(uri_scheme string) AppInfo {
	uri_scheme0 := C.CString(uri_scheme)
	ret0 := C.g_app_info_get_default_for_uri_scheme(uri_scheme0)
	C.free(unsafe.Pointer(uri_scheme0)) /*ch:<stdlib.h>*/
	return wrapAppInfo(ret0)
}

// AppInfoGetFallbackForType is a wrapper around g_app_info_get_fallback_for_type().
func AppInfoGetFallbackForType(content_type string) glib.List {
	content_type0 := (*C.gchar)(C.CString(content_type))
	ret0 := C.g_app_info_get_fallback_for_type(content_type0)
	C.free(unsafe.Pointer(content_type0)) /*ch:<stdlib.h>*/
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapAppInfo(p) }) /*gir:GLib*/
}

// AppInfoGetRecommendedForType is a wrapper around g_app_info_get_recommended_for_type().
func AppInfoGetRecommendedForType(content_type string) glib.List {
	content_type0 := (*C.gchar)(C.CString(content_type))
	ret0 := C.g_app_info_get_recommended_for_type(content_type0)
	C.free(unsafe.Pointer(content_type0)) /*ch:<stdlib.h>*/
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapAppInfo(p) }) /*gir:GLib*/
}

// AppInfoLaunchDefaultForUri is a wrapper around g_app_info_launch_default_for_uri().
func AppInfoLaunchDefaultForUri(uri string, launch_context AppLaunchContext) (bool, error) {
	uri0 := C.CString(uri)
	var err glib.Error
	ret0 := C.g_app_info_launch_default_for_uri(uri0, launch_context.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(uri0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// AppInfoLaunchDefaultForUriAsync is a wrapper around g_app_info_launch_default_for_uri_async().
func AppInfoLaunchDefaultForUriAsync(uri string, launch_context AppLaunchContext, cancellable Cancellable, callback AsyncReadyCallback) {
	uri0 := C.CString(uri)
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_app_info_launch_default_for_uri_async(uri0, launch_context.native(), cancellable.native(), callback0)
	C.free(unsafe.Pointer(uri0)) /*ch:<stdlib.h>*/
}

// AppInfoLaunchDefaultForUriFinish is a wrapper around g_app_info_launch_default_for_uri_finish().
func AppInfoLaunchDefaultForUriFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_app_info_launch_default_for_uri_finish(result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// AppInfoResetTypeAssociations is a wrapper around g_app_info_reset_type_associations().
func AppInfoResetTypeAssociations(content_type string) {
	content_type0 := C.CString(content_type)
	C.g_app_info_reset_type_associations(content_type0)
	C.free(unsafe.Pointer(content_type0)) /*ch:<stdlib.h>*/
}

// Object DesktopAppInfo
type DesktopAppInfo struct {
	gobject.Object
}

func (v DesktopAppInfo) native() *C.GDesktopAppInfo {
	return (*C.GDesktopAppInfo)(v.Ptr)
}
func wrapDesktopAppInfo(p *C.GDesktopAppInfo) (v DesktopAppInfo) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDesktopAppInfo(p unsafe.Pointer) (v DesktopAppInfo) {
	v.Ptr = p
	return
}
func (v DesktopAppInfo) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDesktopAppInfo(p unsafe.Pointer) interface{} {
	return WrapDesktopAppInfo(p)
}
func (v DesktopAppInfo) GetType() gobject.Type {
	return gobject.Type(C.g_desktop_app_info_get_type())
}
func (v DesktopAppInfo) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDesktopAppInfo(unsafe.Pointer(ptr)), nil
	}
}
func (v DesktopAppInfo) AppInfo() AppInfo {
	return WrapAppInfo(v.Ptr)
}

// DesktopAppInfoNew is a wrapper around g_desktop_app_info_new().
func DesktopAppInfoNew(desktop_id string) DesktopAppInfo {
	desktop_id0 := C.CString(desktop_id)
	ret0 := C.g_desktop_app_info_new(desktop_id0)
	C.free(unsafe.Pointer(desktop_id0)) /*ch:<stdlib.h>*/
	return wrapDesktopAppInfo(ret0)
}

// DesktopAppInfoNewFromFilename is a wrapper around g_desktop_app_info_new_from_filename().
func DesktopAppInfoNewFromFilename(filename string) DesktopAppInfo {
	filename0 := C.CString(filename)
	ret0 := C.g_desktop_app_info_new_from_filename(filename0)
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
	return wrapDesktopAppInfo(ret0)
}

// DesktopAppInfoNewFromKeyfile is a wrapper around g_desktop_app_info_new_from_keyfile().
func DesktopAppInfoNewFromKeyfile(key_file glib.KeyFile) DesktopAppInfo {
	ret0 := C.g_desktop_app_info_new_from_keyfile((*C.GKeyFile)(key_file.Ptr))
	return wrapDesktopAppInfo(ret0)
}

// GetActionName is a wrapper around g_desktop_app_info_get_action_name().
func (info DesktopAppInfo) GetActionName(action_name string) string {
	action_name0 := (*C.gchar)(C.CString(action_name))
	ret0 := C.g_desktop_app_info_get_action_name(info.native(), action_name0)
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetBoolean is a wrapper around g_desktop_app_info_get_boolean().
func (info DesktopAppInfo) GetBoolean(key string) bool {
	key0 := C.CString(key)
	ret0 := C.g_desktop_app_info_get_boolean(info.native(), key0)
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetCategories is a wrapper around g_desktop_app_info_get_categories().
func (info DesktopAppInfo) GetCategories() string {
	ret0 := C.g_desktop_app_info_get_categories(info.native())
	ret := C.GoString(ret0)
	return ret
}

// GetFilename is a wrapper around g_desktop_app_info_get_filename().
func (info DesktopAppInfo) GetFilename() string {
	ret0 := C.g_desktop_app_info_get_filename(info.native())
	ret := C.GoString(ret0)
	return ret
}

// GetGenericName is a wrapper around g_desktop_app_info_get_generic_name().
func (info DesktopAppInfo) GetGenericName() string {
	ret0 := C.g_desktop_app_info_get_generic_name(info.native())
	ret := C.GoString(ret0)
	return ret
}

// GetIsHidden is a wrapper around g_desktop_app_info_get_is_hidden().
func (info DesktopAppInfo) GetIsHidden() bool {
	ret0 := C.g_desktop_app_info_get_is_hidden(info.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetNodisplay is a wrapper around g_desktop_app_info_get_nodisplay().
func (info DesktopAppInfo) GetNodisplay() bool {
	ret0 := C.g_desktop_app_info_get_nodisplay(info.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetShowIn is a wrapper around g_desktop_app_info_get_show_in().
func (info DesktopAppInfo) GetShowIn(desktop_env string) bool {
	desktop_env0 := (*C.gchar)(C.CString(desktop_env))
	ret0 := C.g_desktop_app_info_get_show_in(info.native(), desktop_env0)
	C.free(unsafe.Pointer(desktop_env0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))      /*go:.util*/
}

// GetStartupWmClass is a wrapper around g_desktop_app_info_get_startup_wm_class().
func (info DesktopAppInfo) GetStartupWmClass() string {
	ret0 := C.g_desktop_app_info_get_startup_wm_class(info.native())
	ret := C.GoString(ret0)
	return ret
}

// GetString is a wrapper around g_desktop_app_info_get_string().
func (info DesktopAppInfo) GetString(key string) string {
	key0 := C.CString(key)
	ret0 := C.g_desktop_app_info_get_string(info.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// HasKey is a wrapper around g_desktop_app_info_has_key().
func (info DesktopAppInfo) HasKey(key string) bool {
	key0 := C.CString(key)
	ret0 := C.g_desktop_app_info_has_key(info.native(), key0)
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// LaunchAction is a wrapper around g_desktop_app_info_launch_action().
func (info DesktopAppInfo) LaunchAction(action_name string, launch_context AppLaunchContext) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.g_desktop_app_info_launch_action(info.native(), action_name0, launch_context.native())
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// ListActions is a wrapper around g_desktop_app_info_list_actions().
func (info DesktopAppInfo) ListActions() []string {
	ret0 := C.g_desktop_app_info_list_actions(info.native())
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
	}
	return ret
}

// DesktopAppInfoGetImplementations is a wrapper around g_desktop_app_info_get_implementations().
func DesktopAppInfoGetImplementations(interface_ string) glib.List {
	interface0 := (*C.gchar)(C.CString(interface_))
	ret0 := C.g_desktop_app_info_get_implementations(interface0)
	C.free(unsafe.Pointer(interface0)) /*ch:<stdlib.h>*/
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapDesktopAppInfo(p) }) /*gir:GLib*/
}

// Object AppLaunchContext
type AppLaunchContext struct {
	gobject.Object
}

func (v AppLaunchContext) native() *C.GAppLaunchContext {
	return (*C.GAppLaunchContext)(v.Ptr)
}
func wrapAppLaunchContext(p *C.GAppLaunchContext) (v AppLaunchContext) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapAppLaunchContext(p unsafe.Pointer) (v AppLaunchContext) {
	v.Ptr = p
	return
}
func (v AppLaunchContext) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAppLaunchContext(p unsafe.Pointer) interface{} {
	return WrapAppLaunchContext(p)
}
func (v AppLaunchContext) GetType() gobject.Type {
	return gobject.Type(C.g_app_launch_context_get_type())
}
func (v AppLaunchContext) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAppLaunchContext(unsafe.Pointer(ptr)), nil
	}
}

// AppLaunchContextNew is a wrapper around g_app_launch_context_new().
func AppLaunchContextNew() AppLaunchContext {
	ret0 := C.g_app_launch_context_new()
	return wrapAppLaunchContext(ret0)
}

// GetDisplay is a wrapper around g_app_launch_context_get_display().
func (context AppLaunchContext) GetDisplay(info AppInfo, files glib.List) string {
	ret0 := C.g_app_launch_context_get_display(context.native(), info.native(), (*C.GList)(files.Ptr))
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetEnvironment is a wrapper around g_app_launch_context_get_environment().
func (context AppLaunchContext) GetEnvironment() []string {
	ret0 := C.g_app_launch_context_get_environment(context.native())
	var ret0Slice []*C.char
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString(elem)
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetStartupNotifyId is a wrapper around g_app_launch_context_get_startup_notify_id().
func (context AppLaunchContext) GetStartupNotifyId(info AppInfo, files glib.List) string {
	ret0 := C.g_app_launch_context_get_startup_notify_id(context.native(), info.native(), (*C.GList)(files.Ptr))
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// LaunchFailed is a wrapper around g_app_launch_context_launch_failed().
func (context AppLaunchContext) LaunchFailed(startup_notify_id string) {
	startup_notify_id0 := C.CString(startup_notify_id)
	C.g_app_launch_context_launch_failed(context.native(), startup_notify_id0)
	C.free(unsafe.Pointer(startup_notify_id0)) /*ch:<stdlib.h>*/
}

// Setenv is a wrapper around g_app_launch_context_setenv().
func (context AppLaunchContext) Setenv(variable string, value string) {
	variable0 := C.CString(variable)
	value0 := C.CString(value)
	C.g_app_launch_context_setenv(context.native(), variable0, value0)
	C.free(unsafe.Pointer(variable0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(value0))    /*ch:<stdlib.h>*/
}

// Unsetenv is a wrapper around g_app_launch_context_unsetenv().
func (context AppLaunchContext) Unsetenv(variable string) {
	variable0 := C.CString(variable)
	C.g_app_launch_context_unsetenv(context.native(), variable0)
	C.free(unsafe.Pointer(variable0)) /*ch:<stdlib.h>*/
}

// Object Settings
type Settings struct {
	gobject.Object
}

func (v Settings) native() *C.GSettings {
	return (*C.GSettings)(v.Ptr)
}
func wrapSettings(p *C.GSettings) (v Settings) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSettings(p unsafe.Pointer) (v Settings) {
	v.Ptr = p
	return
}
func (v Settings) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSettings(p unsafe.Pointer) interface{} {
	return WrapSettings(p)
}
func (v Settings) GetType() gobject.Type {
	return gobject.Type(C.g_settings_get_type())
}
func (v Settings) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSettings(unsafe.Pointer(ptr)), nil
	}
}

// SettingsNew is a wrapper around g_settings_new().
func SettingsNew(schema_id string) Settings {
	schema_id0 := (*C.gchar)(C.CString(schema_id))
	ret0 := C.g_settings_new(schema_id0)
	C.free(unsafe.Pointer(schema_id0)) /*ch:<stdlib.h>*/
	return wrapSettings(ret0)
}

// SettingsNewFull is a wrapper around g_settings_new_full().
func SettingsNewFull(schema SettingsSchema, backend SettingsBackend, path string) Settings {
	path0 := (*C.gchar)(C.CString(path))
	ret0 := C.g_settings_new_full(schema.native(), backend.native(), path0)
	C.free(unsafe.Pointer(path0)) /*ch:<stdlib.h>*/
	return wrapSettings(ret0)
}

// SettingsNewWithBackend is a wrapper around g_settings_new_with_backend().
func SettingsNewWithBackend(schema_id string, backend SettingsBackend) Settings {
	schema_id0 := (*C.gchar)(C.CString(schema_id))
	ret0 := C.g_settings_new_with_backend(schema_id0, backend.native())
	C.free(unsafe.Pointer(schema_id0)) /*ch:<stdlib.h>*/
	return wrapSettings(ret0)
}

// SettingsNewWithBackendAndPath is a wrapper around g_settings_new_with_backend_and_path().
func SettingsNewWithBackendAndPath(schema_id string, backend SettingsBackend, path string) Settings {
	schema_id0 := (*C.gchar)(C.CString(schema_id))
	path0 := (*C.gchar)(C.CString(path))
	ret0 := C.g_settings_new_with_backend_and_path(schema_id0, backend.native(), path0)
	C.free(unsafe.Pointer(schema_id0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(path0))      /*ch:<stdlib.h>*/
	return wrapSettings(ret0)
}

// SettingsNewWithPath is a wrapper around g_settings_new_with_path().
func SettingsNewWithPath(schema_id string, path string) Settings {
	schema_id0 := (*C.gchar)(C.CString(schema_id))
	path0 := (*C.gchar)(C.CString(path))
	ret0 := C.g_settings_new_with_path(schema_id0, path0)
	C.free(unsafe.Pointer(schema_id0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(path0))      /*ch:<stdlib.h>*/
	return wrapSettings(ret0)
}

// Apply is a wrapper around g_settings_apply().
func (settings Settings) Apply() {
	C.g_settings_apply(settings.native())
}

// Bind is a wrapper around g_settings_bind().
func (settings Settings) Bind(key string, object gobject.Object, property string, flags SettingsBindFlags) {
	key0 := (*C.gchar)(C.CString(key))
	property0 := (*C.gchar)(C.CString(property))
	C.g_settings_bind(settings.native(), key0, C.gpointer((C.gpointer)(object.Ptr)), property0, C.GSettingsBindFlags(flags))
	C.free(unsafe.Pointer(key0))      /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(property0)) /*ch:<stdlib.h>*/
}

// BindWritable is a wrapper around g_settings_bind_writable().
func (settings Settings) BindWritable(key string, object gobject.Object, property string, inverted bool) {
	key0 := (*C.gchar)(C.CString(key))
	property0 := (*C.gchar)(C.CString(property))
	C.g_settings_bind_writable(settings.native(), key0, C.gpointer((C.gpointer)(object.Ptr)), property0, C.gboolean(util.Bool2Int(inverted)) /*go:.util*/)
	C.free(unsafe.Pointer(key0))      /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(property0)) /*ch:<stdlib.h>*/
}

// CreateAction is a wrapper around g_settings_create_action().
func (settings Settings) CreateAction(key string) Action {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_create_action(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return wrapAction(ret0)
}

// Delay is a wrapper around g_settings_delay().
func (settings Settings) Delay() {
	C.g_settings_delay(settings.native())
}

// GetBoolean is a wrapper around g_settings_get_boolean().
func (settings Settings) GetBoolean(key string) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_boolean(settings.native(), key0)
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetChild is a wrapper around g_settings_get_child().
func (settings Settings) GetChild(name string) Settings {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_settings_get_child(settings.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapSettings(ret0)
}

// GetDefaultValue is a wrapper around g_settings_get_default_value().
func (settings Settings) GetDefaultValue(key string) glib.Variant {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_default_value(settings.native(), key0)
	C.free(unsafe.Pointer(key0))                  /*ch:<stdlib.h>*/
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetDouble is a wrapper around g_settings_get_double().
func (settings Settings) GetDouble(key string) float64 {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_double(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return float64(ret0)
}

// GetEnum is a wrapper around g_settings_get_enum().
func (settings Settings) GetEnum(key string) int {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_enum(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return int(ret0)
}

// GetFlags is a wrapper around g_settings_get_flags().
func (settings Settings) GetFlags(key string) uint {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_flags(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return uint(ret0)
}

// GetHasUnapplied is a wrapper around g_settings_get_has_unapplied().
func (settings Settings) GetHasUnapplied() bool {
	ret0 := C.g_settings_get_has_unapplied(settings.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetInt is a wrapper around g_settings_get_int().
func (settings Settings) GetInt(key string) int {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_int(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return int(ret0)
}

// GetInt64 is a wrapper around g_settings_get_int64().
func (settings Settings) GetInt64(key string) int64 {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_int64(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return int64(ret0)
}

// GetString is a wrapper around g_settings_get_string().
func (settings Settings) GetString(key string) string {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_string(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetStrv is a wrapper around g_settings_get_strv().
func (settings Settings) GetStrv(key string) []string {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_strv(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetUint is a wrapper around g_settings_get_uint().
func (settings Settings) GetUint(key string) uint {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_uint(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return uint(ret0)
}

// GetUint64 is a wrapper around g_settings_get_uint64().
func (settings Settings) GetUint64(key string) uint64 {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_uint64(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	return uint64(ret0)
}

// GetUserValue is a wrapper around g_settings_get_user_value().
func (settings Settings) GetUserValue(key string) glib.Variant {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_user_value(settings.native(), key0)
	C.free(unsafe.Pointer(key0))                  /*ch:<stdlib.h>*/
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetValue is a wrapper around g_settings_get_value().
func (settings Settings) GetValue(key string) glib.Variant {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_get_value(settings.native(), key0)
	C.free(unsafe.Pointer(key0))                  /*ch:<stdlib.h>*/
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// IsWritable is a wrapper around g_settings_is_writable().
func (settings Settings) IsWritable(name string) bool {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_settings_is_writable(settings.native(), name0)
	C.free(unsafe.Pointer(name0))   /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ListChildren is a wrapper around g_settings_list_children().
func (settings Settings) ListChildren() []string {
	ret0 := C.g_settings_list_children(settings.native())
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// Reset is a wrapper around g_settings_reset().
func (settings Settings) Reset(key string) {
	key0 := (*C.gchar)(C.CString(key))
	C.g_settings_reset(settings.native(), key0)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
}

// Revert is a wrapper around g_settings_revert().
func (settings Settings) Revert() {
	C.g_settings_revert(settings.native())
}

// SetBoolean is a wrapper around g_settings_set_boolean().
func (settings Settings) SetBoolean(key string, value bool) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_set_boolean(settings.native(), key0, C.gboolean(util.Bool2Int(value)) /*go:.util*/)
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetDouble is a wrapper around g_settings_set_double().
func (settings Settings) SetDouble(key string, value float64) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_set_double(settings.native(), key0, C.gdouble(value))
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetEnum is a wrapper around g_settings_set_enum().
func (settings Settings) SetEnum(key string, value int) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_set_enum(settings.native(), key0, C.gint(value))
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetFlags is a wrapper around g_settings_set_flags().
func (settings Settings) SetFlags(key string, value uint) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_set_flags(settings.native(), key0, C.guint(value))
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetInt is a wrapper around g_settings_set_int().
func (settings Settings) SetInt(key string, value int) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_set_int(settings.native(), key0, C.gint(value))
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetInt64 is a wrapper around g_settings_set_int64().
func (settings Settings) SetInt64(key string, value int64) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_set_int64(settings.native(), key0, C.gint64(value))
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetString is a wrapper around g_settings_set_string().
func (settings Settings) SetString(key string, value string) bool {
	key0 := (*C.gchar)(C.CString(key))
	value0 := (*C.gchar)(C.CString(value))
	ret0 := C.g_settings_set_string(settings.native(), key0, value0)
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(value0))  /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetStrv is a wrapper around g_settings_set_strv().
func (settings Settings) SetStrv(key string, value []string) bool {
	key0 := (*C.gchar)(C.CString(key))
	value0 := make([]*C.gchar, len(value))
	for idx, elemG := range value {
		elem := (*C.gchar)(C.CString(elemG))
		value0[idx] = elem
	}
	var value0Ptr **C.gchar
	if len(value0) > 0 {
		value0Ptr = &value0[0]
	}
	ret0 := C.g_settings_set_strv(settings.native(), key0, value0Ptr)
	C.free(unsafe.Pointer(key0)) /*ch:<stdlib.h>*/
	for _, elem := range value0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetUint is a wrapper around g_settings_set_uint().
func (settings Settings) SetUint(key string, value uint) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_set_uint(settings.native(), key0, C.guint(value))
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetUint64 is a wrapper around g_settings_set_uint64().
func (settings Settings) SetUint64(key string, value uint64) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_set_uint64(settings.native(), key0, C.guint64(value))
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetValue is a wrapper around g_settings_set_value().
func (settings Settings) SetValue(key string, value glib.Variant) bool {
	key0 := (*C.gchar)(C.CString(key))
	ret0 := C.g_settings_set_value(settings.native(), key0, (*C.GVariant)(value.Ptr))
	C.free(unsafe.Pointer(key0))    /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SettingsSync is a wrapper around g_settings_sync().
func SettingsSync() {
	C.g_settings_sync()
}

// SettingsUnbind is a wrapper around g_settings_unbind().
func SettingsUnbind(object gobject.Object, property string) {
	property0 := (*C.gchar)(C.CString(property))
	C.g_settings_unbind(C.gpointer((C.gpointer)(object.Ptr)), property0)
	C.free(unsafe.Pointer(property0)) /*ch:<stdlib.h>*/
}

// Struct SettingsSchema
type SettingsSchema struct {
	Ptr unsafe.Pointer
}

func (v SettingsSchema) native() *C.GSettingsSchema {
	return (*C.GSettingsSchema)(v.Ptr)
}
func wrapSettingsSchema(p *C.GSettingsSchema) SettingsSchema {
	return SettingsSchema{unsafe.Pointer(p)}
}
func WrapSettingsSchema(p unsafe.Pointer) SettingsSchema {
	return SettingsSchema{p}
}
func (v SettingsSchema) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSettingsSchema(p unsafe.Pointer) interface{} {
	return WrapSettingsSchema(p)
}

// GetId is a wrapper around g_settings_schema_get_id().
func (schema SettingsSchema) GetId() string {
	ret0 := C.g_settings_schema_get_id(schema.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetKey is a wrapper around g_settings_schema_get_key().
func (schema SettingsSchema) GetKey(name string) SettingsSchemaKey {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_settings_schema_get_key(schema.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapSettingsSchemaKey(ret0)
}

// GetPath is a wrapper around g_settings_schema_get_path().
func (schema SettingsSchema) GetPath() string {
	ret0 := C.g_settings_schema_get_path(schema.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// HasKey is a wrapper around g_settings_schema_has_key().
func (schema SettingsSchema) HasKey(name string) bool {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_settings_schema_has_key(schema.native(), name0)
	C.free(unsafe.Pointer(name0))   /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ListChildren is a wrapper around g_settings_schema_list_children().
func (schema SettingsSchema) ListChildren() []string {
	ret0 := C.g_settings_schema_list_children(schema.native())
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// ListKeys is a wrapper around g_settings_schema_list_keys().
func (schema SettingsSchema) ListKeys() []string {
	ret0 := C.g_settings_schema_list_keys(schema.native())
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// Ref is a wrapper around g_settings_schema_ref().
func (schema SettingsSchema) Ref() SettingsSchema {
	ret0 := C.g_settings_schema_ref(schema.native())
	return wrapSettingsSchema(ret0)
}

// Unref is a wrapper around g_settings_schema_unref().
func (schema SettingsSchema) Unref() {
	C.g_settings_schema_unref(schema.native())
}

// Struct SettingsBackend
type SettingsBackend struct {
	Ptr unsafe.Pointer
}

func (v SettingsBackend) native() *C.GSettingsBackend {
	return (*C.GSettingsBackend)(v.Ptr)
}
func wrapSettingsBackend(p *C.GSettingsBackend) SettingsBackend {
	return SettingsBackend{unsafe.Pointer(p)}
}
func WrapSettingsBackend(p unsafe.Pointer) SettingsBackend {
	return SettingsBackend{p}
}
func (v SettingsBackend) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSettingsBackend(p unsafe.Pointer) interface{} {
	return WrapSettingsBackend(p)
}

// Interface Action
type Action struct {
	Ptr unsafe.Pointer
}

func (v Action) native() *C.GAction {
	return (*C.GAction)(v.Ptr)
}
func wrapAction(p *C.GAction) Action {
	return Action{unsafe.Pointer(p)}
}
func WrapAction(p unsafe.Pointer) Action {
	return Action{p}
}
func (v Action) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAction(p unsafe.Pointer) interface{} {
	return WrapAction(p)
}
func (v Action) GetType() gobject.Type {
	return gobject.Type(C.g_action_get_type())
}
func (v Action) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAction(unsafe.Pointer(ptr)), nil
	}
}

// Activate is a wrapper around g_action_activate().
func (action Action) Activate(parameter glib.Variant) {
	C.g_action_activate(action.native(), (*C.GVariant)(parameter.Ptr))
}

// ChangeState is a wrapper around g_action_change_state().
func (action Action) ChangeState(value glib.Variant) {
	C.g_action_change_state(action.native(), (*C.GVariant)(value.Ptr))
}

// GetEnabled is a wrapper around g_action_get_enabled().
func (action Action) GetEnabled() bool {
	ret0 := C.g_action_get_enabled(action.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetName is a wrapper around g_action_get_name().
func (action Action) GetName() string {
	ret0 := C.g_action_get_name(action.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetParameterType is a wrapper around g_action_get_parameter_type().
func (action Action) GetParameterType() glib.VariantType {
	ret0 := C.g_action_get_parameter_type(action.native())
	return glib.WrapVariantType(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetState is a wrapper around g_action_get_state().
func (action Action) GetState() glib.Variant {
	ret0 := C.g_action_get_state(action.native())
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetStateHint is a wrapper around g_action_get_state_hint().
func (action Action) GetStateHint() glib.Variant {
	ret0 := C.g_action_get_state_hint(action.native())
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetStateType is a wrapper around g_action_get_state_type().
func (action Action) GetStateType() glib.VariantType {
	ret0 := C.g_action_get_state_type(action.native())
	return glib.WrapVariantType(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// ActionNameIsValid is a wrapper around g_action_name_is_valid().
func ActionNameIsValid(action_name string) bool {
	action_name0 := (*C.gchar)(C.CString(action_name))
	ret0 := C.g_action_name_is_valid(action_name0)
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))      /*go:.util*/
}

// ActionParseDetailedName is a wrapper around g_action_parse_detailed_name().
func ActionParseDetailedName(detailed_name string) (bool, string, glib.Variant, error) {
	detailed_name0 := (*C.gchar)(C.CString(detailed_name))
	var action_name0 *C.gchar
	var target_value0 *C.GVariant
	var err glib.Error
	ret0 := C.g_action_parse_detailed_name(detailed_name0, &action_name0, &target_value0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(detailed_name0)) /*ch:<stdlib.h>*/
	action_name := C.GoString((*C.char)(action_name0))
	defer C.g_free(C.gpointer(action_name0))
	if err.Ptr != nil {
		defer err.Free()
		return false, "", glib.Variant{}, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, action_name, glib.WrapVariant(unsafe.Pointer(target_value0)) /*gir:GLib*/, nil
}

// ActionPrintDetailedName is a wrapper around g_action_print_detailed_name().
func ActionPrintDetailedName(action_name string, target_value glib.Variant) string {
	action_name0 := (*C.gchar)(C.CString(action_name))
	ret0 := C.g_action_print_detailed_name(action_name0, (*C.GVariant)(target_value.Ptr))
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// Interface File
type File struct {
	Ptr unsafe.Pointer
}

func (v File) native() *C.GFile {
	return (*C.GFile)(v.Ptr)
}
func wrapFile(p *C.GFile) File {
	return File{unsafe.Pointer(p)}
}
func WrapFile(p unsafe.Pointer) File {
	return File{p}
}
func (v File) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFile(p unsafe.Pointer) interface{} {
	return WrapFile(p)
}
func (v File) GetType() gobject.Type {
	return gobject.Type(C.g_file_get_type())
}
func (v File) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFile(unsafe.Pointer(ptr)), nil
	}
}

// AppendTo is a wrapper around g_file_append_to().
func (file File) AppendTo(flags FileCreateFlags, cancellable Cancellable) (FileOutputStream, error) {
	var err glib.Error
	ret0 := C.g_file_append_to(file.native(), C.GFileCreateFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileOutputStream{}, err.GoValue()
	}
	return wrapFileOutputStream(ret0), nil
}

// AppendToAsync is a wrapper around g_file_append_to_async().
func (file File) AppendToAsync(flags FileCreateFlags, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_append_to_async(file.native(), C.GFileCreateFlags(flags), C.int(io_priority), cancellable.native(), callback0)
}

// AppendToFinish is a wrapper around g_file_append_to_finish().
func (file File) AppendToFinish(res AsyncResult) (FileOutputStream, error) {
	var err glib.Error
	ret0 := C.g_file_append_to_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileOutputStream{}, err.GoValue()
	}
	return wrapFileOutputStream(ret0), nil
}

// Copy is a wrapper around g_file_copy().
func (source File) Copy(destination File, flags FileCopyFlags, cancellable Cancellable, progress_callback FileProgressCallback) (bool, error) {
	progress_callback0 := (*C.GClosure)(gobject.ClosureNew(progress_callback).Ptr) /*gir:GObject*/
	var err glib.Error
	ret0 := C._g_file_copy(source.native(), destination.native(), C.GFileCopyFlags(flags), cancellable.native(), progress_callback0, (**C.GError)(unsafe.Pointer(&err)))
	C.g_closure_unref(progress_callback0)
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// CopyAsync is a wrapper around g_file_copy_async().
func (source File) CopyAsync(destination File, flags FileCopyFlags, io_priority int, cancellable Cancellable, progress_callback FileProgressCallback, callback AsyncReadyCallback) {
	progress_callback0 := (*C.GClosure)(gobject.ClosureNew(progress_callback).Ptr) /*gir:GObject*/
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr)                   /*gir:GObject*/
	C._g_file_copy_async(source.native(), destination.native(), C.GFileCopyFlags(flags), C.int(io_priority), cancellable.native(), progress_callback0, callback0)
	C.g_closure_unref(progress_callback0)
}

// CopyAttributes is a wrapper around g_file_copy_attributes().
func (source File) CopyAttributes(destination File, flags FileCopyFlags, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_copy_attributes(source.native(), destination.native(), C.GFileCopyFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// CopyFinish is a wrapper around g_file_copy_finish().
func (file File) CopyFinish(res AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_copy_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Create is a wrapper around g_file_create().
func (file File) Create(flags FileCreateFlags, cancellable Cancellable) (FileOutputStream, error) {
	var err glib.Error
	ret0 := C.g_file_create(file.native(), C.GFileCreateFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileOutputStream{}, err.GoValue()
	}
	return wrapFileOutputStream(ret0), nil
}

// CreateAsync is a wrapper around g_file_create_async().
func (file File) CreateAsync(flags FileCreateFlags, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_create_async(file.native(), C.GFileCreateFlags(flags), C.int(io_priority), cancellable.native(), callback0)
}

// CreateFinish is a wrapper around g_file_create_finish().
func (file File) CreateFinish(res AsyncResult) (FileOutputStream, error) {
	var err glib.Error
	ret0 := C.g_file_create_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileOutputStream{}, err.GoValue()
	}
	return wrapFileOutputStream(ret0), nil
}

// CreateReadwrite is a wrapper around g_file_create_readwrite().
func (file File) CreateReadwrite(flags FileCreateFlags, cancellable Cancellable) (FileIOStream, error) {
	var err glib.Error
	ret0 := C.g_file_create_readwrite(file.native(), C.GFileCreateFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileIOStream{}, err.GoValue()
	}
	return wrapFileIOStream(ret0), nil
}

// CreateReadwriteAsync is a wrapper around g_file_create_readwrite_async().
func (file File) CreateReadwriteAsync(flags FileCreateFlags, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_create_readwrite_async(file.native(), C.GFileCreateFlags(flags), C.int(io_priority), cancellable.native(), callback0)
}

// CreateReadwriteFinish is a wrapper around g_file_create_readwrite_finish().
func (file File) CreateReadwriteFinish(res AsyncResult) (FileIOStream, error) {
	var err glib.Error
	ret0 := C.g_file_create_readwrite_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileIOStream{}, err.GoValue()
	}
	return wrapFileIOStream(ret0), nil
}

// Delete is a wrapper around g_file_delete().
func (file File) Delete(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_delete(file.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// DeleteAsync is a wrapper around g_file_delete_async().
func (file File) DeleteAsync(io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_delete_async(file.native(), C.int(io_priority), cancellable.native(), callback0)
}

// DeleteFinish is a wrapper around g_file_delete_finish().
func (file File) DeleteFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_delete_finish(file.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Dup is a wrapper around g_file_dup().
func (file File) Dup() File {
	ret0 := C.g_file_dup(file.native())
	return wrapFile(ret0)
}

// EjectMountableWithOperation is a wrapper around g_file_eject_mountable_with_operation().
func (file File) EjectMountableWithOperation(flags MountUnmountFlags, mount_operation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_eject_mountable_with_operation(file.native(), C.GMountUnmountFlags(flags), mount_operation.native(), cancellable.native(), callback0)
}

// EjectMountableWithOperationFinish is a wrapper around g_file_eject_mountable_with_operation_finish().
func (file File) EjectMountableWithOperationFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_eject_mountable_with_operation_finish(file.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// EnumerateChildren is a wrapper around g_file_enumerate_children().
func (file File) EnumerateChildren(attributes string, flags FileQueryInfoFlags, cancellable Cancellable) (FileEnumerator, error) {
	attributes0 := C.CString(attributes)
	var err glib.Error
	ret0 := C.g_file_enumerate_children(file.native(), attributes0, C.GFileQueryInfoFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return FileEnumerator{}, err.GoValue()
	}
	return wrapFileEnumerator(ret0), nil
}

// EnumerateChildrenAsync is a wrapper around g_file_enumerate_children_async().
func (file File) EnumerateChildrenAsync(attributes string, flags FileQueryInfoFlags, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	attributes0 := C.CString(attributes)
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_enumerate_children_async(file.native(), attributes0, C.GFileQueryInfoFlags(flags), C.int(io_priority), cancellable.native(), callback0)
	C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/
}

// EnumerateChildrenFinish is a wrapper around g_file_enumerate_children_finish().
func (file File) EnumerateChildrenFinish(res AsyncResult) (FileEnumerator, error) {
	var err glib.Error
	ret0 := C.g_file_enumerate_children_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileEnumerator{}, err.GoValue()
	}
	return wrapFileEnumerator(ret0), nil
}

// Equal is a wrapper around g_file_equal().
func (file1 File) Equal(file2 File) bool {
	ret0 := C.g_file_equal(file1.native(), file2.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// FindEnclosingMount is a wrapper around g_file_find_enclosing_mount().
func (file File) FindEnclosingMount(cancellable Cancellable) (Mount, error) {
	var err glib.Error
	ret0 := C.g_file_find_enclosing_mount(file.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return Mount{}, err.GoValue()
	}
	return wrapMount(ret0), nil
}

// FindEnclosingMountAsync is a wrapper around g_file_find_enclosing_mount_async().
func (file File) FindEnclosingMountAsync(io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_find_enclosing_mount_async(file.native(), C.int(io_priority), cancellable.native(), callback0)
}

// FindEnclosingMountFinish is a wrapper around g_file_find_enclosing_mount_finish().
func (file File) FindEnclosingMountFinish(res AsyncResult) (Mount, error) {
	var err glib.Error
	ret0 := C.g_file_find_enclosing_mount_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return Mount{}, err.GoValue()
	}
	return wrapMount(ret0), nil
}

// GetBasename is a wrapper around g_file_get_basename().
func (file File) GetBasename() string {
	ret0 := C.g_file_get_basename(file.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetChild is a wrapper around g_file_get_child().
func (file File) GetChild(name string) File {
	name0 := C.CString(name)
	ret0 := C.g_file_get_child(file.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapFile(ret0)
}

// GetChildForDisplayName is a wrapper around g_file_get_child_for_display_name().
func (file File) GetChildForDisplayName(display_name string) (File, error) {
	display_name0 := C.CString(display_name)
	var err glib.Error
	ret0 := C.g_file_get_child_for_display_name(file.native(), display_name0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(display_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return File{}, err.GoValue()
	}
	return wrapFile(ret0), nil
}

// GetParent is a wrapper around g_file_get_parent().
func (file File) GetParent() File {
	ret0 := C.g_file_get_parent(file.native())
	return wrapFile(ret0)
}

// GetParseName is a wrapper around g_file_get_parse_name().
func (file File) GetParseName() string {
	ret0 := C.g_file_get_parse_name(file.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetPath is a wrapper around g_file_get_path().
func (file File) GetPath() string {
	ret0 := C.g_file_get_path(file.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetRelativePath is a wrapper around g_file_get_relative_path().
func (parent File) GetRelativePath(descendant File) string {
	ret0 := C.g_file_get_relative_path(parent.native(), descendant.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetUri is a wrapper around g_file_get_uri().
func (file File) GetUri() string {
	ret0 := C.g_file_get_uri(file.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetUriScheme is a wrapper around g_file_get_uri_scheme().
func (file File) GetUriScheme() string {
	ret0 := C.g_file_get_uri_scheme(file.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// HasParent is a wrapper around g_file_has_parent().
func (file File) HasParent(parent File) bool {
	ret0 := C.g_file_has_parent(file.native(), parent.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// HasPrefix is a wrapper around g_file_has_prefix().
func (file File) HasPrefix(prefix File) bool {
	ret0 := C.g_file_has_prefix(file.native(), prefix.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// HasUriScheme is a wrapper around g_file_has_uri_scheme().
func (file File) HasUriScheme(uri_scheme string) bool {
	uri_scheme0 := C.CString(uri_scheme)
	ret0 := C.g_file_has_uri_scheme(file.native(), uri_scheme0)
	C.free(unsafe.Pointer(uri_scheme0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))     /*go:.util*/
}

// Hash is a wrapper around g_file_hash().
func (file File) Hash() uint {
	ret0 := C.g_file_hash(C.gconstpointer(file.native()))
	return uint(ret0)
}

// IsNative is a wrapper around g_file_is_native().
func (file File) IsNative() bool {
	ret0 := C.g_file_is_native(file.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// MakeDirectory is a wrapper around g_file_make_directory().
func (file File) MakeDirectory(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_make_directory(file.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// MakeDirectoryAsync is a wrapper around g_file_make_directory_async().
func (file File) MakeDirectoryAsync(io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_make_directory_async(file.native(), C.int(io_priority), cancellable.native(), callback0)
}

// MakeDirectoryFinish is a wrapper around g_file_make_directory_finish().
func (file File) MakeDirectoryFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_make_directory_finish(file.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// MakeDirectoryWithParents is a wrapper around g_file_make_directory_with_parents().
func (file File) MakeDirectoryWithParents(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_make_directory_with_parents(file.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// MakeSymbolicLink is a wrapper around g_file_make_symbolic_link().
func (file File) MakeSymbolicLink(symlink_value string, cancellable Cancellable) (bool, error) {
	symlink_value0 := C.CString(symlink_value)
	var err glib.Error
	ret0 := C.g_file_make_symbolic_link(file.native(), symlink_value0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(symlink_value0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// MeasureDiskUsage is a wrapper around g_file_measure_disk_usage().
func (file File) MeasureDiskUsage(flags FileMeasureFlags, cancellable Cancellable, progress_callback FileMeasureProgressCallback) (bool, uint64, uint64, uint64, error) {
	progress_callback0 := (*C.GClosure)(gobject.ClosureNew(progress_callback).Ptr) /*gir:GObject*/
	var disk_usage0 C.guint64
	var num_dirs0 C.guint64
	var num_files0 C.guint64
	var err glib.Error
	ret0 := C._g_file_measure_disk_usage(file.native(), C.GFileMeasureFlags(flags), cancellable.native(), progress_callback0, &disk_usage0, &num_dirs0, &num_files0, (**C.GError)(unsafe.Pointer(&err)))
	C.g_closure_unref(progress_callback0)
	if err.Ptr != nil {
		defer err.Free()
		return false, 0, 0, 0, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, uint64(disk_usage0), uint64(num_dirs0), uint64(num_files0), nil
}

// MeasureDiskUsageAsync is a wrapper around g_file_measure_disk_usage_async().
func (file File) MeasureDiskUsageAsync(flags FileMeasureFlags, io_priority int, cancellable Cancellable, progress_callback FileMeasureProgressCallback, callback AsyncReadyCallback) {
	progress_callback0 := (*C.GClosure)(gobject.ClosureNew(progress_callback).Ptr) /*gir:GObject*/
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr)                   /*gir:GObject*/
	C._g_file_measure_disk_usage_async(file.native(), C.GFileMeasureFlags(flags), C.gint(io_priority), cancellable.native(), progress_callback0, callback0)
	C.g_closure_unref(progress_callback0)
}

// MeasureDiskUsageFinish is a wrapper around g_file_measure_disk_usage_finish().
func (file File) MeasureDiskUsageFinish(result AsyncResult) (bool, uint64, uint64, uint64, error) {
	var disk_usage0 C.guint64
	var num_dirs0 C.guint64
	var num_files0 C.guint64
	var err glib.Error
	ret0 := C.g_file_measure_disk_usage_finish(file.native(), result.native(), &disk_usage0, &num_dirs0, &num_files0, (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, 0, 0, 0, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, uint64(disk_usage0), uint64(num_dirs0), uint64(num_files0), nil
}

// Monitor is a wrapper around g_file_monitor().
func (file File) Monitor(flags FileMonitorFlags, cancellable Cancellable) (FileMonitor, error) {
	var err glib.Error
	ret0 := C.g_file_monitor(file.native(), C.GFileMonitorFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileMonitor{}, err.GoValue()
	}
	return wrapFileMonitor(ret0), nil
}

// MonitorDirectory is a wrapper around g_file_monitor_directory().
func (file File) MonitorDirectory(flags FileMonitorFlags, cancellable Cancellable) (FileMonitor, error) {
	var err glib.Error
	ret0 := C.g_file_monitor_directory(file.native(), C.GFileMonitorFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileMonitor{}, err.GoValue()
	}
	return wrapFileMonitor(ret0), nil
}

// MonitorFile is a wrapper around g_file_monitor_file().
func (file File) MonitorFile(flags FileMonitorFlags, cancellable Cancellable) (FileMonitor, error) {
	var err glib.Error
	ret0 := C.g_file_monitor_file(file.native(), C.GFileMonitorFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileMonitor{}, err.GoValue()
	}
	return wrapFileMonitor(ret0), nil
}

// MountEnclosingVolume is a wrapper around g_file_mount_enclosing_volume().
func (location File) MountEnclosingVolume(flags MountMountFlags, mount_operation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_mount_enclosing_volume(location.native(), C.GMountMountFlags(flags), mount_operation.native(), cancellable.native(), callback0)
}

// MountEnclosingVolumeFinish is a wrapper around g_file_mount_enclosing_volume_finish().
func (location File) MountEnclosingVolumeFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_mount_enclosing_volume_finish(location.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// MountMountable is a wrapper around g_file_mount_mountable().
func (file File) MountMountable(flags MountMountFlags, mount_operation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_mount_mountable(file.native(), C.GMountMountFlags(flags), mount_operation.native(), cancellable.native(), callback0)
}

// MountMountableFinish is a wrapper around g_file_mount_mountable_finish().
func (file File) MountMountableFinish(result AsyncResult) (File, error) {
	var err glib.Error
	ret0 := C.g_file_mount_mountable_finish(file.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return File{}, err.GoValue()
	}
	return wrapFile(ret0), nil
}

// Move is a wrapper around g_file_move().
func (source File) Move(destination File, flags FileCopyFlags, cancellable Cancellable, progress_callback FileProgressCallback) (bool, error) {
	progress_callback0 := (*C.GClosure)(gobject.ClosureNew(progress_callback).Ptr) /*gir:GObject*/
	var err glib.Error
	ret0 := C._g_file_move(source.native(), destination.native(), C.GFileCopyFlags(flags), cancellable.native(), progress_callback0, (**C.GError)(unsafe.Pointer(&err)))
	C.g_closure_unref(progress_callback0)
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// OpenReadwrite is a wrapper around g_file_open_readwrite().
func (file File) OpenReadwrite(cancellable Cancellable) (FileIOStream, error) {
	var err glib.Error
	ret0 := C.g_file_open_readwrite(file.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileIOStream{}, err.GoValue()
	}
	return wrapFileIOStream(ret0), nil
}

// OpenReadwriteAsync is a wrapper around g_file_open_readwrite_async().
func (file File) OpenReadwriteAsync(io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_open_readwrite_async(file.native(), C.int(io_priority), cancellable.native(), callback0)
}

// OpenReadwriteFinish is a wrapper around g_file_open_readwrite_finish().
func (file File) OpenReadwriteFinish(res AsyncResult) (FileIOStream, error) {
	var err glib.Error
	ret0 := C.g_file_open_readwrite_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileIOStream{}, err.GoValue()
	}
	return wrapFileIOStream(ret0), nil
}

// PollMountable is a wrapper around g_file_poll_mountable().
func (file File) PollMountable(cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_poll_mountable(file.native(), cancellable.native(), callback0)
}

// PollMountableFinish is a wrapper around g_file_poll_mountable_finish().
func (file File) PollMountableFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_poll_mountable_finish(file.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// QueryDefaultHandler is a wrapper around g_file_query_default_handler().
func (file File) QueryDefaultHandler(cancellable Cancellable) (AppInfo, error) {
	var err glib.Error
	ret0 := C.g_file_query_default_handler(file.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return AppInfo{}, err.GoValue()
	}
	return wrapAppInfo(ret0), nil
}

// QueryExists is a wrapper around g_file_query_exists().
func (file File) QueryExists(cancellable Cancellable) bool {
	ret0 := C.g_file_query_exists(file.native(), cancellable.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// QueryFileType is a wrapper around g_file_query_file_type().
func (file File) QueryFileType(flags FileQueryInfoFlags, cancellable Cancellable) FileType {
	ret0 := C.g_file_query_file_type(file.native(), C.GFileQueryInfoFlags(flags), cancellable.native())
	return FileType(ret0)
}

// QueryFilesystemInfo is a wrapper around g_file_query_filesystem_info().
func (file File) QueryFilesystemInfo(attributes string, cancellable Cancellable) (FileInfo, error) {
	attributes0 := C.CString(attributes)
	var err glib.Error
	ret0 := C.g_file_query_filesystem_info(file.native(), attributes0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// QueryFilesystemInfoAsync is a wrapper around g_file_query_filesystem_info_async().
func (file File) QueryFilesystemInfoAsync(attributes string, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	attributes0 := C.CString(attributes)
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_query_filesystem_info_async(file.native(), attributes0, C.int(io_priority), cancellable.native(), callback0)
	C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/
}

// QueryFilesystemInfoFinish is a wrapper around g_file_query_filesystem_info_finish().
func (file File) QueryFilesystemInfoFinish(res AsyncResult) (FileInfo, error) {
	var err glib.Error
	ret0 := C.g_file_query_filesystem_info_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// QueryInfo is a wrapper around g_file_query_info().
func (file File) QueryInfo(attributes string, flags FileQueryInfoFlags, cancellable Cancellable) (FileInfo, error) {
	attributes0 := C.CString(attributes)
	var err glib.Error
	ret0 := C.g_file_query_info(file.native(), attributes0, C.GFileQueryInfoFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// QueryInfoAsync is a wrapper around g_file_query_info_async().
func (file File) QueryInfoAsync(attributes string, flags FileQueryInfoFlags, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	attributes0 := C.CString(attributes)
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_query_info_async(file.native(), attributes0, C.GFileQueryInfoFlags(flags), C.int(io_priority), cancellable.native(), callback0)
	C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/
}

// QueryInfoFinish is a wrapper around g_file_query_info_finish().
func (file File) QueryInfoFinish(res AsyncResult) (FileInfo, error) {
	var err glib.Error
	ret0 := C.g_file_query_info_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// QuerySettableAttributes is a wrapper around g_file_query_settable_attributes().
func (file File) QuerySettableAttributes(cancellable Cancellable) (FileAttributeInfoList, error) {
	var err glib.Error
	ret0 := C.g_file_query_settable_attributes(file.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileAttributeInfoList{}, err.GoValue()
	}
	return wrapFileAttributeInfoList(ret0), nil
}

// QueryWritableNamespaces is a wrapper around g_file_query_writable_namespaces().
func (file File) QueryWritableNamespaces(cancellable Cancellable) (FileAttributeInfoList, error) {
	var err glib.Error
	ret0 := C.g_file_query_writable_namespaces(file.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileAttributeInfoList{}, err.GoValue()
	}
	return wrapFileAttributeInfoList(ret0), nil
}

// Read is a wrapper around g_file_read().
func (file File) Read(cancellable Cancellable) (FileInputStream, error) {
	var err glib.Error
	ret0 := C.g_file_read(file.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileInputStream{}, err.GoValue()
	}
	return wrapFileInputStream(ret0), nil
}

// ReadAsync is a wrapper around g_file_read_async().
func (file File) ReadAsync(io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_read_async(file.native(), C.int(io_priority), cancellable.native(), callback0)
}

// ReadFinish is a wrapper around g_file_read_finish().
func (file File) ReadFinish(res AsyncResult) (FileInputStream, error) {
	var err glib.Error
	ret0 := C.g_file_read_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileInputStream{}, err.GoValue()
	}
	return wrapFileInputStream(ret0), nil
}

// Replace is a wrapper around g_file_replace().
func (file File) Replace(etag string, make_backup bool, flags FileCreateFlags, cancellable Cancellable) (FileOutputStream, error) {
	etag0 := C.CString(etag)
	var err glib.Error
	ret0 := C.g_file_replace(file.native(), etag0, C.gboolean(util.Bool2Int(make_backup)) /*go:.util*/, C.GFileCreateFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(etag0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return FileOutputStream{}, err.GoValue()
	}
	return wrapFileOutputStream(ret0), nil
}

// ReplaceAsync is a wrapper around g_file_replace_async().
func (file File) ReplaceAsync(etag string, make_backup bool, flags FileCreateFlags, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	etag0 := C.CString(etag)
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_replace_async(file.native(), etag0, C.gboolean(util.Bool2Int(make_backup)) /*go:.util*/, C.GFileCreateFlags(flags), C.int(io_priority), cancellable.native(), callback0)
	C.free(unsafe.Pointer(etag0)) /*ch:<stdlib.h>*/
}

// ReplaceFinish is a wrapper around g_file_replace_finish().
func (file File) ReplaceFinish(res AsyncResult) (FileOutputStream, error) {
	var err glib.Error
	ret0 := C.g_file_replace_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileOutputStream{}, err.GoValue()
	}
	return wrapFileOutputStream(ret0), nil
}

// ReplaceReadwrite is a wrapper around g_file_replace_readwrite().
func (file File) ReplaceReadwrite(etag string, make_backup bool, flags FileCreateFlags, cancellable Cancellable) (FileIOStream, error) {
	etag0 := C.CString(etag)
	var err glib.Error
	ret0 := C.g_file_replace_readwrite(file.native(), etag0, C.gboolean(util.Bool2Int(make_backup)) /*go:.util*/, C.GFileCreateFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(etag0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return FileIOStream{}, err.GoValue()
	}
	return wrapFileIOStream(ret0), nil
}

// ReplaceReadwriteFinish is a wrapper around g_file_replace_readwrite_finish().
func (file File) ReplaceReadwriteFinish(res AsyncResult) (FileIOStream, error) {
	var err glib.Error
	ret0 := C.g_file_replace_readwrite_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileIOStream{}, err.GoValue()
	}
	return wrapFileIOStream(ret0), nil
}

// ResolveRelativePath is a wrapper around g_file_resolve_relative_path().
func (file File) ResolveRelativePath(relative_path string) File {
	relative_path0 := C.CString(relative_path)
	ret0 := C.g_file_resolve_relative_path(file.native(), relative_path0)
	C.free(unsafe.Pointer(relative_path0)) /*ch:<stdlib.h>*/
	return wrapFile(ret0)
}

// SetAttribute is a wrapper around g_file_set_attribute().
func (file File) SetAttribute(attribute string, type_ FileAttributeType, value_p unsafe.Pointer, flags FileQueryInfoFlags, cancellable Cancellable) (bool, error) {
	attribute0 := C.CString(attribute)
	var err glib.Error
	ret0 := C.g_file_set_attribute(file.native(), attribute0, C.GFileAttributeType(type_), C.gpointer(value_p), C.GFileQueryInfoFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetAttributeByteString is a wrapper around g_file_set_attribute_byte_string().
func (file File) SetAttributeByteString(attribute string, value string, flags FileQueryInfoFlags, cancellable Cancellable) (bool, error) {
	attribute0 := C.CString(attribute)
	value0 := C.CString(value)
	var err glib.Error
	ret0 := C.g_file_set_attribute_byte_string(file.native(), attribute0, value0, C.GFileQueryInfoFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(value0))     /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetAttributeInt32 is a wrapper around g_file_set_attribute_int32().
func (file File) SetAttributeInt32(attribute string, value int32, flags FileQueryInfoFlags, cancellable Cancellable) (bool, error) {
	attribute0 := C.CString(attribute)
	var err glib.Error
	ret0 := C.g_file_set_attribute_int32(file.native(), attribute0, C.gint32(value), C.GFileQueryInfoFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetAttributeInt64 is a wrapper around g_file_set_attribute_int64().
func (file File) SetAttributeInt64(attribute string, value int64, flags FileQueryInfoFlags, cancellable Cancellable) (bool, error) {
	attribute0 := C.CString(attribute)
	var err glib.Error
	ret0 := C.g_file_set_attribute_int64(file.native(), attribute0, C.gint64(value), C.GFileQueryInfoFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetAttributeString is a wrapper around g_file_set_attribute_string().
func (file File) SetAttributeString(attribute string, value string, flags FileQueryInfoFlags, cancellable Cancellable) (bool, error) {
	attribute0 := C.CString(attribute)
	value0 := C.CString(value)
	var err glib.Error
	ret0 := C.g_file_set_attribute_string(file.native(), attribute0, value0, C.GFileQueryInfoFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(value0))     /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetAttributeUint32 is a wrapper around g_file_set_attribute_uint32().
func (file File) SetAttributeUint32(attribute string, value uint32, flags FileQueryInfoFlags, cancellable Cancellable) (bool, error) {
	attribute0 := C.CString(attribute)
	var err glib.Error
	ret0 := C.g_file_set_attribute_uint32(file.native(), attribute0, C.guint32(value), C.GFileQueryInfoFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetAttributeUint64 is a wrapper around g_file_set_attribute_uint64().
func (file File) SetAttributeUint64(attribute string, value uint64, flags FileQueryInfoFlags, cancellable Cancellable) (bool, error) {
	attribute0 := C.CString(attribute)
	var err glib.Error
	ret0 := C.g_file_set_attribute_uint64(file.native(), attribute0, C.guint64(value), C.GFileQueryInfoFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetAttributesAsync is a wrapper around g_file_set_attributes_async().
func (file File) SetAttributesAsync(info FileInfo, flags FileQueryInfoFlags, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_set_attributes_async(file.native(), info.native(), C.GFileQueryInfoFlags(flags), C.int(io_priority), cancellable.native(), callback0)
}

// SetAttributesFinish is a wrapper around g_file_set_attributes_finish().
func (file File) SetAttributesFinish(result AsyncResult) (bool, FileInfo, error) {
	var info0 *C.GFileInfo
	var err glib.Error
	ret0 := C.g_file_set_attributes_finish(file.native(), result.native(), &info0, (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, FileInfo{}, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, wrapFileInfo(info0), nil
}

// SetAttributesFromInfo is a wrapper around g_file_set_attributes_from_info().
func (file File) SetAttributesFromInfo(info FileInfo, flags FileQueryInfoFlags, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_set_attributes_from_info(file.native(), info.native(), C.GFileQueryInfoFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetDisplayName is a wrapper around g_file_set_display_name().
func (file File) SetDisplayName(display_name string, cancellable Cancellable) (File, error) {
	display_name0 := C.CString(display_name)
	var err glib.Error
	ret0 := C.g_file_set_display_name(file.native(), display_name0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(display_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return File{}, err.GoValue()
	}
	return wrapFile(ret0), nil
}

// SetDisplayNameAsync is a wrapper around g_file_set_display_name_async().
func (file File) SetDisplayNameAsync(display_name string, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	display_name0 := C.CString(display_name)
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_set_display_name_async(file.native(), display_name0, C.int(io_priority), cancellable.native(), callback0)
	C.free(unsafe.Pointer(display_name0)) /*ch:<stdlib.h>*/
}

// SetDisplayNameFinish is a wrapper around g_file_set_display_name_finish().
func (file File) SetDisplayNameFinish(res AsyncResult) (File, error) {
	var err glib.Error
	ret0 := C.g_file_set_display_name_finish(file.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return File{}, err.GoValue()
	}
	return wrapFile(ret0), nil
}

// StartMountable is a wrapper around g_file_start_mountable().
func (file File) StartMountable(flags DriveStartFlags, start_operation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_start_mountable(file.native(), C.GDriveStartFlags(flags), start_operation.native(), cancellable.native(), callback0)
}

// StartMountableFinish is a wrapper around g_file_start_mountable_finish().
func (file File) StartMountableFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_start_mountable_finish(file.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// StopMountable is a wrapper around g_file_stop_mountable().
func (file File) StopMountable(flags MountUnmountFlags, mount_operation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_stop_mountable(file.native(), C.GMountUnmountFlags(flags), mount_operation.native(), cancellable.native(), callback0)
}

// StopMountableFinish is a wrapper around g_file_stop_mountable_finish().
func (file File) StopMountableFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_stop_mountable_finish(file.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SupportsThreadContexts is a wrapper around g_file_supports_thread_contexts().
func (file File) SupportsThreadContexts() bool {
	ret0 := C.g_file_supports_thread_contexts(file.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Trash is a wrapper around g_file_trash().
func (file File) Trash(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_trash(file.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// TrashAsync is a wrapper around g_file_trash_async().
func (file File) TrashAsync(io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_trash_async(file.native(), C.int(io_priority), cancellable.native(), callback0)
}

// TrashFinish is a wrapper around g_file_trash_finish().
func (file File) TrashFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_trash_finish(file.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// UnmountMountableWithOperation is a wrapper around g_file_unmount_mountable_with_operation().
func (file File) UnmountMountableWithOperation(flags MountUnmountFlags, mount_operation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_unmount_mountable_with_operation(file.native(), C.GMountUnmountFlags(flags), mount_operation.native(), cancellable.native(), callback0)
}

// UnmountMountableWithOperationFinish is a wrapper around g_file_unmount_mountable_with_operation_finish().
func (file File) UnmountMountableWithOperationFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_unmount_mountable_with_operation_finish(file.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// FileNewForCommandlineArg is a wrapper around g_file_new_for_commandline_arg().
func FileNewForCommandlineArg(arg string) File {
	arg0 := C.CString(arg)
	ret0 := C.g_file_new_for_commandline_arg(arg0)
	C.free(unsafe.Pointer(arg0)) /*ch:<stdlib.h>*/
	return wrapFile(ret0)
}

// FileNewForCommandlineArgAndCwd is a wrapper around g_file_new_for_commandline_arg_and_cwd().
func FileNewForCommandlineArgAndCwd(arg string, cwd string) File {
	arg0 := (*C.gchar)(C.CString(arg))
	cwd0 := (*C.gchar)(C.CString(cwd))
	ret0 := C.g_file_new_for_commandline_arg_and_cwd(arg0, cwd0)
	C.free(unsafe.Pointer(arg0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(cwd0)) /*ch:<stdlib.h>*/
	return wrapFile(ret0)
}

// FileNewForPath is a wrapper around g_file_new_for_path().
func FileNewForPath(path string) File {
	path0 := C.CString(path)
	ret0 := C.g_file_new_for_path(path0)
	C.free(unsafe.Pointer(path0)) /*ch:<stdlib.h>*/
	return wrapFile(ret0)
}

// FileNewForUri is a wrapper around g_file_new_for_uri().
func FileNewForUri(uri string) File {
	uri0 := C.CString(uri)
	ret0 := C.g_file_new_for_uri(uri0)
	C.free(unsafe.Pointer(uri0)) /*ch:<stdlib.h>*/
	return wrapFile(ret0)
}

// FileNewTmp is a wrapper around g_file_new_tmp().
func FileNewTmp(tmpl string) (File, FileIOStream, error) {
	tmpl0 := C.CString(tmpl)
	var iostream0 *C.GFileIOStream
	var err glib.Error
	ret0 := C.g_file_new_tmp(tmpl0, &iostream0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(tmpl0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return File{}, FileIOStream{}, err.GoValue()
	}
	return wrapFile(ret0), wrapFileIOStream(iostream0), nil
}

// FileParseName is a wrapper around g_file_parse_name().
func FileParseName(parse_name string) File {
	parse_name0 := C.CString(parse_name)
	ret0 := C.g_file_parse_name(parse_name0)
	C.free(unsafe.Pointer(parse_name0)) /*ch:<stdlib.h>*/
	return wrapFile(ret0)
}

// Object Cancellable
type Cancellable struct {
	gobject.Object
}

func (v Cancellable) native() *C.GCancellable {
	return (*C.GCancellable)(v.Ptr)
}
func wrapCancellable(p *C.GCancellable) (v Cancellable) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCancellable(p unsafe.Pointer) (v Cancellable) {
	v.Ptr = p
	return
}
func (v Cancellable) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCancellable(p unsafe.Pointer) interface{} {
	return WrapCancellable(p)
}
func (v Cancellable) GetType() gobject.Type {
	return gobject.Type(C.g_cancellable_get_type())
}
func (v Cancellable) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCancellable(unsafe.Pointer(ptr)), nil
	}
}

// CancellableNew is a wrapper around g_cancellable_new().
func CancellableNew() Cancellable {
	ret0 := C.g_cancellable_new()
	return wrapCancellable(ret0)
}

// Cancel is a wrapper around g_cancellable_cancel().
func (cancellable Cancellable) Cancel() {
	C.g_cancellable_cancel(cancellable.native())
}

// Disconnect is a wrapper around g_cancellable_disconnect().
func (cancellable Cancellable) Disconnect(handler_id uint) {
	C.g_cancellable_disconnect(cancellable.native(), C.gulong(handler_id))
}

// GetFd is a wrapper around g_cancellable_get_fd().
func (cancellable Cancellable) GetFd() int {
	ret0 := C.g_cancellable_get_fd(cancellable.native())
	return int(ret0)
}

// IsCancelled is a wrapper around g_cancellable_is_cancelled().
func (cancellable Cancellable) IsCancelled() bool {
	ret0 := C.g_cancellable_is_cancelled(cancellable.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// MakePollfd is a wrapper around g_cancellable_make_pollfd().
func (cancellable Cancellable) MakePollfd(pollfd glib.PollFD) bool {
	ret0 := C.g_cancellable_make_pollfd(cancellable.native(), (*C.GPollFD)(pollfd.Ptr))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// PopCurrent is a wrapper around g_cancellable_pop_current().
func (cancellable Cancellable) PopCurrent() {
	C.g_cancellable_pop_current(cancellable.native())
}

// PushCurrent is a wrapper around g_cancellable_push_current().
func (cancellable Cancellable) PushCurrent() {
	C.g_cancellable_push_current(cancellable.native())
}

// ReleaseFd is a wrapper around g_cancellable_release_fd().
func (cancellable Cancellable) ReleaseFd() {
	C.g_cancellable_release_fd(cancellable.native())
}

// Reset is a wrapper around g_cancellable_reset().
func (cancellable Cancellable) Reset() {
	C.g_cancellable_reset(cancellable.native())
}

// SetErrorIfCancelled is a wrapper around g_cancellable_set_error_if_cancelled().
func (cancellable Cancellable) SetErrorIfCancelled() (bool, error) {
	var err glib.Error
	ret0 := C.g_cancellable_set_error_if_cancelled(cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SourceNew is a wrapper around g_cancellable_source_new().
func (cancellable Cancellable) SourceNew() glib.Source {
	ret0 := C.g_cancellable_source_new(cancellable.native())
	return glib.WrapSource(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// CancellableGetCurrent is a wrapper around g_cancellable_get_current().
func CancellableGetCurrent() Cancellable {
	ret0 := C.g_cancellable_get_current()
	return wrapCancellable(ret0)
}

// Object OutputStream
type OutputStream struct {
	gobject.Object
}

func (v OutputStream) native() *C.GOutputStream {
	return (*C.GOutputStream)(v.Ptr)
}
func wrapOutputStream(p *C.GOutputStream) (v OutputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapOutputStream(p unsafe.Pointer) (v OutputStream) {
	v.Ptr = p
	return
}
func (v OutputStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapOutputStream(p unsafe.Pointer) interface{} {
	return WrapOutputStream(p)
}
func (v OutputStream) GetType() gobject.Type {
	return gobject.Type(C.g_output_stream_get_type())
}
func (v OutputStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapOutputStream(unsafe.Pointer(ptr)), nil
	}
}

// ClearPending is a wrapper around g_output_stream_clear_pending().
func (stream OutputStream) ClearPending() {
	C.g_output_stream_clear_pending(stream.native())
}

// Close is a wrapper around g_output_stream_close().
func (stream OutputStream) Close(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_output_stream_close(stream.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// CloseAsync is a wrapper around g_output_stream_close_async().
func (stream OutputStream) CloseAsync(io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_output_stream_close_async(stream.native(), C.int(io_priority), cancellable.native(), callback0)
}

// CloseFinish is a wrapper around g_output_stream_close_finish().
func (stream OutputStream) CloseFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_output_stream_close_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Flush is a wrapper around g_output_stream_flush().
func (stream OutputStream) Flush(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_output_stream_flush(stream.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// FlushAsync is a wrapper around g_output_stream_flush_async().
func (stream OutputStream) FlushAsync(io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_output_stream_flush_async(stream.native(), C.int(io_priority), cancellable.native(), callback0)
}

// FlushFinish is a wrapper around g_output_stream_flush_finish().
func (stream OutputStream) FlushFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_output_stream_flush_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// HasPending is a wrapper around g_output_stream_has_pending().
func (stream OutputStream) HasPending() bool {
	ret0 := C.g_output_stream_has_pending(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsClosed is a wrapper around g_output_stream_is_closed().
func (stream OutputStream) IsClosed() bool {
	ret0 := C.g_output_stream_is_closed(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsClosing is a wrapper around g_output_stream_is_closing().
func (stream OutputStream) IsClosing() bool {
	ret0 := C.g_output_stream_is_closing(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetPending is a wrapper around g_output_stream_set_pending().
func (stream OutputStream) SetPending() (bool, error) {
	var err glib.Error
	ret0 := C.g_output_stream_set_pending(stream.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Splice is a wrapper around g_output_stream_splice().
func (stream OutputStream) Splice(source InputStream, flags OutputStreamSpliceFlags, cancellable Cancellable) (int, error) {
	var err glib.Error
	ret0 := C.g_output_stream_splice(stream.native(), source.native(), C.GOutputStreamSpliceFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// SpliceAsync is a wrapper around g_output_stream_splice_async().
func (stream OutputStream) SpliceAsync(source InputStream, flags OutputStreamSpliceFlags, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_output_stream_splice_async(stream.native(), source.native(), C.GOutputStreamSpliceFlags(flags), C.int(io_priority), cancellable.native(), callback0)
}

// SpliceFinish is a wrapper around g_output_stream_splice_finish().
func (stream OutputStream) SpliceFinish(result AsyncResult) (int, error) {
	var err glib.Error
	ret0 := C.g_output_stream_splice_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// WriteAllFinish is a wrapper around g_output_stream_write_all_finish().
func (stream OutputStream) WriteAllFinish(result AsyncResult) (bool, uint, error) {
	var bytes_written0 C.gsize
	var err glib.Error
	ret0 := C.g_output_stream_write_all_finish(stream.native(), result.native(), &bytes_written0, (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, 0, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, uint(bytes_written0), nil
}

// WriteBytes is a wrapper around g_output_stream_write_bytes().
func (stream OutputStream) WriteBytes(bytes glib.Bytes, cancellable Cancellable) (int, error) {
	var err glib.Error
	ret0 := C.g_output_stream_write_bytes(stream.native(), (*C.GBytes)(bytes.Ptr), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// WriteBytesAsync is a wrapper around g_output_stream_write_bytes_async().
func (stream OutputStream) WriteBytesAsync(bytes glib.Bytes, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_output_stream_write_bytes_async(stream.native(), (*C.GBytes)(bytes.Ptr), C.int(io_priority), cancellable.native(), callback0)
}

// WriteBytesFinish is a wrapper around g_output_stream_write_bytes_finish().
func (stream OutputStream) WriteBytesFinish(result AsyncResult) (int, error) {
	var err glib.Error
	ret0 := C.g_output_stream_write_bytes_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// WriteFinish is a wrapper around g_output_stream_write_finish().
func (stream OutputStream) WriteFinish(result AsyncResult) (int, error) {
	var err glib.Error
	ret0 := C.g_output_stream_write_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// Object FileOutputStream
type FileOutputStream struct {
	OutputStream
}

func (v FileOutputStream) native() *C.GFileOutputStream {
	return (*C.GFileOutputStream)(v.Ptr)
}
func wrapFileOutputStream(p *C.GFileOutputStream) (v FileOutputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileOutputStream(p unsafe.Pointer) (v FileOutputStream) {
	v.Ptr = p
	return
}
func (v FileOutputStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFileOutputStream(p unsafe.Pointer) interface{} {
	return WrapFileOutputStream(p)
}
func (v FileOutputStream) GetType() gobject.Type {
	return gobject.Type(C.g_file_output_stream_get_type())
}
func (v FileOutputStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFileOutputStream(unsafe.Pointer(ptr)), nil
	}
}
func (v FileOutputStream) Seekable() Seekable {
	return WrapSeekable(v.Ptr)
}

// GetEtag is a wrapper around g_file_output_stream_get_etag().
func (stream FileOutputStream) GetEtag() string {
	ret0 := C.g_file_output_stream_get_etag(stream.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// QueryInfo is a wrapper around g_file_output_stream_query_info().
func (stream FileOutputStream) QueryInfo(attributes string, cancellable Cancellable) (FileInfo, error) {
	attributes0 := C.CString(attributes)
	var err glib.Error
	ret0 := C.g_file_output_stream_query_info(stream.native(), attributes0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// QueryInfoAsync is a wrapper around g_file_output_stream_query_info_async().
func (stream FileOutputStream) QueryInfoAsync(attributes string, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	attributes0 := C.CString(attributes)
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_output_stream_query_info_async(stream.native(), attributes0, C.int(io_priority), cancellable.native(), callback0)
	C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/
}

// QueryInfoFinish is a wrapper around g_file_output_stream_query_info_finish().
func (stream FileOutputStream) QueryInfoFinish(result AsyncResult) (FileInfo, error) {
	var err glib.Error
	ret0 := C.g_file_output_stream_query_info_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// Interface Seekable
type Seekable struct {
	Ptr unsafe.Pointer
}

func (v Seekable) native() *C.GSeekable {
	return (*C.GSeekable)(v.Ptr)
}
func wrapSeekable(p *C.GSeekable) Seekable {
	return Seekable{unsafe.Pointer(p)}
}
func WrapSeekable(p unsafe.Pointer) Seekable {
	return Seekable{p}
}
func (v Seekable) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSeekable(p unsafe.Pointer) interface{} {
	return WrapSeekable(p)
}
func (v Seekable) GetType() gobject.Type {
	return gobject.Type(C.g_seekable_get_type())
}
func (v Seekable) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSeekable(unsafe.Pointer(ptr)), nil
	}
}

// CanSeek is a wrapper around g_seekable_can_seek().
func (seekable Seekable) CanSeek() bool {
	ret0 := C.g_seekable_can_seek(seekable.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// CanTruncate is a wrapper around g_seekable_can_truncate().
func (seekable Seekable) CanTruncate() bool {
	ret0 := C.g_seekable_can_truncate(seekable.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Seek is a wrapper around g_seekable_seek().
func (seekable Seekable) Seek(offset int64, type_ /*gir:GLib*/ glib.SeekType, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_seekable_seek(seekable.native(), C.goffset(offset), C.GSeekType(type_), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Tell is a wrapper around g_seekable_tell().
func (seekable Seekable) Tell() int64 {
	ret0 := C.g_seekable_tell(seekable.native())
	return int64(ret0)
}

// Truncate is a wrapper around g_seekable_truncate().
func (seekable Seekable) Truncate(offset int64, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_seekable_truncate(seekable.native(), C.goffset(offset), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Object FileInfo
type FileInfo struct {
	gobject.Object
}

func (v FileInfo) native() *C.GFileInfo {
	return (*C.GFileInfo)(v.Ptr)
}
func wrapFileInfo(p *C.GFileInfo) (v FileInfo) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileInfo(p unsafe.Pointer) (v FileInfo) {
	v.Ptr = p
	return
}
func (v FileInfo) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFileInfo(p unsafe.Pointer) interface{} {
	return WrapFileInfo(p)
}
func (v FileInfo) GetType() gobject.Type {
	return gobject.Type(C.g_file_info_get_type())
}
func (v FileInfo) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFileInfo(unsafe.Pointer(ptr)), nil
	}
}

// FileInfoNew is a wrapper around g_file_info_new().
func FileInfoNew() FileInfo {
	ret0 := C.g_file_info_new()
	return wrapFileInfo(ret0)
}

// ClearStatus is a wrapper around g_file_info_clear_status().
func (info FileInfo) ClearStatus() {
	C.g_file_info_clear_status(info.native())
}

// CopyInto is a wrapper around g_file_info_copy_into().
func (src_info FileInfo) CopyInto(dest_info FileInfo) {
	C.g_file_info_copy_into(src_info.native(), dest_info.native())
}

// Dup is a wrapper around g_file_info_dup().
func (other FileInfo) Dup() FileInfo {
	ret0 := C.g_file_info_dup(other.native())
	return wrapFileInfo(ret0)
}

// GetAttributeAsString is a wrapper around g_file_info_get_attribute_as_string().
func (info FileInfo) GetAttributeAsString(attribute string) string {
	attribute0 := C.CString(attribute)
	ret0 := C.g_file_info_get_attribute_as_string(info.native(), attribute0)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetAttributeBoolean is a wrapper around g_file_info_get_attribute_boolean().
func (info FileInfo) GetAttributeBoolean(attribute string) bool {
	attribute0 := C.CString(attribute)
	ret0 := C.g_file_info_get_attribute_boolean(info.native(), attribute0)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))    /*go:.util*/
}

// GetAttributeByteString is a wrapper around g_file_info_get_attribute_byte_string().
func (info FileInfo) GetAttributeByteString(attribute string) string {
	attribute0 := C.CString(attribute)
	ret0 := C.g_file_info_get_attribute_byte_string(info.native(), attribute0)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	ret := C.GoString(ret0)
	return ret
}

// GetAttributeData is a wrapper around g_file_info_get_attribute_data().
func (info FileInfo) GetAttributeData(attribute string) (bool, FileAttributeType, unsafe.Pointer, FileAttributeStatus) {
	attribute0 := C.CString(attribute)
	var type0 C.GFileAttributeType
	var value_pp0 C.gpointer
	var status0 C.GFileAttributeStatus
	ret0 := C.g_file_info_get_attribute_data(info.native(), attribute0, &type0, &value_pp0, &status0)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/, FileAttributeType(type0), unsafe.Pointer(value_pp0), FileAttributeStatus(status0)
}

// GetAttributeInt32 is a wrapper around g_file_info_get_attribute_int32().
func (info FileInfo) GetAttributeInt32(attribute string) int32 {
	attribute0 := C.CString(attribute)
	ret0 := C.g_file_info_get_attribute_int32(info.native(), attribute0)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	return int32(ret0)
}

// GetAttributeInt64 is a wrapper around g_file_info_get_attribute_int64().
func (info FileInfo) GetAttributeInt64(attribute string) int64 {
	attribute0 := C.CString(attribute)
	ret0 := C.g_file_info_get_attribute_int64(info.native(), attribute0)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	return int64(ret0)
}

// GetAttributeObject is a wrapper around g_file_info_get_attribute_object().
func (info FileInfo) GetAttributeObject(attribute string) gobject.Object {
	attribute0 := C.CString(attribute)
	ret0 := C.g_file_info_get_attribute_object(info.native(), attribute0)
	C.free(unsafe.Pointer(attribute0))              /*ch:<stdlib.h>*/
	return gobject.WrapObject(unsafe.Pointer(ret0)) /*gir:GObject*/
}

// GetAttributeStatus is a wrapper around g_file_info_get_attribute_status().
func (info FileInfo) GetAttributeStatus(attribute string) FileAttributeStatus {
	attribute0 := C.CString(attribute)
	ret0 := C.g_file_info_get_attribute_status(info.native(), attribute0)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	return FileAttributeStatus(ret0)
}

// GetAttributeString is a wrapper around g_file_info_get_attribute_string().
func (info FileInfo) GetAttributeString(attribute string) string {
	attribute0 := C.CString(attribute)
	ret0 := C.g_file_info_get_attribute_string(info.native(), attribute0)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	ret := C.GoString(ret0)
	return ret
}

// GetAttributeStringv is a wrapper around g_file_info_get_attribute_stringv().
func (info FileInfo) GetAttributeStringv(attribute string) []string {
	attribute0 := C.CString(attribute)
	ret0 := C.g_file_info_get_attribute_stringv(info.native(), attribute0)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	var ret0Slice []*C.char
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString(elem)
		ret[idx] = elemG
	}
	return ret
}

// GetAttributeType is a wrapper around g_file_info_get_attribute_type().
func (info FileInfo) GetAttributeType(attribute string) FileAttributeType {
	attribute0 := C.CString(attribute)
	ret0 := C.g_file_info_get_attribute_type(info.native(), attribute0)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	return FileAttributeType(ret0)
}

// GetAttributeUint32 is a wrapper around g_file_info_get_attribute_uint32().
func (info FileInfo) GetAttributeUint32(attribute string) uint32 {
	attribute0 := C.CString(attribute)
	ret0 := C.g_file_info_get_attribute_uint32(info.native(), attribute0)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	return uint32(ret0)
}

// GetAttributeUint64 is a wrapper around g_file_info_get_attribute_uint64().
func (info FileInfo) GetAttributeUint64(attribute string) uint64 {
	attribute0 := C.CString(attribute)
	ret0 := C.g_file_info_get_attribute_uint64(info.native(), attribute0)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	return uint64(ret0)
}

// GetContentType is a wrapper around g_file_info_get_content_type().
func (info FileInfo) GetContentType() string {
	ret0 := C.g_file_info_get_content_type(info.native())
	ret := C.GoString(ret0)
	return ret
}

// GetDeletionDate is a wrapper around g_file_info_get_deletion_date().
func (info FileInfo) GetDeletionDate() glib.DateTime {
	ret0 := C.g_file_info_get_deletion_date(info.native())
	return glib.WrapDateTime(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetDisplayName is a wrapper around g_file_info_get_display_name().
func (info FileInfo) GetDisplayName() string {
	ret0 := C.g_file_info_get_display_name(info.native())
	ret := C.GoString(ret0)
	return ret
}

// GetEditName is a wrapper around g_file_info_get_edit_name().
func (info FileInfo) GetEditName() string {
	ret0 := C.g_file_info_get_edit_name(info.native())
	ret := C.GoString(ret0)
	return ret
}

// GetEtag is a wrapper around g_file_info_get_etag().
func (info FileInfo) GetEtag() string {
	ret0 := C.g_file_info_get_etag(info.native())
	ret := C.GoString(ret0)
	return ret
}

// GetFileType is a wrapper around g_file_info_get_file_type().
func (info FileInfo) GetFileType() FileType {
	ret0 := C.g_file_info_get_file_type(info.native())
	return FileType(ret0)
}

// GetIcon is a wrapper around g_file_info_get_icon().
func (info FileInfo) GetIcon() Icon {
	ret0 := C.g_file_info_get_icon(info.native())
	return wrapIcon(ret0)
}

// GetIsBackup is a wrapper around g_file_info_get_is_backup().
func (info FileInfo) GetIsBackup() bool {
	ret0 := C.g_file_info_get_is_backup(info.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIsHidden is a wrapper around g_file_info_get_is_hidden().
func (info FileInfo) GetIsHidden() bool {
	ret0 := C.g_file_info_get_is_hidden(info.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIsSymlink is a wrapper around g_file_info_get_is_symlink().
func (info FileInfo) GetIsSymlink() bool {
	ret0 := C.g_file_info_get_is_symlink(info.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetName is a wrapper around g_file_info_get_name().
func (info FileInfo) GetName() string {
	ret0 := C.g_file_info_get_name(info.native())
	ret := C.GoString(ret0)
	return ret
}

// GetSize is a wrapper around g_file_info_get_size().
func (info FileInfo) GetSize() int64 {
	ret0 := C.g_file_info_get_size(info.native())
	return int64(ret0)
}

// GetSortOrder is a wrapper around g_file_info_get_sort_order().
func (info FileInfo) GetSortOrder() int32 {
	ret0 := C.g_file_info_get_sort_order(info.native())
	return int32(ret0)
}

// GetSymbolicIcon is a wrapper around g_file_info_get_symbolic_icon().
func (info FileInfo) GetSymbolicIcon() Icon {
	ret0 := C.g_file_info_get_symbolic_icon(info.native())
	return wrapIcon(ret0)
}

// GetSymlinkTarget is a wrapper around g_file_info_get_symlink_target().
func (info FileInfo) GetSymlinkTarget() string {
	ret0 := C.g_file_info_get_symlink_target(info.native())
	ret := C.GoString(ret0)
	return ret
}

// HasAttribute is a wrapper around g_file_info_has_attribute().
func (info FileInfo) HasAttribute(attribute string) bool {
	attribute0 := C.CString(attribute)
	ret0 := C.g_file_info_has_attribute(info.native(), attribute0)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))    /*go:.util*/
}

// HasNamespace is a wrapper around g_file_info_has_namespace().
func (info FileInfo) HasNamespace(name_space string) bool {
	name_space0 := C.CString(name_space)
	ret0 := C.g_file_info_has_namespace(info.native(), name_space0)
	C.free(unsafe.Pointer(name_space0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))     /*go:.util*/
}

// ListAttributes is a wrapper around g_file_info_list_attributes().
func (info FileInfo) ListAttributes(name_space string) []string {
	name_space0 := C.CString(name_space)
	ret0 := C.g_file_info_list_attributes(info.native(), name_space0)
	C.free(unsafe.Pointer(name_space0)) /*ch:<stdlib.h>*/
	var ret0Slice []*C.char
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString(elem)
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// RemoveAttribute is a wrapper around g_file_info_remove_attribute().
func (info FileInfo) RemoveAttribute(attribute string) {
	attribute0 := C.CString(attribute)
	C.g_file_info_remove_attribute(info.native(), attribute0)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
}

// SetAttribute is a wrapper around g_file_info_set_attribute().
func (info FileInfo) SetAttribute(attribute string, type_ FileAttributeType, value_p unsafe.Pointer) {
	attribute0 := C.CString(attribute)
	C.g_file_info_set_attribute(info.native(), attribute0, C.GFileAttributeType(type_), C.gpointer(value_p))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
}

// SetAttributeBoolean is a wrapper around g_file_info_set_attribute_boolean().
func (info FileInfo) SetAttributeBoolean(attribute string, attr_value bool) {
	attribute0 := C.CString(attribute)
	C.g_file_info_set_attribute_boolean(info.native(), attribute0, C.gboolean(util.Bool2Int(attr_value)) /*go:.util*/)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
}

// SetAttributeByteString is a wrapper around g_file_info_set_attribute_byte_string().
func (info FileInfo) SetAttributeByteString(attribute string, attr_value string) {
	attribute0 := C.CString(attribute)
	attr_value0 := C.CString(attr_value)
	C.g_file_info_set_attribute_byte_string(info.native(), attribute0, attr_value0)
	C.free(unsafe.Pointer(attribute0))  /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(attr_value0)) /*ch:<stdlib.h>*/
}

// SetAttributeInt32 is a wrapper around g_file_info_set_attribute_int32().
func (info FileInfo) SetAttributeInt32(attribute string, attr_value int32) {
	attribute0 := C.CString(attribute)
	C.g_file_info_set_attribute_int32(info.native(), attribute0, C.gint32(attr_value))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
}

// SetAttributeInt64 is a wrapper around g_file_info_set_attribute_int64().
func (info FileInfo) SetAttributeInt64(attribute string, attr_value int64) {
	attribute0 := C.CString(attribute)
	C.g_file_info_set_attribute_int64(info.native(), attribute0, C.gint64(attr_value))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
}

// SetAttributeMask is a wrapper around g_file_info_set_attribute_mask().
func (info FileInfo) SetAttributeMask(mask FileAttributeMatcher) {
	C.g_file_info_set_attribute_mask(info.native(), mask.native())
}

// SetAttributeObject is a wrapper around g_file_info_set_attribute_object().
func (info FileInfo) SetAttributeObject(attribute string, attr_value gobject.Object) {
	attribute0 := C.CString(attribute)
	C.g_file_info_set_attribute_object(info.native(), attribute0, (*C.GObject)(attr_value.Ptr))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
}

// SetAttributeStatus is a wrapper around g_file_info_set_attribute_status().
func (info FileInfo) SetAttributeStatus(attribute string, status FileAttributeStatus) bool {
	attribute0 := C.CString(attribute)
	ret0 := C.g_file_info_set_attribute_status(info.native(), attribute0, C.GFileAttributeStatus(status))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))    /*go:.util*/
}

// SetAttributeString is a wrapper around g_file_info_set_attribute_string().
func (info FileInfo) SetAttributeString(attribute string, attr_value string) {
	attribute0 := C.CString(attribute)
	attr_value0 := C.CString(attr_value)
	C.g_file_info_set_attribute_string(info.native(), attribute0, attr_value0)
	C.free(unsafe.Pointer(attribute0))  /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(attr_value0)) /*ch:<stdlib.h>*/
}

// SetAttributeStringv is a wrapper around g_file_info_set_attribute_stringv().
func (info FileInfo) SetAttributeStringv(attribute string, attr_value []string) {
	attribute0 := C.CString(attribute)
	attr_value0 := make([]*C.char, len(attr_value))
	for idx, elemG := range attr_value {
		elem := C.CString(elemG)
		attr_value0[idx] = elem
	}
	var attr_value0Ptr **C.char
	if len(attr_value0) > 0 {
		attr_value0Ptr = &attr_value0[0]
	}
	C.g_file_info_set_attribute_stringv(info.native(), attribute0, attr_value0Ptr)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	for _, elem := range attr_value0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
}

// SetAttributeUint32 is a wrapper around g_file_info_set_attribute_uint32().
func (info FileInfo) SetAttributeUint32(attribute string, attr_value uint32) {
	attribute0 := C.CString(attribute)
	C.g_file_info_set_attribute_uint32(info.native(), attribute0, C.guint32(attr_value))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
}

// SetAttributeUint64 is a wrapper around g_file_info_set_attribute_uint64().
func (info FileInfo) SetAttributeUint64(attribute string, attr_value uint64) {
	attribute0 := C.CString(attribute)
	C.g_file_info_set_attribute_uint64(info.native(), attribute0, C.guint64(attr_value))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
}

// SetContentType is a wrapper around g_file_info_set_content_type().
func (info FileInfo) SetContentType(content_type string) {
	content_type0 := C.CString(content_type)
	C.g_file_info_set_content_type(info.native(), content_type0)
	C.free(unsafe.Pointer(content_type0)) /*ch:<stdlib.h>*/
}

// SetDisplayName is a wrapper around g_file_info_set_display_name().
func (info FileInfo) SetDisplayName(display_name string) {
	display_name0 := C.CString(display_name)
	C.g_file_info_set_display_name(info.native(), display_name0)
	C.free(unsafe.Pointer(display_name0)) /*ch:<stdlib.h>*/
}

// SetEditName is a wrapper around g_file_info_set_edit_name().
func (info FileInfo) SetEditName(edit_name string) {
	edit_name0 := C.CString(edit_name)
	C.g_file_info_set_edit_name(info.native(), edit_name0)
	C.free(unsafe.Pointer(edit_name0)) /*ch:<stdlib.h>*/
}

// SetFileType is a wrapper around g_file_info_set_file_type().
func (info FileInfo) SetFileType(type_ FileType) {
	C.g_file_info_set_file_type(info.native(), C.GFileType(type_))
}

// SetIcon is a wrapper around g_file_info_set_icon().
func (info FileInfo) SetIcon(icon Icon) {
	C.g_file_info_set_icon(info.native(), icon.native())
}

// SetIsHidden is a wrapper around g_file_info_set_is_hidden().
func (info FileInfo) SetIsHidden(is_hidden bool) {
	C.g_file_info_set_is_hidden(info.native(), C.gboolean(util.Bool2Int(is_hidden)) /*go:.util*/)
}

// SetIsSymlink is a wrapper around g_file_info_set_is_symlink().
func (info FileInfo) SetIsSymlink(is_symlink bool) {
	C.g_file_info_set_is_symlink(info.native(), C.gboolean(util.Bool2Int(is_symlink)) /*go:.util*/)
}

// SetModificationTime is a wrapper around g_file_info_set_modification_time().
func (info FileInfo) SetModificationTime(mtime glib.TimeVal) {
	C.g_file_info_set_modification_time(info.native(), (*C.GTimeVal)(mtime.Ptr))
}

// SetName is a wrapper around g_file_info_set_name().
func (info FileInfo) SetName(name string) {
	name0 := C.CString(name)
	C.g_file_info_set_name(info.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// SetSize is a wrapper around g_file_info_set_size().
func (info FileInfo) SetSize(size int64) {
	C.g_file_info_set_size(info.native(), C.goffset(size))
}

// SetSortOrder is a wrapper around g_file_info_set_sort_order().
func (info FileInfo) SetSortOrder(sort_order int32) {
	C.g_file_info_set_sort_order(info.native(), C.gint32(sort_order))
}

// SetSymbolicIcon is a wrapper around g_file_info_set_symbolic_icon().
func (info FileInfo) SetSymbolicIcon(icon Icon) {
	C.g_file_info_set_symbolic_icon(info.native(), icon.native())
}

// SetSymlinkTarget is a wrapper around g_file_info_set_symlink_target().
func (info FileInfo) SetSymlinkTarget(symlink_target string) {
	symlink_target0 := C.CString(symlink_target)
	C.g_file_info_set_symlink_target(info.native(), symlink_target0)
	C.free(unsafe.Pointer(symlink_target0)) /*ch:<stdlib.h>*/
}

// UnsetAttributeMask is a wrapper around g_file_info_unset_attribute_mask().
func (info FileInfo) UnsetAttributeMask() {
	C.g_file_info_unset_attribute_mask(info.native())
}

// Object Application
type Application struct {
	gobject.Object
}

func (v Application) native() *C.GApplication {
	return (*C.GApplication)(v.Ptr)
}
func wrapApplication(p *C.GApplication) (v Application) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapApplication(p unsafe.Pointer) (v Application) {
	v.Ptr = p
	return
}
func (v Application) IsNil() bool {
	return v.Ptr == nil
}
func IWrapApplication(p unsafe.Pointer) interface{} {
	return WrapApplication(p)
}
func (v Application) GetType() gobject.Type {
	return gobject.Type(C.g_application_get_type())
}
func (v Application) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapApplication(unsafe.Pointer(ptr)), nil
	}
}
func (v Application) ActionGroup() ActionGroup {
	return WrapActionGroup(v.Ptr)
}
func (v Application) ActionMap() ActionMap {
	return WrapActionMap(v.Ptr)
}

// ApplicationNew is a wrapper around g_application_new().
func ApplicationNew(application_id string, flags ApplicationFlags) Application {
	application_id0 := (*C.gchar)(C.CString(application_id))
	ret0 := C.g_application_new(application_id0, C.GApplicationFlags(flags))
	C.free(unsafe.Pointer(application_id0)) /*ch:<stdlib.h>*/
	return wrapApplication(ret0)
}

// Activate is a wrapper around g_application_activate().
func (application Application) Activate() {
	C.g_application_activate(application.native())
}

// AddOptionGroup is a wrapper around g_application_add_option_group().
func (application Application) AddOptionGroup(group glib.OptionGroup) {
	C.g_application_add_option_group(application.native(), (*C.GOptionGroup)(group.Ptr))
}

// BindBusyProperty is a wrapper around g_application_bind_busy_property().
func (application Application) BindBusyProperty(object gobject.Object, property string) {
	property0 := (*C.gchar)(C.CString(property))
	C.g_application_bind_busy_property(application.native(), C.gpointer((C.gpointer)(object.Ptr)), property0)
	C.free(unsafe.Pointer(property0)) /*ch:<stdlib.h>*/
}

// GetApplicationId is a wrapper around g_application_get_application_id().
func (application Application) GetApplicationId() string {
	ret0 := C.g_application_get_application_id(application.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetDbusConnection is a wrapper around g_application_get_dbus_connection().
func (application Application) GetDbusConnection() DBusConnection {
	ret0 := C.g_application_get_dbus_connection(application.native())
	return wrapDBusConnection(ret0)
}

// GetDbusObjectPath is a wrapper around g_application_get_dbus_object_path().
func (application Application) GetDbusObjectPath() string {
	ret0 := C.g_application_get_dbus_object_path(application.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetFlags is a wrapper around g_application_get_flags().
func (application Application) GetFlags() ApplicationFlags {
	ret0 := C.g_application_get_flags(application.native())
	return ApplicationFlags(ret0)
}

// GetInactivityTimeout is a wrapper around g_application_get_inactivity_timeout().
func (application Application) GetInactivityTimeout() uint {
	ret0 := C.g_application_get_inactivity_timeout(application.native())
	return uint(ret0)
}

// GetIsBusy is a wrapper around g_application_get_is_busy().
func (application Application) GetIsBusy() bool {
	ret0 := C.g_application_get_is_busy(application.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIsRegistered is a wrapper around g_application_get_is_registered().
func (application Application) GetIsRegistered() bool {
	ret0 := C.g_application_get_is_registered(application.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIsRemote is a wrapper around g_application_get_is_remote().
func (application Application) GetIsRemote() bool {
	ret0 := C.g_application_get_is_remote(application.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetResourceBasePath is a wrapper around g_application_get_resource_base_path().
func (application Application) GetResourceBasePath() string {
	ret0 := C.g_application_get_resource_base_path(application.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// Hold is a wrapper around g_application_hold().
func (application Application) Hold() {
	C.g_application_hold(application.native())
}

// MarkBusy is a wrapper around g_application_mark_busy().
func (application Application) MarkBusy() {
	C.g_application_mark_busy(application.native())
}

// Open is a wrapper around g_application_open().
func (application Application) Open(files []File, hint string) {
	files0 := make([]*C.GFile, len(files))
	for idx, elemG := range files {
		files0[idx] = elemG.native()
	}
	var files0Ptr **C.GFile
	if len(files0) > 0 {
		files0Ptr = &files0[0]
	}
	hint0 := (*C.gchar)(C.CString(hint))
	C.g_application_open(application.native(), files0Ptr, C.gint(len(files)), hint0)
	C.free(unsafe.Pointer(hint0)) /*ch:<stdlib.h>*/
}

// Quit is a wrapper around g_application_quit().
func (application Application) Quit() {
	C.g_application_quit(application.native())
}

// Register is a wrapper around g_application_register().
func (application Application) Register(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_application_register(application.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Release is a wrapper around g_application_release().
func (application Application) Release() {
	C.g_application_release(application.native())
}

// Run is a wrapper around g_application_run().
func (application Application) Run(argv []string) int {
	argv0 := make([]*C.char, len(argv))
	for idx, elemG := range argv {
		elem := C.CString(elemG)
		argv0[idx] = elem
	}
	var argv0Ptr **C.char
	if len(argv0) > 0 {
		argv0Ptr = &argv0[0]
	}
	ret0 := C.g_application_run(application.native(), C.int(len(argv)), argv0Ptr)
	for _, elem := range argv0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	return int(ret0)
}

// SendNotification is a wrapper around g_application_send_notification().
func (application Application) SendNotification(id string, notification Notification) {
	id0 := (*C.gchar)(C.CString(id))
	C.g_application_send_notification(application.native(), id0, notification.native())
	C.free(unsafe.Pointer(id0)) /*ch:<stdlib.h>*/
}

// SetApplicationId is a wrapper around g_application_set_application_id().
func (application Application) SetApplicationId(application_id string) {
	application_id0 := (*C.gchar)(C.CString(application_id))
	C.g_application_set_application_id(application.native(), application_id0)
	C.free(unsafe.Pointer(application_id0)) /*ch:<stdlib.h>*/
}

// SetDefault is a wrapper around g_application_set_default().
func (application Application) SetDefault() {
	C.g_application_set_default(application.native())
}

// SetFlags is a wrapper around g_application_set_flags().
func (application Application) SetFlags(flags ApplicationFlags) {
	C.g_application_set_flags(application.native(), C.GApplicationFlags(flags))
}

// SetInactivityTimeout is a wrapper around g_application_set_inactivity_timeout().
func (application Application) SetInactivityTimeout(inactivity_timeout uint) {
	C.g_application_set_inactivity_timeout(application.native(), C.guint(inactivity_timeout))
}

// SetResourceBasePath is a wrapper around g_application_set_resource_base_path().
func (application Application) SetResourceBasePath(resource_path string) {
	resource_path0 := (*C.gchar)(C.CString(resource_path))
	C.g_application_set_resource_base_path(application.native(), resource_path0)
	C.free(unsafe.Pointer(resource_path0)) /*ch:<stdlib.h>*/
}

// UnbindBusyProperty is a wrapper around g_application_unbind_busy_property().
func (application Application) UnbindBusyProperty(object gobject.Object, property string) {
	property0 := (*C.gchar)(C.CString(property))
	C.g_application_unbind_busy_property(application.native(), C.gpointer((C.gpointer)(object.Ptr)), property0)
	C.free(unsafe.Pointer(property0)) /*ch:<stdlib.h>*/
}

// UnmarkBusy is a wrapper around g_application_unmark_busy().
func (application Application) UnmarkBusy() {
	C.g_application_unmark_busy(application.native())
}

// WithdrawNotification is a wrapper around g_application_withdraw_notification().
func (application Application) WithdrawNotification(id string) {
	id0 := (*C.gchar)(C.CString(id))
	C.g_application_withdraw_notification(application.native(), id0)
	C.free(unsafe.Pointer(id0)) /*ch:<stdlib.h>*/
}

// ApplicationGetDefault is a wrapper around g_application_get_default().
func ApplicationGetDefault() Application {
	ret0 := C.g_application_get_default()
	return wrapApplication(ret0)
}

// ApplicationIdIsValid is a wrapper around g_application_id_is_valid().
func ApplicationIdIsValid(application_id string) bool {
	application_id0 := (*C.gchar)(C.CString(application_id))
	ret0 := C.g_application_id_is_valid(application_id0)
	C.free(unsafe.Pointer(application_id0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))         /*go:.util*/
}

// Interface ActionMap
type ActionMap struct {
	Ptr unsafe.Pointer
}

func (v ActionMap) native() *C.GActionMap {
	return (*C.GActionMap)(v.Ptr)
}
func wrapActionMap(p *C.GActionMap) ActionMap {
	return ActionMap{unsafe.Pointer(p)}
}
func WrapActionMap(p unsafe.Pointer) ActionMap {
	return ActionMap{p}
}
func (v ActionMap) IsNil() bool {
	return v.Ptr == nil
}
func IWrapActionMap(p unsafe.Pointer) interface{} {
	return WrapActionMap(p)
}
func (v ActionMap) GetType() gobject.Type {
	return gobject.Type(C.g_action_map_get_type())
}
func (v ActionMap) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapActionMap(unsafe.Pointer(ptr)), nil
	}
}

// AddAction is a wrapper around g_action_map_add_action().
func (action_map ActionMap) AddAction(action Action) {
	C.g_action_map_add_action(action_map.native(), action.native())
}

// LookupAction is a wrapper around g_action_map_lookup_action().
func (action_map ActionMap) LookupAction(action_name string) Action {
	action_name0 := (*C.gchar)(C.CString(action_name))
	ret0 := C.g_action_map_lookup_action(action_map.native(), action_name0)
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
	return wrapAction(ret0)
}

// RemoveAction is a wrapper around g_action_map_remove_action().
func (action_map ActionMap) RemoveAction(action_name string) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.g_action_map_remove_action(action_map.native(), action_name0)
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// Interface ActionGroup
type ActionGroup struct {
	Ptr unsafe.Pointer
}

func (v ActionGroup) native() *C.GActionGroup {
	return (*C.GActionGroup)(v.Ptr)
}
func wrapActionGroup(p *C.GActionGroup) ActionGroup {
	return ActionGroup{unsafe.Pointer(p)}
}
func WrapActionGroup(p unsafe.Pointer) ActionGroup {
	return ActionGroup{p}
}
func (v ActionGroup) IsNil() bool {
	return v.Ptr == nil
}
func IWrapActionGroup(p unsafe.Pointer) interface{} {
	return WrapActionGroup(p)
}
func (v ActionGroup) GetType() gobject.Type {
	return gobject.Type(C.g_action_group_get_type())
}
func (v ActionGroup) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapActionGroup(unsafe.Pointer(ptr)), nil
	}
}

// ActionAdded is a wrapper around g_action_group_action_added().
func (action_group ActionGroup) ActionAdded(action_name string) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.g_action_group_action_added(action_group.native(), action_name0)
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// ActionEnabledChanged is a wrapper around g_action_group_action_enabled_changed().
func (action_group ActionGroup) ActionEnabledChanged(action_name string, enabled bool) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.g_action_group_action_enabled_changed(action_group.native(), action_name0, C.gboolean(util.Bool2Int(enabled)) /*go:.util*/)
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// ActionRemoved is a wrapper around g_action_group_action_removed().
func (action_group ActionGroup) ActionRemoved(action_name string) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.g_action_group_action_removed(action_group.native(), action_name0)
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// ActionStateChanged is a wrapper around g_action_group_action_state_changed().
func (action_group ActionGroup) ActionStateChanged(action_name string, state glib.Variant) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.g_action_group_action_state_changed(action_group.native(), action_name0, (*C.GVariant)(state.Ptr))
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// ActivateAction is a wrapper around g_action_group_activate_action().
func (action_group ActionGroup) ActivateAction(action_name string, parameter glib.Variant) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.g_action_group_activate_action(action_group.native(), action_name0, (*C.GVariant)(parameter.Ptr))
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// ChangeActionState is a wrapper around g_action_group_change_action_state().
func (action_group ActionGroup) ChangeActionState(action_name string, value glib.Variant) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.g_action_group_change_action_state(action_group.native(), action_name0, (*C.GVariant)(value.Ptr))
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// GetActionEnabled is a wrapper around g_action_group_get_action_enabled().
func (action_group ActionGroup) GetActionEnabled(action_name string) bool {
	action_name0 := (*C.gchar)(C.CString(action_name))
	ret0 := C.g_action_group_get_action_enabled(action_group.native(), action_name0)
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))      /*go:.util*/
}

// GetActionParameterType is a wrapper around g_action_group_get_action_parameter_type().
func (action_group ActionGroup) GetActionParameterType(action_name string) glib.VariantType {
	action_name0 := (*C.gchar)(C.CString(action_name))
	ret0 := C.g_action_group_get_action_parameter_type(action_group.native(), action_name0)
	C.free(unsafe.Pointer(action_name0))              /*ch:<stdlib.h>*/
	return glib.WrapVariantType(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetActionState is a wrapper around g_action_group_get_action_state().
func (action_group ActionGroup) GetActionState(action_name string) glib.Variant {
	action_name0 := (*C.gchar)(C.CString(action_name))
	ret0 := C.g_action_group_get_action_state(action_group.native(), action_name0)
	C.free(unsafe.Pointer(action_name0))          /*ch:<stdlib.h>*/
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetActionStateHint is a wrapper around g_action_group_get_action_state_hint().
func (action_group ActionGroup) GetActionStateHint(action_name string) glib.Variant {
	action_name0 := (*C.gchar)(C.CString(action_name))
	ret0 := C.g_action_group_get_action_state_hint(action_group.native(), action_name0)
	C.free(unsafe.Pointer(action_name0))          /*ch:<stdlib.h>*/
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetActionStateType is a wrapper around g_action_group_get_action_state_type().
func (action_group ActionGroup) GetActionStateType(action_name string) glib.VariantType {
	action_name0 := (*C.gchar)(C.CString(action_name))
	ret0 := C.g_action_group_get_action_state_type(action_group.native(), action_name0)
	C.free(unsafe.Pointer(action_name0))              /*ch:<stdlib.h>*/
	return glib.WrapVariantType(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// HasAction is a wrapper around g_action_group_has_action().
func (action_group ActionGroup) HasAction(action_name string) bool {
	action_name0 := (*C.gchar)(C.CString(action_name))
	ret0 := C.g_action_group_has_action(action_group.native(), action_name0)
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))      /*go:.util*/
}

// ListActions is a wrapper around g_action_group_list_actions().
func (action_group ActionGroup) ListActions() []string {
	ret0 := C.g_action_group_list_actions(action_group.native())
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// QueryAction is a wrapper around g_action_group_query_action().
func (action_group ActionGroup) QueryAction(action_name string) (bool, bool, glib.VariantType, glib.VariantType, glib.Variant, glib.Variant) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	var enabled0 C.gboolean
	var parameter_type0 *C.GVariantType
	var state_type0 *C.GVariantType
	var state_hint0 *C.GVariant
	var state0 *C.GVariant
	ret0 := C.g_action_group_query_action(action_group.native(), action_name0, &enabled0, &parameter_type0, &state_type0, &state_hint0, &state0)
	C.free(unsafe.Pointer(action_name0))                                                                                                                                                                                                                                                                                              /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/, util.Int2Bool(int(enabled0)) /*go:.util*/, glib.WrapVariantType(unsafe.Pointer(parameter_type0)) /*gir:GLib*/, glib.WrapVariantType(unsafe.Pointer(state_type0)) /*gir:GLib*/, glib.WrapVariant(unsafe.Pointer(state_hint0)) /*gir:GLib*/, glib.WrapVariant(unsafe.Pointer(state0)) /*gir:GLib*/
}

// Object SimpleActionGroup
type SimpleActionGroup struct {
	gobject.Object
}

func (v SimpleActionGroup) native() *C.GSimpleActionGroup {
	return (*C.GSimpleActionGroup)(v.Ptr)
}
func wrapSimpleActionGroup(p *C.GSimpleActionGroup) (v SimpleActionGroup) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSimpleActionGroup(p unsafe.Pointer) (v SimpleActionGroup) {
	v.Ptr = p
	return
}
func (v SimpleActionGroup) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSimpleActionGroup(p unsafe.Pointer) interface{} {
	return WrapSimpleActionGroup(p)
}
func (v SimpleActionGroup) GetType() gobject.Type {
	return gobject.Type(C.g_simple_action_group_get_type())
}
func (v SimpleActionGroup) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSimpleActionGroup(unsafe.Pointer(ptr)), nil
	}
}
func (v SimpleActionGroup) ActionGroup() ActionGroup {
	return WrapActionGroup(v.Ptr)
}
func (v SimpleActionGroup) ActionMap() ActionMap {
	return WrapActionMap(v.Ptr)
}

// SimpleActionGroupNew is a wrapper around g_simple_action_group_new().
func SimpleActionGroupNew() SimpleActionGroup {
	ret0 := C.g_simple_action_group_new()
	return wrapSimpleActionGroup(ret0)
}

// Object Notification
type Notification struct {
	gobject.Object
}

func (v Notification) native() *C.GNotification {
	return (*C.GNotification)(v.Ptr)
}
func wrapNotification(p *C.GNotification) (v Notification) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapNotification(p unsafe.Pointer) (v Notification) {
	v.Ptr = p
	return
}
func (v Notification) IsNil() bool {
	return v.Ptr == nil
}
func IWrapNotification(p unsafe.Pointer) interface{} {
	return WrapNotification(p)
}
func (v Notification) GetType() gobject.Type {
	return gobject.Type(C.g_notification_get_type())
}
func (v Notification) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapNotification(unsafe.Pointer(ptr)), nil
	}
}

// NotificationNew is a wrapper around g_notification_new().
func NotificationNew(title string) Notification {
	title0 := (*C.gchar)(C.CString(title))
	ret0 := C.g_notification_new(title0)
	C.free(unsafe.Pointer(title0)) /*ch:<stdlib.h>*/
	return wrapNotification(ret0)
}

// AddButton is a wrapper around g_notification_add_button().
func (notification Notification) AddButton(label string, detailed_action string) {
	label0 := (*C.gchar)(C.CString(label))
	detailed_action0 := (*C.gchar)(C.CString(detailed_action))
	C.g_notification_add_button(notification.native(), label0, detailed_action0)
	C.free(unsafe.Pointer(label0))           /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(detailed_action0)) /*ch:<stdlib.h>*/
}

// AddButtonWithTargetValue is a wrapper around g_notification_add_button_with_target_value().
func (notification Notification) AddButtonWithTargetValue(label string, action string, target glib.Variant) {
	label0 := (*C.gchar)(C.CString(label))
	action0 := (*C.gchar)(C.CString(action))
	C.g_notification_add_button_with_target_value(notification.native(), label0, action0, (*C.GVariant)(target.Ptr))
	C.free(unsafe.Pointer(label0))  /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(action0)) /*ch:<stdlib.h>*/
}

// SetBody is a wrapper around g_notification_set_body().
func (notification Notification) SetBody(body string) {
	body0 := (*C.gchar)(C.CString(body))
	C.g_notification_set_body(notification.native(), body0)
	C.free(unsafe.Pointer(body0)) /*ch:<stdlib.h>*/
}

// SetDefaultAction is a wrapper around g_notification_set_default_action().
func (notification Notification) SetDefaultAction(detailed_action string) {
	detailed_action0 := (*C.gchar)(C.CString(detailed_action))
	C.g_notification_set_default_action(notification.native(), detailed_action0)
	C.free(unsafe.Pointer(detailed_action0)) /*ch:<stdlib.h>*/
}

// SetDefaultActionAndTargetValue is a wrapper around g_notification_set_default_action_and_target_value().
func (notification Notification) SetDefaultActionAndTargetValue(action string, target glib.Variant) {
	action0 := (*C.gchar)(C.CString(action))
	C.g_notification_set_default_action_and_target_value(notification.native(), action0, (*C.GVariant)(target.Ptr))
	C.free(unsafe.Pointer(action0)) /*ch:<stdlib.h>*/
}

// SetIcon is a wrapper around g_notification_set_icon().
func (notification Notification) SetIcon(icon Icon) {
	C.g_notification_set_icon(notification.native(), icon.native())
}

// SetPriority is a wrapper around g_notification_set_priority().
func (notification Notification) SetPriority(priority NotificationPriority) {
	C.g_notification_set_priority(notification.native(), C.GNotificationPriority(priority))
}

// SetTitle is a wrapper around g_notification_set_title().
func (notification Notification) SetTitle(title string) {
	title0 := (*C.gchar)(C.CString(title))
	C.g_notification_set_title(notification.native(), title0)
	C.free(unsafe.Pointer(title0)) /*ch:<stdlib.h>*/
}

// Interface Icon
type Icon struct {
	Ptr unsafe.Pointer
}

func (v Icon) native() *C.GIcon {
	return (*C.GIcon)(v.Ptr)
}
func wrapIcon(p *C.GIcon) Icon {
	return Icon{unsafe.Pointer(p)}
}
func WrapIcon(p unsafe.Pointer) Icon {
	return Icon{p}
}
func (v Icon) IsNil() bool {
	return v.Ptr == nil
}
func IWrapIcon(p unsafe.Pointer) interface{} {
	return WrapIcon(p)
}
func (v Icon) GetType() gobject.Type {
	return gobject.Type(C.g_icon_get_type())
}
func (v Icon) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapIcon(unsafe.Pointer(ptr)), nil
	}
}

// Equal is a wrapper around g_icon_equal().
func (icon1 Icon) Equal(icon2 Icon) bool {
	ret0 := C.g_icon_equal(icon1.native(), icon2.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Serialize is a wrapper around g_icon_serialize().
func (icon Icon) Serialize() glib.Variant {
	ret0 := C.g_icon_serialize(icon.native())
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// ToString is a wrapper around g_icon_to_string().
func (icon Icon) ToString() string {
	ret0 := C.g_icon_to_string(icon.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// IconDeserialize is a wrapper around g_icon_deserialize().
func IconDeserialize(value glib.Variant) Icon {
	ret0 := C.g_icon_deserialize((*C.GVariant)(value.Ptr))
	return wrapIcon(ret0)
}

// IconHash is a wrapper around g_icon_hash().
func IconHash(icon unsafe.Pointer) uint {
	ret0 := C.g_icon_hash(C.gconstpointer(icon))
	return uint(ret0)
}

// IconNewForString is a wrapper around g_icon_new_for_string().
func IconNewForString(str string) (Icon, error) {
	str0 := (*C.gchar)(C.CString(str))
	var err glib.Error
	ret0 := C.g_icon_new_for_string(str0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(str0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return Icon{}, err.GoValue()
	}
	return wrapIcon(ret0), nil
}

// Interface AsyncResult
type AsyncResult struct {
	Ptr unsafe.Pointer
}

func (v AsyncResult) native() *C.GAsyncResult {
	return (*C.GAsyncResult)(v.Ptr)
}
func wrapAsyncResult(p *C.GAsyncResult) AsyncResult {
	return AsyncResult{unsafe.Pointer(p)}
}
func WrapAsyncResult(p unsafe.Pointer) AsyncResult {
	return AsyncResult{p}
}
func (v AsyncResult) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAsyncResult(p unsafe.Pointer) interface{} {
	return WrapAsyncResult(p)
}
func (v AsyncResult) GetType() gobject.Type {
	return gobject.Type(C.g_async_result_get_type())
}
func (v AsyncResult) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAsyncResult(unsafe.Pointer(ptr)), nil
	}
}

// GetSourceObject is a wrapper around g_async_result_get_source_object().
func (res AsyncResult) GetSourceObject() gobject.Object {
	ret0 := C.g_async_result_get_source_object(res.native())
	return gobject.WrapObject(unsafe.Pointer(ret0)) /*gir:GObject*/
}

// GetUserData is a wrapper around g_async_result_get_user_data().
func (res AsyncResult) GetUserData() unsafe.Pointer {
	ret0 := C.g_async_result_get_user_data(res.native())
	return unsafe.Pointer(ret0)
}

// IsTagged is a wrapper around g_async_result_is_tagged().
func (res AsyncResult) IsTagged(source_tag unsafe.Pointer) bool {
	ret0 := C.g_async_result_is_tagged(res.native(), C.gpointer(source_tag))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// LegacyPropagateError is a wrapper around g_async_result_legacy_propagate_error().
func (res AsyncResult) LegacyPropagateError() (bool, error) {
	var err glib.Error
	ret0 := C.g_async_result_legacy_propagate_error(res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Object IOStream
type IOStream struct {
	gobject.Object
}

func (v IOStream) native() *C.GIOStream {
	return (*C.GIOStream)(v.Ptr)
}
func wrapIOStream(p *C.GIOStream) (v IOStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapIOStream(p unsafe.Pointer) (v IOStream) {
	v.Ptr = p
	return
}
func (v IOStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapIOStream(p unsafe.Pointer) interface{} {
	return WrapIOStream(p)
}
func (v IOStream) GetType() gobject.Type {
	return gobject.Type(C.g_io_stream_get_type())
}
func (v IOStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapIOStream(unsafe.Pointer(ptr)), nil
	}
}

// ClearPending is a wrapper around g_io_stream_clear_pending().
func (stream IOStream) ClearPending() {
	C.g_io_stream_clear_pending(stream.native())
}

// Close is a wrapper around g_io_stream_close().
func (stream IOStream) Close(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_io_stream_close(stream.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// CloseAsync is a wrapper around g_io_stream_close_async().
func (stream IOStream) CloseAsync(io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_io_stream_close_async(stream.native(), C.int(io_priority), cancellable.native(), callback0)
}

// CloseFinish is a wrapper around g_io_stream_close_finish().
func (stream IOStream) CloseFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_io_stream_close_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// GetInputStream is a wrapper around g_io_stream_get_input_stream().
func (stream IOStream) GetInputStream() InputStream {
	ret0 := C.g_io_stream_get_input_stream(stream.native())
	return wrapInputStream(ret0)
}

// GetOutputStream is a wrapper around g_io_stream_get_output_stream().
func (stream IOStream) GetOutputStream() OutputStream {
	ret0 := C.g_io_stream_get_output_stream(stream.native())
	return wrapOutputStream(ret0)
}

// HasPending is a wrapper around g_io_stream_has_pending().
func (stream IOStream) HasPending() bool {
	ret0 := C.g_io_stream_has_pending(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsClosed is a wrapper around g_io_stream_is_closed().
func (stream IOStream) IsClosed() bool {
	ret0 := C.g_io_stream_is_closed(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetPending is a wrapper around g_io_stream_set_pending().
func (stream IOStream) SetPending() (bool, error) {
	var err glib.Error
	ret0 := C.g_io_stream_set_pending(stream.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SpliceAsync is a wrapper around g_io_stream_splice_async().
func (stream1 IOStream) SpliceAsync(stream2 IOStream, flags IOStreamSpliceFlags, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_io_stream_splice_async(stream1.native(), stream2.native(), C.GIOStreamSpliceFlags(flags), C.int(io_priority), cancellable.native(), callback0)
}

// IOStreamSpliceFinish is a wrapper around g_io_stream_splice_finish().
func IOStreamSpliceFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_io_stream_splice_finish(result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Object InputStream
type InputStream struct {
	gobject.Object
}

func (v InputStream) native() *C.GInputStream {
	return (*C.GInputStream)(v.Ptr)
}
func wrapInputStream(p *C.GInputStream) (v InputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapInputStream(p unsafe.Pointer) (v InputStream) {
	v.Ptr = p
	return
}
func (v InputStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapInputStream(p unsafe.Pointer) interface{} {
	return WrapInputStream(p)
}
func (v InputStream) GetType() gobject.Type {
	return gobject.Type(C.g_input_stream_get_type())
}
func (v InputStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapInputStream(unsafe.Pointer(ptr)), nil
	}
}

// ClearPending is a wrapper around g_input_stream_clear_pending().
func (stream InputStream) ClearPending() {
	C.g_input_stream_clear_pending(stream.native())
}

// Close is a wrapper around g_input_stream_close().
func (stream InputStream) Close(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_input_stream_close(stream.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// CloseAsync is a wrapper around g_input_stream_close_async().
func (stream InputStream) CloseAsync(io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_input_stream_close_async(stream.native(), C.int(io_priority), cancellable.native(), callback0)
}

// CloseFinish is a wrapper around g_input_stream_close_finish().
func (stream InputStream) CloseFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_input_stream_close_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// HasPending is a wrapper around g_input_stream_has_pending().
func (stream InputStream) HasPending() bool {
	ret0 := C.g_input_stream_has_pending(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsClosed is a wrapper around g_input_stream_is_closed().
func (stream InputStream) IsClosed() bool {
	ret0 := C.g_input_stream_is_closed(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ReadAllFinish is a wrapper around g_input_stream_read_all_finish().
func (stream InputStream) ReadAllFinish(result AsyncResult) (bool, uint, error) {
	var bytes_read0 C.gsize
	var err glib.Error
	ret0 := C.g_input_stream_read_all_finish(stream.native(), result.native(), &bytes_read0, (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, 0, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, uint(bytes_read0), nil
}

// ReadBytes is a wrapper around g_input_stream_read_bytes().
func (stream InputStream) ReadBytes(count uint, cancellable Cancellable) (glib.Bytes, error) {
	var err glib.Error
	ret0 := C.g_input_stream_read_bytes(stream.native(), C.gsize(count), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return glib.Bytes{}, err.GoValue()
	}
	return glib.WrapBytes(unsafe.Pointer(ret0)) /*gir:GLib*/, nil
}

// ReadBytesAsync is a wrapper around g_input_stream_read_bytes_async().
func (stream InputStream) ReadBytesAsync(count uint, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_input_stream_read_bytes_async(stream.native(), C.gsize(count), C.int(io_priority), cancellable.native(), callback0)
}

// ReadBytesFinish is a wrapper around g_input_stream_read_bytes_finish().
func (stream InputStream) ReadBytesFinish(result AsyncResult) (glib.Bytes, error) {
	var err glib.Error
	ret0 := C.g_input_stream_read_bytes_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return glib.Bytes{}, err.GoValue()
	}
	return glib.WrapBytes(unsafe.Pointer(ret0)) /*gir:GLib*/, nil
}

// ReadFinish is a wrapper around g_input_stream_read_finish().
func (stream InputStream) ReadFinish(result AsyncResult) (int, error) {
	var err glib.Error
	ret0 := C.g_input_stream_read_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// SetPending is a wrapper around g_input_stream_set_pending().
func (stream InputStream) SetPending() (bool, error) {
	var err glib.Error
	ret0 := C.g_input_stream_set_pending(stream.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Skip is a wrapper around g_input_stream_skip().
func (stream InputStream) Skip(count uint, cancellable Cancellable) (int, error) {
	var err glib.Error
	ret0 := C.g_input_stream_skip(stream.native(), C.gsize(count), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// SkipAsync is a wrapper around g_input_stream_skip_async().
func (stream InputStream) SkipAsync(count uint, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_input_stream_skip_async(stream.native(), C.gsize(count), C.int(io_priority), cancellable.native(), callback0)
}

// SkipFinish is a wrapper around g_input_stream_skip_finish().
func (stream InputStream) SkipFinish(result AsyncResult) (int, error) {
	var err glib.Error
	ret0 := C.g_input_stream_skip_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// Object FileIOStream
type FileIOStream struct {
	IOStream
}

func (v FileIOStream) native() *C.GFileIOStream {
	return (*C.GFileIOStream)(v.Ptr)
}
func wrapFileIOStream(p *C.GFileIOStream) (v FileIOStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileIOStream(p unsafe.Pointer) (v FileIOStream) {
	v.Ptr = p
	return
}
func (v FileIOStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFileIOStream(p unsafe.Pointer) interface{} {
	return WrapFileIOStream(p)
}
func (v FileIOStream) GetType() gobject.Type {
	return gobject.Type(C.g_file_io_stream_get_type())
}
func (v FileIOStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFileIOStream(unsafe.Pointer(ptr)), nil
	}
}
func (v FileIOStream) Seekable() Seekable {
	return WrapSeekable(v.Ptr)
}

// GetEtag is a wrapper around g_file_io_stream_get_etag().
func (stream FileIOStream) GetEtag() string {
	ret0 := C.g_file_io_stream_get_etag(stream.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// QueryInfo is a wrapper around g_file_io_stream_query_info().
func (stream FileIOStream) QueryInfo(attributes string, cancellable Cancellable) (FileInfo, error) {
	attributes0 := C.CString(attributes)
	var err glib.Error
	ret0 := C.g_file_io_stream_query_info(stream.native(), attributes0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// QueryInfoAsync is a wrapper around g_file_io_stream_query_info_async().
func (stream FileIOStream) QueryInfoAsync(attributes string, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	attributes0 := C.CString(attributes)
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_io_stream_query_info_async(stream.native(), attributes0, C.int(io_priority), cancellable.native(), callback0)
	C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/
}

// QueryInfoFinish is a wrapper around g_file_io_stream_query_info_finish().
func (stream FileIOStream) QueryInfoFinish(result AsyncResult) (FileInfo, error) {
	var err glib.Error
	ret0 := C.g_file_io_stream_query_info_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// Object FileEnumerator
type FileEnumerator struct {
	gobject.Object
}

func (v FileEnumerator) native() *C.GFileEnumerator {
	return (*C.GFileEnumerator)(v.Ptr)
}
func wrapFileEnumerator(p *C.GFileEnumerator) (v FileEnumerator) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileEnumerator(p unsafe.Pointer) (v FileEnumerator) {
	v.Ptr = p
	return
}
func (v FileEnumerator) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFileEnumerator(p unsafe.Pointer) interface{} {
	return WrapFileEnumerator(p)
}
func (v FileEnumerator) GetType() gobject.Type {
	return gobject.Type(C.g_file_enumerator_get_type())
}
func (v FileEnumerator) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFileEnumerator(unsafe.Pointer(ptr)), nil
	}
}

// Close is a wrapper around g_file_enumerator_close().
func (enumerator FileEnumerator) Close(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_enumerator_close(enumerator.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// CloseAsync is a wrapper around g_file_enumerator_close_async().
func (enumerator FileEnumerator) CloseAsync(io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_enumerator_close_async(enumerator.native(), C.int(io_priority), cancellable.native(), callback0)
}

// CloseFinish is a wrapper around g_file_enumerator_close_finish().
func (enumerator FileEnumerator) CloseFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_file_enumerator_close_finish(enumerator.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// GetChild is a wrapper around g_file_enumerator_get_child().
func (enumerator FileEnumerator) GetChild(info FileInfo) File {
	ret0 := C.g_file_enumerator_get_child(enumerator.native(), info.native())
	return wrapFile(ret0)
}

// GetContainer is a wrapper around g_file_enumerator_get_container().
func (enumerator FileEnumerator) GetContainer() File {
	ret0 := C.g_file_enumerator_get_container(enumerator.native())
	return wrapFile(ret0)
}

// HasPending is a wrapper around g_file_enumerator_has_pending().
func (enumerator FileEnumerator) HasPending() bool {
	ret0 := C.g_file_enumerator_has_pending(enumerator.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsClosed is a wrapper around g_file_enumerator_is_closed().
func (enumerator FileEnumerator) IsClosed() bool {
	ret0 := C.g_file_enumerator_is_closed(enumerator.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Iterate is a wrapper around g_file_enumerator_iterate().
func (direnum FileEnumerator) Iterate(cancellable Cancellable) (bool, FileInfo, File, error) {
	var out_info0 *C.GFileInfo
	var out_child0 *C.GFile
	var err glib.Error
	ret0 := C.g_file_enumerator_iterate(direnum.native(), &out_info0, &out_child0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, FileInfo{}, File{}, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, wrapFileInfo(out_info0), wrapFile(out_child0), nil
}

// NextFile is a wrapper around g_file_enumerator_next_file().
func (enumerator FileEnumerator) NextFile(cancellable Cancellable) (FileInfo, error) {
	var err glib.Error
	ret0 := C.g_file_enumerator_next_file(enumerator.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// NextFilesAsync is a wrapper around g_file_enumerator_next_files_async().
func (enumerator FileEnumerator) NextFilesAsync(num_files int, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_enumerator_next_files_async(enumerator.native(), C.int(num_files), C.int(io_priority), cancellable.native(), callback0)
}

// NextFilesFinish is a wrapper around g_file_enumerator_next_files_finish().
func (enumerator FileEnumerator) NextFilesFinish(result AsyncResult) (glib.List, error) {
	var err glib.Error
	ret0 := C.g_file_enumerator_next_files_finish(enumerator.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return glib.List{}, err.GoValue()
	}
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapFileInfo(p) }), /*gir:GLib*/ nil
}

// SetPending is a wrapper around g_file_enumerator_set_pending().
func (enumerator FileEnumerator) SetPending(pending bool) {
	C.g_file_enumerator_set_pending(enumerator.native(), C.gboolean(util.Bool2Int(pending)) /*go:.util*/)
}

// Interface Mount
type Mount struct {
	Ptr unsafe.Pointer
}

func (v Mount) native() *C.GMount {
	return (*C.GMount)(v.Ptr)
}
func wrapMount(p *C.GMount) Mount {
	return Mount{unsafe.Pointer(p)}
}
func WrapMount(p unsafe.Pointer) Mount {
	return Mount{p}
}
func (v Mount) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMount(p unsafe.Pointer) interface{} {
	return WrapMount(p)
}
func (v Mount) GetType() gobject.Type {
	return gobject.Type(C.g_mount_get_type())
}
func (v Mount) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMount(unsafe.Pointer(ptr)), nil
	}
}

// CanEject is a wrapper around g_mount_can_eject().
func (mount Mount) CanEject() bool {
	ret0 := C.g_mount_can_eject(mount.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// CanUnmount is a wrapper around g_mount_can_unmount().
func (mount Mount) CanUnmount() bool {
	ret0 := C.g_mount_can_unmount(mount.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// EjectWithOperation is a wrapper around g_mount_eject_with_operation().
func (mount Mount) EjectWithOperation(flags MountUnmountFlags, mount_operation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_mount_eject_with_operation(mount.native(), C.GMountUnmountFlags(flags), mount_operation.native(), cancellable.native(), callback0)
}

// EjectWithOperationFinish is a wrapper around g_mount_eject_with_operation_finish().
func (mount Mount) EjectWithOperationFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_mount_eject_with_operation_finish(mount.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// GetDefaultLocation is a wrapper around g_mount_get_default_location().
func (mount Mount) GetDefaultLocation() File {
	ret0 := C.g_mount_get_default_location(mount.native())
	return wrapFile(ret0)
}

// GetDrive is a wrapper around g_mount_get_drive().
func (mount Mount) GetDrive() Drive {
	ret0 := C.g_mount_get_drive(mount.native())
	return wrapDrive(ret0)
}

// GetIcon is a wrapper around g_mount_get_icon().
func (mount Mount) GetIcon() Icon {
	ret0 := C.g_mount_get_icon(mount.native())
	return wrapIcon(ret0)
}

// GetName is a wrapper around g_mount_get_name().
func (mount Mount) GetName() string {
	ret0 := C.g_mount_get_name(mount.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetRoot is a wrapper around g_mount_get_root().
func (mount Mount) GetRoot() File {
	ret0 := C.g_mount_get_root(mount.native())
	return wrapFile(ret0)
}

// GetSortKey is a wrapper around g_mount_get_sort_key().
func (mount Mount) GetSortKey() string {
	ret0 := C.g_mount_get_sort_key(mount.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetSymbolicIcon is a wrapper around g_mount_get_symbolic_icon().
func (mount Mount) GetSymbolicIcon() Icon {
	ret0 := C.g_mount_get_symbolic_icon(mount.native())
	return wrapIcon(ret0)
}

// GetUuid is a wrapper around g_mount_get_uuid().
func (mount Mount) GetUuid() string {
	ret0 := C.g_mount_get_uuid(mount.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetVolume is a wrapper around g_mount_get_volume().
func (mount Mount) GetVolume() Volume {
	ret0 := C.g_mount_get_volume(mount.native())
	return wrapVolume(ret0)
}

// GuessContentType is a wrapper around g_mount_guess_content_type().
func (mount Mount) GuessContentType(force_rescan bool, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_mount_guess_content_type(mount.native(), C.gboolean(util.Bool2Int(force_rescan)) /*go:.util*/, cancellable.native(), callback0)
}

// GuessContentTypeFinish is a wrapper around g_mount_guess_content_type_finish().
func (mount Mount) GuessContentTypeFinish(result AsyncResult) ([]string, error) {
	var err glib.Error
	ret0 := C.g_mount_guess_content_type_finish(mount.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return nil, err.GoValue()
	}
	return ret, nil
}

// GuessContentTypeSync is a wrapper around g_mount_guess_content_type_sync().
func (mount Mount) GuessContentTypeSync(force_rescan bool, cancellable Cancellable) ([]string, error) {
	var err glib.Error
	ret0 := C.g_mount_guess_content_type_sync(mount.native(), C.gboolean(util.Bool2Int(force_rescan)) /*go:.util*/, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return nil, err.GoValue()
	}
	return ret, nil
}

// IsShadowed is a wrapper around g_mount_is_shadowed().
func (mount Mount) IsShadowed() bool {
	ret0 := C.g_mount_is_shadowed(mount.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Remount is a wrapper around g_mount_remount().
func (mount Mount) Remount(flags MountMountFlags, mount_operation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_mount_remount(mount.native(), C.GMountMountFlags(flags), mount_operation.native(), cancellable.native(), callback0)
}

// RemountFinish is a wrapper around g_mount_remount_finish().
func (mount Mount) RemountFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_mount_remount_finish(mount.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Shadow is a wrapper around g_mount_shadow().
func (mount Mount) Shadow() {
	C.g_mount_shadow(mount.native())
}

// UnmountWithOperation is a wrapper around g_mount_unmount_with_operation().
func (mount Mount) UnmountWithOperation(flags MountUnmountFlags, mount_operation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_mount_unmount_with_operation(mount.native(), C.GMountUnmountFlags(flags), mount_operation.native(), cancellable.native(), callback0)
}

// UnmountWithOperationFinish is a wrapper around g_mount_unmount_with_operation_finish().
func (mount Mount) UnmountWithOperationFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_mount_unmount_with_operation_finish(mount.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Unshadow is a wrapper around g_mount_unshadow().
func (mount Mount) Unshadow() {
	C.g_mount_unshadow(mount.native())
}

// Interface Drive
type Drive struct {
	Ptr unsafe.Pointer
}

func (v Drive) native() *C.GDrive {
	return (*C.GDrive)(v.Ptr)
}
func wrapDrive(p *C.GDrive) Drive {
	return Drive{unsafe.Pointer(p)}
}
func WrapDrive(p unsafe.Pointer) Drive {
	return Drive{p}
}
func (v Drive) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDrive(p unsafe.Pointer) interface{} {
	return WrapDrive(p)
}
func (v Drive) GetType() gobject.Type {
	return gobject.Type(C.g_drive_get_type())
}
func (v Drive) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDrive(unsafe.Pointer(ptr)), nil
	}
}

// CanEject is a wrapper around g_drive_can_eject().
func (drive Drive) CanEject() bool {
	ret0 := C.g_drive_can_eject(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// CanPollForMedia is a wrapper around g_drive_can_poll_for_media().
func (drive Drive) CanPollForMedia() bool {
	ret0 := C.g_drive_can_poll_for_media(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// CanStart is a wrapper around g_drive_can_start().
func (drive Drive) CanStart() bool {
	ret0 := C.g_drive_can_start(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// CanStartDegraded is a wrapper around g_drive_can_start_degraded().
func (drive Drive) CanStartDegraded() bool {
	ret0 := C.g_drive_can_start_degraded(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// CanStop is a wrapper around g_drive_can_stop().
func (drive Drive) CanStop() bool {
	ret0 := C.g_drive_can_stop(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// EjectWithOperation is a wrapper around g_drive_eject_with_operation().
func (drive Drive) EjectWithOperation(flags MountUnmountFlags, mount_operation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_drive_eject_with_operation(drive.native(), C.GMountUnmountFlags(flags), mount_operation.native(), cancellable.native(), callback0)
}

// EjectWithOperationFinish is a wrapper around g_drive_eject_with_operation_finish().
func (drive Drive) EjectWithOperationFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_drive_eject_with_operation_finish(drive.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// EnumerateIdentifiers is a wrapper around g_drive_enumerate_identifiers().
func (drive Drive) EnumerateIdentifiers() []string {
	ret0 := C.g_drive_enumerate_identifiers(drive.native())
	var ret0Slice []*C.char
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString(elem)
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetIcon is a wrapper around g_drive_get_icon().
func (drive Drive) GetIcon() Icon {
	ret0 := C.g_drive_get_icon(drive.native())
	return wrapIcon(ret0)
}

// GetIdentifier is a wrapper around g_drive_get_identifier().
func (drive Drive) GetIdentifier(kind string) string {
	kind0 := C.CString(kind)
	ret0 := C.g_drive_get_identifier(drive.native(), kind0)
	C.free(unsafe.Pointer(kind0)) /*ch:<stdlib.h>*/
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetName is a wrapper around g_drive_get_name().
func (drive Drive) GetName() string {
	ret0 := C.g_drive_get_name(drive.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetSortKey is a wrapper around g_drive_get_sort_key().
func (drive Drive) GetSortKey() string {
	ret0 := C.g_drive_get_sort_key(drive.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetStartStopType is a wrapper around g_drive_get_start_stop_type().
func (drive Drive) GetStartStopType() DriveStartStopType {
	ret0 := C.g_drive_get_start_stop_type(drive.native())
	return DriveStartStopType(ret0)
}

// GetSymbolicIcon is a wrapper around g_drive_get_symbolic_icon().
func (drive Drive) GetSymbolicIcon() Icon {
	ret0 := C.g_drive_get_symbolic_icon(drive.native())
	return wrapIcon(ret0)
}

// GetVolumes is a wrapper around g_drive_get_volumes().
func (drive Drive) GetVolumes() glib.List {
	ret0 := C.g_drive_get_volumes(drive.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapVolume(p) }) /*gir:GLib*/
}

// HasMedia is a wrapper around g_drive_has_media().
func (drive Drive) HasMedia() bool {
	ret0 := C.g_drive_has_media(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// HasVolumes is a wrapper around g_drive_has_volumes().
func (drive Drive) HasVolumes() bool {
	ret0 := C.g_drive_has_volumes(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsMediaCheckAutomatic is a wrapper around g_drive_is_media_check_automatic().
func (drive Drive) IsMediaCheckAutomatic() bool {
	ret0 := C.g_drive_is_media_check_automatic(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsMediaRemovable is a wrapper around g_drive_is_media_removable().
func (drive Drive) IsMediaRemovable() bool {
	ret0 := C.g_drive_is_media_removable(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsRemovable is a wrapper around g_drive_is_removable().
func (drive Drive) IsRemovable() bool {
	ret0 := C.g_drive_is_removable(drive.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// PollForMedia is a wrapper around g_drive_poll_for_media().
func (drive Drive) PollForMedia(cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_drive_poll_for_media(drive.native(), cancellable.native(), callback0)
}

// PollForMediaFinish is a wrapper around g_drive_poll_for_media_finish().
func (drive Drive) PollForMediaFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_drive_poll_for_media_finish(drive.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Start is a wrapper around g_drive_start().
func (drive Drive) Start(flags DriveStartFlags, mount_operation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_drive_start(drive.native(), C.GDriveStartFlags(flags), mount_operation.native(), cancellable.native(), callback0)
}

// StartFinish is a wrapper around g_drive_start_finish().
func (drive Drive) StartFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_drive_start_finish(drive.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Stop is a wrapper around g_drive_stop().
func (drive Drive) Stop(flags MountUnmountFlags, mount_operation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_drive_stop(drive.native(), C.GMountUnmountFlags(flags), mount_operation.native(), cancellable.native(), callback0)
}

// StopFinish is a wrapper around g_drive_stop_finish().
func (drive Drive) StopFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_drive_stop_finish(drive.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Interface Volume
type Volume struct {
	Ptr unsafe.Pointer
}

func (v Volume) native() *C.GVolume {
	return (*C.GVolume)(v.Ptr)
}
func wrapVolume(p *C.GVolume) Volume {
	return Volume{unsafe.Pointer(p)}
}
func WrapVolume(p unsafe.Pointer) Volume {
	return Volume{p}
}
func (v Volume) IsNil() bool {
	return v.Ptr == nil
}
func IWrapVolume(p unsafe.Pointer) interface{} {
	return WrapVolume(p)
}
func (v Volume) GetType() gobject.Type {
	return gobject.Type(C.g_volume_get_type())
}
func (v Volume) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapVolume(unsafe.Pointer(ptr)), nil
	}
}

// CanEject is a wrapper around g_volume_can_eject().
func (volume Volume) CanEject() bool {
	ret0 := C.g_volume_can_eject(volume.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// CanMount is a wrapper around g_volume_can_mount().
func (volume Volume) CanMount() bool {
	ret0 := C.g_volume_can_mount(volume.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// EjectWithOperation is a wrapper around g_volume_eject_with_operation().
func (volume Volume) EjectWithOperation(flags MountUnmountFlags, mount_operation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_volume_eject_with_operation(volume.native(), C.GMountUnmountFlags(flags), mount_operation.native(), cancellable.native(), callback0)
}

// EjectWithOperationFinish is a wrapper around g_volume_eject_with_operation_finish().
func (volume Volume) EjectWithOperationFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_volume_eject_with_operation_finish(volume.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// EnumerateIdentifiers is a wrapper around g_volume_enumerate_identifiers().
func (volume Volume) EnumerateIdentifiers() []string {
	ret0 := C.g_volume_enumerate_identifiers(volume.native())
	var ret0Slice []*C.char
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString(elem)
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetActivationRoot is a wrapper around g_volume_get_activation_root().
func (volume Volume) GetActivationRoot() File {
	ret0 := C.g_volume_get_activation_root(volume.native())
	return wrapFile(ret0)
}

// GetDrive is a wrapper around g_volume_get_drive().
func (volume Volume) GetDrive() Drive {
	ret0 := C.g_volume_get_drive(volume.native())
	return wrapDrive(ret0)
}

// GetIcon is a wrapper around g_volume_get_icon().
func (volume Volume) GetIcon() Icon {
	ret0 := C.g_volume_get_icon(volume.native())
	return wrapIcon(ret0)
}

// GetIdentifier is a wrapper around g_volume_get_identifier().
func (volume Volume) GetIdentifier(kind string) string {
	kind0 := C.CString(kind)
	ret0 := C.g_volume_get_identifier(volume.native(), kind0)
	C.free(unsafe.Pointer(kind0)) /*ch:<stdlib.h>*/
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetMount is a wrapper around g_volume_get_mount().
func (volume Volume) GetMount() Mount {
	ret0 := C.g_volume_get_mount(volume.native())
	return wrapMount(ret0)
}

// GetName is a wrapper around g_volume_get_name().
func (volume Volume) GetName() string {
	ret0 := C.g_volume_get_name(volume.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetSortKey is a wrapper around g_volume_get_sort_key().
func (volume Volume) GetSortKey() string {
	ret0 := C.g_volume_get_sort_key(volume.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetSymbolicIcon is a wrapper around g_volume_get_symbolic_icon().
func (volume Volume) GetSymbolicIcon() Icon {
	ret0 := C.g_volume_get_symbolic_icon(volume.native())
	return wrapIcon(ret0)
}

// GetUuid is a wrapper around g_volume_get_uuid().
func (volume Volume) GetUuid() string {
	ret0 := C.g_volume_get_uuid(volume.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// Mount is a wrapper around g_volume_mount().
func (volume Volume) Mount(flags MountMountFlags, mount_operation MountOperation, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_volume_mount(volume.native(), C.GMountMountFlags(flags), mount_operation.native(), cancellable.native(), callback0)
}

// MountFinish is a wrapper around g_volume_mount_finish().
func (volume Volume) MountFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_volume_mount_finish(volume.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// ShouldAutomount is a wrapper around g_volume_should_automount().
func (volume Volume) ShouldAutomount() bool {
	ret0 := C.g_volume_should_automount(volume.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object FileMonitor
type FileMonitor struct {
	gobject.Object
}

func (v FileMonitor) native() *C.GFileMonitor {
	return (*C.GFileMonitor)(v.Ptr)
}
func wrapFileMonitor(p *C.GFileMonitor) (v FileMonitor) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileMonitor(p unsafe.Pointer) (v FileMonitor) {
	v.Ptr = p
	return
}
func (v FileMonitor) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFileMonitor(p unsafe.Pointer) interface{} {
	return WrapFileMonitor(p)
}
func (v FileMonitor) GetType() gobject.Type {
	return gobject.Type(C.g_file_monitor_get_type())
}
func (v FileMonitor) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFileMonitor(unsafe.Pointer(ptr)), nil
	}
}

// Cancel is a wrapper around g_file_monitor_cancel().
func (monitor FileMonitor) Cancel() bool {
	ret0 := C.g_file_monitor_cancel(monitor.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// EmitEvent is a wrapper around g_file_monitor_emit_event().
func (monitor FileMonitor) EmitEvent(child File, other_file File, event_type FileMonitorEvent) {
	C.g_file_monitor_emit_event(monitor.native(), child.native(), other_file.native(), C.GFileMonitorEvent(event_type))
}

// IsCancelled is a wrapper around g_file_monitor_is_cancelled().
func (monitor FileMonitor) IsCancelled() bool {
	ret0 := C.g_file_monitor_is_cancelled(monitor.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetRateLimit is a wrapper around g_file_monitor_set_rate_limit().
func (monitor FileMonitor) SetRateLimit(limit_msecs int) {
	C.g_file_monitor_set_rate_limit(monitor.native(), C.gint(limit_msecs))
}

// Struct FileAttributeInfoList
type FileAttributeInfoList struct {
	Ptr unsafe.Pointer
}

func (v FileAttributeInfoList) native() *C.GFileAttributeInfoList {
	return (*C.GFileAttributeInfoList)(v.Ptr)
}
func wrapFileAttributeInfoList(p *C.GFileAttributeInfoList) FileAttributeInfoList {
	return FileAttributeInfoList{unsafe.Pointer(p)}
}
func WrapFileAttributeInfoList(p unsafe.Pointer) FileAttributeInfoList {
	return FileAttributeInfoList{p}
}
func (v FileAttributeInfoList) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFileAttributeInfoList(p unsafe.Pointer) interface{} {
	return WrapFileAttributeInfoList(p)
}

// FileAttributeInfoListNew is a wrapper around g_file_attribute_info_list_new().
func FileAttributeInfoListNew() FileAttributeInfoList {
	ret0 := C.g_file_attribute_info_list_new()
	return wrapFileAttributeInfoList(ret0)
}

// Add is a wrapper around g_file_attribute_info_list_add().
func (list FileAttributeInfoList) Add(name string, type_ FileAttributeType, flags FileAttributeInfoFlags) {
	name0 := C.CString(name)
	C.g_file_attribute_info_list_add(list.native(), name0, C.GFileAttributeType(type_), C.GFileAttributeInfoFlags(flags))
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
}

// Dup is a wrapper around g_file_attribute_info_list_dup().
func (list FileAttributeInfoList) Dup() FileAttributeInfoList {
	ret0 := C.g_file_attribute_info_list_dup(list.native())
	return wrapFileAttributeInfoList(ret0)
}

// Lookup is a wrapper around g_file_attribute_info_list_lookup().
func (list FileAttributeInfoList) Lookup(name string) FileAttributeInfo {
	name0 := C.CString(name)
	ret0 := C.g_file_attribute_info_list_lookup(list.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapFileAttributeInfo(ret0)
}

// Ref is a wrapper around g_file_attribute_info_list_ref().
func (list FileAttributeInfoList) Ref() FileAttributeInfoList {
	ret0 := C.g_file_attribute_info_list_ref(list.native())
	return wrapFileAttributeInfoList(ret0)
}

// Unref is a wrapper around g_file_attribute_info_list_unref().
func (list FileAttributeInfoList) Unref() {
	C.g_file_attribute_info_list_unref(list.native())
}

// Struct FileAttributeInfo
type FileAttributeInfo struct {
	Ptr unsafe.Pointer
}

func (v FileAttributeInfo) native() *C.GFileAttributeInfo {
	return (*C.GFileAttributeInfo)(v.Ptr)
}
func wrapFileAttributeInfo(p *C.GFileAttributeInfo) FileAttributeInfo {
	return FileAttributeInfo{unsafe.Pointer(p)}
}
func WrapFileAttributeInfo(p unsafe.Pointer) FileAttributeInfo {
	return FileAttributeInfo{p}
}
func (v FileAttributeInfo) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFileAttributeInfo(p unsafe.Pointer) interface{} {
	return WrapFileAttributeInfo(p)
}

// Object FileInputStream
type FileInputStream struct {
	InputStream
}

func (v FileInputStream) native() *C.GFileInputStream {
	return (*C.GFileInputStream)(v.Ptr)
}
func wrapFileInputStream(p *C.GFileInputStream) (v FileInputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileInputStream(p unsafe.Pointer) (v FileInputStream) {
	v.Ptr = p
	return
}
func (v FileInputStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFileInputStream(p unsafe.Pointer) interface{} {
	return WrapFileInputStream(p)
}
func (v FileInputStream) GetType() gobject.Type {
	return gobject.Type(C.g_file_input_stream_get_type())
}
func (v FileInputStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFileInputStream(unsafe.Pointer(ptr)), nil
	}
}
func (v FileInputStream) Seekable() Seekable {
	return WrapSeekable(v.Ptr)
}

// QueryInfo is a wrapper around g_file_input_stream_query_info().
func (stream FileInputStream) QueryInfo(attributes string, cancellable Cancellable) (FileInfo, error) {
	attributes0 := C.CString(attributes)
	var err glib.Error
	ret0 := C.g_file_input_stream_query_info(stream.native(), attributes0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// QueryInfoAsync is a wrapper around g_file_input_stream_query_info_async().
func (stream FileInputStream) QueryInfoAsync(attributes string, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	attributes0 := C.CString(attributes)
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_file_input_stream_query_info_async(stream.native(), attributes0, C.int(io_priority), cancellable.native(), callback0)
	C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/
}

// QueryInfoFinish is a wrapper around g_file_input_stream_query_info_finish().
func (stream FileInputStream) QueryInfoFinish(result AsyncResult) (FileInfo, error) {
	var err glib.Error
	ret0 := C.g_file_input_stream_query_info_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return FileInfo{}, err.GoValue()
	}
	return wrapFileInfo(ret0), nil
}

// Struct SettingsSchemaKey
type SettingsSchemaKey struct {
	Ptr unsafe.Pointer
}

func (v SettingsSchemaKey) native() *C.GSettingsSchemaKey {
	return (*C.GSettingsSchemaKey)(v.Ptr)
}
func wrapSettingsSchemaKey(p *C.GSettingsSchemaKey) SettingsSchemaKey {
	return SettingsSchemaKey{unsafe.Pointer(p)}
}
func WrapSettingsSchemaKey(p unsafe.Pointer) SettingsSchemaKey {
	return SettingsSchemaKey{p}
}
func (v SettingsSchemaKey) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSettingsSchemaKey(p unsafe.Pointer) interface{} {
	return WrapSettingsSchemaKey(p)
}

// GetDefaultValue is a wrapper around g_settings_schema_key_get_default_value().
func (key SettingsSchemaKey) GetDefaultValue() glib.Variant {
	ret0 := C.g_settings_schema_key_get_default_value(key.native())
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetDescription is a wrapper around g_settings_schema_key_get_description().
func (key SettingsSchemaKey) GetDescription() string {
	ret0 := C.g_settings_schema_key_get_description(key.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetName is a wrapper around g_settings_schema_key_get_name().
func (key SettingsSchemaKey) GetName() string {
	ret0 := C.g_settings_schema_key_get_name(key.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetRange is a wrapper around g_settings_schema_key_get_range().
func (key SettingsSchemaKey) GetRange() glib.Variant {
	ret0 := C.g_settings_schema_key_get_range(key.native())
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetSummary is a wrapper around g_settings_schema_key_get_summary().
func (key SettingsSchemaKey) GetSummary() string {
	ret0 := C.g_settings_schema_key_get_summary(key.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetValueType is a wrapper around g_settings_schema_key_get_value_type().
func (key SettingsSchemaKey) GetValueType() glib.VariantType {
	ret0 := C.g_settings_schema_key_get_value_type(key.native())
	return glib.WrapVariantType(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// RangeCheck is a wrapper around g_settings_schema_key_range_check().
func (key SettingsSchemaKey) RangeCheck(value glib.Variant) bool {
	ret0 := C.g_settings_schema_key_range_check(key.native(), (*C.GVariant)(value.Ptr))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Ref is a wrapper around g_settings_schema_key_ref().
func (key SettingsSchemaKey) Ref() SettingsSchemaKey {
	ret0 := C.g_settings_schema_key_ref(key.native())
	return wrapSettingsSchemaKey(ret0)
}

// Unref is a wrapper around g_settings_schema_key_unref().
func (key SettingsSchemaKey) Unref() {
	C.g_settings_schema_key_unref(key.native())
}

// Struct FileAttributeMatcher
type FileAttributeMatcher struct {
	Ptr unsafe.Pointer
}

func (v FileAttributeMatcher) native() *C.GFileAttributeMatcher {
	return (*C.GFileAttributeMatcher)(v.Ptr)
}
func wrapFileAttributeMatcher(p *C.GFileAttributeMatcher) FileAttributeMatcher {
	return FileAttributeMatcher{unsafe.Pointer(p)}
}
func WrapFileAttributeMatcher(p unsafe.Pointer) FileAttributeMatcher {
	return FileAttributeMatcher{p}
}
func (v FileAttributeMatcher) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFileAttributeMatcher(p unsafe.Pointer) interface{} {
	return WrapFileAttributeMatcher(p)
}

// FileAttributeMatcherNew is a wrapper around g_file_attribute_matcher_new().
func FileAttributeMatcherNew(attributes string) FileAttributeMatcher {
	attributes0 := C.CString(attributes)
	ret0 := C.g_file_attribute_matcher_new(attributes0)
	C.free(unsafe.Pointer(attributes0)) /*ch:<stdlib.h>*/
	return wrapFileAttributeMatcher(ret0)
}

// EnumerateNamespace is a wrapper around g_file_attribute_matcher_enumerate_namespace().
func (matcher FileAttributeMatcher) EnumerateNamespace(ns string) bool {
	ns0 := C.CString(ns)
	ret0 := C.g_file_attribute_matcher_enumerate_namespace(matcher.native(), ns0)
	C.free(unsafe.Pointer(ns0))     /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// EnumerateNext is a wrapper around g_file_attribute_matcher_enumerate_next().
func (matcher FileAttributeMatcher) EnumerateNext() string {
	ret0 := C.g_file_attribute_matcher_enumerate_next(matcher.native())
	ret := C.GoString(ret0)
	return ret
}

// Matches is a wrapper around g_file_attribute_matcher_matches().
func (matcher FileAttributeMatcher) Matches(attribute string) bool {
	attribute0 := C.CString(attribute)
	ret0 := C.g_file_attribute_matcher_matches(matcher.native(), attribute0)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))    /*go:.util*/
}

// MatchesOnly is a wrapper around g_file_attribute_matcher_matches_only().
func (matcher FileAttributeMatcher) MatchesOnly(attribute string) bool {
	attribute0 := C.CString(attribute)
	ret0 := C.g_file_attribute_matcher_matches_only(matcher.native(), attribute0)
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))    /*go:.util*/
}

// Ref is a wrapper around g_file_attribute_matcher_ref().
func (matcher FileAttributeMatcher) Ref() FileAttributeMatcher {
	ret0 := C.g_file_attribute_matcher_ref(matcher.native())
	return wrapFileAttributeMatcher(ret0)
}

// Subtract is a wrapper around g_file_attribute_matcher_subtract().
func (matcher FileAttributeMatcher) Subtract(subtract FileAttributeMatcher) FileAttributeMatcher {
	ret0 := C.g_file_attribute_matcher_subtract(matcher.native(), subtract.native())
	return wrapFileAttributeMatcher(ret0)
}

// ToString is a wrapper around g_file_attribute_matcher_to_string().
func (matcher FileAttributeMatcher) ToString() string {
	ret0 := C.g_file_attribute_matcher_to_string(matcher.native())
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// Unref is a wrapper around g_file_attribute_matcher_unref().
func (matcher FileAttributeMatcher) Unref() {
	C.g_file_attribute_matcher_unref(matcher.native())
}

// Object DBusConnection
type DBusConnection struct {
	gobject.Object
}

func (v DBusConnection) native() *C.GDBusConnection {
	return (*C.GDBusConnection)(v.Ptr)
}
func wrapDBusConnection(p *C.GDBusConnection) (v DBusConnection) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDBusConnection(p unsafe.Pointer) (v DBusConnection) {
	v.Ptr = p
	return
}
func (v DBusConnection) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusConnection(p unsafe.Pointer) interface{} {
	return WrapDBusConnection(p)
}
func (v DBusConnection) GetType() gobject.Type {
	return gobject.Type(C.g_dbus_connection_get_type())
}
func (v DBusConnection) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDBusConnection(unsafe.Pointer(ptr)), nil
	}
}
func (v DBusConnection) AsyncInitable() AsyncInitable {
	return WrapAsyncInitable(v.Ptr)
}
func (v DBusConnection) Initable() Initable {
	return WrapInitable(v.Ptr)
}

// DBusConnectionNewForAddressSync is a wrapper around g_dbus_connection_new_for_address_sync().
func DBusConnectionNewForAddressSync(address string, flags DBusConnectionFlags, observer DBusAuthObserver, cancellable Cancellable) (DBusConnection, error) {
	address0 := (*C.gchar)(C.CString(address))
	var err glib.Error
	ret0 := C.g_dbus_connection_new_for_address_sync(address0, C.GDBusConnectionFlags(flags), observer.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(address0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return DBusConnection{}, err.GoValue()
	}
	return wrapDBusConnection(ret0), nil
}

// DBusConnectionNewSync is a wrapper around g_dbus_connection_new_sync().
func DBusConnectionNewSync(stream IOStream, guid string, flags DBusConnectionFlags, observer DBusAuthObserver, cancellable Cancellable) (DBusConnection, error) {
	guid0 := (*C.gchar)(C.CString(guid))
	var err glib.Error
	ret0 := C.g_dbus_connection_new_sync(stream.native(), guid0, C.GDBusConnectionFlags(flags), observer.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(guid0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return DBusConnection{}, err.GoValue()
	}
	return wrapDBusConnection(ret0), nil
}

// Call is a wrapper around g_dbus_connection_call().
func (connection DBusConnection) Call(bus_name string, object_path string, interface_name string, method_name string, parameters glib.Variant, reply_type glib.VariantType, flags DBusCallFlags, timeout_msec int, cancellable Cancellable, callback AsyncReadyCallback) {
	bus_name0 := (*C.gchar)(C.CString(bus_name))
	object_path0 := (*C.gchar)(C.CString(object_path))
	interface_name0 := (*C.gchar)(C.CString(interface_name))
	method_name0 := (*C.gchar)(C.CString(method_name))
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_dbus_connection_call(connection.native(), bus_name0, object_path0, interface_name0, method_name0, (*C.GVariant)(parameters.Ptr), (*C.GVariantType)(reply_type.Ptr), C.GDBusCallFlags(flags), C.gint(timeout_msec), cancellable.native(), callback0)
	C.free(unsafe.Pointer(bus_name0))       /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(object_path0))    /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(interface_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(method_name0))    /*ch:<stdlib.h>*/
}

// CallFinish is a wrapper around g_dbus_connection_call_finish().
func (connection DBusConnection) CallFinish(res AsyncResult) (glib.Variant, error) {
	var err glib.Error
	ret0 := C.g_dbus_connection_call_finish(connection.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return glib.Variant{}, err.GoValue()
	}
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/, nil
}

// CallSync is a wrapper around g_dbus_connection_call_sync().
func (connection DBusConnection) CallSync(bus_name string, object_path string, interface_name string, method_name string, parameters glib.Variant, reply_type glib.VariantType, flags DBusCallFlags, timeout_msec int, cancellable Cancellable) (glib.Variant, error) {
	bus_name0 := (*C.gchar)(C.CString(bus_name))
	object_path0 := (*C.gchar)(C.CString(object_path))
	interface_name0 := (*C.gchar)(C.CString(interface_name))
	method_name0 := (*C.gchar)(C.CString(method_name))
	var err glib.Error
	ret0 := C.g_dbus_connection_call_sync(connection.native(), bus_name0, object_path0, interface_name0, method_name0, (*C.GVariant)(parameters.Ptr), (*C.GVariantType)(reply_type.Ptr), C.GDBusCallFlags(flags), C.gint(timeout_msec), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(bus_name0))       /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(object_path0))    /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(interface_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(method_name0))    /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return glib.Variant{}, err.GoValue()
	}
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/, nil
}

// CallWithUnixFdList is a wrapper around g_dbus_connection_call_with_unix_fd_list().
func (connection DBusConnection) CallWithUnixFdList(bus_name string, object_path string, interface_name string, method_name string, parameters glib.Variant, reply_type glib.VariantType, flags DBusCallFlags, timeout_msec int, fd_list UnixFDList, cancellable Cancellable, callback AsyncReadyCallback) {
	bus_name0 := (*C.gchar)(C.CString(bus_name))
	object_path0 := (*C.gchar)(C.CString(object_path))
	interface_name0 := (*C.gchar)(C.CString(interface_name))
	method_name0 := (*C.gchar)(C.CString(method_name))
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_dbus_connection_call_with_unix_fd_list(connection.native(), bus_name0, object_path0, interface_name0, method_name0, (*C.GVariant)(parameters.Ptr), (*C.GVariantType)(reply_type.Ptr), C.GDBusCallFlags(flags), C.gint(timeout_msec), fd_list.native(), cancellable.native(), callback0)
	C.free(unsafe.Pointer(bus_name0))       /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(object_path0))    /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(interface_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(method_name0))    /*ch:<stdlib.h>*/
}

// CallWithUnixFdListFinish is a wrapper around g_dbus_connection_call_with_unix_fd_list_finish().
func (connection DBusConnection) CallWithUnixFdListFinish(res AsyncResult) (glib.Variant, UnixFDList, error) {
	var out_fd_list0 *C.GUnixFDList
	var err glib.Error
	ret0 := C.g_dbus_connection_call_with_unix_fd_list_finish(connection.native(), &out_fd_list0, res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return glib.Variant{}, UnixFDList{}, err.GoValue()
	}
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/, wrapUnixFDList(out_fd_list0), nil
}

// CallWithUnixFdListSync is a wrapper around g_dbus_connection_call_with_unix_fd_list_sync().
func (connection DBusConnection) CallWithUnixFdListSync(bus_name string, object_path string, interface_name string, method_name string, parameters glib.Variant, reply_type glib.VariantType, flags DBusCallFlags, timeout_msec int, fd_list UnixFDList, cancellable Cancellable) (glib.Variant, UnixFDList, error) {
	bus_name0 := (*C.gchar)(C.CString(bus_name))
	object_path0 := (*C.gchar)(C.CString(object_path))
	interface_name0 := (*C.gchar)(C.CString(interface_name))
	method_name0 := (*C.gchar)(C.CString(method_name))
	var out_fd_list0 *C.GUnixFDList
	var err glib.Error
	ret0 := C.g_dbus_connection_call_with_unix_fd_list_sync(connection.native(), bus_name0, object_path0, interface_name0, method_name0, (*C.GVariant)(parameters.Ptr), (*C.GVariantType)(reply_type.Ptr), C.GDBusCallFlags(flags), C.gint(timeout_msec), fd_list.native(), &out_fd_list0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(bus_name0))       /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(object_path0))    /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(interface_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(method_name0))    /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return glib.Variant{}, UnixFDList{}, err.GoValue()
	}
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/, wrapUnixFDList(out_fd_list0), nil
}

// Close is a wrapper around g_dbus_connection_close().
func (connection DBusConnection) Close(cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_dbus_connection_close(connection.native(), cancellable.native(), callback0)
}

// CloseFinish is a wrapper around g_dbus_connection_close_finish().
func (connection DBusConnection) CloseFinish(res AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_dbus_connection_close_finish(connection.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// CloseSync is a wrapper around g_dbus_connection_close_sync().
func (connection DBusConnection) CloseSync(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_dbus_connection_close_sync(connection.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// EmitSignal is a wrapper around g_dbus_connection_emit_signal().
func (connection DBusConnection) EmitSignal(destination_bus_name string, object_path string, interface_name string, signal_name string, parameters glib.Variant) (bool, error) {
	destination_bus_name0 := (*C.gchar)(C.CString(destination_bus_name))
	object_path0 := (*C.gchar)(C.CString(object_path))
	interface_name0 := (*C.gchar)(C.CString(interface_name))
	signal_name0 := (*C.gchar)(C.CString(signal_name))
	var err glib.Error
	ret0 := C.g_dbus_connection_emit_signal(connection.native(), destination_bus_name0, object_path0, interface_name0, signal_name0, (*C.GVariant)(parameters.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(destination_bus_name0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(object_path0))          /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(interface_name0))       /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(signal_name0))          /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// ExportActionGroup is a wrapper around g_dbus_connection_export_action_group().
func (connection DBusConnection) ExportActionGroup(object_path string, action_group ActionGroup) (uint, error) {
	object_path0 := (*C.gchar)(C.CString(object_path))
	var err glib.Error
	ret0 := C.g_dbus_connection_export_action_group(connection.native(), object_path0, action_group.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(object_path0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return uint(ret0), nil
}

// ExportMenuModel is a wrapper around g_dbus_connection_export_menu_model().
func (connection DBusConnection) ExportMenuModel(object_path string, menu MenuModel) (uint, error) {
	object_path0 := (*C.gchar)(C.CString(object_path))
	var err glib.Error
	ret0 := C.g_dbus_connection_export_menu_model(connection.native(), object_path0, menu.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(object_path0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return uint(ret0), nil
}

// Flush is a wrapper around g_dbus_connection_flush().
func (connection DBusConnection) Flush(cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_dbus_connection_flush(connection.native(), cancellable.native(), callback0)
}

// FlushFinish is a wrapper around g_dbus_connection_flush_finish().
func (connection DBusConnection) FlushFinish(res AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_dbus_connection_flush_finish(connection.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// FlushSync is a wrapper around g_dbus_connection_flush_sync().
func (connection DBusConnection) FlushSync(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_dbus_connection_flush_sync(connection.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// GetCapabilities is a wrapper around g_dbus_connection_get_capabilities().
func (connection DBusConnection) GetCapabilities() DBusCapabilityFlags {
	ret0 := C.g_dbus_connection_get_capabilities(connection.native())
	return DBusCapabilityFlags(ret0)
}

// GetExitOnClose is a wrapper around g_dbus_connection_get_exit_on_close().
func (connection DBusConnection) GetExitOnClose() bool {
	ret0 := C.g_dbus_connection_get_exit_on_close(connection.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetGuid is a wrapper around g_dbus_connection_get_guid().
func (connection DBusConnection) GetGuid() string {
	ret0 := C.g_dbus_connection_get_guid(connection.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetLastSerial is a wrapper around g_dbus_connection_get_last_serial().
func (connection DBusConnection) GetLastSerial() uint32 {
	ret0 := C.g_dbus_connection_get_last_serial(connection.native())
	return uint32(ret0)
}

// GetPeerCredentials is a wrapper around g_dbus_connection_get_peer_credentials().
func (connection DBusConnection) GetPeerCredentials() Credentials {
	ret0 := C.g_dbus_connection_get_peer_credentials(connection.native())
	return wrapCredentials(ret0)
}

// GetStream is a wrapper around g_dbus_connection_get_stream().
func (connection DBusConnection) GetStream() IOStream {
	ret0 := C.g_dbus_connection_get_stream(connection.native())
	return wrapIOStream(ret0)
}

// GetUniqueName is a wrapper around g_dbus_connection_get_unique_name().
func (connection DBusConnection) GetUniqueName() string {
	ret0 := C.g_dbus_connection_get_unique_name(connection.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// IsClosed is a wrapper around g_dbus_connection_is_closed().
func (connection DBusConnection) IsClosed() bool {
	ret0 := C.g_dbus_connection_is_closed(connection.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// RegisterObjectWithClosures is a wrapper around g_dbus_connection_register_object_with_closures().
func (connection DBusConnection) RegisterObjectWithClosures(object_path string, interface_info DBusInterfaceInfo, method_call_closure gobject.Closure, get_property_closure gobject.Closure, set_property_closure gobject.Closure) (uint, error) {
	object_path0 := (*C.gchar)(C.CString(object_path))
	var err glib.Error
	ret0 := C.g_dbus_connection_register_object_with_closures(connection.native(), object_path0, interface_info.native(), (*C.GClosure)(method_call_closure.Ptr), (*C.GClosure)(get_property_closure.Ptr), (*C.GClosure)(set_property_closure.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(object_path0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return uint(ret0), nil
}

// RemoveFilter is a wrapper around g_dbus_connection_remove_filter().
func (connection DBusConnection) RemoveFilter(filter_id uint) {
	C.g_dbus_connection_remove_filter(connection.native(), C.guint(filter_id))
}

// SendMessageWithReplyFinish is a wrapper around g_dbus_connection_send_message_with_reply_finish().
func (connection DBusConnection) SendMessageWithReplyFinish(res AsyncResult) (DBusMessage, error) {
	var err glib.Error
	ret0 := C.g_dbus_connection_send_message_with_reply_finish(connection.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return DBusMessage{}, err.GoValue()
	}
	return wrapDBusMessage(ret0), nil
}

// SetExitOnClose is a wrapper around g_dbus_connection_set_exit_on_close().
func (connection DBusConnection) SetExitOnClose(exit_on_close bool) {
	C.g_dbus_connection_set_exit_on_close(connection.native(), C.gboolean(util.Bool2Int(exit_on_close)) /*go:.util*/)
}

// SignalUnsubscribe is a wrapper around g_dbus_connection_signal_unsubscribe().
func (connection DBusConnection) SignalUnsubscribe(subscription_id uint) {
	C.g_dbus_connection_signal_unsubscribe(connection.native(), C.guint(subscription_id))
}

// StartMessageProcessing is a wrapper around g_dbus_connection_start_message_processing().
func (connection DBusConnection) StartMessageProcessing() {
	C.g_dbus_connection_start_message_processing(connection.native())
}

// UnexportActionGroup is a wrapper around g_dbus_connection_unexport_action_group().
func (connection DBusConnection) UnexportActionGroup(export_id uint) {
	C.g_dbus_connection_unexport_action_group(connection.native(), C.guint(export_id))
}

// UnexportMenuModel is a wrapper around g_dbus_connection_unexport_menu_model().
func (connection DBusConnection) UnexportMenuModel(export_id uint) {
	C.g_dbus_connection_unexport_menu_model(connection.native(), C.guint(export_id))
}

// UnregisterObject is a wrapper around g_dbus_connection_unregister_object().
func (connection DBusConnection) UnregisterObject(registration_id uint) bool {
	ret0 := C.g_dbus_connection_unregister_object(connection.native(), C.guint(registration_id))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// UnregisterSubtree is a wrapper around g_dbus_connection_unregister_subtree().
func (connection DBusConnection) UnregisterSubtree(registration_id uint) bool {
	ret0 := C.g_dbus_connection_unregister_subtree(connection.native(), C.guint(registration_id))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// DBusConnectionNew is a wrapper around g_dbus_connection_new().
func DBusConnectionNew(stream IOStream, guid string, flags DBusConnectionFlags, observer DBusAuthObserver, cancellable Cancellable, callback AsyncReadyCallback) {
	guid0 := (*C.gchar)(C.CString(guid))
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_dbus_connection_new(stream.native(), guid0, C.GDBusConnectionFlags(flags), observer.native(), cancellable.native(), callback0)
	C.free(unsafe.Pointer(guid0)) /*ch:<stdlib.h>*/
}

// DBusConnectionNewForAddress is a wrapper around g_dbus_connection_new_for_address().
func DBusConnectionNewForAddress(address string, flags DBusConnectionFlags, observer DBusAuthObserver, cancellable Cancellable, callback AsyncReadyCallback) {
	address0 := (*C.gchar)(C.CString(address))
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_dbus_connection_new_for_address(address0, C.GDBusConnectionFlags(flags), observer.native(), cancellable.native(), callback0)
	C.free(unsafe.Pointer(address0)) /*ch:<stdlib.h>*/
}

// Interface AsyncInitable
type AsyncInitable struct {
	Ptr unsafe.Pointer
}

func (v AsyncInitable) native() *C.GAsyncInitable {
	return (*C.GAsyncInitable)(v.Ptr)
}
func wrapAsyncInitable(p *C.GAsyncInitable) AsyncInitable {
	return AsyncInitable{unsafe.Pointer(p)}
}
func WrapAsyncInitable(p unsafe.Pointer) AsyncInitable {
	return AsyncInitable{p}
}
func (v AsyncInitable) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAsyncInitable(p unsafe.Pointer) interface{} {
	return WrapAsyncInitable(p)
}
func (v AsyncInitable) GetType() gobject.Type {
	return gobject.Type(C.g_async_initable_get_type())
}
func (v AsyncInitable) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAsyncInitable(unsafe.Pointer(ptr)), nil
	}
}

// InitAsync is a wrapper around g_async_initable_init_async().
func (initable AsyncInitable) InitAsync(io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_async_initable_init_async(initable.native(), C.int(io_priority), cancellable.native(), callback0)
}

// InitFinish is a wrapper around g_async_initable_init_finish().
func (initable AsyncInitable) InitFinish(res AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_async_initable_init_finish(initable.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// NewFinish is a wrapper around g_async_initable_new_finish().
func (initable AsyncInitable) NewFinish(res AsyncResult) (gobject.Object, error) {
	var err glib.Error
	ret0 := C.g_async_initable_new_finish(initable.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return gobject.Object{}, err.GoValue()
	}
	return gobject.WrapObject(unsafe.Pointer(ret0)) /*gir:GObject*/, nil
}

// Interface Initable
type Initable struct {
	Ptr unsafe.Pointer
}

func (v Initable) native() *C.GInitable {
	return (*C.GInitable)(v.Ptr)
}
func wrapInitable(p *C.GInitable) Initable {
	return Initable{unsafe.Pointer(p)}
}
func WrapInitable(p unsafe.Pointer) Initable {
	return Initable{p}
}
func (v Initable) IsNil() bool {
	return v.Ptr == nil
}
func IWrapInitable(p unsafe.Pointer) interface{} {
	return WrapInitable(p)
}
func (v Initable) GetType() gobject.Type {
	return gobject.Type(C.g_initable_get_type())
}
func (v Initable) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapInitable(unsafe.Pointer(ptr)), nil
	}
}

// Init is a wrapper around g_initable_init().
func (initable Initable) Init(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_initable_init(initable.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Object MountOperation
type MountOperation struct {
	gobject.Object
}

func (v MountOperation) native() *C.GMountOperation {
	return (*C.GMountOperation)(v.Ptr)
}
func wrapMountOperation(p *C.GMountOperation) (v MountOperation) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMountOperation(p unsafe.Pointer) (v MountOperation) {
	v.Ptr = p
	return
}
func (v MountOperation) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMountOperation(p unsafe.Pointer) interface{} {
	return WrapMountOperation(p)
}
func (v MountOperation) GetType() gobject.Type {
	return gobject.Type(C.g_mount_operation_get_type())
}
func (v MountOperation) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMountOperation(unsafe.Pointer(ptr)), nil
	}
}

// MountOperationNew is a wrapper around g_mount_operation_new().
func MountOperationNew() MountOperation {
	ret0 := C.g_mount_operation_new()
	return wrapMountOperation(ret0)
}

// GetAnonymous is a wrapper around g_mount_operation_get_anonymous().
func (op MountOperation) GetAnonymous() bool {
	ret0 := C.g_mount_operation_get_anonymous(op.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetChoice is a wrapper around g_mount_operation_get_choice().
func (op MountOperation) GetChoice() int {
	ret0 := C.g_mount_operation_get_choice(op.native())
	return int(ret0)
}

// GetDomain is a wrapper around g_mount_operation_get_domain().
func (op MountOperation) GetDomain() string {
	ret0 := C.g_mount_operation_get_domain(op.native())
	ret := C.GoString(ret0)
	return ret
}

// GetPassword is a wrapper around g_mount_operation_get_password().
func (op MountOperation) GetPassword() string {
	ret0 := C.g_mount_operation_get_password(op.native())
	ret := C.GoString(ret0)
	return ret
}

// GetPasswordSave is a wrapper around g_mount_operation_get_password_save().
func (op MountOperation) GetPasswordSave() PasswordSave {
	ret0 := C.g_mount_operation_get_password_save(op.native())
	return PasswordSave(ret0)
}

// GetUsername is a wrapper around g_mount_operation_get_username().
func (op MountOperation) GetUsername() string {
	ret0 := C.g_mount_operation_get_username(op.native())
	ret := C.GoString(ret0)
	return ret
}

// Reply is a wrapper around g_mount_operation_reply().
func (op MountOperation) Reply(result MountOperationResult) {
	C.g_mount_operation_reply(op.native(), C.GMountOperationResult(result))
}

// SetAnonymous is a wrapper around g_mount_operation_set_anonymous().
func (op MountOperation) SetAnonymous(anonymous bool) {
	C.g_mount_operation_set_anonymous(op.native(), C.gboolean(util.Bool2Int(anonymous)) /*go:.util*/)
}

// SetChoice is a wrapper around g_mount_operation_set_choice().
func (op MountOperation) SetChoice(choice int) {
	C.g_mount_operation_set_choice(op.native(), C.int(choice))
}

// SetDomain is a wrapper around g_mount_operation_set_domain().
func (op MountOperation) SetDomain(domain string) {
	domain0 := C.CString(domain)
	C.g_mount_operation_set_domain(op.native(), domain0)
	C.free(unsafe.Pointer(domain0)) /*ch:<stdlib.h>*/
}

// SetPassword is a wrapper around g_mount_operation_set_password().
func (op MountOperation) SetPassword(password string) {
	password0 := C.CString(password)
	C.g_mount_operation_set_password(op.native(), password0)
	C.free(unsafe.Pointer(password0)) /*ch:<stdlib.h>*/
}

// SetPasswordSave is a wrapper around g_mount_operation_set_password_save().
func (op MountOperation) SetPasswordSave(save PasswordSave) {
	C.g_mount_operation_set_password_save(op.native(), C.GPasswordSave(save))
}

// SetUsername is a wrapper around g_mount_operation_set_username().
func (op MountOperation) SetUsername(username string) {
	username0 := C.CString(username)
	C.g_mount_operation_set_username(op.native(), username0)
	C.free(unsafe.Pointer(username0)) /*ch:<stdlib.h>*/
}

// Struct ActionEntry
type ActionEntry struct {
	Ptr unsafe.Pointer
}

func (v ActionEntry) native() *C.GActionEntry {
	return (*C.GActionEntry)(v.Ptr)
}
func wrapActionEntry(p *C.GActionEntry) ActionEntry {
	return ActionEntry{unsafe.Pointer(p)}
}
func WrapActionEntry(p unsafe.Pointer) ActionEntry {
	return ActionEntry{p}
}
func (v ActionEntry) IsNil() bool {
	return v.Ptr == nil
}
func IWrapActionEntry(p unsafe.Pointer) interface{} {
	return WrapActionEntry(p)
}

// Struct DBusAnnotationInfo
type DBusAnnotationInfo struct {
	Ptr unsafe.Pointer
}

func (v DBusAnnotationInfo) native() *C.GDBusAnnotationInfo {
	return (*C.GDBusAnnotationInfo)(v.Ptr)
}
func wrapDBusAnnotationInfo(p *C.GDBusAnnotationInfo) DBusAnnotationInfo {
	return DBusAnnotationInfo{unsafe.Pointer(p)}
}
func WrapDBusAnnotationInfo(p unsafe.Pointer) DBusAnnotationInfo {
	return DBusAnnotationInfo{p}
}
func (v DBusAnnotationInfo) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusAnnotationInfo(p unsafe.Pointer) interface{} {
	return WrapDBusAnnotationInfo(p)
}

// Ref is a wrapper around g_dbus_annotation_info_ref().
func (info DBusAnnotationInfo) Ref() DBusAnnotationInfo {
	ret0 := C.g_dbus_annotation_info_ref(info.native())
	return wrapDBusAnnotationInfo(ret0)
}

// Unref is a wrapper around g_dbus_annotation_info_unref().
func (info DBusAnnotationInfo) Unref() {
	C.g_dbus_annotation_info_unref(info.native())
}

// DBusAnnotationInfoLookup is a wrapper around g_dbus_annotation_info_lookup().
func DBusAnnotationInfoLookup(annotations []DBusAnnotationInfo, name string) string {
	annotations0 := make([]*C.GDBusAnnotationInfo, len(annotations))
	for idx, elemG := range annotations {
		annotations0[idx] = elemG.native()
	}
	var annotations0Ptr **C.GDBusAnnotationInfo
	if len(annotations0) > 0 {
		annotations0Ptr = &annotations0[0]
	}
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_dbus_annotation_info_lookup(annotations0Ptr, name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// Struct DBusArgInfo
type DBusArgInfo struct {
	Ptr unsafe.Pointer
}

func (v DBusArgInfo) native() *C.GDBusArgInfo {
	return (*C.GDBusArgInfo)(v.Ptr)
}
func wrapDBusArgInfo(p *C.GDBusArgInfo) DBusArgInfo {
	return DBusArgInfo{unsafe.Pointer(p)}
}
func WrapDBusArgInfo(p unsafe.Pointer) DBusArgInfo {
	return DBusArgInfo{p}
}
func (v DBusArgInfo) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusArgInfo(p unsafe.Pointer) interface{} {
	return WrapDBusArgInfo(p)
}

// Ref is a wrapper around g_dbus_arg_info_ref().
func (info DBusArgInfo) Ref() DBusArgInfo {
	ret0 := C.g_dbus_arg_info_ref(info.native())
	return wrapDBusArgInfo(ret0)
}

// Unref is a wrapper around g_dbus_arg_info_unref().
func (info DBusArgInfo) Unref() {
	C.g_dbus_arg_info_unref(info.native())
}

// Struct DBusErrorEntry
type DBusErrorEntry struct {
	Ptr unsafe.Pointer
}

func (v DBusErrorEntry) native() *C.GDBusErrorEntry {
	return (*C.GDBusErrorEntry)(v.Ptr)
}
func wrapDBusErrorEntry(p *C.GDBusErrorEntry) DBusErrorEntry {
	return DBusErrorEntry{unsafe.Pointer(p)}
}
func WrapDBusErrorEntry(p unsafe.Pointer) DBusErrorEntry {
	return DBusErrorEntry{p}
}
func (v DBusErrorEntry) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusErrorEntry(p unsafe.Pointer) interface{} {
	return WrapDBusErrorEntry(p)
}

// Struct DBusInterfaceInfo
type DBusInterfaceInfo struct {
	Ptr unsafe.Pointer
}

func (v DBusInterfaceInfo) native() *C.GDBusInterfaceInfo {
	return (*C.GDBusInterfaceInfo)(v.Ptr)
}
func wrapDBusInterfaceInfo(p *C.GDBusInterfaceInfo) DBusInterfaceInfo {
	return DBusInterfaceInfo{unsafe.Pointer(p)}
}
func WrapDBusInterfaceInfo(p unsafe.Pointer) DBusInterfaceInfo {
	return DBusInterfaceInfo{p}
}
func (v DBusInterfaceInfo) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusInterfaceInfo(p unsafe.Pointer) interface{} {
	return WrapDBusInterfaceInfo(p)
}

// CacheBuild is a wrapper around g_dbus_interface_info_cache_build().
func (info DBusInterfaceInfo) CacheBuild() {
	C.g_dbus_interface_info_cache_build(info.native())
}

// CacheRelease is a wrapper around g_dbus_interface_info_cache_release().
func (info DBusInterfaceInfo) CacheRelease() {
	C.g_dbus_interface_info_cache_release(info.native())
}

// LookupMethod is a wrapper around g_dbus_interface_info_lookup_method().
func (info DBusInterfaceInfo) LookupMethod(name string) DBusMethodInfo {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_dbus_interface_info_lookup_method(info.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapDBusMethodInfo(ret0)
}

// LookupProperty is a wrapper around g_dbus_interface_info_lookup_property().
func (info DBusInterfaceInfo) LookupProperty(name string) DBusPropertyInfo {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_dbus_interface_info_lookup_property(info.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapDBusPropertyInfo(ret0)
}

// LookupSignal is a wrapper around g_dbus_interface_info_lookup_signal().
func (info DBusInterfaceInfo) LookupSignal(name string) DBusSignalInfo {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_dbus_interface_info_lookup_signal(info.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapDBusSignalInfo(ret0)
}

// Ref is a wrapper around g_dbus_interface_info_ref().
func (info DBusInterfaceInfo) Ref() DBusInterfaceInfo {
	ret0 := C.g_dbus_interface_info_ref(info.native())
	return wrapDBusInterfaceInfo(ret0)
}

// Unref is a wrapper around g_dbus_interface_info_unref().
func (info DBusInterfaceInfo) Unref() {
	C.g_dbus_interface_info_unref(info.native())
}

// Struct DBusMethodInfo
type DBusMethodInfo struct {
	Ptr unsafe.Pointer
}

func (v DBusMethodInfo) native() *C.GDBusMethodInfo {
	return (*C.GDBusMethodInfo)(v.Ptr)
}
func wrapDBusMethodInfo(p *C.GDBusMethodInfo) DBusMethodInfo {
	return DBusMethodInfo{unsafe.Pointer(p)}
}
func WrapDBusMethodInfo(p unsafe.Pointer) DBusMethodInfo {
	return DBusMethodInfo{p}
}
func (v DBusMethodInfo) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusMethodInfo(p unsafe.Pointer) interface{} {
	return WrapDBusMethodInfo(p)
}

// Ref is a wrapper around g_dbus_method_info_ref().
func (info DBusMethodInfo) Ref() DBusMethodInfo {
	ret0 := C.g_dbus_method_info_ref(info.native())
	return wrapDBusMethodInfo(ret0)
}

// Unref is a wrapper around g_dbus_method_info_unref().
func (info DBusMethodInfo) Unref() {
	C.g_dbus_method_info_unref(info.native())
}

// Struct DBusPropertyInfo
type DBusPropertyInfo struct {
	Ptr unsafe.Pointer
}

func (v DBusPropertyInfo) native() *C.GDBusPropertyInfo {
	return (*C.GDBusPropertyInfo)(v.Ptr)
}
func wrapDBusPropertyInfo(p *C.GDBusPropertyInfo) DBusPropertyInfo {
	return DBusPropertyInfo{unsafe.Pointer(p)}
}
func WrapDBusPropertyInfo(p unsafe.Pointer) DBusPropertyInfo {
	return DBusPropertyInfo{p}
}
func (v DBusPropertyInfo) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusPropertyInfo(p unsafe.Pointer) interface{} {
	return WrapDBusPropertyInfo(p)
}

// Ref is a wrapper around g_dbus_property_info_ref().
func (info DBusPropertyInfo) Ref() DBusPropertyInfo {
	ret0 := C.g_dbus_property_info_ref(info.native())
	return wrapDBusPropertyInfo(ret0)
}

// Unref is a wrapper around g_dbus_property_info_unref().
func (info DBusPropertyInfo) Unref() {
	C.g_dbus_property_info_unref(info.native())
}

// Struct DBusInterfaceVTable
type DBusInterfaceVTable struct {
	Ptr unsafe.Pointer
}

func (v DBusInterfaceVTable) native() *C.GDBusInterfaceVTable {
	return (*C.GDBusInterfaceVTable)(v.Ptr)
}
func wrapDBusInterfaceVTable(p *C.GDBusInterfaceVTable) DBusInterfaceVTable {
	return DBusInterfaceVTable{unsafe.Pointer(p)}
}
func WrapDBusInterfaceVTable(p unsafe.Pointer) DBusInterfaceVTable {
	return DBusInterfaceVTable{p}
}
func (v DBusInterfaceVTable) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusInterfaceVTable(p unsafe.Pointer) interface{} {
	return WrapDBusInterfaceVTable(p)
}

// Struct DBusNodeInfo
type DBusNodeInfo struct {
	Ptr unsafe.Pointer
}

func (v DBusNodeInfo) native() *C.GDBusNodeInfo {
	return (*C.GDBusNodeInfo)(v.Ptr)
}
func wrapDBusNodeInfo(p *C.GDBusNodeInfo) DBusNodeInfo {
	return DBusNodeInfo{unsafe.Pointer(p)}
}
func WrapDBusNodeInfo(p unsafe.Pointer) DBusNodeInfo {
	return DBusNodeInfo{p}
}
func (v DBusNodeInfo) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusNodeInfo(p unsafe.Pointer) interface{} {
	return WrapDBusNodeInfo(p)
}

// DBusNodeInfoNewForXml is a wrapper around g_dbus_node_info_new_for_xml().
func DBusNodeInfoNewForXml(xml_data string) (DBusNodeInfo, error) {
	xml_data0 := (*C.gchar)(C.CString(xml_data))
	var err glib.Error
	ret0 := C.g_dbus_node_info_new_for_xml(xml_data0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(xml_data0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return DBusNodeInfo{}, err.GoValue()
	}
	return wrapDBusNodeInfo(ret0), nil
}

// LookupInterface is a wrapper around g_dbus_node_info_lookup_interface().
func (info DBusNodeInfo) LookupInterface(name string) DBusInterfaceInfo {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_dbus_node_info_lookup_interface(info.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapDBusInterfaceInfo(ret0)
}

// Ref is a wrapper around g_dbus_node_info_ref().
func (info DBusNodeInfo) Ref() DBusNodeInfo {
	ret0 := C.g_dbus_node_info_ref(info.native())
	return wrapDBusNodeInfo(ret0)
}

// Unref is a wrapper around g_dbus_node_info_unref().
func (info DBusNodeInfo) Unref() {
	C.g_dbus_node_info_unref(info.native())
}

// Struct DBusSignalInfo
type DBusSignalInfo struct {
	Ptr unsafe.Pointer
}

func (v DBusSignalInfo) native() *C.GDBusSignalInfo {
	return (*C.GDBusSignalInfo)(v.Ptr)
}
func wrapDBusSignalInfo(p *C.GDBusSignalInfo) DBusSignalInfo {
	return DBusSignalInfo{unsafe.Pointer(p)}
}
func WrapDBusSignalInfo(p unsafe.Pointer) DBusSignalInfo {
	return DBusSignalInfo{p}
}
func (v DBusSignalInfo) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusSignalInfo(p unsafe.Pointer) interface{} {
	return WrapDBusSignalInfo(p)
}

// Ref is a wrapper around g_dbus_signal_info_ref().
func (info DBusSignalInfo) Ref() DBusSignalInfo {
	ret0 := C.g_dbus_signal_info_ref(info.native())
	return wrapDBusSignalInfo(ret0)
}

// Unref is a wrapper around g_dbus_signal_info_unref().
func (info DBusSignalInfo) Unref() {
	C.g_dbus_signal_info_unref(info.native())
}

// Struct DBusSubtreeVTable
type DBusSubtreeVTable struct {
	Ptr unsafe.Pointer
}

func (v DBusSubtreeVTable) native() *C.GDBusSubtreeVTable {
	return (*C.GDBusSubtreeVTable)(v.Ptr)
}
func wrapDBusSubtreeVTable(p *C.GDBusSubtreeVTable) DBusSubtreeVTable {
	return DBusSubtreeVTable{unsafe.Pointer(p)}
}
func WrapDBusSubtreeVTable(p unsafe.Pointer) DBusSubtreeVTable {
	return DBusSubtreeVTable{p}
}
func (v DBusSubtreeVTable) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusSubtreeVTable(p unsafe.Pointer) interface{} {
	return WrapDBusSubtreeVTable(p)
}

// Struct InputMessage
type InputMessage struct {
	Ptr unsafe.Pointer
}

func (v InputMessage) native() *C.GInputMessage {
	return (*C.GInputMessage)(v.Ptr)
}
func wrapInputMessage(p *C.GInputMessage) InputMessage {
	return InputMessage{unsafe.Pointer(p)}
}
func WrapInputMessage(p unsafe.Pointer) InputMessage {
	return InputMessage{p}
}
func (v InputMessage) IsNil() bool {
	return v.Ptr == nil
}
func IWrapInputMessage(p unsafe.Pointer) interface{} {
	return WrapInputMessage(p)
}

// Struct InputVector
type InputVector struct {
	Ptr unsafe.Pointer
}

func (v InputVector) native() *C.GInputVector {
	return (*C.GInputVector)(v.Ptr)
}
func wrapInputVector(p *C.GInputVector) InputVector {
	return InputVector{unsafe.Pointer(p)}
}
func WrapInputVector(p unsafe.Pointer) InputVector {
	return InputVector{p}
}
func (v InputVector) IsNil() bool {
	return v.Ptr == nil
}
func IWrapInputVector(p unsafe.Pointer) interface{} {
	return WrapInputVector(p)
}

// Struct Resource
type Resource struct {
	Ptr unsafe.Pointer
}

func (v Resource) native() *C.GResource {
	return (*C.GResource)(v.Ptr)
}
func wrapResource(p *C.GResource) Resource {
	return Resource{unsafe.Pointer(p)}
}
func WrapResource(p unsafe.Pointer) Resource {
	return Resource{p}
}
func (v Resource) IsNil() bool {
	return v.Ptr == nil
}
func IWrapResource(p unsafe.Pointer) interface{} {
	return WrapResource(p)
}

// ResourceNewFromData is a wrapper around g_resource_new_from_data().
func ResourceNewFromData(data glib.Bytes) (Resource, error) {
	var err glib.Error
	ret0 := C.g_resource_new_from_data((*C.GBytes)(data.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return Resource{}, err.GoValue()
	}
	return wrapResource(ret0), nil
}

// ResourcesRegister is a wrapper around g_resources_register().
func (resource Resource) ResourcesRegister() {
	C.g_resources_register(resource.native())
}

// ResourcesUnregister is a wrapper around g_resources_unregister().
func (resource Resource) ResourcesUnregister() {
	C.g_resources_unregister(resource.native())
}

// EnumerateChildren is a wrapper around g_resource_enumerate_children().
func (resource Resource) EnumerateChildren(path string, lookup_flags ResourceLookupFlags) ([]string, error) {
	path0 := C.CString(path)
	var err glib.Error
	ret0 := C.g_resource_enumerate_children(resource.native(), path0, C.GResourceLookupFlags(lookup_flags), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(path0)) /*ch:<stdlib.h>*/
	var ret0Slice []*C.char
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString(elem)
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return nil, err.GoValue()
	}
	return ret, nil
}

// GetInfo is a wrapper around g_resource_get_info().
func (resource Resource) GetInfo(path string, lookup_flags ResourceLookupFlags) (bool, uint, uint32, error) {
	path0 := C.CString(path)
	var size0 C.gsize
	var flags0 C.guint32
	var err glib.Error
	ret0 := C.g_resource_get_info(resource.native(), path0, C.GResourceLookupFlags(lookup_flags), &size0, &flags0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(path0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, 0, 0, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, uint(size0), uint32(flags0), nil
}

// LookupData is a wrapper around g_resource_lookup_data().
func (resource Resource) LookupData(path string, lookup_flags ResourceLookupFlags) (glib.Bytes, error) {
	path0 := C.CString(path)
	var err glib.Error
	ret0 := C.g_resource_lookup_data(resource.native(), path0, C.GResourceLookupFlags(lookup_flags), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(path0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return glib.Bytes{}, err.GoValue()
	}
	return glib.WrapBytes(unsafe.Pointer(ret0)) /*gir:GLib*/, nil
}

// OpenStream is a wrapper around g_resource_open_stream().
func (resource Resource) OpenStream(path string, lookup_flags ResourceLookupFlags) (InputStream, error) {
	path0 := C.CString(path)
	var err glib.Error
	ret0 := C.g_resource_open_stream(resource.native(), path0, C.GResourceLookupFlags(lookup_flags), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(path0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return InputStream{}, err.GoValue()
	}
	return wrapInputStream(ret0), nil
}

// Ref is a wrapper around g_resource_ref().
func (resource Resource) Ref() Resource {
	ret0 := C.g_resource_ref(resource.native())
	return wrapResource(ret0)
}

// Unref is a wrapper around g_resource_unref().
func (resource Resource) Unref() {
	C.g_resource_unref(resource.native())
}

// ResourceLoad is a wrapper around g_resource_load().
func ResourceLoad(filename string) (Resource, error) {
	filename0 := (*C.gchar)(C.CString(filename))
	var err glib.Error
	ret0 := C.g_resource_load(filename0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(filename0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return Resource{}, err.GoValue()
	}
	return wrapResource(ret0), nil
}

// Struct OutputMessage
type OutputMessage struct {
	Ptr unsafe.Pointer
}

func (v OutputMessage) native() *C.GOutputMessage {
	return (*C.GOutputMessage)(v.Ptr)
}
func wrapOutputMessage(p *C.GOutputMessage) OutputMessage {
	return OutputMessage{unsafe.Pointer(p)}
}
func WrapOutputMessage(p unsafe.Pointer) OutputMessage {
	return OutputMessage{p}
}
func (v OutputMessage) IsNil() bool {
	return v.Ptr == nil
}
func IWrapOutputMessage(p unsafe.Pointer) interface{} {
	return WrapOutputMessage(p)
}

// Struct OutputVector
type OutputVector struct {
	Ptr unsafe.Pointer
}

func (v OutputVector) native() *C.GOutputVector {
	return (*C.GOutputVector)(v.Ptr)
}
func wrapOutputVector(p *C.GOutputVector) OutputVector {
	return OutputVector{unsafe.Pointer(p)}
}
func WrapOutputVector(p unsafe.Pointer) OutputVector {
	return OutputVector{p}
}
func (v OutputVector) IsNil() bool {
	return v.Ptr == nil
}
func IWrapOutputVector(p unsafe.Pointer) interface{} {
	return WrapOutputVector(p)
}

// Struct SettingsSchemaSource
type SettingsSchemaSource struct {
	Ptr unsafe.Pointer
}

func (v SettingsSchemaSource) native() *C.GSettingsSchemaSource {
	return (*C.GSettingsSchemaSource)(v.Ptr)
}
func wrapSettingsSchemaSource(p *C.GSettingsSchemaSource) SettingsSchemaSource {
	return SettingsSchemaSource{unsafe.Pointer(p)}
}
func WrapSettingsSchemaSource(p unsafe.Pointer) SettingsSchemaSource {
	return SettingsSchemaSource{p}
}
func (v SettingsSchemaSource) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSettingsSchemaSource(p unsafe.Pointer) interface{} {
	return WrapSettingsSchemaSource(p)
}

// SettingsSchemaSourceNewFromDirectory is a wrapper around g_settings_schema_source_new_from_directory().
func SettingsSchemaSourceNewFromDirectory(directory string, parent SettingsSchemaSource, trusted bool) (SettingsSchemaSource, error) {
	directory0 := (*C.gchar)(C.CString(directory))
	var err glib.Error
	ret0 := C.g_settings_schema_source_new_from_directory(directory0, parent.native(), C.gboolean(util.Bool2Int(trusted)) /*go:.util*/, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(directory0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return SettingsSchemaSource{}, err.GoValue()
	}
	return wrapSettingsSchemaSource(ret0), nil
}

// Lookup is a wrapper around g_settings_schema_source_lookup().
func (source SettingsSchemaSource) Lookup(schema_id string, recursive bool) SettingsSchema {
	schema_id0 := (*C.gchar)(C.CString(schema_id))
	ret0 := C.g_settings_schema_source_lookup(source.native(), schema_id0, C.gboolean(util.Bool2Int(recursive)) /*go:.util*/)
	C.free(unsafe.Pointer(schema_id0)) /*ch:<stdlib.h>*/
	return wrapSettingsSchema(ret0)
}

// Ref is a wrapper around g_settings_schema_source_ref().
func (source SettingsSchemaSource) Ref() SettingsSchemaSource {
	ret0 := C.g_settings_schema_source_ref(source.native())
	return wrapSettingsSchemaSource(ret0)
}

// Unref is a wrapper around g_settings_schema_source_unref().
func (source SettingsSchemaSource) Unref() {
	C.g_settings_schema_source_unref(source.native())
}

// SettingsSchemaSourceGetDefault is a wrapper around g_settings_schema_source_get_default().
func SettingsSchemaSourceGetDefault() SettingsSchemaSource {
	ret0 := C.g_settings_schema_source_get_default()
	return wrapSettingsSchemaSource(ret0)
}

// Struct SrvTarget
type SrvTarget struct {
	Ptr unsafe.Pointer
}

func (v SrvTarget) native() *C.GSrvTarget {
	return (*C.GSrvTarget)(v.Ptr)
}
func wrapSrvTarget(p *C.GSrvTarget) SrvTarget {
	return SrvTarget{unsafe.Pointer(p)}
}
func WrapSrvTarget(p unsafe.Pointer) SrvTarget {
	return SrvTarget{p}
}
func (v SrvTarget) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSrvTarget(p unsafe.Pointer) interface{} {
	return WrapSrvTarget(p)
}

// SrvTargetNew is a wrapper around g_srv_target_new().
func SrvTargetNew(hostname string, port uint16, priority uint16, weight uint16) SrvTarget {
	hostname0 := (*C.gchar)(C.CString(hostname))
	ret0 := C.g_srv_target_new(hostname0, C.guint16(port), C.guint16(priority), C.guint16(weight))
	C.free(unsafe.Pointer(hostname0)) /*ch:<stdlib.h>*/
	return wrapSrvTarget(ret0)
}

// Copy is a wrapper around g_srv_target_copy().
func (target SrvTarget) Copy() SrvTarget {
	ret0 := C.g_srv_target_copy(target.native())
	return wrapSrvTarget(ret0)
}

// Free is a wrapper around g_srv_target_free().
func (target SrvTarget) Free() {
	C.g_srv_target_free(target.native())
}

// GetHostname is a wrapper around g_srv_target_get_hostname().
func (target SrvTarget) GetHostname() string {
	ret0 := C.g_srv_target_get_hostname(target.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetPort is a wrapper around g_srv_target_get_port().
func (target SrvTarget) GetPort() uint16 {
	ret0 := C.g_srv_target_get_port(target.native())
	return uint16(ret0)
}

// GetPriority is a wrapper around g_srv_target_get_priority().
func (target SrvTarget) GetPriority() uint16 {
	ret0 := C.g_srv_target_get_priority(target.native())
	return uint16(ret0)
}

// GetWeight is a wrapper around g_srv_target_get_weight().
func (target SrvTarget) GetWeight() uint16 {
	ret0 := C.g_srv_target_get_weight(target.native())
	return uint16(ret0)
}

// Struct StaticResource
type StaticResource struct {
	Ptr unsafe.Pointer
}

func (v StaticResource) native() *C.GStaticResource {
	return (*C.GStaticResource)(v.Ptr)
}
func wrapStaticResource(p *C.GStaticResource) StaticResource {
	return StaticResource{unsafe.Pointer(p)}
}
func WrapStaticResource(p unsafe.Pointer) StaticResource {
	return StaticResource{p}
}
func (v StaticResource) IsNil() bool {
	return v.Ptr == nil
}
func IWrapStaticResource(p unsafe.Pointer) interface{} {
	return WrapStaticResource(p)
}

// Fini is a wrapper around g_static_resource_fini().
func (static_resource StaticResource) Fini() {
	C.g_static_resource_fini(static_resource.native())
}

// GetResource is a wrapper around g_static_resource_get_resource().
func (static_resource StaticResource) GetResource() Resource {
	ret0 := C.g_static_resource_get_resource(static_resource.native())
	return wrapResource(ret0)
}

// Init is a wrapper around g_static_resource_init().
func (static_resource StaticResource) Init() {
	C.g_static_resource_init(static_resource.native())
}

// Interface Converter
type Converter struct {
	Ptr unsafe.Pointer
}

func (v Converter) native() *C.GConverter {
	return (*C.GConverter)(v.Ptr)
}
func wrapConverter(p *C.GConverter) Converter {
	return Converter{unsafe.Pointer(p)}
}
func WrapConverter(p unsafe.Pointer) Converter {
	return Converter{p}
}
func (v Converter) IsNil() bool {
	return v.Ptr == nil
}
func IWrapConverter(p unsafe.Pointer) interface{} {
	return WrapConverter(p)
}
func (v Converter) GetType() gobject.Type {
	return gobject.Type(C.g_converter_get_type())
}
func (v Converter) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapConverter(unsafe.Pointer(ptr)), nil
	}
}

// Reset is a wrapper around g_converter_reset().
func (converter Converter) Reset() {
	C.g_converter_reset(converter.native())
}

// Interface DBusObjectManager
type DBusObjectManager struct {
	Ptr unsafe.Pointer
}

func (v DBusObjectManager) native() *C.GDBusObjectManager {
	return (*C.GDBusObjectManager)(v.Ptr)
}
func wrapDBusObjectManager(p *C.GDBusObjectManager) DBusObjectManager {
	return DBusObjectManager{unsafe.Pointer(p)}
}
func WrapDBusObjectManager(p unsafe.Pointer) DBusObjectManager {
	return DBusObjectManager{p}
}
func (v DBusObjectManager) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusObjectManager(p unsafe.Pointer) interface{} {
	return WrapDBusObjectManager(p)
}
func (v DBusObjectManager) GetType() gobject.Type {
	return gobject.Type(C.g_dbus_object_manager_get_type())
}
func (v DBusObjectManager) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDBusObjectManager(unsafe.Pointer(ptr)), nil
	}
}

// GetInterface is a wrapper around g_dbus_object_manager_get_interface().
func (manager DBusObjectManager) GetInterface(object_path string, interface_name string) DBusInterface {
	object_path0 := (*C.gchar)(C.CString(object_path))
	interface_name0 := (*C.gchar)(C.CString(interface_name))
	ret0 := C.g_dbus_object_manager_get_interface(manager.native(), object_path0, interface_name0)
	C.free(unsafe.Pointer(object_path0))    /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(interface_name0)) /*ch:<stdlib.h>*/
	return wrapDBusInterface(ret0)
}

// GetObject is a wrapper around g_dbus_object_manager_get_object().
func (manager DBusObjectManager) GetObject(object_path string) DBusObject {
	object_path0 := (*C.gchar)(C.CString(object_path))
	ret0 := C.g_dbus_object_manager_get_object(manager.native(), object_path0)
	C.free(unsafe.Pointer(object_path0)) /*ch:<stdlib.h>*/
	return wrapDBusObject(ret0)
}

// GetObjectPath is a wrapper around g_dbus_object_manager_get_object_path().
func (manager DBusObjectManager) GetObjectPath() string {
	ret0 := C.g_dbus_object_manager_get_object_path(manager.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetObjects is a wrapper around g_dbus_object_manager_get_objects().
func (manager DBusObjectManager) GetObjects() glib.List {
	ret0 := C.g_dbus_object_manager_get_objects(manager.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapDBusObject(p) }) /*gir:GLib*/
}

// Interface DBusInterface
type DBusInterface struct {
	Ptr unsafe.Pointer
}

func (v DBusInterface) native() *C.GDBusInterface {
	return (*C.GDBusInterface)(v.Ptr)
}
func wrapDBusInterface(p *C.GDBusInterface) DBusInterface {
	return DBusInterface{unsafe.Pointer(p)}
}
func WrapDBusInterface(p unsafe.Pointer) DBusInterface {
	return DBusInterface{p}
}
func (v DBusInterface) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusInterface(p unsafe.Pointer) interface{} {
	return WrapDBusInterface(p)
}
func (v DBusInterface) GetType() gobject.Type {
	return gobject.Type(C.g_dbus_interface_get_type())
}
func (v DBusInterface) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDBusInterface(unsafe.Pointer(ptr)), nil
	}
}

// DupObject is a wrapper around g_dbus_interface_dup_object().
func (interface_ DBusInterface) DupObject() DBusObject {
	ret0 := C.g_dbus_interface_dup_object(interface_.native())
	return wrapDBusObject(ret0)
}

// GetInfo is a wrapper around g_dbus_interface_get_info().
func (interface_ DBusInterface) GetInfo() DBusInterfaceInfo {
	ret0 := C.g_dbus_interface_get_info(interface_.native())
	return wrapDBusInterfaceInfo(ret0)
}

// GetObject is a wrapper around g_dbus_interface_get_object().
func (interface_ DBusInterface) GetObject() DBusObject {
	ret0 := C.g_dbus_interface_get_object(interface_.native())
	return wrapDBusObject(ret0)
}

// SetObject is a wrapper around g_dbus_interface_set_object().
func (interface_ DBusInterface) SetObject(object DBusObject) {
	C.g_dbus_interface_set_object(interface_.native(), object.native())
}

// Interface DBusObject
type DBusObject struct {
	Ptr unsafe.Pointer
}

func (v DBusObject) native() *C.GDBusObject {
	return (*C.GDBusObject)(v.Ptr)
}
func wrapDBusObject(p *C.GDBusObject) DBusObject {
	return DBusObject{unsafe.Pointer(p)}
}
func WrapDBusObject(p unsafe.Pointer) DBusObject {
	return DBusObject{p}
}
func (v DBusObject) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusObject(p unsafe.Pointer) interface{} {
	return WrapDBusObject(p)
}
func (v DBusObject) GetType() gobject.Type {
	return gobject.Type(C.g_dbus_object_get_type())
}
func (v DBusObject) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDBusObject(unsafe.Pointer(ptr)), nil
	}
}

// GetInterface is a wrapper around g_dbus_object_get_interface().
func (object DBusObject) GetInterface(interface_name string) DBusInterface {
	interface_name0 := (*C.gchar)(C.CString(interface_name))
	ret0 := C.g_dbus_object_get_interface(object.native(), interface_name0)
	C.free(unsafe.Pointer(interface_name0)) /*ch:<stdlib.h>*/
	return wrapDBusInterface(ret0)
}

// GetInterfaces is a wrapper around g_dbus_object_get_interfaces().
func (object DBusObject) GetInterfaces() glib.List {
	ret0 := C.g_dbus_object_get_interfaces(object.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapDBusInterface(p) }) /*gir:GLib*/
}

// GetObjectPath is a wrapper around g_dbus_object_get_object_path().
func (object DBusObject) GetObjectPath() string {
	ret0 := C.g_dbus_object_get_object_path(object.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// Interface DatagramBased
type DatagramBased struct {
	Ptr unsafe.Pointer
}

func (v DatagramBased) native() *C.GDatagramBased {
	return (*C.GDatagramBased)(v.Ptr)
}
func wrapDatagramBased(p *C.GDatagramBased) DatagramBased {
	return DatagramBased{unsafe.Pointer(p)}
}
func WrapDatagramBased(p unsafe.Pointer) DatagramBased {
	return DatagramBased{p}
}
func (v DatagramBased) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDatagramBased(p unsafe.Pointer) interface{} {
	return WrapDatagramBased(p)
}
func (v DatagramBased) GetType() gobject.Type {
	return gobject.Type(C.g_datagram_based_get_type())
}
func (v DatagramBased) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDatagramBased(unsafe.Pointer(ptr)), nil
	}
}

// ConditionCheck is a wrapper around g_datagram_based_condition_check().
func (datagram_based DatagramBased) ConditionCheck(condition /*gir:GLib*/ glib.IOCondition) /*gir:GLib*/ glib.IOCondition {
	ret0 := C.g_datagram_based_condition_check(datagram_based.native(), C.GIOCondition(condition))
	return /*gir:GLib*/ glib.IOCondition(ret0)
}

// ConditionWait is a wrapper around g_datagram_based_condition_wait().
func (datagram_based DatagramBased) ConditionWait(condition /*gir:GLib*/ glib.IOCondition, timeout int64, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_datagram_based_condition_wait(datagram_based.native(), C.GIOCondition(condition), C.gint64(timeout), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// CreateSource is a wrapper around g_datagram_based_create_source().
func (datagram_based DatagramBased) CreateSource(condition /*gir:GLib*/ glib.IOCondition, cancellable Cancellable) glib.Source {
	ret0 := C.g_datagram_based_create_source(datagram_based.native(), C.GIOCondition(condition), cancellable.native())
	return glib.WrapSource(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// Interface DtlsClientConnection
type DtlsClientConnection struct {
	Ptr unsafe.Pointer
}

func (v DtlsClientConnection) native() *C.GDtlsClientConnection {
	return (*C.GDtlsClientConnection)(v.Ptr)
}
func wrapDtlsClientConnection(p *C.GDtlsClientConnection) DtlsClientConnection {
	return DtlsClientConnection{unsafe.Pointer(p)}
}
func WrapDtlsClientConnection(p unsafe.Pointer) DtlsClientConnection {
	return DtlsClientConnection{p}
}
func (v DtlsClientConnection) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDtlsClientConnection(p unsafe.Pointer) interface{} {
	return WrapDtlsClientConnection(p)
}
func (v DtlsClientConnection) GetType() gobject.Type {
	return gobject.Type(C.g_dtls_client_connection_get_type())
}
func (v DtlsClientConnection) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDtlsClientConnection(unsafe.Pointer(ptr)), nil
	}
}

// GetServerIdentity is a wrapper around g_dtls_client_connection_get_server_identity().
func (conn DtlsClientConnection) GetServerIdentity() SocketConnectable {
	ret0 := C.g_dtls_client_connection_get_server_identity(conn.native())
	return wrapSocketConnectable(ret0)
}

// GetValidationFlags is a wrapper around g_dtls_client_connection_get_validation_flags().
func (conn DtlsClientConnection) GetValidationFlags() TlsCertificateFlags {
	ret0 := C.g_dtls_client_connection_get_validation_flags(conn.native())
	return TlsCertificateFlags(ret0)
}

// SetServerIdentity is a wrapper around g_dtls_client_connection_set_server_identity().
func (conn DtlsClientConnection) SetServerIdentity(identity SocketConnectable) {
	C.g_dtls_client_connection_set_server_identity(conn.native(), identity.native())
}

// SetValidationFlags is a wrapper around g_dtls_client_connection_set_validation_flags().
func (conn DtlsClientConnection) SetValidationFlags(flags TlsCertificateFlags) {
	C.g_dtls_client_connection_set_validation_flags(conn.native(), C.GTlsCertificateFlags(flags))
}

// Interface SocketConnectable
type SocketConnectable struct {
	Ptr unsafe.Pointer
}

func (v SocketConnectable) native() *C.GSocketConnectable {
	return (*C.GSocketConnectable)(v.Ptr)
}
func wrapSocketConnectable(p *C.GSocketConnectable) SocketConnectable {
	return SocketConnectable{unsafe.Pointer(p)}
}
func WrapSocketConnectable(p unsafe.Pointer) SocketConnectable {
	return SocketConnectable{p}
}
func (v SocketConnectable) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSocketConnectable(p unsafe.Pointer) interface{} {
	return WrapSocketConnectable(p)
}
func (v SocketConnectable) GetType() gobject.Type {
	return gobject.Type(C.g_socket_connectable_get_type())
}
func (v SocketConnectable) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSocketConnectable(unsafe.Pointer(ptr)), nil
	}
}

// Enumerate is a wrapper around g_socket_connectable_enumerate().
func (connectable SocketConnectable) Enumerate() SocketAddressEnumerator {
	ret0 := C.g_socket_connectable_enumerate(connectable.native())
	return wrapSocketAddressEnumerator(ret0)
}

// ProxyEnumerate is a wrapper around g_socket_connectable_proxy_enumerate().
func (connectable SocketConnectable) ProxyEnumerate() SocketAddressEnumerator {
	ret0 := C.g_socket_connectable_proxy_enumerate(connectable.native())
	return wrapSocketAddressEnumerator(ret0)
}

// ToString is a wrapper around g_socket_connectable_to_string().
func (connectable SocketConnectable) ToString() string {
	ret0 := C.g_socket_connectable_to_string(connectable.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// Object SocketAddressEnumerator
type SocketAddressEnumerator struct {
	gobject.Object
}

func (v SocketAddressEnumerator) native() *C.GSocketAddressEnumerator {
	return (*C.GSocketAddressEnumerator)(v.Ptr)
}
func wrapSocketAddressEnumerator(p *C.GSocketAddressEnumerator) (v SocketAddressEnumerator) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSocketAddressEnumerator(p unsafe.Pointer) (v SocketAddressEnumerator) {
	v.Ptr = p
	return
}
func (v SocketAddressEnumerator) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSocketAddressEnumerator(p unsafe.Pointer) interface{} {
	return WrapSocketAddressEnumerator(p)
}
func (v SocketAddressEnumerator) GetType() gobject.Type {
	return gobject.Type(C.g_socket_address_enumerator_get_type())
}
func (v SocketAddressEnumerator) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSocketAddressEnumerator(unsafe.Pointer(ptr)), nil
	}
}

// Next is a wrapper around g_socket_address_enumerator_next().
func (enumerator SocketAddressEnumerator) Next(cancellable Cancellable) (SocketAddress, error) {
	var err glib.Error
	ret0 := C.g_socket_address_enumerator_next(enumerator.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return SocketAddress{}, err.GoValue()
	}
	return wrapSocketAddress(ret0), nil
}

// NextAsync is a wrapper around g_socket_address_enumerator_next_async().
func (enumerator SocketAddressEnumerator) NextAsync(cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_socket_address_enumerator_next_async(enumerator.native(), cancellable.native(), callback0)
}

// NextFinish is a wrapper around g_socket_address_enumerator_next_finish().
func (enumerator SocketAddressEnumerator) NextFinish(result AsyncResult) (SocketAddress, error) {
	var err glib.Error
	ret0 := C.g_socket_address_enumerator_next_finish(enumerator.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return SocketAddress{}, err.GoValue()
	}
	return wrapSocketAddress(ret0), nil
}

// Object SocketAddress
type SocketAddress struct {
	gobject.Object
}

func (v SocketAddress) native() *C.GSocketAddress {
	return (*C.GSocketAddress)(v.Ptr)
}
func wrapSocketAddress(p *C.GSocketAddress) (v SocketAddress) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSocketAddress(p unsafe.Pointer) (v SocketAddress) {
	v.Ptr = p
	return
}
func (v SocketAddress) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSocketAddress(p unsafe.Pointer) interface{} {
	return WrapSocketAddress(p)
}
func (v SocketAddress) GetType() gobject.Type {
	return gobject.Type(C.g_socket_address_get_type())
}
func (v SocketAddress) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSocketAddress(unsafe.Pointer(ptr)), nil
	}
}
func (v SocketAddress) SocketConnectable() SocketConnectable {
	return WrapSocketConnectable(v.Ptr)
}

// SocketAddressNewFromNative is a wrapper around g_socket_address_new_from_native().
func SocketAddressNewFromNative(native unsafe.Pointer, len uint) SocketAddress {
	ret0 := C.g_socket_address_new_from_native(C.gpointer(native), C.gsize(len))
	return wrapSocketAddress(ret0)
}

// GetFamily is a wrapper around g_socket_address_get_family().
func (address SocketAddress) GetFamily() SocketFamily {
	ret0 := C.g_socket_address_get_family(address.native())
	return SocketFamily(ret0)
}

// GetNativeSize is a wrapper around g_socket_address_get_native_size().
func (address SocketAddress) GetNativeSize() int {
	ret0 := C.g_socket_address_get_native_size(address.native())
	return int(ret0)
}

// ToNative is a wrapper around g_socket_address_to_native().
func (address SocketAddress) ToNative(dest unsafe.Pointer, destlen uint) (bool, error) {
	var err glib.Error
	ret0 := C.g_socket_address_to_native(address.native(), C.gpointer(dest), C.gsize(destlen), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Interface DtlsConnection
type DtlsConnection struct {
	Ptr unsafe.Pointer
}

func (v DtlsConnection) native() *C.GDtlsConnection {
	return (*C.GDtlsConnection)(v.Ptr)
}
func wrapDtlsConnection(p *C.GDtlsConnection) DtlsConnection {
	return DtlsConnection{unsafe.Pointer(p)}
}
func WrapDtlsConnection(p unsafe.Pointer) DtlsConnection {
	return DtlsConnection{p}
}
func (v DtlsConnection) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDtlsConnection(p unsafe.Pointer) interface{} {
	return WrapDtlsConnection(p)
}
func (v DtlsConnection) GetType() gobject.Type {
	return gobject.Type(C.g_dtls_connection_get_type())
}
func (v DtlsConnection) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDtlsConnection(unsafe.Pointer(ptr)), nil
	}
}

// Close is a wrapper around g_dtls_connection_close().
func (conn DtlsConnection) Close(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_dtls_connection_close(conn.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// CloseAsync is a wrapper around g_dtls_connection_close_async().
func (conn DtlsConnection) CloseAsync(io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_dtls_connection_close_async(conn.native(), C.int(io_priority), cancellable.native(), callback0)
}

// CloseFinish is a wrapper around g_dtls_connection_close_finish().
func (conn DtlsConnection) CloseFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_dtls_connection_close_finish(conn.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// EmitAcceptCertificate is a wrapper around g_dtls_connection_emit_accept_certificate().
func (conn DtlsConnection) EmitAcceptCertificate(peer_cert TlsCertificate, errors TlsCertificateFlags) bool {
	ret0 := C.g_dtls_connection_emit_accept_certificate(conn.native(), peer_cert.native(), C.GTlsCertificateFlags(errors))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetCertificate is a wrapper around g_dtls_connection_get_certificate().
func (conn DtlsConnection) GetCertificate() TlsCertificate {
	ret0 := C.g_dtls_connection_get_certificate(conn.native())
	return wrapTlsCertificate(ret0)
}

// GetDatabase is a wrapper around g_dtls_connection_get_database().
func (conn DtlsConnection) GetDatabase() TlsDatabase {
	ret0 := C.g_dtls_connection_get_database(conn.native())
	return wrapTlsDatabase(ret0)
}

// GetInteraction is a wrapper around g_dtls_connection_get_interaction().
func (conn DtlsConnection) GetInteraction() TlsInteraction {
	ret0 := C.g_dtls_connection_get_interaction(conn.native())
	return wrapTlsInteraction(ret0)
}

// GetPeerCertificate is a wrapper around g_dtls_connection_get_peer_certificate().
func (conn DtlsConnection) GetPeerCertificate() TlsCertificate {
	ret0 := C.g_dtls_connection_get_peer_certificate(conn.native())
	return wrapTlsCertificate(ret0)
}

// GetPeerCertificateErrors is a wrapper around g_dtls_connection_get_peer_certificate_errors().
func (conn DtlsConnection) GetPeerCertificateErrors() TlsCertificateFlags {
	ret0 := C.g_dtls_connection_get_peer_certificate_errors(conn.native())
	return TlsCertificateFlags(ret0)
}

// GetRehandshakeMode is a wrapper around g_dtls_connection_get_rehandshake_mode().
func (conn DtlsConnection) GetRehandshakeMode() TlsRehandshakeMode {
	ret0 := C.g_dtls_connection_get_rehandshake_mode(conn.native())
	return TlsRehandshakeMode(ret0)
}

// GetRequireCloseNotify is a wrapper around g_dtls_connection_get_require_close_notify().
func (conn DtlsConnection) GetRequireCloseNotify() bool {
	ret0 := C.g_dtls_connection_get_require_close_notify(conn.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Handshake is a wrapper around g_dtls_connection_handshake().
func (conn DtlsConnection) Handshake(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_dtls_connection_handshake(conn.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// HandshakeAsync is a wrapper around g_dtls_connection_handshake_async().
func (conn DtlsConnection) HandshakeAsync(io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_dtls_connection_handshake_async(conn.native(), C.int(io_priority), cancellable.native(), callback0)
}

// HandshakeFinish is a wrapper around g_dtls_connection_handshake_finish().
func (conn DtlsConnection) HandshakeFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_dtls_connection_handshake_finish(conn.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetCertificate is a wrapper around g_dtls_connection_set_certificate().
func (conn DtlsConnection) SetCertificate(certificate TlsCertificate) {
	C.g_dtls_connection_set_certificate(conn.native(), certificate.native())
}

// SetDatabase is a wrapper around g_dtls_connection_set_database().
func (conn DtlsConnection) SetDatabase(database TlsDatabase) {
	C.g_dtls_connection_set_database(conn.native(), database.native())
}

// SetInteraction is a wrapper around g_dtls_connection_set_interaction().
func (conn DtlsConnection) SetInteraction(interaction TlsInteraction) {
	C.g_dtls_connection_set_interaction(conn.native(), interaction.native())
}

// SetRehandshakeMode is a wrapper around g_dtls_connection_set_rehandshake_mode().
func (conn DtlsConnection) SetRehandshakeMode(mode TlsRehandshakeMode) {
	C.g_dtls_connection_set_rehandshake_mode(conn.native(), C.GTlsRehandshakeMode(mode))
}

// SetRequireCloseNotify is a wrapper around g_dtls_connection_set_require_close_notify().
func (conn DtlsConnection) SetRequireCloseNotify(require_close_notify bool) {
	C.g_dtls_connection_set_require_close_notify(conn.native(), C.gboolean(util.Bool2Int(require_close_notify)) /*go:.util*/)
}

// Shutdown is a wrapper around g_dtls_connection_shutdown().
func (conn DtlsConnection) Shutdown(shutdown_read bool, shutdown_write bool, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_dtls_connection_shutdown(conn.native(), C.gboolean(util.Bool2Int(shutdown_read)) /*go:.util*/, C.gboolean(util.Bool2Int(shutdown_write)) /*go:.util*/, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// ShutdownAsync is a wrapper around g_dtls_connection_shutdown_async().
func (conn DtlsConnection) ShutdownAsync(shutdown_read bool, shutdown_write bool, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_dtls_connection_shutdown_async(conn.native(), C.gboolean(util.Bool2Int(shutdown_read)) /*go:.util*/, C.gboolean(util.Bool2Int(shutdown_write)) /*go:.util*/, C.int(io_priority), cancellable.native(), callback0)
}

// ShutdownFinish is a wrapper around g_dtls_connection_shutdown_finish().
func (conn DtlsConnection) ShutdownFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_dtls_connection_shutdown_finish(conn.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Object TlsCertificate
type TlsCertificate struct {
	gobject.Object
}

func (v TlsCertificate) native() *C.GTlsCertificate {
	return (*C.GTlsCertificate)(v.Ptr)
}
func wrapTlsCertificate(p *C.GTlsCertificate) (v TlsCertificate) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTlsCertificate(p unsafe.Pointer) (v TlsCertificate) {
	v.Ptr = p
	return
}
func (v TlsCertificate) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTlsCertificate(p unsafe.Pointer) interface{} {
	return WrapTlsCertificate(p)
}
func (v TlsCertificate) GetType() gobject.Type {
	return gobject.Type(C.g_tls_certificate_get_type())
}
func (v TlsCertificate) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTlsCertificate(unsafe.Pointer(ptr)), nil
	}
}

// TlsCertificateNewFromFile is a wrapper around g_tls_certificate_new_from_file().
func TlsCertificateNewFromFile(file string) (TlsCertificate, error) {
	file0 := (*C.gchar)(C.CString(file))
	var err glib.Error
	ret0 := C.g_tls_certificate_new_from_file(file0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(file0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return TlsCertificate{}, err.GoValue()
	}
	return wrapTlsCertificate(ret0), nil
}

// TlsCertificateNewFromFiles is a wrapper around g_tls_certificate_new_from_files().
func TlsCertificateNewFromFiles(cert_file string, key_file string) (TlsCertificate, error) {
	cert_file0 := (*C.gchar)(C.CString(cert_file))
	key_file0 := (*C.gchar)(C.CString(key_file))
	var err glib.Error
	ret0 := C.g_tls_certificate_new_from_files(cert_file0, key_file0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(cert_file0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(key_file0))  /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return TlsCertificate{}, err.GoValue()
	}
	return wrapTlsCertificate(ret0), nil
}

// TlsCertificateNewFromPem is a wrapper around g_tls_certificate_new_from_pem().
func TlsCertificateNewFromPem(data string, length int) (TlsCertificate, error) {
	data0 := (*C.gchar)(C.CString(data))
	var err glib.Error
	ret0 := C.g_tls_certificate_new_from_pem(data0, C.gssize(length), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(data0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return TlsCertificate{}, err.GoValue()
	}
	return wrapTlsCertificate(ret0), nil
}

// GetIssuer is a wrapper around g_tls_certificate_get_issuer().
func (cert TlsCertificate) GetIssuer() TlsCertificate {
	ret0 := C.g_tls_certificate_get_issuer(cert.native())
	return wrapTlsCertificate(ret0)
}

// IsSame is a wrapper around g_tls_certificate_is_same().
func (cert_one TlsCertificate) IsSame(cert_two TlsCertificate) bool {
	ret0 := C.g_tls_certificate_is_same(cert_one.native(), cert_two.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Verify is a wrapper around g_tls_certificate_verify().
func (cert TlsCertificate) Verify(identity SocketConnectable, trusted_ca TlsCertificate) TlsCertificateFlags {
	ret0 := C.g_tls_certificate_verify(cert.native(), identity.native(), trusted_ca.native())
	return TlsCertificateFlags(ret0)
}

// TlsCertificateListNewFromFile is a wrapper around g_tls_certificate_list_new_from_file().
func TlsCertificateListNewFromFile(file string) (glib.List, error) {
	file0 := (*C.gchar)(C.CString(file))
	var err glib.Error
	ret0 := C.g_tls_certificate_list_new_from_file(file0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(file0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return glib.List{}, err.GoValue()
	}
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapTlsCertificate(p) }), /*gir:GLib*/ nil
}

// Interface DtlsServerConnection
type DtlsServerConnection struct {
	Ptr unsafe.Pointer
}

func (v DtlsServerConnection) native() *C.GDtlsServerConnection {
	return (*C.GDtlsServerConnection)(v.Ptr)
}
func wrapDtlsServerConnection(p *C.GDtlsServerConnection) DtlsServerConnection {
	return DtlsServerConnection{unsafe.Pointer(p)}
}
func WrapDtlsServerConnection(p unsafe.Pointer) DtlsServerConnection {
	return DtlsServerConnection{p}
}
func (v DtlsServerConnection) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDtlsServerConnection(p unsafe.Pointer) interface{} {
	return WrapDtlsServerConnection(p)
}
func (v DtlsServerConnection) GetType() gobject.Type {
	return gobject.Type(C.g_dtls_server_connection_get_type())
}
func (v DtlsServerConnection) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDtlsServerConnection(unsafe.Pointer(ptr)), nil
	}
}

// Interface FileDescriptorBased
type FileDescriptorBased struct {
	Ptr unsafe.Pointer
}

func (v FileDescriptorBased) native() *C.GFileDescriptorBased {
	return (*C.GFileDescriptorBased)(v.Ptr)
}
func wrapFileDescriptorBased(p *C.GFileDescriptorBased) FileDescriptorBased {
	return FileDescriptorBased{unsafe.Pointer(p)}
}
func WrapFileDescriptorBased(p unsafe.Pointer) FileDescriptorBased {
	return FileDescriptorBased{p}
}
func (v FileDescriptorBased) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFileDescriptorBased(p unsafe.Pointer) interface{} {
	return WrapFileDescriptorBased(p)
}
func (v FileDescriptorBased) GetType() gobject.Type {
	return gobject.Type(C.g_file_descriptor_based_get_type())
}
func (v FileDescriptorBased) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFileDescriptorBased(unsafe.Pointer(ptr)), nil
	}
}

// GetFd is a wrapper around g_file_descriptor_based_get_fd().
func (fd_based FileDescriptorBased) GetFd() int {
	ret0 := C.g_file_descriptor_based_get_fd(fd_based.native())
	return int(ret0)
}

// Interface ListModel
type ListModel struct {
	Ptr unsafe.Pointer
}

func (v ListModel) native() *C.GListModel {
	return (*C.GListModel)(v.Ptr)
}
func wrapListModel(p *C.GListModel) ListModel {
	return ListModel{unsafe.Pointer(p)}
}
func WrapListModel(p unsafe.Pointer) ListModel {
	return ListModel{p}
}
func (v ListModel) IsNil() bool {
	return v.Ptr == nil
}
func IWrapListModel(p unsafe.Pointer) interface{} {
	return WrapListModel(p)
}
func (v ListModel) GetType() gobject.Type {
	return gobject.Type(C.g_list_model_get_type())
}
func (v ListModel) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapListModel(unsafe.Pointer(ptr)), nil
	}
}

// GetItem is a wrapper around g_list_model_get_item().
func (list ListModel) GetItem(position uint) gobject.Object {
	ret0 := C.g_list_model_get_item(list.native(), C.guint(position))
	return gobject.WrapObject(unsafe.Pointer(ret0)) /*gir:GObject*/
}

// GetNItems is a wrapper around g_list_model_get_n_items().
func (list ListModel) GetNItems() uint {
	ret0 := C.g_list_model_get_n_items(list.native())
	return uint(ret0)
}

// GetObject is a wrapper around g_list_model_get_object().
func (list ListModel) GetObject(position uint) gobject.Object {
	ret0 := C.g_list_model_get_object(list.native(), C.guint(position))
	return gobject.WrapObject(unsafe.Pointer(ret0)) /*gir:GObject*/
}

// ItemsChanged is a wrapper around g_list_model_items_changed().
func (list ListModel) ItemsChanged(position uint, removed uint, added uint) {
	C.g_list_model_items_changed(list.native(), C.guint(position), C.guint(removed), C.guint(added))
}

// Interface LoadableIcon
type LoadableIcon struct {
	Ptr unsafe.Pointer
}

func (v LoadableIcon) native() *C.GLoadableIcon {
	return (*C.GLoadableIcon)(v.Ptr)
}
func wrapLoadableIcon(p *C.GLoadableIcon) LoadableIcon {
	return LoadableIcon{unsafe.Pointer(p)}
}
func WrapLoadableIcon(p unsafe.Pointer) LoadableIcon {
	return LoadableIcon{p}
}
func (v LoadableIcon) IsNil() bool {
	return v.Ptr == nil
}
func IWrapLoadableIcon(p unsafe.Pointer) interface{} {
	return WrapLoadableIcon(p)
}
func (v LoadableIcon) GetType() gobject.Type {
	return gobject.Type(C.g_loadable_icon_get_type())
}
func (v LoadableIcon) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapLoadableIcon(unsafe.Pointer(ptr)), nil
	}
}

// Load is a wrapper around g_loadable_icon_load().
func (icon LoadableIcon) Load(size int, cancellable Cancellable) (InputStream, string, error) {
	var type0 *C.char
	var err glib.Error
	ret0 := C.g_loadable_icon_load(icon.native(), C.int(size), &type0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	type_ := C.GoString(type0)
	defer C.g_free(C.gpointer(type0))
	if err.Ptr != nil {
		defer err.Free()
		return InputStream{}, "", err.GoValue()
	}
	return wrapInputStream(ret0), type_, nil
}

// LoadAsync is a wrapper around g_loadable_icon_load_async().
func (icon LoadableIcon) LoadAsync(size int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_loadable_icon_load_async(icon.native(), C.int(size), cancellable.native(), callback0)
}

// LoadFinish is a wrapper around g_loadable_icon_load_finish().
func (icon LoadableIcon) LoadFinish(res AsyncResult) (InputStream, string, error) {
	var type0 *C.char
	var err glib.Error
	ret0 := C.g_loadable_icon_load_finish(icon.native(), res.native(), &type0, (**C.GError)(unsafe.Pointer(&err)))
	type_ := C.GoString(type0)
	defer C.g_free(C.gpointer(type0))
	if err.Ptr != nil {
		defer err.Free()
		return InputStream{}, "", err.GoValue()
	}
	return wrapInputStream(ret0), type_, nil
}

// Interface NetworkMonitor
type NetworkMonitor struct {
	Ptr unsafe.Pointer
}

func (v NetworkMonitor) native() *C.GNetworkMonitor {
	return (*C.GNetworkMonitor)(v.Ptr)
}
func wrapNetworkMonitor(p *C.GNetworkMonitor) NetworkMonitor {
	return NetworkMonitor{unsafe.Pointer(p)}
}
func WrapNetworkMonitor(p unsafe.Pointer) NetworkMonitor {
	return NetworkMonitor{p}
}
func (v NetworkMonitor) IsNil() bool {
	return v.Ptr == nil
}
func IWrapNetworkMonitor(p unsafe.Pointer) interface{} {
	return WrapNetworkMonitor(p)
}
func (v NetworkMonitor) GetType() gobject.Type {
	return gobject.Type(C.g_network_monitor_get_type())
}
func (v NetworkMonitor) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapNetworkMonitor(unsafe.Pointer(ptr)), nil
	}
}

// CanReach is a wrapper around g_network_monitor_can_reach().
func (monitor NetworkMonitor) CanReach(connectable SocketConnectable, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_network_monitor_can_reach(monitor.native(), connectable.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// CanReachAsync is a wrapper around g_network_monitor_can_reach_async().
func (monitor NetworkMonitor) CanReachAsync(connectable SocketConnectable, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_network_monitor_can_reach_async(monitor.native(), connectable.native(), cancellable.native(), callback0)
}

// CanReachFinish is a wrapper around g_network_monitor_can_reach_finish().
func (monitor NetworkMonitor) CanReachFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_network_monitor_can_reach_finish(monitor.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// GetConnectivity is a wrapper around g_network_monitor_get_connectivity().
func (monitor NetworkMonitor) GetConnectivity() NetworkConnectivity {
	ret0 := C.g_network_monitor_get_connectivity(monitor.native())
	return NetworkConnectivity(ret0)
}

// GetNetworkAvailable is a wrapper around g_network_monitor_get_network_available().
func (monitor NetworkMonitor) GetNetworkAvailable() bool {
	ret0 := C.g_network_monitor_get_network_available(monitor.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetNetworkMetered is a wrapper around g_network_monitor_get_network_metered().
func (monitor NetworkMonitor) GetNetworkMetered() bool {
	ret0 := C.g_network_monitor_get_network_metered(monitor.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// NetworkMonitorGetDefault is a wrapper around g_network_monitor_get_default().
func NetworkMonitorGetDefault() NetworkMonitor {
	ret0 := C.g_network_monitor_get_default()
	return wrapNetworkMonitor(ret0)
}

// Interface PollableInputStream
type PollableInputStream struct {
	Ptr unsafe.Pointer
}

func (v PollableInputStream) native() *C.GPollableInputStream {
	return (*C.GPollableInputStream)(v.Ptr)
}
func wrapPollableInputStream(p *C.GPollableInputStream) PollableInputStream {
	return PollableInputStream{unsafe.Pointer(p)}
}
func WrapPollableInputStream(p unsafe.Pointer) PollableInputStream {
	return PollableInputStream{p}
}
func (v PollableInputStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPollableInputStream(p unsafe.Pointer) interface{} {
	return WrapPollableInputStream(p)
}
func (v PollableInputStream) GetType() gobject.Type {
	return gobject.Type(C.g_pollable_input_stream_get_type())
}
func (v PollableInputStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPollableInputStream(unsafe.Pointer(ptr)), nil
	}
}

// CanPoll is a wrapper around g_pollable_input_stream_can_poll().
func (stream PollableInputStream) CanPoll() bool {
	ret0 := C.g_pollable_input_stream_can_poll(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// CreateSource is a wrapper around g_pollable_input_stream_create_source().
func (stream PollableInputStream) CreateSource(cancellable Cancellable) glib.Source {
	ret0 := C.g_pollable_input_stream_create_source(stream.native(), cancellable.native())
	return glib.WrapSource(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// IsReadable is a wrapper around g_pollable_input_stream_is_readable().
func (stream PollableInputStream) IsReadable() bool {
	ret0 := C.g_pollable_input_stream_is_readable(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Interface PollableOutputStream
type PollableOutputStream struct {
	Ptr unsafe.Pointer
}

func (v PollableOutputStream) native() *C.GPollableOutputStream {
	return (*C.GPollableOutputStream)(v.Ptr)
}
func wrapPollableOutputStream(p *C.GPollableOutputStream) PollableOutputStream {
	return PollableOutputStream{unsafe.Pointer(p)}
}
func WrapPollableOutputStream(p unsafe.Pointer) PollableOutputStream {
	return PollableOutputStream{p}
}
func (v PollableOutputStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPollableOutputStream(p unsafe.Pointer) interface{} {
	return WrapPollableOutputStream(p)
}
func (v PollableOutputStream) GetType() gobject.Type {
	return gobject.Type(C.g_pollable_output_stream_get_type())
}
func (v PollableOutputStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPollableOutputStream(unsafe.Pointer(ptr)), nil
	}
}

// CanPoll is a wrapper around g_pollable_output_stream_can_poll().
func (stream PollableOutputStream) CanPoll() bool {
	ret0 := C.g_pollable_output_stream_can_poll(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// CreateSource is a wrapper around g_pollable_output_stream_create_source().
func (stream PollableOutputStream) CreateSource(cancellable Cancellable) glib.Source {
	ret0 := C.g_pollable_output_stream_create_source(stream.native(), cancellable.native())
	return glib.WrapSource(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// IsWritable is a wrapper around g_pollable_output_stream_is_writable().
func (stream PollableOutputStream) IsWritable() bool {
	ret0 := C.g_pollable_output_stream_is_writable(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Interface Proxy
type Proxy struct {
	Ptr unsafe.Pointer
}

func (v Proxy) native() *C.GProxy {
	return (*C.GProxy)(v.Ptr)
}
func wrapProxy(p *C.GProxy) Proxy {
	return Proxy{unsafe.Pointer(p)}
}
func WrapProxy(p unsafe.Pointer) Proxy {
	return Proxy{p}
}
func (v Proxy) IsNil() bool {
	return v.Ptr == nil
}
func IWrapProxy(p unsafe.Pointer) interface{} {
	return WrapProxy(p)
}
func (v Proxy) GetType() gobject.Type {
	return gobject.Type(C.g_proxy_get_type())
}
func (v Proxy) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapProxy(unsafe.Pointer(ptr)), nil
	}
}

// Connect is a wrapper around g_proxy_connect().
func (proxy Proxy) Connect(connection IOStream, proxy_address ProxyAddress, cancellable Cancellable) (IOStream, error) {
	var err glib.Error
	ret0 := C.g_proxy_connect(proxy.native(), connection.native(), proxy_address.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return IOStream{}, err.GoValue()
	}
	return wrapIOStream(ret0), nil
}

// ConnectAsync is a wrapper around g_proxy_connect_async().
func (proxy Proxy) ConnectAsync(connection IOStream, proxy_address ProxyAddress, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_proxy_connect_async(proxy.native(), connection.native(), proxy_address.native(), cancellable.native(), callback0)
}

// ConnectFinish is a wrapper around g_proxy_connect_finish().
func (proxy Proxy) ConnectFinish(result AsyncResult) (IOStream, error) {
	var err glib.Error
	ret0 := C.g_proxy_connect_finish(proxy.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return IOStream{}, err.GoValue()
	}
	return wrapIOStream(ret0), nil
}

// SupportsHostname is a wrapper around g_proxy_supports_hostname().
func (proxy Proxy) SupportsHostname() bool {
	ret0 := C.g_proxy_supports_hostname(proxy.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ProxyGetDefaultForProtocol is a wrapper around g_proxy_get_default_for_protocol().
func ProxyGetDefaultForProtocol(protocol string) Proxy {
	protocol0 := (*C.gchar)(C.CString(protocol))
	ret0 := C.g_proxy_get_default_for_protocol(protocol0)
	C.free(unsafe.Pointer(protocol0)) /*ch:<stdlib.h>*/
	return wrapProxy(ret0)
}

// Object ProxyAddress
type ProxyAddress struct {
	InetSocketAddress
}

func (v ProxyAddress) native() *C.GProxyAddress {
	return (*C.GProxyAddress)(v.Ptr)
}
func wrapProxyAddress(p *C.GProxyAddress) (v ProxyAddress) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapProxyAddress(p unsafe.Pointer) (v ProxyAddress) {
	v.Ptr = p
	return
}
func (v ProxyAddress) IsNil() bool {
	return v.Ptr == nil
}
func IWrapProxyAddress(p unsafe.Pointer) interface{} {
	return WrapProxyAddress(p)
}
func (v ProxyAddress) GetType() gobject.Type {
	return gobject.Type(C.g_proxy_address_get_type())
}
func (v ProxyAddress) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapProxyAddress(unsafe.Pointer(ptr)), nil
	}
}
func (v ProxyAddress) SocketConnectable() SocketConnectable {
	return WrapSocketConnectable(v.Ptr)
}

// ProxyAddressNew is a wrapper around g_proxy_address_new().
func ProxyAddressNew(inetaddr InetAddress, port uint16, protocol string, dest_hostname string, dest_port uint16, username string, password string) SocketAddress {
	protocol0 := (*C.gchar)(C.CString(protocol))
	dest_hostname0 := (*C.gchar)(C.CString(dest_hostname))
	username0 := (*C.gchar)(C.CString(username))
	password0 := (*C.gchar)(C.CString(password))
	ret0 := C.g_proxy_address_new(inetaddr.native(), C.guint16(port), protocol0, dest_hostname0, C.guint16(dest_port), username0, password0)
	C.free(unsafe.Pointer(protocol0))      /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(dest_hostname0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(username0))      /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(password0))      /*ch:<stdlib.h>*/
	return wrapSocketAddress(ret0)
}

// GetDestinationHostname is a wrapper around g_proxy_address_get_destination_hostname().
func (proxy ProxyAddress) GetDestinationHostname() string {
	ret0 := C.g_proxy_address_get_destination_hostname(proxy.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetDestinationPort is a wrapper around g_proxy_address_get_destination_port().
func (proxy ProxyAddress) GetDestinationPort() uint16 {
	ret0 := C.g_proxy_address_get_destination_port(proxy.native())
	return uint16(ret0)
}

// GetDestinationProtocol is a wrapper around g_proxy_address_get_destination_protocol().
func (proxy ProxyAddress) GetDestinationProtocol() string {
	ret0 := C.g_proxy_address_get_destination_protocol(proxy.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetPassword is a wrapper around g_proxy_address_get_password().
func (proxy ProxyAddress) GetPassword() string {
	ret0 := C.g_proxy_address_get_password(proxy.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetProtocol is a wrapper around g_proxy_address_get_protocol().
func (proxy ProxyAddress) GetProtocol() string {
	ret0 := C.g_proxy_address_get_protocol(proxy.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetUri is a wrapper around g_proxy_address_get_uri().
func (proxy ProxyAddress) GetUri() string {
	ret0 := C.g_proxy_address_get_uri(proxy.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetUsername is a wrapper around g_proxy_address_get_username().
func (proxy ProxyAddress) GetUsername() string {
	ret0 := C.g_proxy_address_get_username(proxy.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// Object InetSocketAddress
type InetSocketAddress struct {
	SocketAddress
}

func (v InetSocketAddress) native() *C.GInetSocketAddress {
	return (*C.GInetSocketAddress)(v.Ptr)
}
func wrapInetSocketAddress(p *C.GInetSocketAddress) (v InetSocketAddress) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapInetSocketAddress(p unsafe.Pointer) (v InetSocketAddress) {
	v.Ptr = p
	return
}
func (v InetSocketAddress) IsNil() bool {
	return v.Ptr == nil
}
func IWrapInetSocketAddress(p unsafe.Pointer) interface{} {
	return WrapInetSocketAddress(p)
}
func (v InetSocketAddress) GetType() gobject.Type {
	return gobject.Type(C.g_inet_socket_address_get_type())
}
func (v InetSocketAddress) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapInetSocketAddress(unsafe.Pointer(ptr)), nil
	}
}
func (v InetSocketAddress) SocketConnectable() SocketConnectable {
	return WrapSocketConnectable(v.Ptr)
}

// InetSocketAddressNew is a wrapper around g_inet_socket_address_new().
func InetSocketAddressNew(address InetAddress, port uint16) SocketAddress {
	ret0 := C.g_inet_socket_address_new(address.native(), C.guint16(port))
	return wrapSocketAddress(ret0)
}

// InetSocketAddressNewFromString is a wrapper around g_inet_socket_address_new_from_string().
func InetSocketAddressNewFromString(address string, port uint) SocketAddress {
	address0 := C.CString(address)
	ret0 := C.g_inet_socket_address_new_from_string(address0, C.guint(port))
	C.free(unsafe.Pointer(address0)) /*ch:<stdlib.h>*/
	return wrapSocketAddress(ret0)
}

// GetAddress is a wrapper around g_inet_socket_address_get_address().
func (address InetSocketAddress) GetAddress() InetAddress {
	ret0 := C.g_inet_socket_address_get_address(address.native())
	return wrapInetAddress(ret0)
}

// GetFlowinfo is a wrapper around g_inet_socket_address_get_flowinfo().
func (address InetSocketAddress) GetFlowinfo() uint32 {
	ret0 := C.g_inet_socket_address_get_flowinfo(address.native())
	return uint32(ret0)
}

// GetPort is a wrapper around g_inet_socket_address_get_port().
func (address InetSocketAddress) GetPort() uint16 {
	ret0 := C.g_inet_socket_address_get_port(address.native())
	return uint16(ret0)
}

// GetScopeId is a wrapper around g_inet_socket_address_get_scope_id().
func (address InetSocketAddress) GetScopeId() uint32 {
	ret0 := C.g_inet_socket_address_get_scope_id(address.native())
	return uint32(ret0)
}

// Object InetAddress
type InetAddress struct {
	gobject.Object
}

func (v InetAddress) native() *C.GInetAddress {
	return (*C.GInetAddress)(v.Ptr)
}
func wrapInetAddress(p *C.GInetAddress) (v InetAddress) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapInetAddress(p unsafe.Pointer) (v InetAddress) {
	v.Ptr = p
	return
}
func (v InetAddress) IsNil() bool {
	return v.Ptr == nil
}
func IWrapInetAddress(p unsafe.Pointer) interface{} {
	return WrapInetAddress(p)
}
func (v InetAddress) GetType() gobject.Type {
	return gobject.Type(C.g_inet_address_get_type())
}
func (v InetAddress) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapInetAddress(unsafe.Pointer(ptr)), nil
	}
}

// InetAddressNewAny is a wrapper around g_inet_address_new_any().
func InetAddressNewAny(family SocketFamily) InetAddress {
	ret0 := C.g_inet_address_new_any(C.GSocketFamily(family))
	return wrapInetAddress(ret0)
}

// InetAddressNewFromBytes is a wrapper around g_inet_address_new_from_bytes().
func InetAddressNewFromBytes(bytes []uint8, family SocketFamily) InetAddress {
	bytes0 := make([]C.guint8, len(bytes))
	for idx, elemG := range bytes {
		bytes0[idx] = C.guint8(elemG)
	}
	var bytes0Ptr *C.guint8
	if len(bytes0) > 0 {
		bytes0Ptr = &bytes0[0]
	}
	ret0 := C.g_inet_address_new_from_bytes(bytes0Ptr, C.GSocketFamily(family))
	return wrapInetAddress(ret0)
}

// InetAddressNewFromString is a wrapper around g_inet_address_new_from_string().
func InetAddressNewFromString(string string) InetAddress {
	string0 := (*C.gchar)(C.CString(string))
	ret0 := C.g_inet_address_new_from_string(string0)
	C.free(unsafe.Pointer(string0)) /*ch:<stdlib.h>*/
	return wrapInetAddress(ret0)
}

// InetAddressNewLoopback is a wrapper around g_inet_address_new_loopback().
func InetAddressNewLoopback(family SocketFamily) InetAddress {
	ret0 := C.g_inet_address_new_loopback(C.GSocketFamily(family))
	return wrapInetAddress(ret0)
}

// Equal is a wrapper around g_inet_address_equal().
func (address InetAddress) Equal(other_address InetAddress) bool {
	ret0 := C.g_inet_address_equal(address.native(), other_address.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetFamily is a wrapper around g_inet_address_get_family().
func (address InetAddress) GetFamily() SocketFamily {
	ret0 := C.g_inet_address_get_family(address.native())
	return SocketFamily(ret0)
}

// GetIsAny is a wrapper around g_inet_address_get_is_any().
func (address InetAddress) GetIsAny() bool {
	ret0 := C.g_inet_address_get_is_any(address.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIsLinkLocal is a wrapper around g_inet_address_get_is_link_local().
func (address InetAddress) GetIsLinkLocal() bool {
	ret0 := C.g_inet_address_get_is_link_local(address.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIsLoopback is a wrapper around g_inet_address_get_is_loopback().
func (address InetAddress) GetIsLoopback() bool {
	ret0 := C.g_inet_address_get_is_loopback(address.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIsMcGlobal is a wrapper around g_inet_address_get_is_mc_global().
func (address InetAddress) GetIsMcGlobal() bool {
	ret0 := C.g_inet_address_get_is_mc_global(address.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIsMcLinkLocal is a wrapper around g_inet_address_get_is_mc_link_local().
func (address InetAddress) GetIsMcLinkLocal() bool {
	ret0 := C.g_inet_address_get_is_mc_link_local(address.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIsMcNodeLocal is a wrapper around g_inet_address_get_is_mc_node_local().
func (address InetAddress) GetIsMcNodeLocal() bool {
	ret0 := C.g_inet_address_get_is_mc_node_local(address.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIsMcOrgLocal is a wrapper around g_inet_address_get_is_mc_org_local().
func (address InetAddress) GetIsMcOrgLocal() bool {
	ret0 := C.g_inet_address_get_is_mc_org_local(address.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIsMcSiteLocal is a wrapper around g_inet_address_get_is_mc_site_local().
func (address InetAddress) GetIsMcSiteLocal() bool {
	ret0 := C.g_inet_address_get_is_mc_site_local(address.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIsMulticast is a wrapper around g_inet_address_get_is_multicast().
func (address InetAddress) GetIsMulticast() bool {
	ret0 := C.g_inet_address_get_is_multicast(address.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIsSiteLocal is a wrapper around g_inet_address_get_is_site_local().
func (address InetAddress) GetIsSiteLocal() bool {
	ret0 := C.g_inet_address_get_is_site_local(address.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetNativeSize is a wrapper around g_inet_address_get_native_size().
func (address InetAddress) GetNativeSize() uint {
	ret0 := C.g_inet_address_get_native_size(address.native())
	return uint(ret0)
}

// ToString is a wrapper around g_inet_address_to_string().
func (address InetAddress) ToString() string {
	ret0 := C.g_inet_address_to_string(address.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// Interface ProxyResolver
type ProxyResolver struct {
	Ptr unsafe.Pointer
}

func (v ProxyResolver) native() *C.GProxyResolver {
	return (*C.GProxyResolver)(v.Ptr)
}
func wrapProxyResolver(p *C.GProxyResolver) ProxyResolver {
	return ProxyResolver{unsafe.Pointer(p)}
}
func WrapProxyResolver(p unsafe.Pointer) ProxyResolver {
	return ProxyResolver{p}
}
func (v ProxyResolver) IsNil() bool {
	return v.Ptr == nil
}
func IWrapProxyResolver(p unsafe.Pointer) interface{} {
	return WrapProxyResolver(p)
}
func (v ProxyResolver) GetType() gobject.Type {
	return gobject.Type(C.g_proxy_resolver_get_type())
}
func (v ProxyResolver) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapProxyResolver(unsafe.Pointer(ptr)), nil
	}
}

// IsSupported is a wrapper around g_proxy_resolver_is_supported().
func (resolver ProxyResolver) IsSupported() bool {
	ret0 := C.g_proxy_resolver_is_supported(resolver.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Lookup is a wrapper around g_proxy_resolver_lookup().
func (resolver ProxyResolver) Lookup(uri string, cancellable Cancellable) ([]string, error) {
	uri0 := (*C.gchar)(C.CString(uri))
	var err glib.Error
	ret0 := C.g_proxy_resolver_lookup(resolver.native(), uri0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(uri0)) /*ch:<stdlib.h>*/
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return nil, err.GoValue()
	}
	return ret, nil
}

// LookupAsync is a wrapper around g_proxy_resolver_lookup_async().
func (resolver ProxyResolver) LookupAsync(uri string, cancellable Cancellable, callback AsyncReadyCallback) {
	uri0 := (*C.gchar)(C.CString(uri))
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_proxy_resolver_lookup_async(resolver.native(), uri0, cancellable.native(), callback0)
	C.free(unsafe.Pointer(uri0)) /*ch:<stdlib.h>*/
}

// LookupFinish is a wrapper around g_proxy_resolver_lookup_finish().
func (resolver ProxyResolver) LookupFinish(result AsyncResult) ([]string, error) {
	var err glib.Error
	ret0 := C.g_proxy_resolver_lookup_finish(resolver.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return nil, err.GoValue()
	}
	return ret, nil
}

// ProxyResolverGetDefault is a wrapper around g_proxy_resolver_get_default().
func ProxyResolverGetDefault() ProxyResolver {
	ret0 := C.g_proxy_resolver_get_default()
	return wrapProxyResolver(ret0)
}

// Interface RemoteActionGroup
type RemoteActionGroup struct {
	Ptr unsafe.Pointer
}

func (v RemoteActionGroup) native() *C.GRemoteActionGroup {
	return (*C.GRemoteActionGroup)(v.Ptr)
}
func wrapRemoteActionGroup(p *C.GRemoteActionGroup) RemoteActionGroup {
	return RemoteActionGroup{unsafe.Pointer(p)}
}
func WrapRemoteActionGroup(p unsafe.Pointer) RemoteActionGroup {
	return RemoteActionGroup{p}
}
func (v RemoteActionGroup) IsNil() bool {
	return v.Ptr == nil
}
func IWrapRemoteActionGroup(p unsafe.Pointer) interface{} {
	return WrapRemoteActionGroup(p)
}
func (v RemoteActionGroup) GetType() gobject.Type {
	return gobject.Type(C.g_remote_action_group_get_type())
}
func (v RemoteActionGroup) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapRemoteActionGroup(unsafe.Pointer(ptr)), nil
	}
}

// ActivateActionFull is a wrapper around g_remote_action_group_activate_action_full().
func (remote RemoteActionGroup) ActivateActionFull(action_name string, parameter glib.Variant, platform_data glib.Variant) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.g_remote_action_group_activate_action_full(remote.native(), action_name0, (*C.GVariant)(parameter.Ptr), (*C.GVariant)(platform_data.Ptr))
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// ChangeActionStateFull is a wrapper around g_remote_action_group_change_action_state_full().
func (remote RemoteActionGroup) ChangeActionStateFull(action_name string, value glib.Variant, platform_data glib.Variant) {
	action_name0 := (*C.gchar)(C.CString(action_name))
	C.g_remote_action_group_change_action_state_full(remote.native(), action_name0, (*C.GVariant)(value.Ptr), (*C.GVariant)(platform_data.Ptr))
	C.free(unsafe.Pointer(action_name0)) /*ch:<stdlib.h>*/
}

// Interface TlsBackend
type TlsBackend struct {
	Ptr unsafe.Pointer
}

func (v TlsBackend) native() *C.GTlsBackend {
	return (*C.GTlsBackend)(v.Ptr)
}
func wrapTlsBackend(p *C.GTlsBackend) TlsBackend {
	return TlsBackend{unsafe.Pointer(p)}
}
func WrapTlsBackend(p unsafe.Pointer) TlsBackend {
	return TlsBackend{p}
}
func (v TlsBackend) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTlsBackend(p unsafe.Pointer) interface{} {
	return WrapTlsBackend(p)
}
func (v TlsBackend) GetType() gobject.Type {
	return gobject.Type(C.g_tls_backend_get_type())
}
func (v TlsBackend) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTlsBackend(unsafe.Pointer(ptr)), nil
	}
}

// GetDefaultDatabase is a wrapper around g_tls_backend_get_default_database().
func (backend TlsBackend) GetDefaultDatabase() TlsDatabase {
	ret0 := C.g_tls_backend_get_default_database(backend.native())
	return wrapTlsDatabase(ret0)
}

// SupportsDtls is a wrapper around g_tls_backend_supports_dtls().
func (backend TlsBackend) SupportsDtls() bool {
	ret0 := C.g_tls_backend_supports_dtls(backend.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SupportsTls is a wrapper around g_tls_backend_supports_tls().
func (backend TlsBackend) SupportsTls() bool {
	ret0 := C.g_tls_backend_supports_tls(backend.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// TlsBackendGetDefault is a wrapper around g_tls_backend_get_default().
func TlsBackendGetDefault() TlsBackend {
	ret0 := C.g_tls_backend_get_default()
	return wrapTlsBackend(ret0)
}

// Object TlsDatabase
type TlsDatabase struct {
	gobject.Object
}

func (v TlsDatabase) native() *C.GTlsDatabase {
	return (*C.GTlsDatabase)(v.Ptr)
}
func wrapTlsDatabase(p *C.GTlsDatabase) (v TlsDatabase) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTlsDatabase(p unsafe.Pointer) (v TlsDatabase) {
	v.Ptr = p
	return
}
func (v TlsDatabase) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTlsDatabase(p unsafe.Pointer) interface{} {
	return WrapTlsDatabase(p)
}
func (v TlsDatabase) GetType() gobject.Type {
	return gobject.Type(C.g_tls_database_get_type())
}
func (v TlsDatabase) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTlsDatabase(unsafe.Pointer(ptr)), nil
	}
}

// CreateCertificateHandle is a wrapper around g_tls_database_create_certificate_handle().
func (self TlsDatabase) CreateCertificateHandle(certificate TlsCertificate) string {
	ret0 := C.g_tls_database_create_certificate_handle(self.native(), certificate.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// LookupCertificateForHandle is a wrapper around g_tls_database_lookup_certificate_for_handle().
func (self TlsDatabase) LookupCertificateForHandle(handle string, interaction TlsInteraction, flags TlsDatabaseLookupFlags, cancellable Cancellable) (TlsCertificate, error) {
	handle0 := (*C.gchar)(C.CString(handle))
	var err glib.Error
	ret0 := C.g_tls_database_lookup_certificate_for_handle(self.native(), handle0, interaction.native(), C.GTlsDatabaseLookupFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(handle0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return TlsCertificate{}, err.GoValue()
	}
	return wrapTlsCertificate(ret0), nil
}

// LookupCertificateForHandleAsync is a wrapper around g_tls_database_lookup_certificate_for_handle_async().
func (self TlsDatabase) LookupCertificateForHandleAsync(handle string, interaction TlsInteraction, flags TlsDatabaseLookupFlags, cancellable Cancellable, callback AsyncReadyCallback) {
	handle0 := (*C.gchar)(C.CString(handle))
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_tls_database_lookup_certificate_for_handle_async(self.native(), handle0, interaction.native(), C.GTlsDatabaseLookupFlags(flags), cancellable.native(), callback0)
	C.free(unsafe.Pointer(handle0)) /*ch:<stdlib.h>*/
}

// LookupCertificateForHandleFinish is a wrapper around g_tls_database_lookup_certificate_for_handle_finish().
func (self TlsDatabase) LookupCertificateForHandleFinish(result AsyncResult) (TlsCertificate, error) {
	var err glib.Error
	ret0 := C.g_tls_database_lookup_certificate_for_handle_finish(self.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return TlsCertificate{}, err.GoValue()
	}
	return wrapTlsCertificate(ret0), nil
}

// LookupCertificateIssuer is a wrapper around g_tls_database_lookup_certificate_issuer().
func (self TlsDatabase) LookupCertificateIssuer(certificate TlsCertificate, interaction TlsInteraction, flags TlsDatabaseLookupFlags, cancellable Cancellable) (TlsCertificate, error) {
	var err glib.Error
	ret0 := C.g_tls_database_lookup_certificate_issuer(self.native(), certificate.native(), interaction.native(), C.GTlsDatabaseLookupFlags(flags), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return TlsCertificate{}, err.GoValue()
	}
	return wrapTlsCertificate(ret0), nil
}

// LookupCertificateIssuerAsync is a wrapper around g_tls_database_lookup_certificate_issuer_async().
func (self TlsDatabase) LookupCertificateIssuerAsync(certificate TlsCertificate, interaction TlsInteraction, flags TlsDatabaseLookupFlags, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_tls_database_lookup_certificate_issuer_async(self.native(), certificate.native(), interaction.native(), C.GTlsDatabaseLookupFlags(flags), cancellable.native(), callback0)
}

// LookupCertificateIssuerFinish is a wrapper around g_tls_database_lookup_certificate_issuer_finish().
func (self TlsDatabase) LookupCertificateIssuerFinish(result AsyncResult) (TlsCertificate, error) {
	var err glib.Error
	ret0 := C.g_tls_database_lookup_certificate_issuer_finish(self.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return TlsCertificate{}, err.GoValue()
	}
	return wrapTlsCertificate(ret0), nil
}

// LookupCertificatesIssuedByFinish is a wrapper around g_tls_database_lookup_certificates_issued_by_finish().
func (self TlsDatabase) LookupCertificatesIssuedByFinish(result AsyncResult) (glib.List, error) {
	var err glib.Error
	ret0 := C.g_tls_database_lookup_certificates_issued_by_finish(self.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return glib.List{}, err.GoValue()
	}
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapTlsCertificate(p) }), /*gir:GLib*/ nil
}

// VerifyChainAsync is a wrapper around g_tls_database_verify_chain_async().
func (self TlsDatabase) VerifyChainAsync(chain TlsCertificate, purpose string, identity SocketConnectable, interaction TlsInteraction, flags TlsDatabaseVerifyFlags, cancellable Cancellable, callback AsyncReadyCallback) {
	purpose0 := (*C.gchar)(C.CString(purpose))
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_tls_database_verify_chain_async(self.native(), chain.native(), purpose0, identity.native(), interaction.native(), C.GTlsDatabaseVerifyFlags(flags), cancellable.native(), callback0)
	C.free(unsafe.Pointer(purpose0)) /*ch:<stdlib.h>*/
}

// Object TlsInteraction
type TlsInteraction struct {
	gobject.Object
}

func (v TlsInteraction) native() *C.GTlsInteraction {
	return (*C.GTlsInteraction)(v.Ptr)
}
func wrapTlsInteraction(p *C.GTlsInteraction) (v TlsInteraction) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTlsInteraction(p unsafe.Pointer) (v TlsInteraction) {
	v.Ptr = p
	return
}
func (v TlsInteraction) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTlsInteraction(p unsafe.Pointer) interface{} {
	return WrapTlsInteraction(p)
}
func (v TlsInteraction) GetType() gobject.Type {
	return gobject.Type(C.g_tls_interaction_get_type())
}
func (v TlsInteraction) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTlsInteraction(unsafe.Pointer(ptr)), nil
	}
}

// AskPasswordAsync is a wrapper around g_tls_interaction_ask_password_async().
func (interaction TlsInteraction) AskPasswordAsync(password TlsPassword, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_tls_interaction_ask_password_async(interaction.native(), password.native(), cancellable.native(), callback0)
}

// RequestCertificateAsync is a wrapper around g_tls_interaction_request_certificate_async().
func (interaction TlsInteraction) RequestCertificateAsync(connection TlsConnection, flags TlsCertificateRequestFlags, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_tls_interaction_request_certificate_async(interaction.native(), connection.native(), C.GTlsCertificateRequestFlags(flags), cancellable.native(), callback0)
}

// Object TlsPassword
type TlsPassword struct {
	gobject.Object
}

func (v TlsPassword) native() *C.GTlsPassword {
	return (*C.GTlsPassword)(v.Ptr)
}
func wrapTlsPassword(p *C.GTlsPassword) (v TlsPassword) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTlsPassword(p unsafe.Pointer) (v TlsPassword) {
	v.Ptr = p
	return
}
func (v TlsPassword) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTlsPassword(p unsafe.Pointer) interface{} {
	return WrapTlsPassword(p)
}
func (v TlsPassword) GetType() gobject.Type {
	return gobject.Type(C.g_tls_password_get_type())
}
func (v TlsPassword) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTlsPassword(unsafe.Pointer(ptr)), nil
	}
}

// TlsPasswordNew is a wrapper around g_tls_password_new().
func TlsPasswordNew(flags TlsPasswordFlags, description string) TlsPassword {
	description0 := (*C.gchar)(C.CString(description))
	ret0 := C.g_tls_password_new(C.GTlsPasswordFlags(flags), description0)
	C.free(unsafe.Pointer(description0)) /*ch:<stdlib.h>*/
	return wrapTlsPassword(ret0)
}

// GetDescription is a wrapper around g_tls_password_get_description().
func (password TlsPassword) GetDescription() string {
	ret0 := C.g_tls_password_get_description(password.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetFlags is a wrapper around g_tls_password_get_flags().
func (password TlsPassword) GetFlags() TlsPasswordFlags {
	ret0 := C.g_tls_password_get_flags(password.native())
	return TlsPasswordFlags(ret0)
}

// GetWarning is a wrapper around g_tls_password_get_warning().
func (password TlsPassword) GetWarning() string {
	ret0 := C.g_tls_password_get_warning(password.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// SetDescription is a wrapper around g_tls_password_set_description().
func (password TlsPassword) SetDescription(description string) {
	description0 := (*C.gchar)(C.CString(description))
	C.g_tls_password_set_description(password.native(), description0)
	C.free(unsafe.Pointer(description0)) /*ch:<stdlib.h>*/
}

// SetFlags is a wrapper around g_tls_password_set_flags().
func (password TlsPassword) SetFlags(flags TlsPasswordFlags) {
	C.g_tls_password_set_flags(password.native(), C.GTlsPasswordFlags(flags))
}

// SetWarning is a wrapper around g_tls_password_set_warning().
func (password TlsPassword) SetWarning(warning string) {
	warning0 := (*C.gchar)(C.CString(warning))
	C.g_tls_password_set_warning(password.native(), warning0)
	C.free(unsafe.Pointer(warning0)) /*ch:<stdlib.h>*/
}

// Object TlsConnection
type TlsConnection struct {
	IOStream
}

func (v TlsConnection) native() *C.GTlsConnection {
	return (*C.GTlsConnection)(v.Ptr)
}
func wrapTlsConnection(p *C.GTlsConnection) (v TlsConnection) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTlsConnection(p unsafe.Pointer) (v TlsConnection) {
	v.Ptr = p
	return
}
func (v TlsConnection) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTlsConnection(p unsafe.Pointer) interface{} {
	return WrapTlsConnection(p)
}
func (v TlsConnection) GetType() gobject.Type {
	return gobject.Type(C.g_tls_connection_get_type())
}
func (v TlsConnection) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTlsConnection(unsafe.Pointer(ptr)), nil
	}
}

// EmitAcceptCertificate is a wrapper around g_tls_connection_emit_accept_certificate().
func (conn TlsConnection) EmitAcceptCertificate(peer_cert TlsCertificate, errors TlsCertificateFlags) bool {
	ret0 := C.g_tls_connection_emit_accept_certificate(conn.native(), peer_cert.native(), C.GTlsCertificateFlags(errors))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetCertificate is a wrapper around g_tls_connection_get_certificate().
func (conn TlsConnection) GetCertificate() TlsCertificate {
	ret0 := C.g_tls_connection_get_certificate(conn.native())
	return wrapTlsCertificate(ret0)
}

// GetDatabase is a wrapper around g_tls_connection_get_database().
func (conn TlsConnection) GetDatabase() TlsDatabase {
	ret0 := C.g_tls_connection_get_database(conn.native())
	return wrapTlsDatabase(ret0)
}

// GetInteraction is a wrapper around g_tls_connection_get_interaction().
func (conn TlsConnection) GetInteraction() TlsInteraction {
	ret0 := C.g_tls_connection_get_interaction(conn.native())
	return wrapTlsInteraction(ret0)
}

// GetPeerCertificate is a wrapper around g_tls_connection_get_peer_certificate().
func (conn TlsConnection) GetPeerCertificate() TlsCertificate {
	ret0 := C.g_tls_connection_get_peer_certificate(conn.native())
	return wrapTlsCertificate(ret0)
}

// GetPeerCertificateErrors is a wrapper around g_tls_connection_get_peer_certificate_errors().
func (conn TlsConnection) GetPeerCertificateErrors() TlsCertificateFlags {
	ret0 := C.g_tls_connection_get_peer_certificate_errors(conn.native())
	return TlsCertificateFlags(ret0)
}

// GetRehandshakeMode is a wrapper around g_tls_connection_get_rehandshake_mode().
func (conn TlsConnection) GetRehandshakeMode() TlsRehandshakeMode {
	ret0 := C.g_tls_connection_get_rehandshake_mode(conn.native())
	return TlsRehandshakeMode(ret0)
}

// GetRequireCloseNotify is a wrapper around g_tls_connection_get_require_close_notify().
func (conn TlsConnection) GetRequireCloseNotify() bool {
	ret0 := C.g_tls_connection_get_require_close_notify(conn.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Handshake is a wrapper around g_tls_connection_handshake().
func (conn TlsConnection) Handshake(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_tls_connection_handshake(conn.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// HandshakeAsync is a wrapper around g_tls_connection_handshake_async().
func (conn TlsConnection) HandshakeAsync(io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_tls_connection_handshake_async(conn.native(), C.int(io_priority), cancellable.native(), callback0)
}

// HandshakeFinish is a wrapper around g_tls_connection_handshake_finish().
func (conn TlsConnection) HandshakeFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_tls_connection_handshake_finish(conn.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetCertificate is a wrapper around g_tls_connection_set_certificate().
func (conn TlsConnection) SetCertificate(certificate TlsCertificate) {
	C.g_tls_connection_set_certificate(conn.native(), certificate.native())
}

// SetDatabase is a wrapper around g_tls_connection_set_database().
func (conn TlsConnection) SetDatabase(database TlsDatabase) {
	C.g_tls_connection_set_database(conn.native(), database.native())
}

// SetInteraction is a wrapper around g_tls_connection_set_interaction().
func (conn TlsConnection) SetInteraction(interaction TlsInteraction) {
	C.g_tls_connection_set_interaction(conn.native(), interaction.native())
}

// SetRehandshakeMode is a wrapper around g_tls_connection_set_rehandshake_mode().
func (conn TlsConnection) SetRehandshakeMode(mode TlsRehandshakeMode) {
	C.g_tls_connection_set_rehandshake_mode(conn.native(), C.GTlsRehandshakeMode(mode))
}

// SetRequireCloseNotify is a wrapper around g_tls_connection_set_require_close_notify().
func (conn TlsConnection) SetRequireCloseNotify(require_close_notify bool) {
	C.g_tls_connection_set_require_close_notify(conn.native(), C.gboolean(util.Bool2Int(require_close_notify)) /*go:.util*/)
}

// Interface TlsClientConnection
type TlsClientConnection struct {
	Ptr unsafe.Pointer
}

func (v TlsClientConnection) native() *C.GTlsClientConnection {
	return (*C.GTlsClientConnection)(v.Ptr)
}
func wrapTlsClientConnection(p *C.GTlsClientConnection) TlsClientConnection {
	return TlsClientConnection{unsafe.Pointer(p)}
}
func WrapTlsClientConnection(p unsafe.Pointer) TlsClientConnection {
	return TlsClientConnection{p}
}
func (v TlsClientConnection) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTlsClientConnection(p unsafe.Pointer) interface{} {
	return WrapTlsClientConnection(p)
}
func (v TlsClientConnection) GetType() gobject.Type {
	return gobject.Type(C.g_tls_client_connection_get_type())
}
func (v TlsClientConnection) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTlsClientConnection(unsafe.Pointer(ptr)), nil
	}
}

// CopySessionState is a wrapper around g_tls_client_connection_copy_session_state().
func (conn TlsClientConnection) CopySessionState(source TlsClientConnection) {
	C.g_tls_client_connection_copy_session_state(conn.native(), source.native())
}

// GetServerIdentity is a wrapper around g_tls_client_connection_get_server_identity().
func (conn TlsClientConnection) GetServerIdentity() SocketConnectable {
	ret0 := C.g_tls_client_connection_get_server_identity(conn.native())
	return wrapSocketConnectable(ret0)
}

// GetUseSsl3 is a wrapper around g_tls_client_connection_get_use_ssl3().
func (conn TlsClientConnection) GetUseSsl3() bool {
	ret0 := C.g_tls_client_connection_get_use_ssl3(conn.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetValidationFlags is a wrapper around g_tls_client_connection_get_validation_flags().
func (conn TlsClientConnection) GetValidationFlags() TlsCertificateFlags {
	ret0 := C.g_tls_client_connection_get_validation_flags(conn.native())
	return TlsCertificateFlags(ret0)
}

// SetServerIdentity is a wrapper around g_tls_client_connection_set_server_identity().
func (conn TlsClientConnection) SetServerIdentity(identity SocketConnectable) {
	C.g_tls_client_connection_set_server_identity(conn.native(), identity.native())
}

// SetUseSsl3 is a wrapper around g_tls_client_connection_set_use_ssl3().
func (conn TlsClientConnection) SetUseSsl3(use_ssl3 bool) {
	C.g_tls_client_connection_set_use_ssl3(conn.native(), C.gboolean(util.Bool2Int(use_ssl3)) /*go:.util*/)
}

// SetValidationFlags is a wrapper around g_tls_client_connection_set_validation_flags().
func (conn TlsClientConnection) SetValidationFlags(flags TlsCertificateFlags) {
	C.g_tls_client_connection_set_validation_flags(conn.native(), C.GTlsCertificateFlags(flags))
}

// Interface TlsFileDatabase
type TlsFileDatabase struct {
	Ptr unsafe.Pointer
}

func (v TlsFileDatabase) native() *C.GTlsFileDatabase {
	return (*C.GTlsFileDatabase)(v.Ptr)
}
func wrapTlsFileDatabase(p *C.GTlsFileDatabase) TlsFileDatabase {
	return TlsFileDatabase{unsafe.Pointer(p)}
}
func WrapTlsFileDatabase(p unsafe.Pointer) TlsFileDatabase {
	return TlsFileDatabase{p}
}
func (v TlsFileDatabase) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTlsFileDatabase(p unsafe.Pointer) interface{} {
	return WrapTlsFileDatabase(p)
}
func (v TlsFileDatabase) GetType() gobject.Type {
	return gobject.Type(C.g_tls_file_database_get_type())
}
func (v TlsFileDatabase) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTlsFileDatabase(unsafe.Pointer(ptr)), nil
	}
}

// Interface TlsServerConnection
type TlsServerConnection struct {
	Ptr unsafe.Pointer
}

func (v TlsServerConnection) native() *C.GTlsServerConnection {
	return (*C.GTlsServerConnection)(v.Ptr)
}
func wrapTlsServerConnection(p *C.GTlsServerConnection) TlsServerConnection {
	return TlsServerConnection{unsafe.Pointer(p)}
}
func WrapTlsServerConnection(p unsafe.Pointer) TlsServerConnection {
	return TlsServerConnection{p}
}
func (v TlsServerConnection) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTlsServerConnection(p unsafe.Pointer) interface{} {
	return WrapTlsServerConnection(p)
}
func (v TlsServerConnection) GetType() gobject.Type {
	return gobject.Type(C.g_tls_server_connection_get_type())
}
func (v TlsServerConnection) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTlsServerConnection(unsafe.Pointer(ptr)), nil
	}
}

// Object AppInfoMonitor
type AppInfoMonitor struct {
	gobject.Object
}

func (v AppInfoMonitor) native() *C.GAppInfoMonitor {
	return (*C.GAppInfoMonitor)(v.Ptr)
}
func wrapAppInfoMonitor(p *C.GAppInfoMonitor) (v AppInfoMonitor) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapAppInfoMonitor(p unsafe.Pointer) (v AppInfoMonitor) {
	v.Ptr = p
	return
}
func (v AppInfoMonitor) IsNil() bool {
	return v.Ptr == nil
}
func IWrapAppInfoMonitor(p unsafe.Pointer) interface{} {
	return WrapAppInfoMonitor(p)
}
func (v AppInfoMonitor) GetType() gobject.Type {
	return gobject.Type(C.g_app_info_monitor_get_type())
}
func (v AppInfoMonitor) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapAppInfoMonitor(unsafe.Pointer(ptr)), nil
	}
}

// AppInfoMonitorGet is a wrapper around g_app_info_monitor_get().
func AppInfoMonitorGet() AppInfoMonitor {
	ret0 := C.g_app_info_monitor_get()
	return wrapAppInfoMonitor(ret0)
}

// Object ApplicationCommandLine
type ApplicationCommandLine struct {
	gobject.Object
}

func (v ApplicationCommandLine) native() *C.GApplicationCommandLine {
	return (*C.GApplicationCommandLine)(v.Ptr)
}
func wrapApplicationCommandLine(p *C.GApplicationCommandLine) (v ApplicationCommandLine) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapApplicationCommandLine(p unsafe.Pointer) (v ApplicationCommandLine) {
	v.Ptr = p
	return
}
func (v ApplicationCommandLine) IsNil() bool {
	return v.Ptr == nil
}
func IWrapApplicationCommandLine(p unsafe.Pointer) interface{} {
	return WrapApplicationCommandLine(p)
}
func (v ApplicationCommandLine) GetType() gobject.Type {
	return gobject.Type(C.g_application_command_line_get_type())
}
func (v ApplicationCommandLine) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapApplicationCommandLine(unsafe.Pointer(ptr)), nil
	}
}

// CreateFileForArg is a wrapper around g_application_command_line_create_file_for_arg().
func (cmdline ApplicationCommandLine) CreateFileForArg(arg string) File {
	arg0 := (*C.gchar)(C.CString(arg))
	ret0 := C.g_application_command_line_create_file_for_arg(cmdline.native(), arg0)
	C.free(unsafe.Pointer(arg0)) /*ch:<stdlib.h>*/
	return wrapFile(ret0)
}

// GetArguments is a wrapper around g_application_command_line_get_arguments().
func (cmdline ApplicationCommandLine) GetArguments() []string {
	var argc0 C.int
	ret0 := C.g_application_command_line_get_arguments(cmdline.native(), &argc0)
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(argc0)) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetCwd is a wrapper around g_application_command_line_get_cwd().
func (cmdline ApplicationCommandLine) GetCwd() string {
	ret0 := C.g_application_command_line_get_cwd(cmdline.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetEnviron is a wrapper around g_application_command_line_get_environ().
func (cmdline ApplicationCommandLine) GetEnviron() []string {
	ret0 := C.g_application_command_line_get_environ(cmdline.native())
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
	}
	return ret
}

// GetExitStatus is a wrapper around g_application_command_line_get_exit_status().
func (cmdline ApplicationCommandLine) GetExitStatus() int {
	ret0 := C.g_application_command_line_get_exit_status(cmdline.native())
	return int(ret0)
}

// GetIsRemote is a wrapper around g_application_command_line_get_is_remote().
func (cmdline ApplicationCommandLine) GetIsRemote() bool {
	ret0 := C.g_application_command_line_get_is_remote(cmdline.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetPlatformData is a wrapper around g_application_command_line_get_platform_data().
func (cmdline ApplicationCommandLine) GetPlatformData() glib.Variant {
	ret0 := C.g_application_command_line_get_platform_data(cmdline.native())
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetStdin is a wrapper around g_application_command_line_get_stdin().
func (cmdline ApplicationCommandLine) GetStdin() InputStream {
	ret0 := C.g_application_command_line_get_stdin(cmdline.native())
	return wrapInputStream(ret0)
}

// Getenv is a wrapper around g_application_command_line_getenv().
func (cmdline ApplicationCommandLine) Getenv(name string) string {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_application_command_line_getenv(cmdline.native(), name0)
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// SetExitStatus is a wrapper around g_application_command_line_set_exit_status().
func (cmdline ApplicationCommandLine) SetExitStatus(exit_status int) {
	C.g_application_command_line_set_exit_status(cmdline.native(), C.int(exit_status))
}

// Object BufferedInputStream
type BufferedInputStream struct {
	FilterInputStream
}

func (v BufferedInputStream) native() *C.GBufferedInputStream {
	return (*C.GBufferedInputStream)(v.Ptr)
}
func wrapBufferedInputStream(p *C.GBufferedInputStream) (v BufferedInputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapBufferedInputStream(p unsafe.Pointer) (v BufferedInputStream) {
	v.Ptr = p
	return
}
func (v BufferedInputStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapBufferedInputStream(p unsafe.Pointer) interface{} {
	return WrapBufferedInputStream(p)
}
func (v BufferedInputStream) GetType() gobject.Type {
	return gobject.Type(C.g_buffered_input_stream_get_type())
}
func (v BufferedInputStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapBufferedInputStream(unsafe.Pointer(ptr)), nil
	}
}
func (v BufferedInputStream) Seekable() Seekable {
	return WrapSeekable(v.Ptr)
}

// BufferedInputStreamNew is a wrapper around g_buffered_input_stream_new().
func BufferedInputStreamNew(base_stream InputStream) InputStream {
	ret0 := C.g_buffered_input_stream_new(base_stream.native())
	return wrapInputStream(ret0)
}

// BufferedInputStreamNewSized is a wrapper around g_buffered_input_stream_new_sized().
func BufferedInputStreamNewSized(base_stream InputStream, size uint) InputStream {
	ret0 := C.g_buffered_input_stream_new_sized(base_stream.native(), C.gsize(size))
	return wrapInputStream(ret0)
}

// Fill is a wrapper around g_buffered_input_stream_fill().
func (stream BufferedInputStream) Fill(count int, cancellable Cancellable) (int, error) {
	var err glib.Error
	ret0 := C.g_buffered_input_stream_fill(stream.native(), C.gssize(count), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// FillAsync is a wrapper around g_buffered_input_stream_fill_async().
func (stream BufferedInputStream) FillAsync(count int, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_buffered_input_stream_fill_async(stream.native(), C.gssize(count), C.int(io_priority), cancellable.native(), callback0)
}

// FillFinish is a wrapper around g_buffered_input_stream_fill_finish().
func (stream BufferedInputStream) FillFinish(result AsyncResult) (int, error) {
	var err glib.Error
	ret0 := C.g_buffered_input_stream_fill_finish(stream.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// GetAvailable is a wrapper around g_buffered_input_stream_get_available().
func (stream BufferedInputStream) GetAvailable() uint {
	ret0 := C.g_buffered_input_stream_get_available(stream.native())
	return uint(ret0)
}

// GetBufferSize is a wrapper around g_buffered_input_stream_get_buffer_size().
func (stream BufferedInputStream) GetBufferSize() uint {
	ret0 := C.g_buffered_input_stream_get_buffer_size(stream.native())
	return uint(ret0)
}

// ReadByte is a wrapper around g_buffered_input_stream_read_byte().
func (stream BufferedInputStream) ReadByte(cancellable Cancellable) (int, error) {
	var err glib.Error
	ret0 := C.g_buffered_input_stream_read_byte(stream.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// SetBufferSize is a wrapper around g_buffered_input_stream_set_buffer_size().
func (stream BufferedInputStream) SetBufferSize(size uint) {
	C.g_buffered_input_stream_set_buffer_size(stream.native(), C.gsize(size))
}

// Object FilterInputStream
type FilterInputStream struct {
	InputStream
}

func (v FilterInputStream) native() *C.GFilterInputStream {
	return (*C.GFilterInputStream)(v.Ptr)
}
func wrapFilterInputStream(p *C.GFilterInputStream) (v FilterInputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFilterInputStream(p unsafe.Pointer) (v FilterInputStream) {
	v.Ptr = p
	return
}
func (v FilterInputStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFilterInputStream(p unsafe.Pointer) interface{} {
	return WrapFilterInputStream(p)
}
func (v FilterInputStream) GetType() gobject.Type {
	return gobject.Type(C.g_filter_input_stream_get_type())
}
func (v FilterInputStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFilterInputStream(unsafe.Pointer(ptr)), nil
	}
}

// GetBaseStream is a wrapper around g_filter_input_stream_get_base_stream().
func (stream FilterInputStream) GetBaseStream() InputStream {
	ret0 := C.g_filter_input_stream_get_base_stream(stream.native())
	return wrapInputStream(ret0)
}

// GetCloseBaseStream is a wrapper around g_filter_input_stream_get_close_base_stream().
func (stream FilterInputStream) GetCloseBaseStream() bool {
	ret0 := C.g_filter_input_stream_get_close_base_stream(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetCloseBaseStream is a wrapper around g_filter_input_stream_set_close_base_stream().
func (stream FilterInputStream) SetCloseBaseStream(close_base bool) {
	C.g_filter_input_stream_set_close_base_stream(stream.native(), C.gboolean(util.Bool2Int(close_base)) /*go:.util*/)
}

// Object BufferedOutputStream
type BufferedOutputStream struct {
	FilterOutputStream
}

func (v BufferedOutputStream) native() *C.GBufferedOutputStream {
	return (*C.GBufferedOutputStream)(v.Ptr)
}
func wrapBufferedOutputStream(p *C.GBufferedOutputStream) (v BufferedOutputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapBufferedOutputStream(p unsafe.Pointer) (v BufferedOutputStream) {
	v.Ptr = p
	return
}
func (v BufferedOutputStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapBufferedOutputStream(p unsafe.Pointer) interface{} {
	return WrapBufferedOutputStream(p)
}
func (v BufferedOutputStream) GetType() gobject.Type {
	return gobject.Type(C.g_buffered_output_stream_get_type())
}
func (v BufferedOutputStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapBufferedOutputStream(unsafe.Pointer(ptr)), nil
	}
}
func (v BufferedOutputStream) Seekable() Seekable {
	return WrapSeekable(v.Ptr)
}

// BufferedOutputStreamNew is a wrapper around g_buffered_output_stream_new().
func BufferedOutputStreamNew(base_stream OutputStream) OutputStream {
	ret0 := C.g_buffered_output_stream_new(base_stream.native())
	return wrapOutputStream(ret0)
}

// BufferedOutputStreamNewSized is a wrapper around g_buffered_output_stream_new_sized().
func BufferedOutputStreamNewSized(base_stream OutputStream, size uint) OutputStream {
	ret0 := C.g_buffered_output_stream_new_sized(base_stream.native(), C.gsize(size))
	return wrapOutputStream(ret0)
}

// GetAutoGrow is a wrapper around g_buffered_output_stream_get_auto_grow().
func (stream BufferedOutputStream) GetAutoGrow() bool {
	ret0 := C.g_buffered_output_stream_get_auto_grow(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetBufferSize is a wrapper around g_buffered_output_stream_get_buffer_size().
func (stream BufferedOutputStream) GetBufferSize() uint {
	ret0 := C.g_buffered_output_stream_get_buffer_size(stream.native())
	return uint(ret0)
}

// SetAutoGrow is a wrapper around g_buffered_output_stream_set_auto_grow().
func (stream BufferedOutputStream) SetAutoGrow(auto_grow bool) {
	C.g_buffered_output_stream_set_auto_grow(stream.native(), C.gboolean(util.Bool2Int(auto_grow)) /*go:.util*/)
}

// SetBufferSize is a wrapper around g_buffered_output_stream_set_buffer_size().
func (stream BufferedOutputStream) SetBufferSize(size uint) {
	C.g_buffered_output_stream_set_buffer_size(stream.native(), C.gsize(size))
}

// Object FilterOutputStream
type FilterOutputStream struct {
	OutputStream
}

func (v FilterOutputStream) native() *C.GFilterOutputStream {
	return (*C.GFilterOutputStream)(v.Ptr)
}
func wrapFilterOutputStream(p *C.GFilterOutputStream) (v FilterOutputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFilterOutputStream(p unsafe.Pointer) (v FilterOutputStream) {
	v.Ptr = p
	return
}
func (v FilterOutputStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFilterOutputStream(p unsafe.Pointer) interface{} {
	return WrapFilterOutputStream(p)
}
func (v FilterOutputStream) GetType() gobject.Type {
	return gobject.Type(C.g_filter_output_stream_get_type())
}
func (v FilterOutputStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFilterOutputStream(unsafe.Pointer(ptr)), nil
	}
}

// GetBaseStream is a wrapper around g_filter_output_stream_get_base_stream().
func (stream FilterOutputStream) GetBaseStream() OutputStream {
	ret0 := C.g_filter_output_stream_get_base_stream(stream.native())
	return wrapOutputStream(ret0)
}

// GetCloseBaseStream is a wrapper around g_filter_output_stream_get_close_base_stream().
func (stream FilterOutputStream) GetCloseBaseStream() bool {
	ret0 := C.g_filter_output_stream_get_close_base_stream(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetCloseBaseStream is a wrapper around g_filter_output_stream_set_close_base_stream().
func (stream FilterOutputStream) SetCloseBaseStream(close_base bool) {
	C.g_filter_output_stream_set_close_base_stream(stream.native(), C.gboolean(util.Bool2Int(close_base)) /*go:.util*/)
}

// Object BytesIcon
type BytesIcon struct {
	gobject.Object
}

func (v BytesIcon) native() *C.GBytesIcon {
	return (*C.GBytesIcon)(v.Ptr)
}
func wrapBytesIcon(p *C.GBytesIcon) (v BytesIcon) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapBytesIcon(p unsafe.Pointer) (v BytesIcon) {
	v.Ptr = p
	return
}
func (v BytesIcon) IsNil() bool {
	return v.Ptr == nil
}
func IWrapBytesIcon(p unsafe.Pointer) interface{} {
	return WrapBytesIcon(p)
}
func (v BytesIcon) GetType() gobject.Type {
	return gobject.Type(C.g_bytes_icon_get_type())
}
func (v BytesIcon) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapBytesIcon(unsafe.Pointer(ptr)), nil
	}
}
func (v BytesIcon) Icon() Icon {
	return WrapIcon(v.Ptr)
}
func (v BytesIcon) LoadableIcon() LoadableIcon {
	return WrapLoadableIcon(v.Ptr)
}

// GetBytes is a wrapper around g_bytes_icon_get_bytes().
func (icon BytesIcon) GetBytes() glib.Bytes {
	ret0 := C.g_bytes_icon_get_bytes(icon.native())
	return glib.WrapBytes(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// Object CharsetConverter
type CharsetConverter struct {
	gobject.Object
}

func (v CharsetConverter) native() *C.GCharsetConverter {
	return (*C.GCharsetConverter)(v.Ptr)
}
func wrapCharsetConverter(p *C.GCharsetConverter) (v CharsetConverter) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCharsetConverter(p unsafe.Pointer) (v CharsetConverter) {
	v.Ptr = p
	return
}
func (v CharsetConverter) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCharsetConverter(p unsafe.Pointer) interface{} {
	return WrapCharsetConverter(p)
}
func (v CharsetConverter) GetType() gobject.Type {
	return gobject.Type(C.g_charset_converter_get_type())
}
func (v CharsetConverter) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCharsetConverter(unsafe.Pointer(ptr)), nil
	}
}
func (v CharsetConverter) Converter() Converter {
	return WrapConverter(v.Ptr)
}
func (v CharsetConverter) Initable() Initable {
	return WrapInitable(v.Ptr)
}

// CharsetConverterNew is a wrapper around g_charset_converter_new().
func CharsetConverterNew(to_charset string, from_charset string) (CharsetConverter, error) {
	to_charset0 := (*C.gchar)(C.CString(to_charset))
	from_charset0 := (*C.gchar)(C.CString(from_charset))
	var err glib.Error
	ret0 := C.g_charset_converter_new(to_charset0, from_charset0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(to_charset0))   /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(from_charset0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return CharsetConverter{}, err.GoValue()
	}
	return wrapCharsetConverter(ret0), nil
}

// GetNumFallbacks is a wrapper around g_charset_converter_get_num_fallbacks().
func (converter CharsetConverter) GetNumFallbacks() uint {
	ret0 := C.g_charset_converter_get_num_fallbacks(converter.native())
	return uint(ret0)
}

// GetUseFallback is a wrapper around g_charset_converter_get_use_fallback().
func (converter CharsetConverter) GetUseFallback() bool {
	ret0 := C.g_charset_converter_get_use_fallback(converter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetUseFallback is a wrapper around g_charset_converter_set_use_fallback().
func (converter CharsetConverter) SetUseFallback(use_fallback bool) {
	C.g_charset_converter_set_use_fallback(converter.native(), C.gboolean(util.Bool2Int(use_fallback)) /*go:.util*/)
}

// Object ConverterInputStream
type ConverterInputStream struct {
	FilterInputStream
}

func (v ConverterInputStream) native() *C.GConverterInputStream {
	return (*C.GConverterInputStream)(v.Ptr)
}
func wrapConverterInputStream(p *C.GConverterInputStream) (v ConverterInputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapConverterInputStream(p unsafe.Pointer) (v ConverterInputStream) {
	v.Ptr = p
	return
}
func (v ConverterInputStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapConverterInputStream(p unsafe.Pointer) interface{} {
	return WrapConverterInputStream(p)
}
func (v ConverterInputStream) GetType() gobject.Type {
	return gobject.Type(C.g_converter_input_stream_get_type())
}
func (v ConverterInputStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapConverterInputStream(unsafe.Pointer(ptr)), nil
	}
}
func (v ConverterInputStream) PollableInputStream() PollableInputStream {
	return WrapPollableInputStream(v.Ptr)
}

// ConverterInputStreamNew is a wrapper around g_converter_input_stream_new().
func ConverterInputStreamNew(base_stream InputStream, converter Converter) InputStream {
	ret0 := C.g_converter_input_stream_new(base_stream.native(), converter.native())
	return wrapInputStream(ret0)
}

// GetConverter is a wrapper around g_converter_input_stream_get_converter().
func (converter_stream ConverterInputStream) GetConverter() Converter {
	ret0 := C.g_converter_input_stream_get_converter(converter_stream.native())
	return wrapConverter(ret0)
}

// Object ConverterOutputStream
type ConverterOutputStream struct {
	FilterOutputStream
}

func (v ConverterOutputStream) native() *C.GConverterOutputStream {
	return (*C.GConverterOutputStream)(v.Ptr)
}
func wrapConverterOutputStream(p *C.GConverterOutputStream) (v ConverterOutputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapConverterOutputStream(p unsafe.Pointer) (v ConverterOutputStream) {
	v.Ptr = p
	return
}
func (v ConverterOutputStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapConverterOutputStream(p unsafe.Pointer) interface{} {
	return WrapConverterOutputStream(p)
}
func (v ConverterOutputStream) GetType() gobject.Type {
	return gobject.Type(C.g_converter_output_stream_get_type())
}
func (v ConverterOutputStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapConverterOutputStream(unsafe.Pointer(ptr)), nil
	}
}
func (v ConverterOutputStream) PollableOutputStream() PollableOutputStream {
	return WrapPollableOutputStream(v.Ptr)
}

// ConverterOutputStreamNew is a wrapper around g_converter_output_stream_new().
func ConverterOutputStreamNew(base_stream OutputStream, converter Converter) OutputStream {
	ret0 := C.g_converter_output_stream_new(base_stream.native(), converter.native())
	return wrapOutputStream(ret0)
}

// GetConverter is a wrapper around g_converter_output_stream_get_converter().
func (converter_stream ConverterOutputStream) GetConverter() Converter {
	ret0 := C.g_converter_output_stream_get_converter(converter_stream.native())
	return wrapConverter(ret0)
}

// Object Credentials
type Credentials struct {
	gobject.Object
}

func (v Credentials) native() *C.GCredentials {
	return (*C.GCredentials)(v.Ptr)
}
func wrapCredentials(p *C.GCredentials) (v Credentials) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapCredentials(p unsafe.Pointer) (v Credentials) {
	v.Ptr = p
	return
}
func (v Credentials) IsNil() bool {
	return v.Ptr == nil
}
func IWrapCredentials(p unsafe.Pointer) interface{} {
	return WrapCredentials(p)
}
func (v Credentials) GetType() gobject.Type {
	return gobject.Type(C.g_credentials_get_type())
}
func (v Credentials) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapCredentials(unsafe.Pointer(ptr)), nil
	}
}

// CredentialsNew is a wrapper around g_credentials_new().
func CredentialsNew() Credentials {
	ret0 := C.g_credentials_new()
	return wrapCredentials(ret0)
}

// GetNative is a wrapper around g_credentials_get_native().
func (credentials Credentials) GetNative(native_type CredentialsType) unsafe.Pointer {
	ret0 := C.g_credentials_get_native(credentials.native(), C.GCredentialsType(native_type))
	return unsafe.Pointer(ret0)
}

// IsSameUser is a wrapper around g_credentials_is_same_user().
func (credentials Credentials) IsSameUser(other_credentials Credentials) (bool, error) {
	var err glib.Error
	ret0 := C.g_credentials_is_same_user(credentials.native(), other_credentials.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetNative is a wrapper around g_credentials_set_native().
func (credentials Credentials) SetNative(native_type CredentialsType, native unsafe.Pointer) {
	C.g_credentials_set_native(credentials.native(), C.GCredentialsType(native_type), C.gpointer(native))
}

// ToString is a wrapper around g_credentials_to_string().
func (credentials Credentials) ToString() string {
	ret0 := C.g_credentials_to_string(credentials.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// Object DBusActionGroup
type DBusActionGroup struct {
	gobject.Object
}

func (v DBusActionGroup) native() *C.GDBusActionGroup {
	return (*C.GDBusActionGroup)(v.Ptr)
}
func wrapDBusActionGroup(p *C.GDBusActionGroup) (v DBusActionGroup) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDBusActionGroup(p unsafe.Pointer) (v DBusActionGroup) {
	v.Ptr = p
	return
}
func (v DBusActionGroup) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusActionGroup(p unsafe.Pointer) interface{} {
	return WrapDBusActionGroup(p)
}
func (v DBusActionGroup) GetType() gobject.Type {
	return gobject.Type(C.g_dbus_action_group_get_type())
}
func (v DBusActionGroup) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDBusActionGroup(unsafe.Pointer(ptr)), nil
	}
}
func (v DBusActionGroup) ActionGroup() ActionGroup {
	return WrapActionGroup(v.Ptr)
}
func (v DBusActionGroup) RemoteActionGroup() RemoteActionGroup {
	return WrapRemoteActionGroup(v.Ptr)
}

// DBusActionGroupGet is a wrapper around g_dbus_action_group_get().
func DBusActionGroupGet(connection DBusConnection, bus_name string, object_path string) DBusActionGroup {
	bus_name0 := (*C.gchar)(C.CString(bus_name))
	object_path0 := (*C.gchar)(C.CString(object_path))
	ret0 := C.g_dbus_action_group_get(connection.native(), bus_name0, object_path0)
	C.free(unsafe.Pointer(bus_name0))    /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(object_path0)) /*ch:<stdlib.h>*/
	return wrapDBusActionGroup(ret0)
}

// Object DBusAuthObserver
type DBusAuthObserver struct {
	gobject.Object
}

func (v DBusAuthObserver) native() *C.GDBusAuthObserver {
	return (*C.GDBusAuthObserver)(v.Ptr)
}
func wrapDBusAuthObserver(p *C.GDBusAuthObserver) (v DBusAuthObserver) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDBusAuthObserver(p unsafe.Pointer) (v DBusAuthObserver) {
	v.Ptr = p
	return
}
func (v DBusAuthObserver) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusAuthObserver(p unsafe.Pointer) interface{} {
	return WrapDBusAuthObserver(p)
}
func (v DBusAuthObserver) GetType() gobject.Type {
	return gobject.Type(C.g_dbus_auth_observer_get_type())
}
func (v DBusAuthObserver) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDBusAuthObserver(unsafe.Pointer(ptr)), nil
	}
}

// DBusAuthObserverNew is a wrapper around g_dbus_auth_observer_new().
func DBusAuthObserverNew() DBusAuthObserver {
	ret0 := C.g_dbus_auth_observer_new()
	return wrapDBusAuthObserver(ret0)
}

// AllowMechanism is a wrapper around g_dbus_auth_observer_allow_mechanism().
func (observer DBusAuthObserver) AllowMechanism(mechanism string) bool {
	mechanism0 := (*C.gchar)(C.CString(mechanism))
	ret0 := C.g_dbus_auth_observer_allow_mechanism(observer.native(), mechanism0)
	C.free(unsafe.Pointer(mechanism0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))    /*go:.util*/
}

// AuthorizeAuthenticatedPeer is a wrapper around g_dbus_auth_observer_authorize_authenticated_peer().
func (observer DBusAuthObserver) AuthorizeAuthenticatedPeer(stream IOStream, credentials Credentials) bool {
	ret0 := C.g_dbus_auth_observer_authorize_authenticated_peer(observer.native(), stream.native(), credentials.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object DBusInterfaceSkeleton
type DBusInterfaceSkeleton struct {
	gobject.Object
}

func (v DBusInterfaceSkeleton) native() *C.GDBusInterfaceSkeleton {
	return (*C.GDBusInterfaceSkeleton)(v.Ptr)
}
func wrapDBusInterfaceSkeleton(p *C.GDBusInterfaceSkeleton) (v DBusInterfaceSkeleton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDBusInterfaceSkeleton(p unsafe.Pointer) (v DBusInterfaceSkeleton) {
	v.Ptr = p
	return
}
func (v DBusInterfaceSkeleton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusInterfaceSkeleton(p unsafe.Pointer) interface{} {
	return WrapDBusInterfaceSkeleton(p)
}
func (v DBusInterfaceSkeleton) GetType() gobject.Type {
	return gobject.Type(C.g_dbus_interface_skeleton_get_type())
}
func (v DBusInterfaceSkeleton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDBusInterfaceSkeleton(unsafe.Pointer(ptr)), nil
	}
}
func (v DBusInterfaceSkeleton) DBusInterface() DBusInterface {
	return WrapDBusInterface(v.Ptr)
}

// Export is a wrapper around g_dbus_interface_skeleton_export().
func (interface_ DBusInterfaceSkeleton) Export(connection DBusConnection, object_path string) (bool, error) {
	object_path0 := (*C.gchar)(C.CString(object_path))
	var err glib.Error
	ret0 := C.g_dbus_interface_skeleton_export(interface_.native(), connection.native(), object_path0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(object_path0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Flush is a wrapper around g_dbus_interface_skeleton_flush().
func (interface_ DBusInterfaceSkeleton) Flush() {
	C.g_dbus_interface_skeleton_flush(interface_.native())
}

// GetConnection is a wrapper around g_dbus_interface_skeleton_get_connection().
func (interface_ DBusInterfaceSkeleton) GetConnection() DBusConnection {
	ret0 := C.g_dbus_interface_skeleton_get_connection(interface_.native())
	return wrapDBusConnection(ret0)
}

// GetConnections is a wrapper around g_dbus_interface_skeleton_get_connections().
func (interface_ DBusInterfaceSkeleton) GetConnections() glib.List {
	ret0 := C.g_dbus_interface_skeleton_get_connections(interface_.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapDBusConnection(p) }) /*gir:GLib*/
}

// GetFlags is a wrapper around g_dbus_interface_skeleton_get_flags().
func (interface_ DBusInterfaceSkeleton) GetFlags() DBusInterfaceSkeletonFlags {
	ret0 := C.g_dbus_interface_skeleton_get_flags(interface_.native())
	return DBusInterfaceSkeletonFlags(ret0)
}

// GetInfo is a wrapper around g_dbus_interface_skeleton_get_info().
func (interface_ DBusInterfaceSkeleton) GetInfo() DBusInterfaceInfo {
	ret0 := C.g_dbus_interface_skeleton_get_info(interface_.native())
	return wrapDBusInterfaceInfo(ret0)
}

// GetObjectPath is a wrapper around g_dbus_interface_skeleton_get_object_path().
func (interface_ DBusInterfaceSkeleton) GetObjectPath() string {
	ret0 := C.g_dbus_interface_skeleton_get_object_path(interface_.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetProperties is a wrapper around g_dbus_interface_skeleton_get_properties().
func (interface_ DBusInterfaceSkeleton) GetProperties() glib.Variant {
	ret0 := C.g_dbus_interface_skeleton_get_properties(interface_.native())
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetVtable is a wrapper around g_dbus_interface_skeleton_get_vtable().
func (interface_ DBusInterfaceSkeleton) GetVtable() DBusInterfaceVTable {
	ret0 := C.g_dbus_interface_skeleton_get_vtable(interface_.native())
	return wrapDBusInterfaceVTable(ret0)
}

// HasConnection is a wrapper around g_dbus_interface_skeleton_has_connection().
func (interface_ DBusInterfaceSkeleton) HasConnection(connection DBusConnection) bool {
	ret0 := C.g_dbus_interface_skeleton_has_connection(interface_.native(), connection.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetFlags is a wrapper around g_dbus_interface_skeleton_set_flags().
func (interface_ DBusInterfaceSkeleton) SetFlags(flags DBusInterfaceSkeletonFlags) {
	C.g_dbus_interface_skeleton_set_flags(interface_.native(), C.GDBusInterfaceSkeletonFlags(flags))
}

// Unexport is a wrapper around g_dbus_interface_skeleton_unexport().
func (interface_ DBusInterfaceSkeleton) Unexport() {
	C.g_dbus_interface_skeleton_unexport(interface_.native())
}

// UnexportFromConnection is a wrapper around g_dbus_interface_skeleton_unexport_from_connection().
func (interface_ DBusInterfaceSkeleton) UnexportFromConnection(connection DBusConnection) {
	C.g_dbus_interface_skeleton_unexport_from_connection(interface_.native(), connection.native())
}

// Object DBusMenuModel
type DBusMenuModel struct {
	MenuModel
}

func (v DBusMenuModel) native() *C.GDBusMenuModel {
	return (*C.GDBusMenuModel)(v.Ptr)
}
func wrapDBusMenuModel(p *C.GDBusMenuModel) (v DBusMenuModel) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDBusMenuModel(p unsafe.Pointer) (v DBusMenuModel) {
	v.Ptr = p
	return
}
func (v DBusMenuModel) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusMenuModel(p unsafe.Pointer) interface{} {
	return WrapDBusMenuModel(p)
}
func (v DBusMenuModel) GetType() gobject.Type {
	return gobject.Type(C.g_dbus_menu_model_get_type())
}
func (v DBusMenuModel) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDBusMenuModel(unsafe.Pointer(ptr)), nil
	}
}

// DBusMenuModelGet is a wrapper around g_dbus_menu_model_get().
func DBusMenuModelGet(connection DBusConnection, bus_name string, object_path string) DBusMenuModel {
	bus_name0 := (*C.gchar)(C.CString(bus_name))
	object_path0 := (*C.gchar)(C.CString(object_path))
	ret0 := C.g_dbus_menu_model_get(connection.native(), bus_name0, object_path0)
	C.free(unsafe.Pointer(bus_name0))    /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(object_path0)) /*ch:<stdlib.h>*/
	return wrapDBusMenuModel(ret0)
}

// Object MenuModel
type MenuModel struct {
	gobject.Object
}

func (v MenuModel) native() *C.GMenuModel {
	return (*C.GMenuModel)(v.Ptr)
}
func wrapMenuModel(p *C.GMenuModel) (v MenuModel) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMenuModel(p unsafe.Pointer) (v MenuModel) {
	v.Ptr = p
	return
}
func (v MenuModel) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMenuModel(p unsafe.Pointer) interface{} {
	return WrapMenuModel(p)
}
func (v MenuModel) GetType() gobject.Type {
	return gobject.Type(C.g_menu_model_get_type())
}
func (v MenuModel) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMenuModel(unsafe.Pointer(ptr)), nil
	}
}

// GetItemAttributeValue is a wrapper around g_menu_model_get_item_attribute_value().
func (model MenuModel) GetItemAttributeValue(item_index int, attribute string, expected_type glib.VariantType) glib.Variant {
	attribute0 := (*C.gchar)(C.CString(attribute))
	ret0 := C.g_menu_model_get_item_attribute_value(model.native(), C.gint(item_index), attribute0, (*C.GVariantType)(expected_type.Ptr))
	C.free(unsafe.Pointer(attribute0))            /*ch:<stdlib.h>*/
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetItemLink is a wrapper around g_menu_model_get_item_link().
func (model MenuModel) GetItemLink(item_index int, link string) MenuModel {
	link0 := (*C.gchar)(C.CString(link))
	ret0 := C.g_menu_model_get_item_link(model.native(), C.gint(item_index), link0)
	C.free(unsafe.Pointer(link0)) /*ch:<stdlib.h>*/
	return wrapMenuModel(ret0)
}

// GetNItems is a wrapper around g_menu_model_get_n_items().
func (model MenuModel) GetNItems() int {
	ret0 := C.g_menu_model_get_n_items(model.native())
	return int(ret0)
}

// IsMutable is a wrapper around g_menu_model_is_mutable().
func (model MenuModel) IsMutable() bool {
	ret0 := C.g_menu_model_is_mutable(model.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ItemsChanged is a wrapper around g_menu_model_items_changed().
func (model MenuModel) ItemsChanged(position int, removed int, added int) {
	C.g_menu_model_items_changed(model.native(), C.gint(position), C.gint(removed), C.gint(added))
}

// IterateItemAttributes is a wrapper around g_menu_model_iterate_item_attributes().
func (model MenuModel) IterateItemAttributes(item_index int) MenuAttributeIter {
	ret0 := C.g_menu_model_iterate_item_attributes(model.native(), C.gint(item_index))
	return wrapMenuAttributeIter(ret0)
}

// IterateItemLinks is a wrapper around g_menu_model_iterate_item_links().
func (model MenuModel) IterateItemLinks(item_index int) MenuLinkIter {
	ret0 := C.g_menu_model_iterate_item_links(model.native(), C.gint(item_index))
	return wrapMenuLinkIter(ret0)
}

// Object MenuAttributeIter
type MenuAttributeIter struct {
	gobject.Object
}

func (v MenuAttributeIter) native() *C.GMenuAttributeIter {
	return (*C.GMenuAttributeIter)(v.Ptr)
}
func wrapMenuAttributeIter(p *C.GMenuAttributeIter) (v MenuAttributeIter) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMenuAttributeIter(p unsafe.Pointer) (v MenuAttributeIter) {
	v.Ptr = p
	return
}
func (v MenuAttributeIter) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMenuAttributeIter(p unsafe.Pointer) interface{} {
	return WrapMenuAttributeIter(p)
}
func (v MenuAttributeIter) GetType() gobject.Type {
	return gobject.Type(C.g_menu_attribute_iter_get_type())
}
func (v MenuAttributeIter) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMenuAttributeIter(unsafe.Pointer(ptr)), nil
	}
}

// GetName is a wrapper around g_menu_attribute_iter_get_name().
func (iter MenuAttributeIter) GetName() string {
	ret0 := C.g_menu_attribute_iter_get_name(iter.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetNext is a wrapper around g_menu_attribute_iter_get_next().
func (iter MenuAttributeIter) GetNext() (bool, string, glib.Variant) {
	var out_name0 *C.gchar
	var value0 *C.GVariant
	ret0 := C.g_menu_attribute_iter_get_next(iter.native(), &out_name0, &value0)
	out_name := C.GoString((*C.char)(out_name0))
	defer C.g_free(C.gpointer(out_name0))
	return util.Int2Bool(int(ret0)) /*go:.util*/, out_name, glib.WrapVariant(unsafe.Pointer(value0)) /*gir:GLib*/
}

// GetValue is a wrapper around g_menu_attribute_iter_get_value().
func (iter MenuAttributeIter) GetValue() glib.Variant {
	ret0 := C.g_menu_attribute_iter_get_value(iter.native())
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// Next is a wrapper around g_menu_attribute_iter_next().
func (iter MenuAttributeIter) Next() bool {
	ret0 := C.g_menu_attribute_iter_next(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object MenuLinkIter
type MenuLinkIter struct {
	gobject.Object
}

func (v MenuLinkIter) native() *C.GMenuLinkIter {
	return (*C.GMenuLinkIter)(v.Ptr)
}
func wrapMenuLinkIter(p *C.GMenuLinkIter) (v MenuLinkIter) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMenuLinkIter(p unsafe.Pointer) (v MenuLinkIter) {
	v.Ptr = p
	return
}
func (v MenuLinkIter) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMenuLinkIter(p unsafe.Pointer) interface{} {
	return WrapMenuLinkIter(p)
}
func (v MenuLinkIter) GetType() gobject.Type {
	return gobject.Type(C.g_menu_link_iter_get_type())
}
func (v MenuLinkIter) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMenuLinkIter(unsafe.Pointer(ptr)), nil
	}
}

// GetName is a wrapper around g_menu_link_iter_get_name().
func (iter MenuLinkIter) GetName() string {
	ret0 := C.g_menu_link_iter_get_name(iter.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetNext is a wrapper around g_menu_link_iter_get_next().
func (iter MenuLinkIter) GetNext() (bool, string, MenuModel) {
	var out_link0 *C.gchar
	var value0 *C.GMenuModel
	ret0 := C.g_menu_link_iter_get_next(iter.native(), &out_link0, &value0)
	out_link := C.GoString((*C.char)(out_link0))
	defer C.g_free(C.gpointer(out_link0))
	return util.Int2Bool(int(ret0)) /*go:.util*/, out_link, wrapMenuModel(value0)
}

// GetValue is a wrapper around g_menu_link_iter_get_value().
func (iter MenuLinkIter) GetValue() MenuModel {
	ret0 := C.g_menu_link_iter_get_value(iter.native())
	return wrapMenuModel(ret0)
}

// Next is a wrapper around g_menu_link_iter_next().
func (iter MenuLinkIter) Next() bool {
	ret0 := C.g_menu_link_iter_next(iter.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object DBusMessage
type DBusMessage struct {
	gobject.Object
}

func (v DBusMessage) native() *C.GDBusMessage {
	return (*C.GDBusMessage)(v.Ptr)
}
func wrapDBusMessage(p *C.GDBusMessage) (v DBusMessage) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDBusMessage(p unsafe.Pointer) (v DBusMessage) {
	v.Ptr = p
	return
}
func (v DBusMessage) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusMessage(p unsafe.Pointer) interface{} {
	return WrapDBusMessage(p)
}
func (v DBusMessage) GetType() gobject.Type {
	return gobject.Type(C.g_dbus_message_get_type())
}
func (v DBusMessage) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDBusMessage(unsafe.Pointer(ptr)), nil
	}
}

// DBusMessageNew is a wrapper around g_dbus_message_new().
func DBusMessageNew() DBusMessage {
	ret0 := C.g_dbus_message_new()
	return wrapDBusMessage(ret0)
}

// DBusMessageNewFromBlob is a wrapper around g_dbus_message_new_from_blob().
func DBusMessageNewFromBlob(blob []byte, capabilities DBusCapabilityFlags) (DBusMessage, error) {
	blob0 := make([]C.guchar, len(blob))
	for idx, elemG := range blob {
		blob0[idx] = C.guchar(elemG)
	}
	var blob0Ptr *C.guchar
	if len(blob0) > 0 {
		blob0Ptr = &blob0[0]
	}
	var err glib.Error
	ret0 := C.g_dbus_message_new_from_blob(blob0Ptr, C.gsize(len(blob)), C.GDBusCapabilityFlags(capabilities), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return DBusMessage{}, err.GoValue()
	}
	return wrapDBusMessage(ret0), nil
}

// DBusMessageNewMethodCall is a wrapper around g_dbus_message_new_method_call().
func DBusMessageNewMethodCall(name string, path string, interface_ string, method string) DBusMessage {
	name0 := (*C.gchar)(C.CString(name))
	path0 := (*C.gchar)(C.CString(path))
	interface_0 := (*C.gchar)(C.CString(interface_))
	method0 := (*C.gchar)(C.CString(method))
	ret0 := C.g_dbus_message_new_method_call(name0, path0, interface_0, method0)
	C.free(unsafe.Pointer(name0))       /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(path0))       /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(interface_0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(method0))     /*ch:<stdlib.h>*/
	return wrapDBusMessage(ret0)
}

// DBusMessageNewSignal is a wrapper around g_dbus_message_new_signal().
func DBusMessageNewSignal(path string, interface_ string, signal string) DBusMessage {
	path0 := (*C.gchar)(C.CString(path))
	interface_0 := (*C.gchar)(C.CString(interface_))
	signal0 := (*C.gchar)(C.CString(signal))
	ret0 := C.g_dbus_message_new_signal(path0, interface_0, signal0)
	C.free(unsafe.Pointer(path0))       /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(interface_0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(signal0))     /*ch:<stdlib.h>*/
	return wrapDBusMessage(ret0)
}

// Copy is a wrapper around g_dbus_message_copy().
func (message DBusMessage) Copy() (DBusMessage, error) {
	var err glib.Error
	ret0 := C.g_dbus_message_copy(message.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return DBusMessage{}, err.GoValue()
	}
	return wrapDBusMessage(ret0), nil
}

// GetArg0 is a wrapper around g_dbus_message_get_arg0().
func (message DBusMessage) GetArg0() string {
	ret0 := C.g_dbus_message_get_arg0(message.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetBody is a wrapper around g_dbus_message_get_body().
func (message DBusMessage) GetBody() glib.Variant {
	ret0 := C.g_dbus_message_get_body(message.native())
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetByteOrder is a wrapper around g_dbus_message_get_byte_order().
func (message DBusMessage) GetByteOrder() DBusMessageByteOrder {
	ret0 := C.g_dbus_message_get_byte_order(message.native())
	return DBusMessageByteOrder(ret0)
}

// GetDestination is a wrapper around g_dbus_message_get_destination().
func (message DBusMessage) GetDestination() string {
	ret0 := C.g_dbus_message_get_destination(message.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetErrorName is a wrapper around g_dbus_message_get_error_name().
func (message DBusMessage) GetErrorName() string {
	ret0 := C.g_dbus_message_get_error_name(message.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetFlags is a wrapper around g_dbus_message_get_flags().
func (message DBusMessage) GetFlags() DBusMessageFlags {
	ret0 := C.g_dbus_message_get_flags(message.native())
	return DBusMessageFlags(ret0)
}

// GetHeader is a wrapper around g_dbus_message_get_header().
func (message DBusMessage) GetHeader(header_field DBusMessageHeaderField) glib.Variant {
	ret0 := C.g_dbus_message_get_header(message.native(), C.GDBusMessageHeaderField(header_field))
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetHeaderFields is a wrapper around g_dbus_message_get_header_fields().
func (message DBusMessage) GetHeaderFields() []byte {
	ret0 := C.g_dbus_message_get_header_fields(message.native())
	var ret0Slice []C.guchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(C.guchar(0))) /*go:.util*/) /*go:.util*/
	ret := make([]byte, len(ret0Slice))
	for idx, elem := range ret0Slice {
		ret[idx] = byte(elem)
	}
	return ret
}

// GetInterface is a wrapper around g_dbus_message_get_interface().
func (message DBusMessage) GetInterface() string {
	ret0 := C.g_dbus_message_get_interface(message.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetLocked is a wrapper around g_dbus_message_get_locked().
func (message DBusMessage) GetLocked() bool {
	ret0 := C.g_dbus_message_get_locked(message.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetMember is a wrapper around g_dbus_message_get_member().
func (message DBusMessage) GetMember() string {
	ret0 := C.g_dbus_message_get_member(message.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetMessageType is a wrapper around g_dbus_message_get_message_type().
func (message DBusMessage) GetMessageType() DBusMessageType {
	ret0 := C.g_dbus_message_get_message_type(message.native())
	return DBusMessageType(ret0)
}

// GetNumUnixFds is a wrapper around g_dbus_message_get_num_unix_fds().
func (message DBusMessage) GetNumUnixFds() uint32 {
	ret0 := C.g_dbus_message_get_num_unix_fds(message.native())
	return uint32(ret0)
}

// GetPath is a wrapper around g_dbus_message_get_path().
func (message DBusMessage) GetPath() string {
	ret0 := C.g_dbus_message_get_path(message.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetReplySerial is a wrapper around g_dbus_message_get_reply_serial().
func (message DBusMessage) GetReplySerial() uint32 {
	ret0 := C.g_dbus_message_get_reply_serial(message.native())
	return uint32(ret0)
}

// GetSender is a wrapper around g_dbus_message_get_sender().
func (message DBusMessage) GetSender() string {
	ret0 := C.g_dbus_message_get_sender(message.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetSerial is a wrapper around g_dbus_message_get_serial().
func (message DBusMessage) GetSerial() uint32 {
	ret0 := C.g_dbus_message_get_serial(message.native())
	return uint32(ret0)
}

// GetSignature is a wrapper around g_dbus_message_get_signature().
func (message DBusMessage) GetSignature() string {
	ret0 := C.g_dbus_message_get_signature(message.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetUnixFdList is a wrapper around g_dbus_message_get_unix_fd_list().
func (message DBusMessage) GetUnixFdList() UnixFDList {
	ret0 := C.g_dbus_message_get_unix_fd_list(message.native())
	return wrapUnixFDList(ret0)
}

// Lock is a wrapper around g_dbus_message_lock().
func (message DBusMessage) Lock() {
	C.g_dbus_message_lock(message.native())
}

// NewMethodErrorLiteral is a wrapper around g_dbus_message_new_method_error_literal().
func (method_call_message DBusMessage) NewMethodErrorLiteral(error_name string, error_message string) DBusMessage {
	error_name0 := (*C.gchar)(C.CString(error_name))
	error_message0 := (*C.gchar)(C.CString(error_message))
	ret0 := C.g_dbus_message_new_method_error_literal(method_call_message.native(), error_name0, error_message0)
	C.free(unsafe.Pointer(error_name0))    /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(error_message0)) /*ch:<stdlib.h>*/
	return wrapDBusMessage(ret0)
}

// NewMethodReply is a wrapper around g_dbus_message_new_method_reply().
func (method_call_message DBusMessage) NewMethodReply() DBusMessage {
	ret0 := C.g_dbus_message_new_method_reply(method_call_message.native())
	return wrapDBusMessage(ret0)
}

// Print is a wrapper around g_dbus_message_print().
func (message DBusMessage) Print(indent uint) string {
	ret0 := C.g_dbus_message_print(message.native(), C.guint(indent))
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// SetBody is a wrapper around g_dbus_message_set_body().
func (message DBusMessage) SetBody(body glib.Variant) {
	C.g_dbus_message_set_body(message.native(), (*C.GVariant)(body.Ptr))
}

// SetByteOrder is a wrapper around g_dbus_message_set_byte_order().
func (message DBusMessage) SetByteOrder(byte_order DBusMessageByteOrder) {
	C.g_dbus_message_set_byte_order(message.native(), C.GDBusMessageByteOrder(byte_order))
}

// SetDestination is a wrapper around g_dbus_message_set_destination().
func (message DBusMessage) SetDestination(value string) {
	value0 := (*C.gchar)(C.CString(value))
	C.g_dbus_message_set_destination(message.native(), value0)
	C.free(unsafe.Pointer(value0)) /*ch:<stdlib.h>*/
}

// SetErrorName is a wrapper around g_dbus_message_set_error_name().
func (message DBusMessage) SetErrorName(value string) {
	value0 := (*C.gchar)(C.CString(value))
	C.g_dbus_message_set_error_name(message.native(), value0)
	C.free(unsafe.Pointer(value0)) /*ch:<stdlib.h>*/
}

// SetFlags is a wrapper around g_dbus_message_set_flags().
func (message DBusMessage) SetFlags(flags DBusMessageFlags) {
	C.g_dbus_message_set_flags(message.native(), C.GDBusMessageFlags(flags))
}

// SetHeader is a wrapper around g_dbus_message_set_header().
func (message DBusMessage) SetHeader(header_field DBusMessageHeaderField, value glib.Variant) {
	C.g_dbus_message_set_header(message.native(), C.GDBusMessageHeaderField(header_field), (*C.GVariant)(value.Ptr))
}

// SetInterface is a wrapper around g_dbus_message_set_interface().
func (message DBusMessage) SetInterface(value string) {
	value0 := (*C.gchar)(C.CString(value))
	C.g_dbus_message_set_interface(message.native(), value0)
	C.free(unsafe.Pointer(value0)) /*ch:<stdlib.h>*/
}

// SetMember is a wrapper around g_dbus_message_set_member().
func (message DBusMessage) SetMember(value string) {
	value0 := (*C.gchar)(C.CString(value))
	C.g_dbus_message_set_member(message.native(), value0)
	C.free(unsafe.Pointer(value0)) /*ch:<stdlib.h>*/
}

// SetMessageType is a wrapper around g_dbus_message_set_message_type().
func (message DBusMessage) SetMessageType(type_ DBusMessageType) {
	C.g_dbus_message_set_message_type(message.native(), C.GDBusMessageType(type_))
}

// SetNumUnixFds is a wrapper around g_dbus_message_set_num_unix_fds().
func (message DBusMessage) SetNumUnixFds(value uint32) {
	C.g_dbus_message_set_num_unix_fds(message.native(), C.guint32(value))
}

// SetPath is a wrapper around g_dbus_message_set_path().
func (message DBusMessage) SetPath(value string) {
	value0 := (*C.gchar)(C.CString(value))
	C.g_dbus_message_set_path(message.native(), value0)
	C.free(unsafe.Pointer(value0)) /*ch:<stdlib.h>*/
}

// SetReplySerial is a wrapper around g_dbus_message_set_reply_serial().
func (message DBusMessage) SetReplySerial(value uint32) {
	C.g_dbus_message_set_reply_serial(message.native(), C.guint32(value))
}

// SetSender is a wrapper around g_dbus_message_set_sender().
func (message DBusMessage) SetSender(value string) {
	value0 := (*C.gchar)(C.CString(value))
	C.g_dbus_message_set_sender(message.native(), value0)
	C.free(unsafe.Pointer(value0)) /*ch:<stdlib.h>*/
}

// SetSerial is a wrapper around g_dbus_message_set_serial().
func (message DBusMessage) SetSerial(serial uint32) {
	C.g_dbus_message_set_serial(message.native(), C.guint32(serial))
}

// SetSignature is a wrapper around g_dbus_message_set_signature().
func (message DBusMessage) SetSignature(value string) {
	value0 := (*C.gchar)(C.CString(value))
	C.g_dbus_message_set_signature(message.native(), value0)
	C.free(unsafe.Pointer(value0)) /*ch:<stdlib.h>*/
}

// SetUnixFdList is a wrapper around g_dbus_message_set_unix_fd_list().
func (message DBusMessage) SetUnixFdList(fd_list UnixFDList) {
	C.g_dbus_message_set_unix_fd_list(message.native(), fd_list.native())
}

// ToBlob is a wrapper around g_dbus_message_to_blob().
func (message DBusMessage) ToBlob(capabilities DBusCapabilityFlags) ([]byte, error) {
	var out_size0 C.gsize
	var err glib.Error
	ret0 := C.g_dbus_message_to_blob(message.native(), &out_size0, C.GDBusCapabilityFlags(capabilities), (**C.GError)(unsafe.Pointer(&err)))
	var ret0Slice []C.guchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(out_size0)) /*go:.util*/
	ret := make([]byte, len(ret0Slice))
	for idx, elem := range ret0Slice {
		ret[idx] = byte(elem)
	}
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return nil, err.GoValue()
	}
	return ret, nil
}

// ToGerror is a wrapper around g_dbus_message_to_gerror().
func (message DBusMessage) ToGerror() (bool, error) {
	var err glib.Error
	ret0 := C.g_dbus_message_to_gerror(message.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// DBusMessageBytesNeeded is a wrapper around g_dbus_message_bytes_needed().
func DBusMessageBytesNeeded(blob []byte) (int, error) {
	blob0 := make([]C.guchar, len(blob))
	for idx, elemG := range blob {
		blob0[idx] = C.guchar(elemG)
	}
	var blob0Ptr *C.guchar
	if len(blob0) > 0 {
		blob0Ptr = &blob0[0]
	}
	var err glib.Error
	ret0 := C.g_dbus_message_bytes_needed(blob0Ptr, C.gsize(len(blob)), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// Object UnixFDList
type UnixFDList struct {
	gobject.Object
}

func (v UnixFDList) native() *C.GUnixFDList {
	return (*C.GUnixFDList)(v.Ptr)
}
func wrapUnixFDList(p *C.GUnixFDList) (v UnixFDList) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapUnixFDList(p unsafe.Pointer) (v UnixFDList) {
	v.Ptr = p
	return
}
func (v UnixFDList) IsNil() bool {
	return v.Ptr == nil
}
func IWrapUnixFDList(p unsafe.Pointer) interface{} {
	return WrapUnixFDList(p)
}
func (v UnixFDList) GetType() gobject.Type {
	return gobject.Type(C.g_unix_fd_list_get_type())
}
func (v UnixFDList) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapUnixFDList(unsafe.Pointer(ptr)), nil
	}
}

// UnixFDListNew is a wrapper around g_unix_fd_list_new().
func UnixFDListNew() UnixFDList {
	ret0 := C.g_unix_fd_list_new()
	return wrapUnixFDList(ret0)
}

// UnixFDListNewFromArray is a wrapper around g_unix_fd_list_new_from_array().
func UnixFDListNewFromArray(fds []int) UnixFDList {
	fds0 := make([]C.gint, len(fds))
	for idx, elemG := range fds {
		fds0[idx] = C.gint(elemG)
	}
	var fds0Ptr *C.gint
	if len(fds0) > 0 {
		fds0Ptr = &fds0[0]
	}
	ret0 := C.g_unix_fd_list_new_from_array(fds0Ptr, C.gint(len(fds)))
	return wrapUnixFDList(ret0)
}

// Append is a wrapper around g_unix_fd_list_append().
func (list UnixFDList) Append(fd int) (int, error) {
	var err glib.Error
	ret0 := C.g_unix_fd_list_append(list.native(), C.gint(fd), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// Get is a wrapper around g_unix_fd_list_get().
func (list UnixFDList) Get(index_ int) (int, error) {
	var err glib.Error
	ret0 := C.g_unix_fd_list_get(list.native(), C.gint(index_), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// GetLength is a wrapper around g_unix_fd_list_get_length().
func (list UnixFDList) GetLength() int {
	ret0 := C.g_unix_fd_list_get_length(list.native())
	return int(ret0)
}

// PeekFds is a wrapper around g_unix_fd_list_peek_fds().
func (list UnixFDList) PeekFds() []int {
	var length0 C.gint
	ret0 := C.g_unix_fd_list_peek_fds(list.native(), &length0)
	var ret0Slice []C.gint
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(length0)) /*go:.util*/
	ret := make([]int, len(ret0Slice))
	for idx, elem := range ret0Slice {
		ret[idx] = int(elem)
	}
	return ret
}

// StealFds is a wrapper around g_unix_fd_list_steal_fds().
func (list UnixFDList) StealFds() []int {
	var length0 C.gint
	ret0 := C.g_unix_fd_list_steal_fds(list.native(), &length0)
	var ret0Slice []C.gint
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(length0)) /*go:.util*/
	ret := make([]int, len(ret0Slice))
	for idx, elem := range ret0Slice {
		ret[idx] = int(elem)
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// Object DBusMethodInvocation
type DBusMethodInvocation struct {
	gobject.Object
}

func (v DBusMethodInvocation) native() *C.GDBusMethodInvocation {
	return (*C.GDBusMethodInvocation)(v.Ptr)
}
func wrapDBusMethodInvocation(p *C.GDBusMethodInvocation) (v DBusMethodInvocation) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDBusMethodInvocation(p unsafe.Pointer) (v DBusMethodInvocation) {
	v.Ptr = p
	return
}
func (v DBusMethodInvocation) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusMethodInvocation(p unsafe.Pointer) interface{} {
	return WrapDBusMethodInvocation(p)
}
func (v DBusMethodInvocation) GetType() gobject.Type {
	return gobject.Type(C.g_dbus_method_invocation_get_type())
}
func (v DBusMethodInvocation) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDBusMethodInvocation(unsafe.Pointer(ptr)), nil
	}
}

// GetConnection is a wrapper around g_dbus_method_invocation_get_connection().
func (invocation DBusMethodInvocation) GetConnection() DBusConnection {
	ret0 := C.g_dbus_method_invocation_get_connection(invocation.native())
	return wrapDBusConnection(ret0)
}

// GetInterfaceName is a wrapper around g_dbus_method_invocation_get_interface_name().
func (invocation DBusMethodInvocation) GetInterfaceName() string {
	ret0 := C.g_dbus_method_invocation_get_interface_name(invocation.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetMessage is a wrapper around g_dbus_method_invocation_get_message().
func (invocation DBusMethodInvocation) GetMessage() DBusMessage {
	ret0 := C.g_dbus_method_invocation_get_message(invocation.native())
	return wrapDBusMessage(ret0)
}

// GetMethodInfo is a wrapper around g_dbus_method_invocation_get_method_info().
func (invocation DBusMethodInvocation) GetMethodInfo() DBusMethodInfo {
	ret0 := C.g_dbus_method_invocation_get_method_info(invocation.native())
	return wrapDBusMethodInfo(ret0)
}

// GetMethodName is a wrapper around g_dbus_method_invocation_get_method_name().
func (invocation DBusMethodInvocation) GetMethodName() string {
	ret0 := C.g_dbus_method_invocation_get_method_name(invocation.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetObjectPath is a wrapper around g_dbus_method_invocation_get_object_path().
func (invocation DBusMethodInvocation) GetObjectPath() string {
	ret0 := C.g_dbus_method_invocation_get_object_path(invocation.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetParameters is a wrapper around g_dbus_method_invocation_get_parameters().
func (invocation DBusMethodInvocation) GetParameters() glib.Variant {
	ret0 := C.g_dbus_method_invocation_get_parameters(invocation.native())
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetPropertyInfo is a wrapper around g_dbus_method_invocation_get_property_info().
func (invocation DBusMethodInvocation) GetPropertyInfo() DBusPropertyInfo {
	ret0 := C.g_dbus_method_invocation_get_property_info(invocation.native())
	return wrapDBusPropertyInfo(ret0)
}

// GetSender is a wrapper around g_dbus_method_invocation_get_sender().
func (invocation DBusMethodInvocation) GetSender() string {
	ret0 := C.g_dbus_method_invocation_get_sender(invocation.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetUserData is a wrapper around g_dbus_method_invocation_get_user_data().
func (invocation DBusMethodInvocation) GetUserData() unsafe.Pointer {
	ret0 := C.g_dbus_method_invocation_get_user_data(invocation.native())
	return unsafe.Pointer(ret0)
}

// ReturnDbusError is a wrapper around g_dbus_method_invocation_return_dbus_error().
func (invocation DBusMethodInvocation) ReturnDbusError(error_name string, error_message string) {
	error_name0 := (*C.gchar)(C.CString(error_name))
	error_message0 := (*C.gchar)(C.CString(error_message))
	C.g_dbus_method_invocation_return_dbus_error(invocation.native(), error_name0, error_message0)
	C.free(unsafe.Pointer(error_name0))    /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(error_message0)) /*ch:<stdlib.h>*/
}

// ReturnErrorLiteral is a wrapper around g_dbus_method_invocation_return_error_literal().
func (invocation DBusMethodInvocation) ReturnErrorLiteral(domain /*gir:GLib*/ glib.Quark, code int, message string) {
	message0 := (*C.gchar)(C.CString(message))
	C.g_dbus_method_invocation_return_error_literal(invocation.native(), C.GQuark(domain), C.gint(code), message0)
	C.free(unsafe.Pointer(message0)) /*ch:<stdlib.h>*/
}

// ReturnGerror is a wrapper around g_dbus_method_invocation_return_gerror().
func (invocation DBusMethodInvocation) ReturnGerror(error glib.Error) {
	C.g_dbus_method_invocation_return_gerror(invocation.native(), (*C.GError)(error.Ptr))
}

// ReturnValue is a wrapper around g_dbus_method_invocation_return_value().
func (invocation DBusMethodInvocation) ReturnValue(parameters glib.Variant) {
	C.g_dbus_method_invocation_return_value(invocation.native(), (*C.GVariant)(parameters.Ptr))
}

// ReturnValueWithUnixFdList is a wrapper around g_dbus_method_invocation_return_value_with_unix_fd_list().
func (invocation DBusMethodInvocation) ReturnValueWithUnixFdList(parameters glib.Variant, fd_list UnixFDList) {
	C.g_dbus_method_invocation_return_value_with_unix_fd_list(invocation.native(), (*C.GVariant)(parameters.Ptr), fd_list.native())
}

// TakeError is a wrapper around g_dbus_method_invocation_take_error().
func (invocation DBusMethodInvocation) TakeError(error glib.Error) {
	C.g_dbus_method_invocation_take_error(invocation.native(), (*C.GError)(error.Ptr))
}

// Object DBusObjectManagerClient
type DBusObjectManagerClient struct {
	gobject.Object
}

func (v DBusObjectManagerClient) native() *C.GDBusObjectManagerClient {
	return (*C.GDBusObjectManagerClient)(v.Ptr)
}
func wrapDBusObjectManagerClient(p *C.GDBusObjectManagerClient) (v DBusObjectManagerClient) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDBusObjectManagerClient(p unsafe.Pointer) (v DBusObjectManagerClient) {
	v.Ptr = p
	return
}
func (v DBusObjectManagerClient) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusObjectManagerClient(p unsafe.Pointer) interface{} {
	return WrapDBusObjectManagerClient(p)
}
func (v DBusObjectManagerClient) GetType() gobject.Type {
	return gobject.Type(C.g_dbus_object_manager_client_get_type())
}
func (v DBusObjectManagerClient) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDBusObjectManagerClient(unsafe.Pointer(ptr)), nil
	}
}
func (v DBusObjectManagerClient) AsyncInitable() AsyncInitable {
	return WrapAsyncInitable(v.Ptr)
}
func (v DBusObjectManagerClient) DBusObjectManager() DBusObjectManager {
	return WrapDBusObjectManager(v.Ptr)
}
func (v DBusObjectManagerClient) Initable() Initable {
	return WrapInitable(v.Ptr)
}

// GetConnection is a wrapper around g_dbus_object_manager_client_get_connection().
func (manager DBusObjectManagerClient) GetConnection() DBusConnection {
	ret0 := C.g_dbus_object_manager_client_get_connection(manager.native())
	return wrapDBusConnection(ret0)
}

// GetFlags is a wrapper around g_dbus_object_manager_client_get_flags().
func (manager DBusObjectManagerClient) GetFlags() DBusObjectManagerClientFlags {
	ret0 := C.g_dbus_object_manager_client_get_flags(manager.native())
	return DBusObjectManagerClientFlags(ret0)
}

// GetName is a wrapper around g_dbus_object_manager_client_get_name().
func (manager DBusObjectManagerClient) GetName() string {
	ret0 := C.g_dbus_object_manager_client_get_name(manager.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetNameOwner is a wrapper around g_dbus_object_manager_client_get_name_owner().
func (manager DBusObjectManagerClient) GetNameOwner() string {
	ret0 := C.g_dbus_object_manager_client_get_name_owner(manager.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// Object DBusObjectManagerServer
type DBusObjectManagerServer struct {
	gobject.Object
}

func (v DBusObjectManagerServer) native() *C.GDBusObjectManagerServer {
	return (*C.GDBusObjectManagerServer)(v.Ptr)
}
func wrapDBusObjectManagerServer(p *C.GDBusObjectManagerServer) (v DBusObjectManagerServer) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDBusObjectManagerServer(p unsafe.Pointer) (v DBusObjectManagerServer) {
	v.Ptr = p
	return
}
func (v DBusObjectManagerServer) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusObjectManagerServer(p unsafe.Pointer) interface{} {
	return WrapDBusObjectManagerServer(p)
}
func (v DBusObjectManagerServer) GetType() gobject.Type {
	return gobject.Type(C.g_dbus_object_manager_server_get_type())
}
func (v DBusObjectManagerServer) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDBusObjectManagerServer(unsafe.Pointer(ptr)), nil
	}
}
func (v DBusObjectManagerServer) DBusObjectManager() DBusObjectManager {
	return WrapDBusObjectManager(v.Ptr)
}

// DBusObjectManagerServerNew is a wrapper around g_dbus_object_manager_server_new().
func DBusObjectManagerServerNew(object_path string) DBusObjectManagerServer {
	object_path0 := (*C.gchar)(C.CString(object_path))
	ret0 := C.g_dbus_object_manager_server_new(object_path0)
	C.free(unsafe.Pointer(object_path0)) /*ch:<stdlib.h>*/
	return wrapDBusObjectManagerServer(ret0)
}

// Export is a wrapper around g_dbus_object_manager_server_export().
func (manager DBusObjectManagerServer) Export(object DBusObjectSkeleton) {
	C.g_dbus_object_manager_server_export(manager.native(), object.native())
}

// ExportUniquely is a wrapper around g_dbus_object_manager_server_export_uniquely().
func (manager DBusObjectManagerServer) ExportUniquely(object DBusObjectSkeleton) {
	C.g_dbus_object_manager_server_export_uniquely(manager.native(), object.native())
}

// GetConnection is a wrapper around g_dbus_object_manager_server_get_connection().
func (manager DBusObjectManagerServer) GetConnection() DBusConnection {
	ret0 := C.g_dbus_object_manager_server_get_connection(manager.native())
	return wrapDBusConnection(ret0)
}

// IsExported is a wrapper around g_dbus_object_manager_server_is_exported().
func (manager DBusObjectManagerServer) IsExported(object DBusObjectSkeleton) bool {
	ret0 := C.g_dbus_object_manager_server_is_exported(manager.native(), object.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetConnection is a wrapper around g_dbus_object_manager_server_set_connection().
func (manager DBusObjectManagerServer) SetConnection(connection DBusConnection) {
	C.g_dbus_object_manager_server_set_connection(manager.native(), connection.native())
}

// Unexport is a wrapper around g_dbus_object_manager_server_unexport().
func (manager DBusObjectManagerServer) Unexport(object_path string) bool {
	object_path0 := (*C.gchar)(C.CString(object_path))
	ret0 := C.g_dbus_object_manager_server_unexport(manager.native(), object_path0)
	C.free(unsafe.Pointer(object_path0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0))      /*go:.util*/
}

// Object DBusObjectSkeleton
type DBusObjectSkeleton struct {
	gobject.Object
}

func (v DBusObjectSkeleton) native() *C.GDBusObjectSkeleton {
	return (*C.GDBusObjectSkeleton)(v.Ptr)
}
func wrapDBusObjectSkeleton(p *C.GDBusObjectSkeleton) (v DBusObjectSkeleton) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDBusObjectSkeleton(p unsafe.Pointer) (v DBusObjectSkeleton) {
	v.Ptr = p
	return
}
func (v DBusObjectSkeleton) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusObjectSkeleton(p unsafe.Pointer) interface{} {
	return WrapDBusObjectSkeleton(p)
}
func (v DBusObjectSkeleton) GetType() gobject.Type {
	return gobject.Type(C.g_dbus_object_skeleton_get_type())
}
func (v DBusObjectSkeleton) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDBusObjectSkeleton(unsafe.Pointer(ptr)), nil
	}
}
func (v DBusObjectSkeleton) DBusObject() DBusObject {
	return WrapDBusObject(v.Ptr)
}

// DBusObjectSkeletonNew is a wrapper around g_dbus_object_skeleton_new().
func DBusObjectSkeletonNew(object_path string) DBusObjectSkeleton {
	object_path0 := (*C.gchar)(C.CString(object_path))
	ret0 := C.g_dbus_object_skeleton_new(object_path0)
	C.free(unsafe.Pointer(object_path0)) /*ch:<stdlib.h>*/
	return wrapDBusObjectSkeleton(ret0)
}

// AddInterface is a wrapper around g_dbus_object_skeleton_add_interface().
func (object DBusObjectSkeleton) AddInterface(interface_ DBusInterfaceSkeleton) {
	C.g_dbus_object_skeleton_add_interface(object.native(), interface_.native())
}

// Flush is a wrapper around g_dbus_object_skeleton_flush().
func (object DBusObjectSkeleton) Flush() {
	C.g_dbus_object_skeleton_flush(object.native())
}

// RemoveInterface is a wrapper around g_dbus_object_skeleton_remove_interface().
func (object DBusObjectSkeleton) RemoveInterface(interface_ DBusInterfaceSkeleton) {
	C.g_dbus_object_skeleton_remove_interface(object.native(), interface_.native())
}

// RemoveInterfaceByName is a wrapper around g_dbus_object_skeleton_remove_interface_by_name().
func (object DBusObjectSkeleton) RemoveInterfaceByName(interface_name string) {
	interface_name0 := (*C.gchar)(C.CString(interface_name))
	C.g_dbus_object_skeleton_remove_interface_by_name(object.native(), interface_name0)
	C.free(unsafe.Pointer(interface_name0)) /*ch:<stdlib.h>*/
}

// SetObjectPath is a wrapper around g_dbus_object_skeleton_set_object_path().
func (object DBusObjectSkeleton) SetObjectPath(object_path string) {
	object_path0 := (*C.gchar)(C.CString(object_path))
	C.g_dbus_object_skeleton_set_object_path(object.native(), object_path0)
	C.free(unsafe.Pointer(object_path0)) /*ch:<stdlib.h>*/
}

// Object DBusObjectProxy
type DBusObjectProxy struct {
	gobject.Object
}

func (v DBusObjectProxy) native() *C.GDBusObjectProxy {
	return (*C.GDBusObjectProxy)(v.Ptr)
}
func wrapDBusObjectProxy(p *C.GDBusObjectProxy) (v DBusObjectProxy) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDBusObjectProxy(p unsafe.Pointer) (v DBusObjectProxy) {
	v.Ptr = p
	return
}
func (v DBusObjectProxy) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusObjectProxy(p unsafe.Pointer) interface{} {
	return WrapDBusObjectProxy(p)
}
func (v DBusObjectProxy) GetType() gobject.Type {
	return gobject.Type(C.g_dbus_object_proxy_get_type())
}
func (v DBusObjectProxy) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDBusObjectProxy(unsafe.Pointer(ptr)), nil
	}
}
func (v DBusObjectProxy) DBusObject() DBusObject {
	return WrapDBusObject(v.Ptr)
}

// DBusObjectProxyNew is a wrapper around g_dbus_object_proxy_new().
func DBusObjectProxyNew(connection DBusConnection, object_path string) DBusObjectProxy {
	object_path0 := (*C.gchar)(C.CString(object_path))
	ret0 := C.g_dbus_object_proxy_new(connection.native(), object_path0)
	C.free(unsafe.Pointer(object_path0)) /*ch:<stdlib.h>*/
	return wrapDBusObjectProxy(ret0)
}

// GetConnection is a wrapper around g_dbus_object_proxy_get_connection().
func (proxy DBusObjectProxy) GetConnection() DBusConnection {
	ret0 := C.g_dbus_object_proxy_get_connection(proxy.native())
	return wrapDBusConnection(ret0)
}

// Object DBusProxy
type DBusProxy struct {
	gobject.Object
}

func (v DBusProxy) native() *C.GDBusProxy {
	return (*C.GDBusProxy)(v.Ptr)
}
func wrapDBusProxy(p *C.GDBusProxy) (v DBusProxy) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDBusProxy(p unsafe.Pointer) (v DBusProxy) {
	v.Ptr = p
	return
}
func (v DBusProxy) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusProxy(p unsafe.Pointer) interface{} {
	return WrapDBusProxy(p)
}
func (v DBusProxy) GetType() gobject.Type {
	return gobject.Type(C.g_dbus_proxy_get_type())
}
func (v DBusProxy) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDBusProxy(unsafe.Pointer(ptr)), nil
	}
}
func (v DBusProxy) AsyncInitable() AsyncInitable {
	return WrapAsyncInitable(v.Ptr)
}
func (v DBusProxy) DBusInterface() DBusInterface {
	return WrapDBusInterface(v.Ptr)
}
func (v DBusProxy) Initable() Initable {
	return WrapInitable(v.Ptr)
}

// DBusProxyNewFinish is a wrapper around g_dbus_proxy_new_finish().
func DBusProxyNewFinish(res AsyncResult) (DBusProxy, error) {
	var err glib.Error
	ret0 := C.g_dbus_proxy_new_finish(res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return DBusProxy{}, err.GoValue()
	}
	return wrapDBusProxy(ret0), nil
}

// DBusProxyNewForBusFinish is a wrapper around g_dbus_proxy_new_for_bus_finish().
func DBusProxyNewForBusFinish(res AsyncResult) (DBusProxy, error) {
	var err glib.Error
	ret0 := C.g_dbus_proxy_new_for_bus_finish(res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return DBusProxy{}, err.GoValue()
	}
	return wrapDBusProxy(ret0), nil
}

// DBusProxyNewForBusSync is a wrapper around g_dbus_proxy_new_for_bus_sync().
func DBusProxyNewForBusSync(bus_type BusType, flags DBusProxyFlags, info DBusInterfaceInfo, name string, object_path string, interface_name string, cancellable Cancellable) (DBusProxy, error) {
	name0 := (*C.gchar)(C.CString(name))
	object_path0 := (*C.gchar)(C.CString(object_path))
	interface_name0 := (*C.gchar)(C.CString(interface_name))
	var err glib.Error
	ret0 := C.g_dbus_proxy_new_for_bus_sync(C.GBusType(bus_type), C.GDBusProxyFlags(flags), info.native(), name0, object_path0, interface_name0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(name0))           /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(object_path0))    /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(interface_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return DBusProxy{}, err.GoValue()
	}
	return wrapDBusProxy(ret0), nil
}

// DBusProxyNewSync is a wrapper around g_dbus_proxy_new_sync().
func DBusProxyNewSync(connection DBusConnection, flags DBusProxyFlags, info DBusInterfaceInfo, name string, object_path string, interface_name string, cancellable Cancellable) (DBusProxy, error) {
	name0 := (*C.gchar)(C.CString(name))
	object_path0 := (*C.gchar)(C.CString(object_path))
	interface_name0 := (*C.gchar)(C.CString(interface_name))
	var err glib.Error
	ret0 := C.g_dbus_proxy_new_sync(connection.native(), C.GDBusProxyFlags(flags), info.native(), name0, object_path0, interface_name0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(name0))           /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(object_path0))    /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(interface_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return DBusProxy{}, err.GoValue()
	}
	return wrapDBusProxy(ret0), nil
}

// Call is a wrapper around g_dbus_proxy_call().
func (proxy DBusProxy) Call(method_name string, parameters glib.Variant, flags DBusCallFlags, timeout_msec int, cancellable Cancellable, callback AsyncReadyCallback) {
	method_name0 := (*C.gchar)(C.CString(method_name))
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_dbus_proxy_call(proxy.native(), method_name0, (*C.GVariant)(parameters.Ptr), C.GDBusCallFlags(flags), C.gint(timeout_msec), cancellable.native(), callback0)
	C.free(unsafe.Pointer(method_name0)) /*ch:<stdlib.h>*/
}

// CallFinish is a wrapper around g_dbus_proxy_call_finish().
func (proxy DBusProxy) CallFinish(res AsyncResult) (glib.Variant, error) {
	var err glib.Error
	ret0 := C.g_dbus_proxy_call_finish(proxy.native(), res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return glib.Variant{}, err.GoValue()
	}
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/, nil
}

// CallSync is a wrapper around g_dbus_proxy_call_sync().
func (proxy DBusProxy) CallSync(method_name string, parameters glib.Variant, flags DBusCallFlags, timeout_msec int, cancellable Cancellable) (glib.Variant, error) {
	method_name0 := (*C.gchar)(C.CString(method_name))
	var err glib.Error
	ret0 := C.g_dbus_proxy_call_sync(proxy.native(), method_name0, (*C.GVariant)(parameters.Ptr), C.GDBusCallFlags(flags), C.gint(timeout_msec), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(method_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return glib.Variant{}, err.GoValue()
	}
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/, nil
}

// CallWithUnixFdList is a wrapper around g_dbus_proxy_call_with_unix_fd_list().
func (proxy DBusProxy) CallWithUnixFdList(method_name string, parameters glib.Variant, flags DBusCallFlags, timeout_msec int, fd_list UnixFDList, cancellable Cancellable, callback AsyncReadyCallback) {
	method_name0 := (*C.gchar)(C.CString(method_name))
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_dbus_proxy_call_with_unix_fd_list(proxy.native(), method_name0, (*C.GVariant)(parameters.Ptr), C.GDBusCallFlags(flags), C.gint(timeout_msec), fd_list.native(), cancellable.native(), callback0)
	C.free(unsafe.Pointer(method_name0)) /*ch:<stdlib.h>*/
}

// CallWithUnixFdListFinish is a wrapper around g_dbus_proxy_call_with_unix_fd_list_finish().
func (proxy DBusProxy) CallWithUnixFdListFinish(res AsyncResult) (glib.Variant, UnixFDList, error) {
	var out_fd_list0 *C.GUnixFDList
	var err glib.Error
	ret0 := C.g_dbus_proxy_call_with_unix_fd_list_finish(proxy.native(), &out_fd_list0, res.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return glib.Variant{}, UnixFDList{}, err.GoValue()
	}
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/, wrapUnixFDList(out_fd_list0), nil
}

// CallWithUnixFdListSync is a wrapper around g_dbus_proxy_call_with_unix_fd_list_sync().
func (proxy DBusProxy) CallWithUnixFdListSync(method_name string, parameters glib.Variant, flags DBusCallFlags, timeout_msec int, fd_list UnixFDList, cancellable Cancellable) (glib.Variant, UnixFDList, error) {
	method_name0 := (*C.gchar)(C.CString(method_name))
	var out_fd_list0 *C.GUnixFDList
	var err glib.Error
	ret0 := C.g_dbus_proxy_call_with_unix_fd_list_sync(proxy.native(), method_name0, (*C.GVariant)(parameters.Ptr), C.GDBusCallFlags(flags), C.gint(timeout_msec), fd_list.native(), &out_fd_list0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(method_name0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return glib.Variant{}, UnixFDList{}, err.GoValue()
	}
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/, wrapUnixFDList(out_fd_list0), nil
}

// GetCachedProperty is a wrapper around g_dbus_proxy_get_cached_property().
func (proxy DBusProxy) GetCachedProperty(property_name string) glib.Variant {
	property_name0 := (*C.gchar)(C.CString(property_name))
	ret0 := C.g_dbus_proxy_get_cached_property(proxy.native(), property_name0)
	C.free(unsafe.Pointer(property_name0))        /*ch:<stdlib.h>*/
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetCachedPropertyNames is a wrapper around g_dbus_proxy_get_cached_property_names().
func (proxy DBusProxy) GetCachedPropertyNames() []string {
	ret0 := C.g_dbus_proxy_get_cached_property_names(proxy.native())
	var ret0Slice []*C.gchar
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString((*C.char)(elem))
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetConnection is a wrapper around g_dbus_proxy_get_connection().
func (proxy DBusProxy) GetConnection() DBusConnection {
	ret0 := C.g_dbus_proxy_get_connection(proxy.native())
	return wrapDBusConnection(ret0)
}

// GetDefaultTimeout is a wrapper around g_dbus_proxy_get_default_timeout().
func (proxy DBusProxy) GetDefaultTimeout() int {
	ret0 := C.g_dbus_proxy_get_default_timeout(proxy.native())
	return int(ret0)
}

// GetFlags is a wrapper around g_dbus_proxy_get_flags().
func (proxy DBusProxy) GetFlags() DBusProxyFlags {
	ret0 := C.g_dbus_proxy_get_flags(proxy.native())
	return DBusProxyFlags(ret0)
}

// GetInterfaceInfo is a wrapper around g_dbus_proxy_get_interface_info().
func (proxy DBusProxy) GetInterfaceInfo() DBusInterfaceInfo {
	ret0 := C.g_dbus_proxy_get_interface_info(proxy.native())
	return wrapDBusInterfaceInfo(ret0)
}

// GetInterfaceName is a wrapper around g_dbus_proxy_get_interface_name().
func (proxy DBusProxy) GetInterfaceName() string {
	ret0 := C.g_dbus_proxy_get_interface_name(proxy.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetName is a wrapper around g_dbus_proxy_get_name().
func (proxy DBusProxy) GetName() string {
	ret0 := C.g_dbus_proxy_get_name(proxy.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetNameOwner is a wrapper around g_dbus_proxy_get_name_owner().
func (proxy DBusProxy) GetNameOwner() string {
	ret0 := C.g_dbus_proxy_get_name_owner(proxy.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetObjectPath is a wrapper around g_dbus_proxy_get_object_path().
func (proxy DBusProxy) GetObjectPath() string {
	ret0 := C.g_dbus_proxy_get_object_path(proxy.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// SetCachedProperty is a wrapper around g_dbus_proxy_set_cached_property().
func (proxy DBusProxy) SetCachedProperty(property_name string, value glib.Variant) {
	property_name0 := (*C.gchar)(C.CString(property_name))
	C.g_dbus_proxy_set_cached_property(proxy.native(), property_name0, (*C.GVariant)(value.Ptr))
	C.free(unsafe.Pointer(property_name0)) /*ch:<stdlib.h>*/
}

// SetDefaultTimeout is a wrapper around g_dbus_proxy_set_default_timeout().
func (proxy DBusProxy) SetDefaultTimeout(timeout_msec int) {
	C.g_dbus_proxy_set_default_timeout(proxy.native(), C.gint(timeout_msec))
}

// SetInterfaceInfo is a wrapper around g_dbus_proxy_set_interface_info().
func (proxy DBusProxy) SetInterfaceInfo(info DBusInterfaceInfo) {
	C.g_dbus_proxy_set_interface_info(proxy.native(), info.native())
}

// DBusProxyNew is a wrapper around g_dbus_proxy_new().
func DBusProxyNew(connection DBusConnection, flags DBusProxyFlags, info DBusInterfaceInfo, name string, object_path string, interface_name string, cancellable Cancellable, callback AsyncReadyCallback) {
	name0 := (*C.gchar)(C.CString(name))
	object_path0 := (*C.gchar)(C.CString(object_path))
	interface_name0 := (*C.gchar)(C.CString(interface_name))
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_dbus_proxy_new(connection.native(), C.GDBusProxyFlags(flags), info.native(), name0, object_path0, interface_name0, cancellable.native(), callback0)
	C.free(unsafe.Pointer(name0))           /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(object_path0))    /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(interface_name0)) /*ch:<stdlib.h>*/
}

// DBusProxyNewForBus is a wrapper around g_dbus_proxy_new_for_bus().
func DBusProxyNewForBus(bus_type BusType, flags DBusProxyFlags, info DBusInterfaceInfo, name string, object_path string, interface_name string, cancellable Cancellable, callback AsyncReadyCallback) {
	name0 := (*C.gchar)(C.CString(name))
	object_path0 := (*C.gchar)(C.CString(object_path))
	interface_name0 := (*C.gchar)(C.CString(interface_name))
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_dbus_proxy_new_for_bus(C.GBusType(bus_type), C.GDBusProxyFlags(flags), info.native(), name0, object_path0, interface_name0, cancellable.native(), callback0)
	C.free(unsafe.Pointer(name0))           /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(object_path0))    /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(interface_name0)) /*ch:<stdlib.h>*/
}

// Object DBusServer
type DBusServer struct {
	gobject.Object
}

func (v DBusServer) native() *C.GDBusServer {
	return (*C.GDBusServer)(v.Ptr)
}
func wrapDBusServer(p *C.GDBusServer) (v DBusServer) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDBusServer(p unsafe.Pointer) (v DBusServer) {
	v.Ptr = p
	return
}
func (v DBusServer) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDBusServer(p unsafe.Pointer) interface{} {
	return WrapDBusServer(p)
}
func (v DBusServer) GetType() gobject.Type {
	return gobject.Type(C.g_dbus_server_get_type())
}
func (v DBusServer) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDBusServer(unsafe.Pointer(ptr)), nil
	}
}
func (v DBusServer) Initable() Initable {
	return WrapInitable(v.Ptr)
}

// DBusServerNewSync is a wrapper around g_dbus_server_new_sync().
func DBusServerNewSync(address string, flags DBusServerFlags, guid string, observer DBusAuthObserver, cancellable Cancellable) (DBusServer, error) {
	address0 := (*C.gchar)(C.CString(address))
	guid0 := (*C.gchar)(C.CString(guid))
	var err glib.Error
	ret0 := C.g_dbus_server_new_sync(address0, C.GDBusServerFlags(flags), guid0, observer.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(address0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(guid0))    /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return DBusServer{}, err.GoValue()
	}
	return wrapDBusServer(ret0), nil
}

// GetClientAddress is a wrapper around g_dbus_server_get_client_address().
func (server DBusServer) GetClientAddress() string {
	ret0 := C.g_dbus_server_get_client_address(server.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetFlags is a wrapper around g_dbus_server_get_flags().
func (server DBusServer) GetFlags() DBusServerFlags {
	ret0 := C.g_dbus_server_get_flags(server.native())
	return DBusServerFlags(ret0)
}

// GetGuid is a wrapper around g_dbus_server_get_guid().
func (server DBusServer) GetGuid() string {
	ret0 := C.g_dbus_server_get_guid(server.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// IsActive is a wrapper around g_dbus_server_is_active().
func (server DBusServer) IsActive() bool {
	ret0 := C.g_dbus_server_is_active(server.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Start is a wrapper around g_dbus_server_start().
func (server DBusServer) Start() {
	C.g_dbus_server_start(server.native())
}

// Stop is a wrapper around g_dbus_server_stop().
func (server DBusServer) Stop() {
	C.g_dbus_server_stop(server.native())
}

// Object DataInputStream
type DataInputStream struct {
	BufferedInputStream
}

func (v DataInputStream) native() *C.GDataInputStream {
	return (*C.GDataInputStream)(v.Ptr)
}
func wrapDataInputStream(p *C.GDataInputStream) (v DataInputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDataInputStream(p unsafe.Pointer) (v DataInputStream) {
	v.Ptr = p
	return
}
func (v DataInputStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDataInputStream(p unsafe.Pointer) interface{} {
	return WrapDataInputStream(p)
}
func (v DataInputStream) GetType() gobject.Type {
	return gobject.Type(C.g_data_input_stream_get_type())
}
func (v DataInputStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDataInputStream(unsafe.Pointer(ptr)), nil
	}
}
func (v DataInputStream) Seekable() Seekable {
	return WrapSeekable(v.Ptr)
}

// DataInputStreamNew is a wrapper around g_data_input_stream_new().
func DataInputStreamNew(base_stream InputStream) DataInputStream {
	ret0 := C.g_data_input_stream_new(base_stream.native())
	return wrapDataInputStream(ret0)
}

// GetByteOrder is a wrapper around g_data_input_stream_get_byte_order().
func (stream DataInputStream) GetByteOrder() DataStreamByteOrder {
	ret0 := C.g_data_input_stream_get_byte_order(stream.native())
	return DataStreamByteOrder(ret0)
}

// GetNewlineType is a wrapper around g_data_input_stream_get_newline_type().
func (stream DataInputStream) GetNewlineType() DataStreamNewlineType {
	ret0 := C.g_data_input_stream_get_newline_type(stream.native())
	return DataStreamNewlineType(ret0)
}

// ReadByte is a wrapper around g_data_input_stream_read_byte().
func (stream DataInputStream) ReadByte(cancellable Cancellable) (byte, error) {
	var err glib.Error
	ret0 := C.g_data_input_stream_read_byte(stream.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return byte(ret0), nil
}

// ReadInt16 is a wrapper around g_data_input_stream_read_int16().
func (stream DataInputStream) ReadInt16(cancellable Cancellable) (int16, error) {
	var err glib.Error
	ret0 := C.g_data_input_stream_read_int16(stream.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int16(ret0), nil
}

// ReadInt32 is a wrapper around g_data_input_stream_read_int32().
func (stream DataInputStream) ReadInt32(cancellable Cancellable) (int32, error) {
	var err glib.Error
	ret0 := C.g_data_input_stream_read_int32(stream.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int32(ret0), nil
}

// ReadInt64 is a wrapper around g_data_input_stream_read_int64().
func (stream DataInputStream) ReadInt64(cancellable Cancellable) (int64, error) {
	var err glib.Error
	ret0 := C.g_data_input_stream_read_int64(stream.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int64(ret0), nil
}

// ReadLineAsync is a wrapper around g_data_input_stream_read_line_async().
func (stream DataInputStream) ReadLineAsync(io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_data_input_stream_read_line_async(stream.native(), C.gint(io_priority), cancellable.native(), callback0)
}

// ReadLineFinishUtf8 is a wrapper around g_data_input_stream_read_line_finish_utf8().
func (stream DataInputStream) ReadLineFinishUtf8(result AsyncResult) (string, uint, error) {
	var length0 C.gsize
	var err glib.Error
	ret0 := C.g_data_input_stream_read_line_finish_utf8(stream.native(), result.native(), &length0, (**C.GError)(unsafe.Pointer(&err)))
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return "", 0, err.GoValue()
	}
	return ret, uint(length0), nil
}

// ReadLineUtf8 is a wrapper around g_data_input_stream_read_line_utf8().
func (stream DataInputStream) ReadLineUtf8(cancellable Cancellable) (string, uint, error) {
	var length0 C.gsize
	var err glib.Error
	ret0 := C.g_data_input_stream_read_line_utf8(stream.native(), &length0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return "", 0, err.GoValue()
	}
	return ret, uint(length0), nil
}

// ReadUint16 is a wrapper around g_data_input_stream_read_uint16().
func (stream DataInputStream) ReadUint16(cancellable Cancellable) (uint16, error) {
	var err glib.Error
	ret0 := C.g_data_input_stream_read_uint16(stream.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return uint16(ret0), nil
}

// ReadUint32 is a wrapper around g_data_input_stream_read_uint32().
func (stream DataInputStream) ReadUint32(cancellable Cancellable) (uint32, error) {
	var err glib.Error
	ret0 := C.g_data_input_stream_read_uint32(stream.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return uint32(ret0), nil
}

// ReadUint64 is a wrapper around g_data_input_stream_read_uint64().
func (stream DataInputStream) ReadUint64(cancellable Cancellable) (uint64, error) {
	var err glib.Error
	ret0 := C.g_data_input_stream_read_uint64(stream.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return uint64(ret0), nil
}

// ReadUntil is a wrapper around g_data_input_stream_read_until().
func (stream DataInputStream) ReadUntil(stop_chars string, cancellable Cancellable) (string, uint, error) {
	stop_chars0 := (*C.gchar)(C.CString(stop_chars))
	var length0 C.gsize
	var err glib.Error
	ret0 := C.g_data_input_stream_read_until(stream.native(), stop_chars0, &length0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(stop_chars0)) /*ch:<stdlib.h>*/
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return "", 0, err.GoValue()
	}
	return ret, uint(length0), nil
}

// ReadUntilAsync is a wrapper around g_data_input_stream_read_until_async().
func (stream DataInputStream) ReadUntilAsync(stop_chars string, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	stop_chars0 := (*C.gchar)(C.CString(stop_chars))
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_data_input_stream_read_until_async(stream.native(), stop_chars0, C.gint(io_priority), cancellable.native(), callback0)
	C.free(unsafe.Pointer(stop_chars0)) /*ch:<stdlib.h>*/
}

// ReadUntilFinish is a wrapper around g_data_input_stream_read_until_finish().
func (stream DataInputStream) ReadUntilFinish(result AsyncResult) (string, uint, error) {
	var length0 C.gsize
	var err glib.Error
	ret0 := C.g_data_input_stream_read_until_finish(stream.native(), result.native(), &length0, (**C.GError)(unsafe.Pointer(&err)))
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return "", 0, err.GoValue()
	}
	return ret, uint(length0), nil
}

// ReadUpto is a wrapper around g_data_input_stream_read_upto().
func (stream DataInputStream) ReadUpto(stop_chars string, stop_chars_len int, cancellable Cancellable) (string, uint, error) {
	stop_chars0 := (*C.gchar)(C.CString(stop_chars))
	var length0 C.gsize
	var err glib.Error
	ret0 := C.g_data_input_stream_read_upto(stream.native(), stop_chars0, C.gssize(stop_chars_len), &length0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(stop_chars0)) /*ch:<stdlib.h>*/
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return "", 0, err.GoValue()
	}
	return ret, uint(length0), nil
}

// ReadUptoAsync is a wrapper around g_data_input_stream_read_upto_async().
func (stream DataInputStream) ReadUptoAsync(stop_chars string, stop_chars_len int, io_priority int, cancellable Cancellable, callback AsyncReadyCallback) {
	stop_chars0 := (*C.gchar)(C.CString(stop_chars))
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_data_input_stream_read_upto_async(stream.native(), stop_chars0, C.gssize(stop_chars_len), C.gint(io_priority), cancellable.native(), callback0)
	C.free(unsafe.Pointer(stop_chars0)) /*ch:<stdlib.h>*/
}

// ReadUptoFinish is a wrapper around g_data_input_stream_read_upto_finish().
func (stream DataInputStream) ReadUptoFinish(result AsyncResult) (string, uint, error) {
	var length0 C.gsize
	var err glib.Error
	ret0 := C.g_data_input_stream_read_upto_finish(stream.native(), result.native(), &length0, (**C.GError)(unsafe.Pointer(&err)))
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return "", 0, err.GoValue()
	}
	return ret, uint(length0), nil
}

// SetByteOrder is a wrapper around g_data_input_stream_set_byte_order().
func (stream DataInputStream) SetByteOrder(order DataStreamByteOrder) {
	C.g_data_input_stream_set_byte_order(stream.native(), C.GDataStreamByteOrder(order))
}

// SetNewlineType is a wrapper around g_data_input_stream_set_newline_type().
func (stream DataInputStream) SetNewlineType(type_ DataStreamNewlineType) {
	C.g_data_input_stream_set_newline_type(stream.native(), C.GDataStreamNewlineType(type_))
}

// Object DataOutputStream
type DataOutputStream struct {
	FilterOutputStream
}

func (v DataOutputStream) native() *C.GDataOutputStream {
	return (*C.GDataOutputStream)(v.Ptr)
}
func wrapDataOutputStream(p *C.GDataOutputStream) (v DataOutputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapDataOutputStream(p unsafe.Pointer) (v DataOutputStream) {
	v.Ptr = p
	return
}
func (v DataOutputStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapDataOutputStream(p unsafe.Pointer) interface{} {
	return WrapDataOutputStream(p)
}
func (v DataOutputStream) GetType() gobject.Type {
	return gobject.Type(C.g_data_output_stream_get_type())
}
func (v DataOutputStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapDataOutputStream(unsafe.Pointer(ptr)), nil
	}
}
func (v DataOutputStream) Seekable() Seekable {
	return WrapSeekable(v.Ptr)
}

// DataOutputStreamNew is a wrapper around g_data_output_stream_new().
func DataOutputStreamNew(base_stream OutputStream) DataOutputStream {
	ret0 := C.g_data_output_stream_new(base_stream.native())
	return wrapDataOutputStream(ret0)
}

// GetByteOrder is a wrapper around g_data_output_stream_get_byte_order().
func (stream DataOutputStream) GetByteOrder() DataStreamByteOrder {
	ret0 := C.g_data_output_stream_get_byte_order(stream.native())
	return DataStreamByteOrder(ret0)
}

// PutByte is a wrapper around g_data_output_stream_put_byte().
func (stream DataOutputStream) PutByte(data byte, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_data_output_stream_put_byte(stream.native(), C.guchar(data), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// PutInt16 is a wrapper around g_data_output_stream_put_int16().
func (stream DataOutputStream) PutInt16(data int16, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_data_output_stream_put_int16(stream.native(), C.gint16(data), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// PutInt32 is a wrapper around g_data_output_stream_put_int32().
func (stream DataOutputStream) PutInt32(data int32, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_data_output_stream_put_int32(stream.native(), C.gint32(data), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// PutInt64 is a wrapper around g_data_output_stream_put_int64().
func (stream DataOutputStream) PutInt64(data int64, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_data_output_stream_put_int64(stream.native(), C.gint64(data), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// PutString is a wrapper around g_data_output_stream_put_string().
func (stream DataOutputStream) PutString(str string, cancellable Cancellable) (bool, error) {
	str0 := C.CString(str)
	var err glib.Error
	ret0 := C.g_data_output_stream_put_string(stream.native(), str0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(str0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// PutUint16 is a wrapper around g_data_output_stream_put_uint16().
func (stream DataOutputStream) PutUint16(data uint16, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_data_output_stream_put_uint16(stream.native(), C.guint16(data), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// PutUint32 is a wrapper around g_data_output_stream_put_uint32().
func (stream DataOutputStream) PutUint32(data uint32, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_data_output_stream_put_uint32(stream.native(), C.guint32(data), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// PutUint64 is a wrapper around g_data_output_stream_put_uint64().
func (stream DataOutputStream) PutUint64(data uint64, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_data_output_stream_put_uint64(stream.native(), C.guint64(data), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetByteOrder is a wrapper around g_data_output_stream_set_byte_order().
func (stream DataOutputStream) SetByteOrder(order DataStreamByteOrder) {
	C.g_data_output_stream_set_byte_order(stream.native(), C.GDataStreamByteOrder(order))
}

// Object Emblem
type Emblem struct {
	gobject.Object
}

func (v Emblem) native() *C.GEmblem {
	return (*C.GEmblem)(v.Ptr)
}
func wrapEmblem(p *C.GEmblem) (v Emblem) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapEmblem(p unsafe.Pointer) (v Emblem) {
	v.Ptr = p
	return
}
func (v Emblem) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEmblem(p unsafe.Pointer) interface{} {
	return WrapEmblem(p)
}
func (v Emblem) GetType() gobject.Type {
	return gobject.Type(C.g_emblem_get_type())
}
func (v Emblem) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapEmblem(unsafe.Pointer(ptr)), nil
	}
}
func (v Emblem) Icon() Icon {
	return WrapIcon(v.Ptr)
}

// EmblemNew is a wrapper around g_emblem_new().
func EmblemNew(icon Icon) Emblem {
	ret0 := C.g_emblem_new(icon.native())
	return wrapEmblem(ret0)
}

// EmblemNewWithOrigin is a wrapper around g_emblem_new_with_origin().
func EmblemNewWithOrigin(icon Icon, origin EmblemOrigin) Emblem {
	ret0 := C.g_emblem_new_with_origin(icon.native(), C.GEmblemOrigin(origin))
	return wrapEmblem(ret0)
}

// GetIcon is a wrapper around g_emblem_get_icon().
func (emblem Emblem) GetIcon() Icon {
	ret0 := C.g_emblem_get_icon(emblem.native())
	return wrapIcon(ret0)
}

// GetOrigin is a wrapper around g_emblem_get_origin().
func (emblem Emblem) GetOrigin() EmblemOrigin {
	ret0 := C.g_emblem_get_origin(emblem.native())
	return EmblemOrigin(ret0)
}

// Object EmblemedIcon
type EmblemedIcon struct {
	gobject.Object
}

func (v EmblemedIcon) native() *C.GEmblemedIcon {
	return (*C.GEmblemedIcon)(v.Ptr)
}
func wrapEmblemedIcon(p *C.GEmblemedIcon) (v EmblemedIcon) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapEmblemedIcon(p unsafe.Pointer) (v EmblemedIcon) {
	v.Ptr = p
	return
}
func (v EmblemedIcon) IsNil() bool {
	return v.Ptr == nil
}
func IWrapEmblemedIcon(p unsafe.Pointer) interface{} {
	return WrapEmblemedIcon(p)
}
func (v EmblemedIcon) GetType() gobject.Type {
	return gobject.Type(C.g_emblemed_icon_get_type())
}
func (v EmblemedIcon) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapEmblemedIcon(unsafe.Pointer(ptr)), nil
	}
}
func (v EmblemedIcon) Icon() Icon {
	return WrapIcon(v.Ptr)
}

// AddEmblem is a wrapper around g_emblemed_icon_add_emblem().
func (emblemed EmblemedIcon) AddEmblem(emblem Emblem) {
	C.g_emblemed_icon_add_emblem(emblemed.native(), emblem.native())
}

// ClearEmblems is a wrapper around g_emblemed_icon_clear_emblems().
func (emblemed EmblemedIcon) ClearEmblems() {
	C.g_emblemed_icon_clear_emblems(emblemed.native())
}

// GetEmblems is a wrapper around g_emblemed_icon_get_emblems().
func (emblemed EmblemedIcon) GetEmblems() glib.List {
	ret0 := C.g_emblemed_icon_get_emblems(emblemed.native())
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapEmblem(p) }) /*gir:GLib*/
}

// GetIcon is a wrapper around g_emblemed_icon_get_icon().
func (emblemed EmblemedIcon) GetIcon() Icon {
	ret0 := C.g_emblemed_icon_get_icon(emblemed.native())
	return wrapIcon(ret0)
}

// Object FileIcon
type FileIcon struct {
	gobject.Object
}

func (v FileIcon) native() *C.GFileIcon {
	return (*C.GFileIcon)(v.Ptr)
}
func wrapFileIcon(p *C.GFileIcon) (v FileIcon) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFileIcon(p unsafe.Pointer) (v FileIcon) {
	v.Ptr = p
	return
}
func (v FileIcon) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFileIcon(p unsafe.Pointer) interface{} {
	return WrapFileIcon(p)
}
func (v FileIcon) GetType() gobject.Type {
	return gobject.Type(C.g_file_icon_get_type())
}
func (v FileIcon) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFileIcon(unsafe.Pointer(ptr)), nil
	}
}
func (v FileIcon) Icon() Icon {
	return WrapIcon(v.Ptr)
}
func (v FileIcon) LoadableIcon() LoadableIcon {
	return WrapLoadableIcon(v.Ptr)
}

// GetFile is a wrapper around g_file_icon_get_file().
func (icon FileIcon) GetFile() File {
	ret0 := C.g_file_icon_get_file(icon.native())
	return wrapFile(ret0)
}

// Object FilenameCompleter
type FilenameCompleter struct {
	gobject.Object
}

func (v FilenameCompleter) native() *C.GFilenameCompleter {
	return (*C.GFilenameCompleter)(v.Ptr)
}
func wrapFilenameCompleter(p *C.GFilenameCompleter) (v FilenameCompleter) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapFilenameCompleter(p unsafe.Pointer) (v FilenameCompleter) {
	v.Ptr = p
	return
}
func (v FilenameCompleter) IsNil() bool {
	return v.Ptr == nil
}
func IWrapFilenameCompleter(p unsafe.Pointer) interface{} {
	return WrapFilenameCompleter(p)
}
func (v FilenameCompleter) GetType() gobject.Type {
	return gobject.Type(C.g_filename_completer_get_type())
}
func (v FilenameCompleter) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapFilenameCompleter(unsafe.Pointer(ptr)), nil
	}
}

// FilenameCompleterNew is a wrapper around g_filename_completer_new().
func FilenameCompleterNew() FilenameCompleter {
	ret0 := C.g_filename_completer_new()
	return wrapFilenameCompleter(ret0)
}

// GetCompletionSuffix is a wrapper around g_filename_completer_get_completion_suffix().
func (completer FilenameCompleter) GetCompletionSuffix(initial_text string) string {
	initial_text0 := C.CString(initial_text)
	ret0 := C.g_filename_completer_get_completion_suffix(completer.native(), initial_text0)
	C.free(unsafe.Pointer(initial_text0)) /*ch:<stdlib.h>*/
	ret := C.GoString(ret0)
	C.g_free(C.gpointer(ret0))
	return ret
}

// GetCompletions is a wrapper around g_filename_completer_get_completions().
func (completer FilenameCompleter) GetCompletions(initial_text string) []string {
	initial_text0 := C.CString(initial_text)
	ret0 := C.g_filename_completer_get_completions(completer.native(), initial_text0)
	C.free(unsafe.Pointer(initial_text0)) /*ch:<stdlib.h>*/
	var ret0Slice []*C.char
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0),
		util.GetZeroTermArrayLen(unsafe.Pointer(ret0), unsafe.Sizeof(uintptr(0))) /*go:.util*/) /*go:.util*/
	ret := make([]string, len(ret0Slice))
	for idx, elem := range ret0Slice {
		elemG := C.GoString(elem)
		ret[idx] = elemG
		C.g_free(C.gpointer(elem))
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// SetDirsOnly is a wrapper around g_filename_completer_set_dirs_only().
func (completer FilenameCompleter) SetDirsOnly(dirs_only bool) {
	C.g_filename_completer_set_dirs_only(completer.native(), C.gboolean(util.Bool2Int(dirs_only)) /*go:.util*/)
}

// Object IOModule
type IOModule struct {
	gobject.TypeModule
}

func (v IOModule) native() *C.GIOModule {
	return (*C.GIOModule)(v.Ptr)
}
func wrapIOModule(p *C.GIOModule) (v IOModule) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapIOModule(p unsafe.Pointer) (v IOModule) {
	v.Ptr = p
	return
}
func (v IOModule) IsNil() bool {
	return v.Ptr == nil
}
func IWrapIOModule(p unsafe.Pointer) interface{} {
	return WrapIOModule(p)
}
func (v IOModule) GetType() gobject.Type {
	return gobject.Type(C.g_io_module_get_type())
}
func (v IOModule) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapIOModule(unsafe.Pointer(ptr)), nil
	}
}
func (v IOModule) TypePlugin() gobject.TypePlugin {
	return gobject.WrapTypePlugin(v.Ptr) /*gir:GObject*/
}

// Object InetAddressMask
type InetAddressMask struct {
	gobject.Object
}

func (v InetAddressMask) native() *C.GInetAddressMask {
	return (*C.GInetAddressMask)(v.Ptr)
}
func wrapInetAddressMask(p *C.GInetAddressMask) (v InetAddressMask) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapInetAddressMask(p unsafe.Pointer) (v InetAddressMask) {
	v.Ptr = p
	return
}
func (v InetAddressMask) IsNil() bool {
	return v.Ptr == nil
}
func IWrapInetAddressMask(p unsafe.Pointer) interface{} {
	return WrapInetAddressMask(p)
}
func (v InetAddressMask) GetType() gobject.Type {
	return gobject.Type(C.g_inet_address_mask_get_type())
}
func (v InetAddressMask) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapInetAddressMask(unsafe.Pointer(ptr)), nil
	}
}
func (v InetAddressMask) Initable() Initable {
	return WrapInitable(v.Ptr)
}

// InetAddressMaskNew is a wrapper around g_inet_address_mask_new().
func InetAddressMaskNew(addr InetAddress, length uint) (InetAddressMask, error) {
	var err glib.Error
	ret0 := C.g_inet_address_mask_new(addr.native(), C.guint(length), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return InetAddressMask{}, err.GoValue()
	}
	return wrapInetAddressMask(ret0), nil
}

// InetAddressMaskNewFromString is a wrapper around g_inet_address_mask_new_from_string().
func InetAddressMaskNewFromString(mask_string string) (InetAddressMask, error) {
	mask_string0 := (*C.gchar)(C.CString(mask_string))
	var err glib.Error
	ret0 := C.g_inet_address_mask_new_from_string(mask_string0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(mask_string0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return InetAddressMask{}, err.GoValue()
	}
	return wrapInetAddressMask(ret0), nil
}

// Equal is a wrapper around g_inet_address_mask_equal().
func (mask InetAddressMask) Equal(mask2 InetAddressMask) bool {
	ret0 := C.g_inet_address_mask_equal(mask.native(), mask2.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetAddress is a wrapper around g_inet_address_mask_get_address().
func (mask InetAddressMask) GetAddress() InetAddress {
	ret0 := C.g_inet_address_mask_get_address(mask.native())
	return wrapInetAddress(ret0)
}

// GetFamily is a wrapper around g_inet_address_mask_get_family().
func (mask InetAddressMask) GetFamily() SocketFamily {
	ret0 := C.g_inet_address_mask_get_family(mask.native())
	return SocketFamily(ret0)
}

// GetLength is a wrapper around g_inet_address_mask_get_length().
func (mask InetAddressMask) GetLength() uint {
	ret0 := C.g_inet_address_mask_get_length(mask.native())
	return uint(ret0)
}

// Matches is a wrapper around g_inet_address_mask_matches().
func (mask InetAddressMask) Matches(address InetAddress) bool {
	ret0 := C.g_inet_address_mask_matches(mask.native(), address.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ToString is a wrapper around g_inet_address_mask_to_string().
func (mask InetAddressMask) ToString() string {
	ret0 := C.g_inet_address_mask_to_string(mask.native())
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	return ret
}

// Object ListStore
type ListStore struct {
	gobject.Object
}

func (v ListStore) native() *C.GListStore {
	return (*C.GListStore)(v.Ptr)
}
func wrapListStore(p *C.GListStore) (v ListStore) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapListStore(p unsafe.Pointer) (v ListStore) {
	v.Ptr = p
	return
}
func (v ListStore) IsNil() bool {
	return v.Ptr == nil
}
func IWrapListStore(p unsafe.Pointer) interface{} {
	return WrapListStore(p)
}
func (v ListStore) GetType() gobject.Type {
	return gobject.Type(C.g_list_store_get_type())
}
func (v ListStore) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapListStore(unsafe.Pointer(ptr)), nil
	}
}
func (v ListStore) ListModel() ListModel {
	return WrapListModel(v.Ptr)
}

// Append is a wrapper around g_list_store_append().
func (store ListStore) Append(item gobject.Object) {
	C.g_list_store_append(store.native(), C.gpointer((C.gpointer)(item.Ptr)))
}

// Insert is a wrapper around g_list_store_insert().
func (store ListStore) Insert(position uint, item gobject.Object) {
	C.g_list_store_insert(store.native(), C.guint(position), C.gpointer((C.gpointer)(item.Ptr)))
}

// Remove is a wrapper around g_list_store_remove().
func (store ListStore) Remove(position uint) {
	C.g_list_store_remove(store.native(), C.guint(position))
}

// RemoveAll is a wrapper around g_list_store_remove_all().
func (store ListStore) RemoveAll() {
	C.g_list_store_remove_all(store.native())
}

// Splice is a wrapper around g_list_store_splice().
func (store ListStore) Splice(position uint, n_removals uint, additions []gobject.Object) {
	additions0 := make([]C.gpointer, len(additions))
	for idx, elemG := range additions {
		additions0[idx] = C.gpointer((C.gpointer)(elemG.Ptr))
	}
	var additions0Ptr *C.gpointer
	if len(additions0) > 0 {
		additions0Ptr = &additions0[0]
	}
	C.g_list_store_splice(store.native(), C.guint(position), C.guint(n_removals), additions0Ptr, C.guint(len(additions)))
}

// Object MemoryInputStream
type MemoryInputStream struct {
	InputStream
}

func (v MemoryInputStream) native() *C.GMemoryInputStream {
	return (*C.GMemoryInputStream)(v.Ptr)
}
func wrapMemoryInputStream(p *C.GMemoryInputStream) (v MemoryInputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMemoryInputStream(p unsafe.Pointer) (v MemoryInputStream) {
	v.Ptr = p
	return
}
func (v MemoryInputStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMemoryInputStream(p unsafe.Pointer) interface{} {
	return WrapMemoryInputStream(p)
}
func (v MemoryInputStream) GetType() gobject.Type {
	return gobject.Type(C.g_memory_input_stream_get_type())
}
func (v MemoryInputStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMemoryInputStream(unsafe.Pointer(ptr)), nil
	}
}
func (v MemoryInputStream) PollableInputStream() PollableInputStream {
	return WrapPollableInputStream(v.Ptr)
}
func (v MemoryInputStream) Seekable() Seekable {
	return WrapSeekable(v.Ptr)
}

// MemoryInputStreamNew is a wrapper around g_memory_input_stream_new().
func MemoryInputStreamNew() InputStream {
	ret0 := C.g_memory_input_stream_new()
	return wrapInputStream(ret0)
}

// MemoryInputStreamNewFromBytes is a wrapper around g_memory_input_stream_new_from_bytes().
func MemoryInputStreamNewFromBytes(bytes glib.Bytes) InputStream {
	ret0 := C.g_memory_input_stream_new_from_bytes((*C.GBytes)(bytes.Ptr))
	return wrapInputStream(ret0)
}

// AddBytes is a wrapper around g_memory_input_stream_add_bytes().
func (stream MemoryInputStream) AddBytes(bytes glib.Bytes) {
	C.g_memory_input_stream_add_bytes(stream.native(), (*C.GBytes)(bytes.Ptr))
}

// Object MemoryOutputStream
type MemoryOutputStream struct {
	OutputStream
}

func (v MemoryOutputStream) native() *C.GMemoryOutputStream {
	return (*C.GMemoryOutputStream)(v.Ptr)
}
func wrapMemoryOutputStream(p *C.GMemoryOutputStream) (v MemoryOutputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMemoryOutputStream(p unsafe.Pointer) (v MemoryOutputStream) {
	v.Ptr = p
	return
}
func (v MemoryOutputStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMemoryOutputStream(p unsafe.Pointer) interface{} {
	return WrapMemoryOutputStream(p)
}
func (v MemoryOutputStream) GetType() gobject.Type {
	return gobject.Type(C.g_memory_output_stream_get_type())
}
func (v MemoryOutputStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMemoryOutputStream(unsafe.Pointer(ptr)), nil
	}
}
func (v MemoryOutputStream) PollableOutputStream() PollableOutputStream {
	return WrapPollableOutputStream(v.Ptr)
}
func (v MemoryOutputStream) Seekable() Seekable {
	return WrapSeekable(v.Ptr)
}

// MemoryOutputStreamNewResizable is a wrapper around g_memory_output_stream_new_resizable().
func MemoryOutputStreamNewResizable() OutputStream {
	ret0 := C.g_memory_output_stream_new_resizable()
	return wrapOutputStream(ret0)
}

// GetData is a wrapper around g_memory_output_stream_get_data().
func (ostream MemoryOutputStream) GetData() unsafe.Pointer {
	ret0 := C.g_memory_output_stream_get_data(ostream.native())
	return unsafe.Pointer(ret0)
}

// GetDataSize is a wrapper around g_memory_output_stream_get_data_size().
func (ostream MemoryOutputStream) GetDataSize() uint {
	ret0 := C.g_memory_output_stream_get_data_size(ostream.native())
	return uint(ret0)
}

// GetSize is a wrapper around g_memory_output_stream_get_size().
func (ostream MemoryOutputStream) GetSize() uint {
	ret0 := C.g_memory_output_stream_get_size(ostream.native())
	return uint(ret0)
}

// StealAsBytes is a wrapper around g_memory_output_stream_steal_as_bytes().
func (ostream MemoryOutputStream) StealAsBytes() glib.Bytes {
	ret0 := C.g_memory_output_stream_steal_as_bytes(ostream.native())
	return glib.WrapBytes(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// StealData is a wrapper around g_memory_output_stream_steal_data().
func (ostream MemoryOutputStream) StealData() unsafe.Pointer {
	ret0 := C.g_memory_output_stream_steal_data(ostream.native())
	return unsafe.Pointer(ret0)
}

// Object Menu
type Menu struct {
	MenuModel
}

func (v Menu) native() *C.GMenu {
	return (*C.GMenu)(v.Ptr)
}
func wrapMenu(p *C.GMenu) (v Menu) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMenu(p unsafe.Pointer) (v Menu) {
	v.Ptr = p
	return
}
func (v Menu) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMenu(p unsafe.Pointer) interface{} {
	return WrapMenu(p)
}
func (v Menu) GetType() gobject.Type {
	return gobject.Type(C.g_menu_get_type())
}
func (v Menu) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMenu(unsafe.Pointer(ptr)), nil
	}
}

// MenuNew is a wrapper around g_menu_new().
func MenuNew() Menu {
	ret0 := C.g_menu_new()
	return wrapMenu(ret0)
}

// Append is a wrapper around g_menu_append().
func (menu Menu) Append(label string, detailed_action string) {
	label0 := (*C.gchar)(C.CString(label))
	detailed_action0 := (*C.gchar)(C.CString(detailed_action))
	C.g_menu_append(menu.native(), label0, detailed_action0)
	C.free(unsafe.Pointer(label0))           /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(detailed_action0)) /*ch:<stdlib.h>*/
}

// AppendItem is a wrapper around g_menu_append_item().
func (menu Menu) AppendItem(item MenuItem) {
	C.g_menu_append_item(menu.native(), item.native())
}

// AppendSection is a wrapper around g_menu_append_section().
func (menu Menu) AppendSection(label string, section MenuModel) {
	label0 := (*C.gchar)(C.CString(label))
	C.g_menu_append_section(menu.native(), label0, section.native())
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
}

// AppendSubmenu is a wrapper around g_menu_append_submenu().
func (menu Menu) AppendSubmenu(label string, submenu MenuModel) {
	label0 := (*C.gchar)(C.CString(label))
	C.g_menu_append_submenu(menu.native(), label0, submenu.native())
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
}

// Freeze is a wrapper around g_menu_freeze().
func (menu Menu) Freeze() {
	C.g_menu_freeze(menu.native())
}

// Insert is a wrapper around g_menu_insert().
func (menu Menu) Insert(position int, label string, detailed_action string) {
	label0 := (*C.gchar)(C.CString(label))
	detailed_action0 := (*C.gchar)(C.CString(detailed_action))
	C.g_menu_insert(menu.native(), C.gint(position), label0, detailed_action0)
	C.free(unsafe.Pointer(label0))           /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(detailed_action0)) /*ch:<stdlib.h>*/
}

// InsertItem is a wrapper around g_menu_insert_item().
func (menu Menu) InsertItem(position int, item MenuItem) {
	C.g_menu_insert_item(menu.native(), C.gint(position), item.native())
}

// InsertSection is a wrapper around g_menu_insert_section().
func (menu Menu) InsertSection(position int, label string, section MenuModel) {
	label0 := (*C.gchar)(C.CString(label))
	C.g_menu_insert_section(menu.native(), C.gint(position), label0, section.native())
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
}

// InsertSubmenu is a wrapper around g_menu_insert_submenu().
func (menu Menu) InsertSubmenu(position int, label string, submenu MenuModel) {
	label0 := (*C.gchar)(C.CString(label))
	C.g_menu_insert_submenu(menu.native(), C.gint(position), label0, submenu.native())
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
}

// Prepend is a wrapper around g_menu_prepend().
func (menu Menu) Prepend(label string, detailed_action string) {
	label0 := (*C.gchar)(C.CString(label))
	detailed_action0 := (*C.gchar)(C.CString(detailed_action))
	C.g_menu_prepend(menu.native(), label0, detailed_action0)
	C.free(unsafe.Pointer(label0))           /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(detailed_action0)) /*ch:<stdlib.h>*/
}

// PrependItem is a wrapper around g_menu_prepend_item().
func (menu Menu) PrependItem(item MenuItem) {
	C.g_menu_prepend_item(menu.native(), item.native())
}

// PrependSection is a wrapper around g_menu_prepend_section().
func (menu Menu) PrependSection(label string, section MenuModel) {
	label0 := (*C.gchar)(C.CString(label))
	C.g_menu_prepend_section(menu.native(), label0, section.native())
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
}

// PrependSubmenu is a wrapper around g_menu_prepend_submenu().
func (menu Menu) PrependSubmenu(label string, submenu MenuModel) {
	label0 := (*C.gchar)(C.CString(label))
	C.g_menu_prepend_submenu(menu.native(), label0, submenu.native())
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
}

// Remove is a wrapper around g_menu_remove().
func (menu Menu) Remove(position int) {
	C.g_menu_remove(menu.native(), C.gint(position))
}

// RemoveAll is a wrapper around g_menu_remove_all().
func (menu Menu) RemoveAll() {
	C.g_menu_remove_all(menu.native())
}

// Object MenuItem
type MenuItem struct {
	gobject.Object
}

func (v MenuItem) native() *C.GMenuItem {
	return (*C.GMenuItem)(v.Ptr)
}
func wrapMenuItem(p *C.GMenuItem) (v MenuItem) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapMenuItem(p unsafe.Pointer) (v MenuItem) {
	v.Ptr = p
	return
}
func (v MenuItem) IsNil() bool {
	return v.Ptr == nil
}
func IWrapMenuItem(p unsafe.Pointer) interface{} {
	return WrapMenuItem(p)
}
func (v MenuItem) GetType() gobject.Type {
	return gobject.Type(C.g_menu_item_get_type())
}
func (v MenuItem) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapMenuItem(unsafe.Pointer(ptr)), nil
	}
}

// MenuItemNew is a wrapper around g_menu_item_new().
func MenuItemNew(label string, detailed_action string) MenuItem {
	label0 := (*C.gchar)(C.CString(label))
	detailed_action0 := (*C.gchar)(C.CString(detailed_action))
	ret0 := C.g_menu_item_new(label0, detailed_action0)
	C.free(unsafe.Pointer(label0))           /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(detailed_action0)) /*ch:<stdlib.h>*/
	return wrapMenuItem(ret0)
}

// MenuItemNewFromModel is a wrapper around g_menu_item_new_from_model().
func MenuItemNewFromModel(model MenuModel, item_index int) MenuItem {
	ret0 := C.g_menu_item_new_from_model(model.native(), C.gint(item_index))
	return wrapMenuItem(ret0)
}

// MenuItemNewSection is a wrapper around g_menu_item_new_section().
func MenuItemNewSection(label string, section MenuModel) MenuItem {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.g_menu_item_new_section(label0, section.native())
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapMenuItem(ret0)
}

// MenuItemNewSubmenu is a wrapper around g_menu_item_new_submenu().
func MenuItemNewSubmenu(label string, submenu MenuModel) MenuItem {
	label0 := (*C.gchar)(C.CString(label))
	ret0 := C.g_menu_item_new_submenu(label0, submenu.native())
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
	return wrapMenuItem(ret0)
}

// GetAttributeValue is a wrapper around g_menu_item_get_attribute_value().
func (menu_item MenuItem) GetAttributeValue(attribute string, expected_type glib.VariantType) glib.Variant {
	attribute0 := (*C.gchar)(C.CString(attribute))
	ret0 := C.g_menu_item_get_attribute_value(menu_item.native(), attribute0, (*C.GVariantType)(expected_type.Ptr))
	C.free(unsafe.Pointer(attribute0))            /*ch:<stdlib.h>*/
	return glib.WrapVariant(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetLink is a wrapper around g_menu_item_get_link().
func (menu_item MenuItem) GetLink(link string) MenuModel {
	link0 := (*C.gchar)(C.CString(link))
	ret0 := C.g_menu_item_get_link(menu_item.native(), link0)
	C.free(unsafe.Pointer(link0)) /*ch:<stdlib.h>*/
	return wrapMenuModel(ret0)
}

// SetActionAndTargetValue is a wrapper around g_menu_item_set_action_and_target_value().
func (menu_item MenuItem) SetActionAndTargetValue(action string, target_value glib.Variant) {
	action0 := (*C.gchar)(C.CString(action))
	C.g_menu_item_set_action_and_target_value(menu_item.native(), action0, (*C.GVariant)(target_value.Ptr))
	C.free(unsafe.Pointer(action0)) /*ch:<stdlib.h>*/
}

// SetAttributeValue is a wrapper around g_menu_item_set_attribute_value().
func (menu_item MenuItem) SetAttributeValue(attribute string, value glib.Variant) {
	attribute0 := (*C.gchar)(C.CString(attribute))
	C.g_menu_item_set_attribute_value(menu_item.native(), attribute0, (*C.GVariant)(value.Ptr))
	C.free(unsafe.Pointer(attribute0)) /*ch:<stdlib.h>*/
}

// SetDetailedAction is a wrapper around g_menu_item_set_detailed_action().
func (menu_item MenuItem) SetDetailedAction(detailed_action string) {
	detailed_action0 := (*C.gchar)(C.CString(detailed_action))
	C.g_menu_item_set_detailed_action(menu_item.native(), detailed_action0)
	C.free(unsafe.Pointer(detailed_action0)) /*ch:<stdlib.h>*/
}

// SetIcon is a wrapper around g_menu_item_set_icon().
func (menu_item MenuItem) SetIcon(icon Icon) {
	C.g_menu_item_set_icon(menu_item.native(), icon.native())
}

// SetLabel is a wrapper around g_menu_item_set_label().
func (menu_item MenuItem) SetLabel(label string) {
	label0 := (*C.gchar)(C.CString(label))
	C.g_menu_item_set_label(menu_item.native(), label0)
	C.free(unsafe.Pointer(label0)) /*ch:<stdlib.h>*/
}

// SetLink is a wrapper around g_menu_item_set_link().
func (menu_item MenuItem) SetLink(link string, model MenuModel) {
	link0 := (*C.gchar)(C.CString(link))
	C.g_menu_item_set_link(menu_item.native(), link0, model.native())
	C.free(unsafe.Pointer(link0)) /*ch:<stdlib.h>*/
}

// SetSection is a wrapper around g_menu_item_set_section().
func (menu_item MenuItem) SetSection(section MenuModel) {
	C.g_menu_item_set_section(menu_item.native(), section.native())
}

// SetSubmenu is a wrapper around g_menu_item_set_submenu().
func (menu_item MenuItem) SetSubmenu(submenu MenuModel) {
	C.g_menu_item_set_submenu(menu_item.native(), submenu.native())
}

// Object NativeVolumeMonitor
type NativeVolumeMonitor struct {
	VolumeMonitor
}

func (v NativeVolumeMonitor) native() *C.GNativeVolumeMonitor {
	return (*C.GNativeVolumeMonitor)(v.Ptr)
}
func wrapNativeVolumeMonitor(p *C.GNativeVolumeMonitor) (v NativeVolumeMonitor) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapNativeVolumeMonitor(p unsafe.Pointer) (v NativeVolumeMonitor) {
	v.Ptr = p
	return
}
func (v NativeVolumeMonitor) IsNil() bool {
	return v.Ptr == nil
}
func IWrapNativeVolumeMonitor(p unsafe.Pointer) interface{} {
	return WrapNativeVolumeMonitor(p)
}
func (v NativeVolumeMonitor) GetType() gobject.Type {
	return gobject.Type(C.g_native_volume_monitor_get_type())
}
func (v NativeVolumeMonitor) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapNativeVolumeMonitor(unsafe.Pointer(ptr)), nil
	}
}

// Object NetworkAddress
type NetworkAddress struct {
	gobject.Object
}

func (v NetworkAddress) native() *C.GNetworkAddress {
	return (*C.GNetworkAddress)(v.Ptr)
}
func wrapNetworkAddress(p *C.GNetworkAddress) (v NetworkAddress) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapNetworkAddress(p unsafe.Pointer) (v NetworkAddress) {
	v.Ptr = p
	return
}
func (v NetworkAddress) IsNil() bool {
	return v.Ptr == nil
}
func IWrapNetworkAddress(p unsafe.Pointer) interface{} {
	return WrapNetworkAddress(p)
}
func (v NetworkAddress) GetType() gobject.Type {
	return gobject.Type(C.g_network_address_get_type())
}
func (v NetworkAddress) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapNetworkAddress(unsafe.Pointer(ptr)), nil
	}
}
func (v NetworkAddress) SocketConnectable() SocketConnectable {
	return WrapSocketConnectable(v.Ptr)
}

// Object VolumeMonitor
type VolumeMonitor struct {
	gobject.Object
}

func (v VolumeMonitor) native() *C.GVolumeMonitor {
	return (*C.GVolumeMonitor)(v.Ptr)
}
func wrapVolumeMonitor(p *C.GVolumeMonitor) (v VolumeMonitor) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapVolumeMonitor(p unsafe.Pointer) (v VolumeMonitor) {
	v.Ptr = p
	return
}
func (v VolumeMonitor) IsNil() bool {
	return v.Ptr == nil
}
func IWrapVolumeMonitor(p unsafe.Pointer) interface{} {
	return WrapVolumeMonitor(p)
}
func (v VolumeMonitor) GetType() gobject.Type {
	return gobject.Type(C.g_volume_monitor_get_type())
}
func (v VolumeMonitor) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapVolumeMonitor(unsafe.Pointer(ptr)), nil
	}
}

// Object NetworkService
type NetworkService struct {
	gobject.Object
}

func (v NetworkService) native() *C.GNetworkService {
	return (*C.GNetworkService)(v.Ptr)
}
func wrapNetworkService(p *C.GNetworkService) (v NetworkService) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapNetworkService(p unsafe.Pointer) (v NetworkService) {
	v.Ptr = p
	return
}
func (v NetworkService) IsNil() bool {
	return v.Ptr == nil
}
func IWrapNetworkService(p unsafe.Pointer) interface{} {
	return WrapNetworkService(p)
}
func (v NetworkService) GetType() gobject.Type {
	return gobject.Type(C.g_network_service_get_type())
}
func (v NetworkService) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapNetworkService(unsafe.Pointer(ptr)), nil
	}
}
func (v NetworkService) SocketConnectable() SocketConnectable {
	return WrapSocketConnectable(v.Ptr)
}

// Object Permission
type Permission struct {
	gobject.Object
}

func (v Permission) native() *C.GPermission {
	return (*C.GPermission)(v.Ptr)
}
func wrapPermission(p *C.GPermission) (v Permission) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapPermission(p unsafe.Pointer) (v Permission) {
	v.Ptr = p
	return
}
func (v Permission) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPermission(p unsafe.Pointer) interface{} {
	return WrapPermission(p)
}
func (v Permission) GetType() gobject.Type {
	return gobject.Type(C.g_permission_get_type())
}
func (v Permission) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPermission(unsafe.Pointer(ptr)), nil
	}
}

// Acquire is a wrapper around g_permission_acquire().
func (permission Permission) Acquire(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_permission_acquire(permission.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// AcquireAsync is a wrapper around g_permission_acquire_async().
func (permission Permission) AcquireAsync(cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_permission_acquire_async(permission.native(), cancellable.native(), callback0)
}

// AcquireFinish is a wrapper around g_permission_acquire_finish().
func (permission Permission) AcquireFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_permission_acquire_finish(permission.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// GetAllowed is a wrapper around g_permission_get_allowed().
func (permission Permission) GetAllowed() bool {
	ret0 := C.g_permission_get_allowed(permission.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetCanAcquire is a wrapper around g_permission_get_can_acquire().
func (permission Permission) GetCanAcquire() bool {
	ret0 := C.g_permission_get_can_acquire(permission.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetCanRelease is a wrapper around g_permission_get_can_release().
func (permission Permission) GetCanRelease() bool {
	ret0 := C.g_permission_get_can_release(permission.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ImplUpdate is a wrapper around g_permission_impl_update().
func (permission Permission) ImplUpdate(allowed bool, can_acquire bool, can_release bool) {
	C.g_permission_impl_update(permission.native(), C.gboolean(util.Bool2Int(allowed)) /*go:.util*/, C.gboolean(util.Bool2Int(can_acquire)) /*go:.util*/, C.gboolean(util.Bool2Int(can_release)) /*go:.util*/)
}

// Release is a wrapper around g_permission_release().
func (permission Permission) Release(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_permission_release(permission.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// ReleaseAsync is a wrapper around g_permission_release_async().
func (permission Permission) ReleaseAsync(cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_permission_release_async(permission.native(), cancellable.native(), callback0)
}

// ReleaseFinish is a wrapper around g_permission_release_finish().
func (permission Permission) ReleaseFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_permission_release_finish(permission.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Object PropertyAction
type PropertyAction struct {
	gobject.Object
}

func (v PropertyAction) native() *C.GPropertyAction {
	return (*C.GPropertyAction)(v.Ptr)
}
func wrapPropertyAction(p *C.GPropertyAction) (v PropertyAction) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapPropertyAction(p unsafe.Pointer) (v PropertyAction) {
	v.Ptr = p
	return
}
func (v PropertyAction) IsNil() bool {
	return v.Ptr == nil
}
func IWrapPropertyAction(p unsafe.Pointer) interface{} {
	return WrapPropertyAction(p)
}
func (v PropertyAction) GetType() gobject.Type {
	return gobject.Type(C.g_property_action_get_type())
}
func (v PropertyAction) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapPropertyAction(unsafe.Pointer(ptr)), nil
	}
}
func (v PropertyAction) Action() Action {
	return WrapAction(v.Ptr)
}

// PropertyActionNew is a wrapper around g_property_action_new().
func PropertyActionNew(name string, object gobject.Object, property_name string) PropertyAction {
	name0 := (*C.gchar)(C.CString(name))
	property_name0 := (*C.gchar)(C.CString(property_name))
	ret0 := C.g_property_action_new(name0, C.gpointer((C.gpointer)(object.Ptr)), property_name0)
	C.free(unsafe.Pointer(name0))          /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(property_name0)) /*ch:<stdlib.h>*/
	return wrapPropertyAction(ret0)
}

// Object ProxyAddressEnumerator
type ProxyAddressEnumerator struct {
	SocketAddressEnumerator
}

func (v ProxyAddressEnumerator) native() *C.GProxyAddressEnumerator {
	return (*C.GProxyAddressEnumerator)(v.Ptr)
}
func wrapProxyAddressEnumerator(p *C.GProxyAddressEnumerator) (v ProxyAddressEnumerator) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapProxyAddressEnumerator(p unsafe.Pointer) (v ProxyAddressEnumerator) {
	v.Ptr = p
	return
}
func (v ProxyAddressEnumerator) IsNil() bool {
	return v.Ptr == nil
}
func IWrapProxyAddressEnumerator(p unsafe.Pointer) interface{} {
	return WrapProxyAddressEnumerator(p)
}
func (v ProxyAddressEnumerator) GetType() gobject.Type {
	return gobject.Type(C.g_proxy_address_enumerator_get_type())
}
func (v ProxyAddressEnumerator) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapProxyAddressEnumerator(unsafe.Pointer(ptr)), nil
	}
}

// Object Resolver
type Resolver struct {
	gobject.Object
}

func (v Resolver) native() *C.GResolver {
	return (*C.GResolver)(v.Ptr)
}
func wrapResolver(p *C.GResolver) (v Resolver) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapResolver(p unsafe.Pointer) (v Resolver) {
	v.Ptr = p
	return
}
func (v Resolver) IsNil() bool {
	return v.Ptr == nil
}
func IWrapResolver(p unsafe.Pointer) interface{} {
	return WrapResolver(p)
}
func (v Resolver) GetType() gobject.Type {
	return gobject.Type(C.g_resolver_get_type())
}
func (v Resolver) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapResolver(unsafe.Pointer(ptr)), nil
	}
}

// LookupByAddress is a wrapper around g_resolver_lookup_by_address().
func (resolver Resolver) LookupByAddress(address InetAddress, cancellable Cancellable) (string, error) {
	var err glib.Error
	ret0 := C.g_resolver_lookup_by_address(resolver.native(), address.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return "", err.GoValue()
	}
	return ret, nil
}

// LookupByAddressAsync is a wrapper around g_resolver_lookup_by_address_async().
func (resolver Resolver) LookupByAddressAsync(address InetAddress, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_resolver_lookup_by_address_async(resolver.native(), address.native(), cancellable.native(), callback0)
}

// LookupByAddressFinish is a wrapper around g_resolver_lookup_by_address_finish().
func (resolver Resolver) LookupByAddressFinish(result AsyncResult) (string, error) {
	var err glib.Error
	ret0 := C.g_resolver_lookup_by_address_finish(resolver.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	ret := C.GoString((*C.char)(ret0))
	C.g_free(C.gpointer(ret0))
	if err.Ptr != nil {
		defer err.Free()
		return "", err.GoValue()
	}
	return ret, nil
}

// LookupByName is a wrapper around g_resolver_lookup_by_name().
func (resolver Resolver) LookupByName(hostname string, cancellable Cancellable) (glib.List, error) {
	hostname0 := (*C.gchar)(C.CString(hostname))
	var err glib.Error
	ret0 := C.g_resolver_lookup_by_name(resolver.native(), hostname0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(hostname0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return glib.List{}, err.GoValue()
	}
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapInetAddress(p) }), /*gir:GLib*/ nil
}

// LookupByNameAsync is a wrapper around g_resolver_lookup_by_name_async().
func (resolver Resolver) LookupByNameAsync(hostname string, cancellable Cancellable, callback AsyncReadyCallback) {
	hostname0 := (*C.gchar)(C.CString(hostname))
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_resolver_lookup_by_name_async(resolver.native(), hostname0, cancellable.native(), callback0)
	C.free(unsafe.Pointer(hostname0)) /*ch:<stdlib.h>*/
}

// LookupByNameFinish is a wrapper around g_resolver_lookup_by_name_finish().
func (resolver Resolver) LookupByNameFinish(result AsyncResult) (glib.List, error) {
	var err glib.Error
	ret0 := C.g_resolver_lookup_by_name_finish(resolver.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return glib.List{}, err.GoValue()
	}
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapInetAddress(p) }), /*gir:GLib*/ nil
}

// LookupRecordsAsync is a wrapper around g_resolver_lookup_records_async().
func (resolver Resolver) LookupRecordsAsync(rrname string, record_type ResolverRecordType, cancellable Cancellable, callback AsyncReadyCallback) {
	rrname0 := (*C.gchar)(C.CString(rrname))
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_resolver_lookup_records_async(resolver.native(), rrname0, C.GResolverRecordType(record_type), cancellable.native(), callback0)
	C.free(unsafe.Pointer(rrname0)) /*ch:<stdlib.h>*/
}

// LookupService is a wrapper around g_resolver_lookup_service().
func (resolver Resolver) LookupService(service string, protocol string, domain string, cancellable Cancellable) (glib.List, error) {
	service0 := (*C.gchar)(C.CString(service))
	protocol0 := (*C.gchar)(C.CString(protocol))
	domain0 := (*C.gchar)(C.CString(domain))
	var err glib.Error
	ret0 := C.g_resolver_lookup_service(resolver.native(), service0, protocol0, domain0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(service0))  /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(protocol0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(domain0))   /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return glib.List{}, err.GoValue()
	}
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapSrvTarget(p) }), /*gir:GLib*/ nil
}

// LookupServiceAsync is a wrapper around g_resolver_lookup_service_async().
func (resolver Resolver) LookupServiceAsync(service string, protocol string, domain string, cancellable Cancellable, callback AsyncReadyCallback) {
	service0 := (*C.gchar)(C.CString(service))
	protocol0 := (*C.gchar)(C.CString(protocol))
	domain0 := (*C.gchar)(C.CString(domain))
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_resolver_lookup_service_async(resolver.native(), service0, protocol0, domain0, cancellable.native(), callback0)
	C.free(unsafe.Pointer(service0))  /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(protocol0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(domain0))   /*ch:<stdlib.h>*/
}

// LookupServiceFinish is a wrapper around g_resolver_lookup_service_finish().
func (resolver Resolver) LookupServiceFinish(result AsyncResult) (glib.List, error) {
	var err glib.Error
	ret0 := C.g_resolver_lookup_service_finish(resolver.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return glib.List{}, err.GoValue()
	}
	return glib.WrapList(unsafe.Pointer(ret0),
		func(p unsafe.Pointer) interface{} { return WrapSrvTarget(p) }), /*gir:GLib*/ nil
}

// SetDefault is a wrapper around g_resolver_set_default().
func (resolver Resolver) SetDefault() {
	C.g_resolver_set_default(resolver.native())
}

// ResolverFreeAddresses is a wrapper around g_resolver_free_addresses().
func ResolverFreeAddresses(addresses glib.List) {
	C.g_resolver_free_addresses((*C.GList)(addresses.Ptr))
}

// ResolverFreeTargets is a wrapper around g_resolver_free_targets().
func ResolverFreeTargets(targets glib.List) {
	C.g_resolver_free_targets((*C.GList)(targets.Ptr))
}

// ResolverGetDefault is a wrapper around g_resolver_get_default().
func ResolverGetDefault() Resolver {
	ret0 := C.g_resolver_get_default()
	return wrapResolver(ret0)
}

// Object SimpleAction
type SimpleAction struct {
	gobject.Object
}

func (v SimpleAction) native() *C.GSimpleAction {
	return (*C.GSimpleAction)(v.Ptr)
}
func wrapSimpleAction(p *C.GSimpleAction) (v SimpleAction) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSimpleAction(p unsafe.Pointer) (v SimpleAction) {
	v.Ptr = p
	return
}
func (v SimpleAction) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSimpleAction(p unsafe.Pointer) interface{} {
	return WrapSimpleAction(p)
}
func (v SimpleAction) GetType() gobject.Type {
	return gobject.Type(C.g_simple_action_get_type())
}
func (v SimpleAction) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSimpleAction(unsafe.Pointer(ptr)), nil
	}
}
func (v SimpleAction) Action() Action {
	return WrapAction(v.Ptr)
}

// SimpleActionNew is a wrapper around g_simple_action_new().
func SimpleActionNew(name string, parameter_type glib.VariantType) SimpleAction {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_simple_action_new(name0, (*C.GVariantType)(parameter_type.Ptr))
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapSimpleAction(ret0)
}

// SimpleActionNewStateful is a wrapper around g_simple_action_new_stateful().
func SimpleActionNewStateful(name string, parameter_type glib.VariantType, state glib.Variant) SimpleAction {
	name0 := (*C.gchar)(C.CString(name))
	ret0 := C.g_simple_action_new_stateful(name0, (*C.GVariantType)(parameter_type.Ptr), (*C.GVariant)(state.Ptr))
	C.free(unsafe.Pointer(name0)) /*ch:<stdlib.h>*/
	return wrapSimpleAction(ret0)
}

// SetEnabled is a wrapper around g_simple_action_set_enabled().
func (simple SimpleAction) SetEnabled(enabled bool) {
	C.g_simple_action_set_enabled(simple.native(), C.gboolean(util.Bool2Int(enabled)) /*go:.util*/)
}

// SetState is a wrapper around g_simple_action_set_state().
func (simple SimpleAction) SetState(value glib.Variant) {
	C.g_simple_action_set_state(simple.native(), (*C.GVariant)(value.Ptr))
}

// SetStateHint is a wrapper around g_simple_action_set_state_hint().
func (simple SimpleAction) SetStateHint(state_hint glib.Variant) {
	C.g_simple_action_set_state_hint(simple.native(), (*C.GVariant)(state_hint.Ptr))
}

// Object SimpleAsyncResult
type SimpleAsyncResult struct {
	gobject.Object
}

func (v SimpleAsyncResult) native() *C.GSimpleAsyncResult {
	return (*C.GSimpleAsyncResult)(v.Ptr)
}
func wrapSimpleAsyncResult(p *C.GSimpleAsyncResult) (v SimpleAsyncResult) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSimpleAsyncResult(p unsafe.Pointer) (v SimpleAsyncResult) {
	v.Ptr = p
	return
}
func (v SimpleAsyncResult) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSimpleAsyncResult(p unsafe.Pointer) interface{} {
	return WrapSimpleAsyncResult(p)
}
func (v SimpleAsyncResult) GetType() gobject.Type {
	return gobject.Type(C.g_simple_async_result_get_type())
}
func (v SimpleAsyncResult) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSimpleAsyncResult(unsafe.Pointer(ptr)), nil
	}
}
func (v SimpleAsyncResult) AsyncResult() AsyncResult {
	return WrapAsyncResult(v.Ptr)
}

// Object SimpleIOStream
type SimpleIOStream struct {
	IOStream
}

func (v SimpleIOStream) native() *C.GSimpleIOStream {
	return (*C.GSimpleIOStream)(v.Ptr)
}
func wrapSimpleIOStream(p *C.GSimpleIOStream) (v SimpleIOStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSimpleIOStream(p unsafe.Pointer) (v SimpleIOStream) {
	v.Ptr = p
	return
}
func (v SimpleIOStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSimpleIOStream(p unsafe.Pointer) interface{} {
	return WrapSimpleIOStream(p)
}
func (v SimpleIOStream) GetType() gobject.Type {
	return gobject.Type(C.g_simple_io_stream_get_type())
}
func (v SimpleIOStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSimpleIOStream(unsafe.Pointer(ptr)), nil
	}
}

// SimpleIOStreamNew is a wrapper around g_simple_io_stream_new().
func SimpleIOStreamNew(input_stream InputStream, output_stream OutputStream) IOStream {
	ret0 := C.g_simple_io_stream_new(input_stream.native(), output_stream.native())
	return wrapIOStream(ret0)
}

// Object SimplePermission
type SimplePermission struct {
	Permission
}

func (v SimplePermission) native() *C.GSimplePermission {
	return (*C.GSimplePermission)(v.Ptr)
}
func wrapSimplePermission(p *C.GSimplePermission) (v SimplePermission) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSimplePermission(p unsafe.Pointer) (v SimplePermission) {
	v.Ptr = p
	return
}
func (v SimplePermission) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSimplePermission(p unsafe.Pointer) interface{} {
	return WrapSimplePermission(p)
}
func (v SimplePermission) GetType() gobject.Type {
	return gobject.Type(C.g_simple_permission_get_type())
}
func (v SimplePermission) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSimplePermission(unsafe.Pointer(ptr)), nil
	}
}

// SimplePermissionNew is a wrapper around g_simple_permission_new().
func SimplePermissionNew(allowed bool) Permission {
	ret0 := C.g_simple_permission_new(C.gboolean(util.Bool2Int(allowed)) /*go:.util*/)
	return wrapPermission(ret0)
}

// Object SimpleProxyResolver
type SimpleProxyResolver struct {
	gobject.Object
}

func (v SimpleProxyResolver) native() *C.GSimpleProxyResolver {
	return (*C.GSimpleProxyResolver)(v.Ptr)
}
func wrapSimpleProxyResolver(p *C.GSimpleProxyResolver) (v SimpleProxyResolver) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSimpleProxyResolver(p unsafe.Pointer) (v SimpleProxyResolver) {
	v.Ptr = p
	return
}
func (v SimpleProxyResolver) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSimpleProxyResolver(p unsafe.Pointer) interface{} {
	return WrapSimpleProxyResolver(p)
}
func (v SimpleProxyResolver) GetType() gobject.Type {
	return gobject.Type(C.g_simple_proxy_resolver_get_type())
}
func (v SimpleProxyResolver) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSimpleProxyResolver(unsafe.Pointer(ptr)), nil
	}
}
func (v SimpleProxyResolver) ProxyResolver() ProxyResolver {
	return WrapProxyResolver(v.Ptr)
}

// SetDefaultProxy is a wrapper around g_simple_proxy_resolver_set_default_proxy().
func (resolver SimpleProxyResolver) SetDefaultProxy(default_proxy string) {
	default_proxy0 := (*C.gchar)(C.CString(default_proxy))
	C.g_simple_proxy_resolver_set_default_proxy(resolver.native(), default_proxy0)
	C.free(unsafe.Pointer(default_proxy0)) /*ch:<stdlib.h>*/
}

// SetUriProxy is a wrapper around g_simple_proxy_resolver_set_uri_proxy().
func (resolver SimpleProxyResolver) SetUriProxy(uri_scheme string, proxy string) {
	uri_scheme0 := (*C.gchar)(C.CString(uri_scheme))
	proxy0 := (*C.gchar)(C.CString(proxy))
	C.g_simple_proxy_resolver_set_uri_proxy(resolver.native(), uri_scheme0, proxy0)
	C.free(unsafe.Pointer(uri_scheme0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(proxy0))      /*ch:<stdlib.h>*/
}

// Object Socket
type Socket struct {
	gobject.Object
}

func (v Socket) native() *C.GSocket {
	return (*C.GSocket)(v.Ptr)
}
func wrapSocket(p *C.GSocket) (v Socket) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSocket(p unsafe.Pointer) (v Socket) {
	v.Ptr = p
	return
}
func (v Socket) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSocket(p unsafe.Pointer) interface{} {
	return WrapSocket(p)
}
func (v Socket) GetType() gobject.Type {
	return gobject.Type(C.g_socket_get_type())
}
func (v Socket) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSocket(unsafe.Pointer(ptr)), nil
	}
}
func (v Socket) DatagramBased() DatagramBased {
	return WrapDatagramBased(v.Ptr)
}
func (v Socket) Initable() Initable {
	return WrapInitable(v.Ptr)
}

// SocketNew is a wrapper around g_socket_new().
func SocketNew(family SocketFamily, type_ SocketType, protocol SocketProtocol) (Socket, error) {
	var err glib.Error
	ret0 := C.g_socket_new(C.GSocketFamily(family), C.GSocketType(type_), C.GSocketProtocol(protocol), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return Socket{}, err.GoValue()
	}
	return wrapSocket(ret0), nil
}

// SocketNewFromFd is a wrapper around g_socket_new_from_fd().
func SocketNewFromFd(fd int) (Socket, error) {
	var err glib.Error
	ret0 := C.g_socket_new_from_fd(C.gint(fd), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return Socket{}, err.GoValue()
	}
	return wrapSocket(ret0), nil
}

// Accept is a wrapper around g_socket_accept().
func (socket Socket) Accept(cancellable Cancellable) (Socket, error) {
	var err glib.Error
	ret0 := C.g_socket_accept(socket.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return Socket{}, err.GoValue()
	}
	return wrapSocket(ret0), nil
}

// Bind is a wrapper around g_socket_bind().
func (socket Socket) Bind(address SocketAddress, allow_reuse bool) (bool, error) {
	var err glib.Error
	ret0 := C.g_socket_bind(socket.native(), address.native(), C.gboolean(util.Bool2Int(allow_reuse)) /*go:.util*/, (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// CheckConnectResult is a wrapper around g_socket_check_connect_result().
func (socket Socket) CheckConnectResult() (bool, error) {
	var err glib.Error
	ret0 := C.g_socket_check_connect_result(socket.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Close is a wrapper around g_socket_close().
func (socket Socket) Close() (bool, error) {
	var err glib.Error
	ret0 := C.g_socket_close(socket.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// ConditionCheck is a wrapper around g_socket_condition_check().
func (socket Socket) ConditionCheck(condition /*gir:GLib*/ glib.IOCondition) /*gir:GLib*/ glib.IOCondition {
	ret0 := C.g_socket_condition_check(socket.native(), C.GIOCondition(condition))
	return /*gir:GLib*/ glib.IOCondition(ret0)
}

// ConditionTimedWait is a wrapper around g_socket_condition_timed_wait().
func (socket Socket) ConditionTimedWait(condition /*gir:GLib*/ glib.IOCondition, timeout int64, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_socket_condition_timed_wait(socket.native(), C.GIOCondition(condition), C.gint64(timeout), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// ConditionWait is a wrapper around g_socket_condition_wait().
func (socket Socket) ConditionWait(condition /*gir:GLib*/ glib.IOCondition, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_socket_condition_wait(socket.native(), C.GIOCondition(condition), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Connect is a wrapper around g_socket_connect().
func (socket Socket) Connect(address SocketAddress, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_socket_connect(socket.native(), address.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// ConnectionFactoryCreateConnection is a wrapper around g_socket_connection_factory_create_connection().
func (socket Socket) ConnectionFactoryCreateConnection() SocketConnection {
	ret0 := C.g_socket_connection_factory_create_connection(socket.native())
	return wrapSocketConnection(ret0)
}

// CreateSource is a wrapper around g_socket_create_source().
func (socket Socket) CreateSource(condition /*gir:GLib*/ glib.IOCondition, cancellable Cancellable) glib.Source {
	ret0 := C.g_socket_create_source(socket.native(), C.GIOCondition(condition), cancellable.native())
	return glib.WrapSource(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetAvailableBytes is a wrapper around g_socket_get_available_bytes().
func (socket Socket) GetAvailableBytes() int {
	ret0 := C.g_socket_get_available_bytes(socket.native())
	return int(ret0)
}

// GetBlocking is a wrapper around g_socket_get_blocking().
func (socket Socket) GetBlocking() bool {
	ret0 := C.g_socket_get_blocking(socket.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetBroadcast is a wrapper around g_socket_get_broadcast().
func (socket Socket) GetBroadcast() bool {
	ret0 := C.g_socket_get_broadcast(socket.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetCredentials is a wrapper around g_socket_get_credentials().
func (socket Socket) GetCredentials() (Credentials, error) {
	var err glib.Error
	ret0 := C.g_socket_get_credentials(socket.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return Credentials{}, err.GoValue()
	}
	return wrapCredentials(ret0), nil
}

// GetFamily is a wrapper around g_socket_get_family().
func (socket Socket) GetFamily() SocketFamily {
	ret0 := C.g_socket_get_family(socket.native())
	return SocketFamily(ret0)
}

// GetFd is a wrapper around g_socket_get_fd().
func (socket Socket) GetFd() int {
	ret0 := C.g_socket_get_fd(socket.native())
	return int(ret0)
}

// GetKeepalive is a wrapper around g_socket_get_keepalive().
func (socket Socket) GetKeepalive() bool {
	ret0 := C.g_socket_get_keepalive(socket.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetListenBacklog is a wrapper around g_socket_get_listen_backlog().
func (socket Socket) GetListenBacklog() int {
	ret0 := C.g_socket_get_listen_backlog(socket.native())
	return int(ret0)
}

// GetLocalAddress is a wrapper around g_socket_get_local_address().
func (socket Socket) GetLocalAddress() (SocketAddress, error) {
	var err glib.Error
	ret0 := C.g_socket_get_local_address(socket.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return SocketAddress{}, err.GoValue()
	}
	return wrapSocketAddress(ret0), nil
}

// GetMulticastLoopback is a wrapper around g_socket_get_multicast_loopback().
func (socket Socket) GetMulticastLoopback() bool {
	ret0 := C.g_socket_get_multicast_loopback(socket.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetMulticastTtl is a wrapper around g_socket_get_multicast_ttl().
func (socket Socket) GetMulticastTtl() uint {
	ret0 := C.g_socket_get_multicast_ttl(socket.native())
	return uint(ret0)
}

// GetOption is a wrapper around g_socket_get_option().
func (socket Socket) GetOption(level int, optname int) (bool, int, error) {
	var value0 C.gint
	var err glib.Error
	ret0 := C.g_socket_get_option(socket.native(), C.gint(level), C.gint(optname), &value0, (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, 0, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, int(value0), nil
}

// GetProtocol is a wrapper around g_socket_get_protocol().
func (socket Socket) GetProtocol() SocketProtocol {
	ret0 := C.g_socket_get_protocol(socket.native())
	return SocketProtocol(ret0)
}

// GetRemoteAddress is a wrapper around g_socket_get_remote_address().
func (socket Socket) GetRemoteAddress() (SocketAddress, error) {
	var err glib.Error
	ret0 := C.g_socket_get_remote_address(socket.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return SocketAddress{}, err.GoValue()
	}
	return wrapSocketAddress(ret0), nil
}

// GetSocketType is a wrapper around g_socket_get_socket_type().
func (socket Socket) GetSocketType() SocketType {
	ret0 := C.g_socket_get_socket_type(socket.native())
	return SocketType(ret0)
}

// GetTimeout is a wrapper around g_socket_get_timeout().
func (socket Socket) GetTimeout() uint {
	ret0 := C.g_socket_get_timeout(socket.native())
	return uint(ret0)
}

// GetTtl is a wrapper around g_socket_get_ttl().
func (socket Socket) GetTtl() uint {
	ret0 := C.g_socket_get_ttl(socket.native())
	return uint(ret0)
}

// IsClosed is a wrapper around g_socket_is_closed().
func (socket Socket) IsClosed() bool {
	ret0 := C.g_socket_is_closed(socket.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// IsConnected is a wrapper around g_socket_is_connected().
func (socket Socket) IsConnected() bool {
	ret0 := C.g_socket_is_connected(socket.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// JoinMulticastGroup is a wrapper around g_socket_join_multicast_group().
func (socket Socket) JoinMulticastGroup(group InetAddress, source_specific bool, iface string) (bool, error) {
	iface0 := (*C.gchar)(C.CString(iface))
	var err glib.Error
	ret0 := C.g_socket_join_multicast_group(socket.native(), group.native(), C.gboolean(util.Bool2Int(source_specific)) /*go:.util*/, iface0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(iface0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// LeaveMulticastGroup is a wrapper around g_socket_leave_multicast_group().
func (socket Socket) LeaveMulticastGroup(group InetAddress, source_specific bool, iface string) (bool, error) {
	iface0 := (*C.gchar)(C.CString(iface))
	var err glib.Error
	ret0 := C.g_socket_leave_multicast_group(socket.native(), group.native(), C.gboolean(util.Bool2Int(source_specific)) /*go:.util*/, iface0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(iface0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Listen is a wrapper around g_socket_listen().
func (socket Socket) Listen() (bool, error) {
	var err glib.Error
	ret0 := C.g_socket_listen(socket.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Receive is a wrapper around g_socket_receive().
func (socket Socket) Receive(buffer []byte, cancellable Cancellable) (int, error) {
	buffer0 := make([]C.gchar, len(buffer))
	for idx, elemG := range buffer {
		buffer0[idx] = C.gchar(elemG)
	}
	var buffer0Ptr *C.gchar
	if len(buffer0) > 0 {
		buffer0Ptr = &buffer0[0]
	}
	var err glib.Error
	ret0 := C.g_socket_receive(socket.native(), buffer0Ptr, C.gsize(len(buffer)), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// ReceiveFrom is a wrapper around g_socket_receive_from().
func (socket Socket) ReceiveFrom(buffer []byte, cancellable Cancellable) (int, SocketAddress, error) {
	var address0 *C.GSocketAddress
	buffer0 := make([]C.gchar, len(buffer))
	for idx, elemG := range buffer {
		buffer0[idx] = C.gchar(elemG)
	}
	var buffer0Ptr *C.gchar
	if len(buffer0) > 0 {
		buffer0Ptr = &buffer0[0]
	}
	var err glib.Error
	ret0 := C.g_socket_receive_from(socket.native(), &address0, buffer0Ptr, C.gsize(len(buffer)), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, SocketAddress{}, err.GoValue()
	}
	return int(ret0), wrapSocketAddress(address0), nil
}

// ReceiveWithBlocking is a wrapper around g_socket_receive_with_blocking().
func (socket Socket) ReceiveWithBlocking(buffer []byte, blocking bool, cancellable Cancellable) (int, error) {
	buffer0 := make([]C.gchar, len(buffer))
	for idx, elemG := range buffer {
		buffer0[idx] = C.gchar(elemG)
	}
	var buffer0Ptr *C.gchar
	if len(buffer0) > 0 {
		buffer0Ptr = &buffer0[0]
	}
	var err glib.Error
	ret0 := C.g_socket_receive_with_blocking(socket.native(), buffer0Ptr, C.gsize(len(buffer)), C.gboolean(util.Bool2Int(blocking)) /*go:.util*/, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// Send is a wrapper around g_socket_send().
func (socket Socket) Send(buffer []byte, cancellable Cancellable) (int, error) {
	buffer0 := make([]C.gchar, len(buffer))
	for idx, elemG := range buffer {
		buffer0[idx] = C.gchar(elemG)
	}
	var buffer0Ptr *C.gchar
	if len(buffer0) > 0 {
		buffer0Ptr = &buffer0[0]
	}
	var err glib.Error
	ret0 := C.g_socket_send(socket.native(), buffer0Ptr, C.gsize(len(buffer)), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// SendTo is a wrapper around g_socket_send_to().
func (socket Socket) SendTo(address SocketAddress, buffer []byte, cancellable Cancellable) (int, error) {
	buffer0 := make([]C.gchar, len(buffer))
	for idx, elemG := range buffer {
		buffer0[idx] = C.gchar(elemG)
	}
	var buffer0Ptr *C.gchar
	if len(buffer0) > 0 {
		buffer0Ptr = &buffer0[0]
	}
	var err glib.Error
	ret0 := C.g_socket_send_to(socket.native(), address.native(), buffer0Ptr, C.gsize(len(buffer)), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// SendWithBlocking is a wrapper around g_socket_send_with_blocking().
func (socket Socket) SendWithBlocking(buffer []byte, blocking bool, cancellable Cancellable) (int, error) {
	buffer0 := make([]C.gchar, len(buffer))
	for idx, elemG := range buffer {
		buffer0[idx] = C.gchar(elemG)
	}
	var buffer0Ptr *C.gchar
	if len(buffer0) > 0 {
		buffer0Ptr = &buffer0[0]
	}
	var err glib.Error
	ret0 := C.g_socket_send_with_blocking(socket.native(), buffer0Ptr, C.gsize(len(buffer)), C.gboolean(util.Bool2Int(blocking)) /*go:.util*/, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// SetBlocking is a wrapper around g_socket_set_blocking().
func (socket Socket) SetBlocking(blocking bool) {
	C.g_socket_set_blocking(socket.native(), C.gboolean(util.Bool2Int(blocking)) /*go:.util*/)
}

// SetBroadcast is a wrapper around g_socket_set_broadcast().
func (socket Socket) SetBroadcast(broadcast bool) {
	C.g_socket_set_broadcast(socket.native(), C.gboolean(util.Bool2Int(broadcast)) /*go:.util*/)
}

// SetKeepalive is a wrapper around g_socket_set_keepalive().
func (socket Socket) SetKeepalive(keepalive bool) {
	C.g_socket_set_keepalive(socket.native(), C.gboolean(util.Bool2Int(keepalive)) /*go:.util*/)
}

// SetListenBacklog is a wrapper around g_socket_set_listen_backlog().
func (socket Socket) SetListenBacklog(backlog int) {
	C.g_socket_set_listen_backlog(socket.native(), C.gint(backlog))
}

// SetMulticastLoopback is a wrapper around g_socket_set_multicast_loopback().
func (socket Socket) SetMulticastLoopback(loopback bool) {
	C.g_socket_set_multicast_loopback(socket.native(), C.gboolean(util.Bool2Int(loopback)) /*go:.util*/)
}

// SetMulticastTtl is a wrapper around g_socket_set_multicast_ttl().
func (socket Socket) SetMulticastTtl(ttl uint) {
	C.g_socket_set_multicast_ttl(socket.native(), C.guint(ttl))
}

// SetOption is a wrapper around g_socket_set_option().
func (socket Socket) SetOption(level int, optname int, value int) (bool, error) {
	var err glib.Error
	ret0 := C.g_socket_set_option(socket.native(), C.gint(level), C.gint(optname), C.gint(value), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SetTimeout is a wrapper around g_socket_set_timeout().
func (socket Socket) SetTimeout(timeout uint) {
	C.g_socket_set_timeout(socket.native(), C.guint(timeout))
}

// SetTtl is a wrapper around g_socket_set_ttl().
func (socket Socket) SetTtl(ttl uint) {
	C.g_socket_set_ttl(socket.native(), C.guint(ttl))
}

// Shutdown is a wrapper around g_socket_shutdown().
func (socket Socket) Shutdown(shutdown_read bool, shutdown_write bool) (bool, error) {
	var err glib.Error
	ret0 := C.g_socket_shutdown(socket.native(), C.gboolean(util.Bool2Int(shutdown_read)) /*go:.util*/, C.gboolean(util.Bool2Int(shutdown_write)) /*go:.util*/, (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SpeaksIpv4 is a wrapper around g_socket_speaks_ipv4().
func (socket Socket) SpeaksIpv4() bool {
	ret0 := C.g_socket_speaks_ipv4(socket.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object SocketConnection
type SocketConnection struct {
	IOStream
}

func (v SocketConnection) native() *C.GSocketConnection {
	return (*C.GSocketConnection)(v.Ptr)
}
func wrapSocketConnection(p *C.GSocketConnection) (v SocketConnection) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSocketConnection(p unsafe.Pointer) (v SocketConnection) {
	v.Ptr = p
	return
}
func (v SocketConnection) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSocketConnection(p unsafe.Pointer) interface{} {
	return WrapSocketConnection(p)
}
func (v SocketConnection) GetType() gobject.Type {
	return gobject.Type(C.g_socket_connection_get_type())
}
func (v SocketConnection) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSocketConnection(unsafe.Pointer(ptr)), nil
	}
}

// Connect is a wrapper around g_socket_connection_connect().
func (connection SocketConnection) Connect(address SocketAddress, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_socket_connection_connect(connection.native(), address.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// ConnectAsync is a wrapper around g_socket_connection_connect_async().
func (connection SocketConnection) ConnectAsync(address SocketAddress, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_socket_connection_connect_async(connection.native(), address.native(), cancellable.native(), callback0)
}

// ConnectFinish is a wrapper around g_socket_connection_connect_finish().
func (connection SocketConnection) ConnectFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_socket_connection_connect_finish(connection.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// GetLocalAddress is a wrapper around g_socket_connection_get_local_address().
func (connection SocketConnection) GetLocalAddress() (SocketAddress, error) {
	var err glib.Error
	ret0 := C.g_socket_connection_get_local_address(connection.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return SocketAddress{}, err.GoValue()
	}
	return wrapSocketAddress(ret0), nil
}

// GetRemoteAddress is a wrapper around g_socket_connection_get_remote_address().
func (connection SocketConnection) GetRemoteAddress() (SocketAddress, error) {
	var err glib.Error
	ret0 := C.g_socket_connection_get_remote_address(connection.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return SocketAddress{}, err.GoValue()
	}
	return wrapSocketAddress(ret0), nil
}

// GetSocket is a wrapper around g_socket_connection_get_socket().
func (connection SocketConnection) GetSocket() Socket {
	ret0 := C.g_socket_connection_get_socket(connection.native())
	return wrapSocket(ret0)
}

// IsConnected is a wrapper around g_socket_connection_is_connected().
func (connection SocketConnection) IsConnected() bool {
	ret0 := C.g_socket_connection_is_connected(connection.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object SocketClient
type SocketClient struct {
	gobject.Object
}

func (v SocketClient) native() *C.GSocketClient {
	return (*C.GSocketClient)(v.Ptr)
}
func wrapSocketClient(p *C.GSocketClient) (v SocketClient) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSocketClient(p unsafe.Pointer) (v SocketClient) {
	v.Ptr = p
	return
}
func (v SocketClient) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSocketClient(p unsafe.Pointer) interface{} {
	return WrapSocketClient(p)
}
func (v SocketClient) GetType() gobject.Type {
	return gobject.Type(C.g_socket_client_get_type())
}
func (v SocketClient) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSocketClient(unsafe.Pointer(ptr)), nil
	}
}

// SocketClientNew is a wrapper around g_socket_client_new().
func SocketClientNew() SocketClient {
	ret0 := C.g_socket_client_new()
	return wrapSocketClient(ret0)
}

// AddApplicationProxy is a wrapper around g_socket_client_add_application_proxy().
func (client SocketClient) AddApplicationProxy(protocol string) {
	protocol0 := (*C.gchar)(C.CString(protocol))
	C.g_socket_client_add_application_proxy(client.native(), protocol0)
	C.free(unsafe.Pointer(protocol0)) /*ch:<stdlib.h>*/
}

// Connect is a wrapper around g_socket_client_connect().
func (client SocketClient) Connect(connectable SocketConnectable, cancellable Cancellable) (SocketConnection, error) {
	var err glib.Error
	ret0 := C.g_socket_client_connect(client.native(), connectable.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return SocketConnection{}, err.GoValue()
	}
	return wrapSocketConnection(ret0), nil
}

// ConnectAsync is a wrapper around g_socket_client_connect_async().
func (client SocketClient) ConnectAsync(connectable SocketConnectable, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_socket_client_connect_async(client.native(), connectable.native(), cancellable.native(), callback0)
}

// ConnectFinish is a wrapper around g_socket_client_connect_finish().
func (client SocketClient) ConnectFinish(result AsyncResult) (SocketConnection, error) {
	var err glib.Error
	ret0 := C.g_socket_client_connect_finish(client.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return SocketConnection{}, err.GoValue()
	}
	return wrapSocketConnection(ret0), nil
}

// ConnectToHost is a wrapper around g_socket_client_connect_to_host().
func (client SocketClient) ConnectToHost(host_and_port string, default_port uint16, cancellable Cancellable) (SocketConnection, error) {
	host_and_port0 := (*C.gchar)(C.CString(host_and_port))
	var err glib.Error
	ret0 := C.g_socket_client_connect_to_host(client.native(), host_and_port0, C.guint16(default_port), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(host_and_port0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return SocketConnection{}, err.GoValue()
	}
	return wrapSocketConnection(ret0), nil
}

// ConnectToHostAsync is a wrapper around g_socket_client_connect_to_host_async().
func (client SocketClient) ConnectToHostAsync(host_and_port string, default_port uint16, cancellable Cancellable, callback AsyncReadyCallback) {
	host_and_port0 := (*C.gchar)(C.CString(host_and_port))
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_socket_client_connect_to_host_async(client.native(), host_and_port0, C.guint16(default_port), cancellable.native(), callback0)
	C.free(unsafe.Pointer(host_and_port0)) /*ch:<stdlib.h>*/
}

// ConnectToHostFinish is a wrapper around g_socket_client_connect_to_host_finish().
func (client SocketClient) ConnectToHostFinish(result AsyncResult) (SocketConnection, error) {
	var err glib.Error
	ret0 := C.g_socket_client_connect_to_host_finish(client.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return SocketConnection{}, err.GoValue()
	}
	return wrapSocketConnection(ret0), nil
}

// ConnectToService is a wrapper around g_socket_client_connect_to_service().
func (client SocketClient) ConnectToService(domain string, service string, cancellable Cancellable) (SocketConnection, error) {
	domain0 := (*C.gchar)(C.CString(domain))
	service0 := (*C.gchar)(C.CString(service))
	var err glib.Error
	ret0 := C.g_socket_client_connect_to_service(client.native(), domain0, service0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(domain0))  /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(service0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return SocketConnection{}, err.GoValue()
	}
	return wrapSocketConnection(ret0), nil
}

// ConnectToServiceAsync is a wrapper around g_socket_client_connect_to_service_async().
func (client SocketClient) ConnectToServiceAsync(domain string, service string, cancellable Cancellable, callback AsyncReadyCallback) {
	domain0 := (*C.gchar)(C.CString(domain))
	service0 := (*C.gchar)(C.CString(service))
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_socket_client_connect_to_service_async(client.native(), domain0, service0, cancellable.native(), callback0)
	C.free(unsafe.Pointer(domain0))  /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(service0)) /*ch:<stdlib.h>*/
}

// ConnectToServiceFinish is a wrapper around g_socket_client_connect_to_service_finish().
func (client SocketClient) ConnectToServiceFinish(result AsyncResult) (SocketConnection, error) {
	var err glib.Error
	ret0 := C.g_socket_client_connect_to_service_finish(client.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return SocketConnection{}, err.GoValue()
	}
	return wrapSocketConnection(ret0), nil
}

// ConnectToUri is a wrapper around g_socket_client_connect_to_uri().
func (client SocketClient) ConnectToUri(uri string, default_port uint16, cancellable Cancellable) (SocketConnection, error) {
	uri0 := (*C.gchar)(C.CString(uri))
	var err glib.Error
	ret0 := C.g_socket_client_connect_to_uri(client.native(), uri0, C.guint16(default_port), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(uri0)) /*ch:<stdlib.h>*/
	if err.Ptr != nil {
		defer err.Free()
		return SocketConnection{}, err.GoValue()
	}
	return wrapSocketConnection(ret0), nil
}

// ConnectToUriAsync is a wrapper around g_socket_client_connect_to_uri_async().
func (client SocketClient) ConnectToUriAsync(uri string, default_port uint16, cancellable Cancellable, callback AsyncReadyCallback) {
	uri0 := (*C.gchar)(C.CString(uri))
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_socket_client_connect_to_uri_async(client.native(), uri0, C.guint16(default_port), cancellable.native(), callback0)
	C.free(unsafe.Pointer(uri0)) /*ch:<stdlib.h>*/
}

// ConnectToUriFinish is a wrapper around g_socket_client_connect_to_uri_finish().
func (client SocketClient) ConnectToUriFinish(result AsyncResult) (SocketConnection, error) {
	var err glib.Error
	ret0 := C.g_socket_client_connect_to_uri_finish(client.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return SocketConnection{}, err.GoValue()
	}
	return wrapSocketConnection(ret0), nil
}

// GetEnableProxy is a wrapper around g_socket_client_get_enable_proxy().
func (client SocketClient) GetEnableProxy() bool {
	ret0 := C.g_socket_client_get_enable_proxy(client.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetFamily is a wrapper around g_socket_client_get_family().
func (client SocketClient) GetFamily() SocketFamily {
	ret0 := C.g_socket_client_get_family(client.native())
	return SocketFamily(ret0)
}

// GetLocalAddress is a wrapper around g_socket_client_get_local_address().
func (client SocketClient) GetLocalAddress() SocketAddress {
	ret0 := C.g_socket_client_get_local_address(client.native())
	return wrapSocketAddress(ret0)
}

// GetProtocol is a wrapper around g_socket_client_get_protocol().
func (client SocketClient) GetProtocol() SocketProtocol {
	ret0 := C.g_socket_client_get_protocol(client.native())
	return SocketProtocol(ret0)
}

// GetProxyResolver is a wrapper around g_socket_client_get_proxy_resolver().
func (client SocketClient) GetProxyResolver() ProxyResolver {
	ret0 := C.g_socket_client_get_proxy_resolver(client.native())
	return wrapProxyResolver(ret0)
}

// GetSocketType is a wrapper around g_socket_client_get_socket_type().
func (client SocketClient) GetSocketType() SocketType {
	ret0 := C.g_socket_client_get_socket_type(client.native())
	return SocketType(ret0)
}

// GetTimeout is a wrapper around g_socket_client_get_timeout().
func (client SocketClient) GetTimeout() uint {
	ret0 := C.g_socket_client_get_timeout(client.native())
	return uint(ret0)
}

// GetTls is a wrapper around g_socket_client_get_tls().
func (client SocketClient) GetTls() bool {
	ret0 := C.g_socket_client_get_tls(client.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetTlsValidationFlags is a wrapper around g_socket_client_get_tls_validation_flags().
func (client SocketClient) GetTlsValidationFlags() TlsCertificateFlags {
	ret0 := C.g_socket_client_get_tls_validation_flags(client.native())
	return TlsCertificateFlags(ret0)
}

// SetEnableProxy is a wrapper around g_socket_client_set_enable_proxy().
func (client SocketClient) SetEnableProxy(enable bool) {
	C.g_socket_client_set_enable_proxy(client.native(), C.gboolean(util.Bool2Int(enable)) /*go:.util*/)
}

// SetFamily is a wrapper around g_socket_client_set_family().
func (client SocketClient) SetFamily(family SocketFamily) {
	C.g_socket_client_set_family(client.native(), C.GSocketFamily(family))
}

// SetLocalAddress is a wrapper around g_socket_client_set_local_address().
func (client SocketClient) SetLocalAddress(address SocketAddress) {
	C.g_socket_client_set_local_address(client.native(), address.native())
}

// SetProtocol is a wrapper around g_socket_client_set_protocol().
func (client SocketClient) SetProtocol(protocol SocketProtocol) {
	C.g_socket_client_set_protocol(client.native(), C.GSocketProtocol(protocol))
}

// SetProxyResolver is a wrapper around g_socket_client_set_proxy_resolver().
func (client SocketClient) SetProxyResolver(proxy_resolver ProxyResolver) {
	C.g_socket_client_set_proxy_resolver(client.native(), proxy_resolver.native())
}

// SetSocketType is a wrapper around g_socket_client_set_socket_type().
func (client SocketClient) SetSocketType(type_ SocketType) {
	C.g_socket_client_set_socket_type(client.native(), C.GSocketType(type_))
}

// SetTimeout is a wrapper around g_socket_client_set_timeout().
func (client SocketClient) SetTimeout(timeout uint) {
	C.g_socket_client_set_timeout(client.native(), C.guint(timeout))
}

// SetTls is a wrapper around g_socket_client_set_tls().
func (client SocketClient) SetTls(tls bool) {
	C.g_socket_client_set_tls(client.native(), C.gboolean(util.Bool2Int(tls)) /*go:.util*/)
}

// SetTlsValidationFlags is a wrapper around g_socket_client_set_tls_validation_flags().
func (client SocketClient) SetTlsValidationFlags(flags TlsCertificateFlags) {
	C.g_socket_client_set_tls_validation_flags(client.native(), C.GTlsCertificateFlags(flags))
}

// Object SocketControlMessage
type SocketControlMessage struct {
	gobject.Object
}

func (v SocketControlMessage) native() *C.GSocketControlMessage {
	return (*C.GSocketControlMessage)(v.Ptr)
}
func wrapSocketControlMessage(p *C.GSocketControlMessage) (v SocketControlMessage) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSocketControlMessage(p unsafe.Pointer) (v SocketControlMessage) {
	v.Ptr = p
	return
}
func (v SocketControlMessage) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSocketControlMessage(p unsafe.Pointer) interface{} {
	return WrapSocketControlMessage(p)
}
func (v SocketControlMessage) GetType() gobject.Type {
	return gobject.Type(C.g_socket_control_message_get_type())
}
func (v SocketControlMessage) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSocketControlMessage(unsafe.Pointer(ptr)), nil
	}
}

// GetLevel is a wrapper around g_socket_control_message_get_level().
func (message SocketControlMessage) GetLevel() int {
	ret0 := C.g_socket_control_message_get_level(message.native())
	return int(ret0)
}

// GetMsgType is a wrapper around g_socket_control_message_get_msg_type().
func (message SocketControlMessage) GetMsgType() int {
	ret0 := C.g_socket_control_message_get_msg_type(message.native())
	return int(ret0)
}

// GetSize is a wrapper around g_socket_control_message_get_size().
func (message SocketControlMessage) GetSize() uint {
	ret0 := C.g_socket_control_message_get_size(message.native())
	return uint(ret0)
}

// Serialize is a wrapper around g_socket_control_message_serialize().
func (message SocketControlMessage) Serialize(data unsafe.Pointer) {
	C.g_socket_control_message_serialize(message.native(), C.gpointer(data))
}

// Object SocketListener
type SocketListener struct {
	gobject.Object
}

func (v SocketListener) native() *C.GSocketListener {
	return (*C.GSocketListener)(v.Ptr)
}
func wrapSocketListener(p *C.GSocketListener) (v SocketListener) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSocketListener(p unsafe.Pointer) (v SocketListener) {
	v.Ptr = p
	return
}
func (v SocketListener) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSocketListener(p unsafe.Pointer) interface{} {
	return WrapSocketListener(p)
}
func (v SocketListener) GetType() gobject.Type {
	return gobject.Type(C.g_socket_listener_get_type())
}
func (v SocketListener) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSocketListener(unsafe.Pointer(ptr)), nil
	}
}

// SocketListenerNew is a wrapper around g_socket_listener_new().
func SocketListenerNew() SocketListener {
	ret0 := C.g_socket_listener_new()
	return wrapSocketListener(ret0)
}

// Accept is a wrapper around g_socket_listener_accept().
func (listener SocketListener) Accept(cancellable Cancellable) (SocketConnection, gobject.Object, error) {
	var source_object0 *C.GObject
	var err glib.Error
	ret0 := C.g_socket_listener_accept(listener.native(), &source_object0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return SocketConnection{}, gobject.Object{}, err.GoValue()
	}
	return wrapSocketConnection(ret0), gobject.WrapObject(unsafe.Pointer(source_object0)) /*gir:GObject*/, nil
}

// AcceptAsync is a wrapper around g_socket_listener_accept_async().
func (listener SocketListener) AcceptAsync(cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_socket_listener_accept_async(listener.native(), cancellable.native(), callback0)
}

// AcceptFinish is a wrapper around g_socket_listener_accept_finish().
func (listener SocketListener) AcceptFinish(result AsyncResult) (SocketConnection, gobject.Object, error) {
	var source_object0 *C.GObject
	var err glib.Error
	ret0 := C.g_socket_listener_accept_finish(listener.native(), result.native(), &source_object0, (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return SocketConnection{}, gobject.Object{}, err.GoValue()
	}
	return wrapSocketConnection(ret0), gobject.WrapObject(unsafe.Pointer(source_object0)) /*gir:GObject*/, nil
}

// AcceptSocket is a wrapper around g_socket_listener_accept_socket().
func (listener SocketListener) AcceptSocket(cancellable Cancellable) (Socket, gobject.Object, error) {
	var source_object0 *C.GObject
	var err glib.Error
	ret0 := C.g_socket_listener_accept_socket(listener.native(), &source_object0, cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return Socket{}, gobject.Object{}, err.GoValue()
	}
	return wrapSocket(ret0), gobject.WrapObject(unsafe.Pointer(source_object0)) /*gir:GObject*/, nil
}

// AcceptSocketAsync is a wrapper around g_socket_listener_accept_socket_async().
func (listener SocketListener) AcceptSocketAsync(cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_socket_listener_accept_socket_async(listener.native(), cancellable.native(), callback0)
}

// AcceptSocketFinish is a wrapper around g_socket_listener_accept_socket_finish().
func (listener SocketListener) AcceptSocketFinish(result AsyncResult) (Socket, gobject.Object, error) {
	var source_object0 *C.GObject
	var err glib.Error
	ret0 := C.g_socket_listener_accept_socket_finish(listener.native(), result.native(), &source_object0, (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return Socket{}, gobject.Object{}, err.GoValue()
	}
	return wrapSocket(ret0), gobject.WrapObject(unsafe.Pointer(source_object0)) /*gir:GObject*/, nil
}

// AddAddress is a wrapper around g_socket_listener_add_address().
func (listener SocketListener) AddAddress(address SocketAddress, type_ SocketType, protocol SocketProtocol, source_object gobject.Object) (bool, SocketAddress, error) {
	var effective_address0 *C.GSocketAddress
	var err glib.Error
	ret0 := C.g_socket_listener_add_address(listener.native(), address.native(), C.GSocketType(type_), C.GSocketProtocol(protocol), (*C.GObject)(source_object.Ptr), &effective_address0, (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, SocketAddress{}, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, wrapSocketAddress(effective_address0), nil
}

// AddAnyInetPort is a wrapper around g_socket_listener_add_any_inet_port().
func (listener SocketListener) AddAnyInetPort(source_object gobject.Object) (uint16, error) {
	var err glib.Error
	ret0 := C.g_socket_listener_add_any_inet_port(listener.native(), (*C.GObject)(source_object.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return uint16(ret0), nil
}

// AddInetPort is a wrapper around g_socket_listener_add_inet_port().
func (listener SocketListener) AddInetPort(port uint16, source_object gobject.Object) (bool, error) {
	var err glib.Error
	ret0 := C.g_socket_listener_add_inet_port(listener.native(), C.guint16(port), (*C.GObject)(source_object.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// AddSocket is a wrapper around g_socket_listener_add_socket().
func (listener SocketListener) AddSocket(socket Socket, source_object gobject.Object) (bool, error) {
	var err glib.Error
	ret0 := C.g_socket_listener_add_socket(listener.native(), socket.native(), (*C.GObject)(source_object.Ptr), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Close is a wrapper around g_socket_listener_close().
func (listener SocketListener) Close() {
	C.g_socket_listener_close(listener.native())
}

// SetBacklog is a wrapper around g_socket_listener_set_backlog().
func (listener SocketListener) SetBacklog(listen_backlog int) {
	C.g_socket_listener_set_backlog(listener.native(), C.int(listen_backlog))
}

// Object SocketService
type SocketService struct {
	SocketListener
}

func (v SocketService) native() *C.GSocketService {
	return (*C.GSocketService)(v.Ptr)
}
func wrapSocketService(p *C.GSocketService) (v SocketService) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSocketService(p unsafe.Pointer) (v SocketService) {
	v.Ptr = p
	return
}
func (v SocketService) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSocketService(p unsafe.Pointer) interface{} {
	return WrapSocketService(p)
}
func (v SocketService) GetType() gobject.Type {
	return gobject.Type(C.g_socket_service_get_type())
}
func (v SocketService) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSocketService(unsafe.Pointer(ptr)), nil
	}
}

// SocketServiceNew is a wrapper around g_socket_service_new().
func SocketServiceNew() SocketService {
	ret0 := C.g_socket_service_new()
	return wrapSocketService(ret0)
}

// IsActive is a wrapper around g_socket_service_is_active().
func (service SocketService) IsActive() bool {
	ret0 := C.g_socket_service_is_active(service.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Start is a wrapper around g_socket_service_start().
func (service SocketService) Start() {
	C.g_socket_service_start(service.native())
}

// Stop is a wrapper around g_socket_service_stop().
func (service SocketService) Stop() {
	C.g_socket_service_stop(service.native())
}

// Object Subprocess
type Subprocess struct {
	gobject.Object
}

func (v Subprocess) native() *C.GSubprocess {
	return (*C.GSubprocess)(v.Ptr)
}
func wrapSubprocess(p *C.GSubprocess) (v Subprocess) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSubprocess(p unsafe.Pointer) (v Subprocess) {
	v.Ptr = p
	return
}
func (v Subprocess) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSubprocess(p unsafe.Pointer) interface{} {
	return WrapSubprocess(p)
}
func (v Subprocess) GetType() gobject.Type {
	return gobject.Type(C.g_subprocess_get_type())
}
func (v Subprocess) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSubprocess(unsafe.Pointer(ptr)), nil
	}
}
func (v Subprocess) Initable() Initable {
	return WrapInitable(v.Ptr)
}

// SubprocessNewv is a wrapper around g_subprocess_newv().
func SubprocessNewv(argv []string, flags SubprocessFlags) (Subprocess, error) {
	argv0 := make([]*C.gchar, len(argv))
	for idx, elemG := range argv {
		elem := (*C.gchar)(C.CString(elemG))
		argv0[idx] = elem
	}
	var argv0Ptr **C.gchar
	if len(argv0) > 0 {
		argv0Ptr = &argv0[0]
	}
	var err glib.Error
	ret0 := C.g_subprocess_newv(argv0Ptr, C.GSubprocessFlags(flags), (**C.GError)(unsafe.Pointer(&err)))
	for _, elem := range argv0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	if err.Ptr != nil {
		defer err.Free()
		return Subprocess{}, err.GoValue()
	}
	return wrapSubprocess(ret0), nil
}

// Communicate is a wrapper around g_subprocess_communicate().
func (subprocess Subprocess) Communicate(stdin_buf glib.Bytes, cancellable Cancellable) (bool, glib.Bytes, glib.Bytes, error) {
	var stdout_buf0 *C.GBytes
	var stderr_buf0 *C.GBytes
	var err glib.Error
	ret0 := C.g_subprocess_communicate(subprocess.native(), (*C.GBytes)(stdin_buf.Ptr), cancellable.native(), &stdout_buf0, &stderr_buf0, (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, glib.Bytes{}, glib.Bytes{}, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, glib.WrapBytes(unsafe.Pointer(stdout_buf0)) /*gir:GLib*/, glib.WrapBytes(unsafe.Pointer(stderr_buf0)) /*gir:GLib*/, nil
}

// CommunicateAsync is a wrapper around g_subprocess_communicate_async().
func (subprocess Subprocess) CommunicateAsync(stdin_buf glib.Bytes, cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_subprocess_communicate_async(subprocess.native(), (*C.GBytes)(stdin_buf.Ptr), cancellable.native(), callback0)
}

// CommunicateFinish is a wrapper around g_subprocess_communicate_finish().
func (subprocess Subprocess) CommunicateFinish(result AsyncResult) (bool, glib.Bytes, glib.Bytes, error) {
	var stdout_buf0 *C.GBytes
	var stderr_buf0 *C.GBytes
	var err glib.Error
	ret0 := C.g_subprocess_communicate_finish(subprocess.native(), result.native(), &stdout_buf0, &stderr_buf0, (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, glib.Bytes{}, glib.Bytes{}, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, glib.WrapBytes(unsafe.Pointer(stdout_buf0)) /*gir:GLib*/, glib.WrapBytes(unsafe.Pointer(stderr_buf0)) /*gir:GLib*/, nil
}

// CommunicateUtf8 is a wrapper around g_subprocess_communicate_utf8().
func (subprocess Subprocess) CommunicateUtf8(stdin_buf string, cancellable Cancellable) (bool, string, string, error) {
	stdin_buf0 := C.CString(stdin_buf)
	var stdout_buf0 *C.char
	var stderr_buf0 *C.char
	var err glib.Error
	ret0 := C.g_subprocess_communicate_utf8(subprocess.native(), stdin_buf0, cancellable.native(), &stdout_buf0, &stderr_buf0, (**C.GError)(unsafe.Pointer(&err)))
	C.free(unsafe.Pointer(stdin_buf0)) /*ch:<stdlib.h>*/
	stdout_buf := C.GoString(stdout_buf0)
	defer C.g_free(C.gpointer(stdout_buf0))
	stderr_buf := C.GoString(stderr_buf0)
	defer C.g_free(C.gpointer(stderr_buf0))
	if err.Ptr != nil {
		defer err.Free()
		return false, "", "", err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, stdout_buf, stderr_buf, nil
}

// CommunicateUtf8Async is a wrapper around g_subprocess_communicate_utf8_async().
func (subprocess Subprocess) CommunicateUtf8Async(stdin_buf string, cancellable Cancellable, callback AsyncReadyCallback) {
	stdin_buf0 := C.CString(stdin_buf)
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_subprocess_communicate_utf8_async(subprocess.native(), stdin_buf0, cancellable.native(), callback0)
	C.free(unsafe.Pointer(stdin_buf0)) /*ch:<stdlib.h>*/
}

// CommunicateUtf8Finish is a wrapper around g_subprocess_communicate_utf8_finish().
func (subprocess Subprocess) CommunicateUtf8Finish(result AsyncResult) (bool, string, string, error) {
	var stdout_buf0 *C.char
	var stderr_buf0 *C.char
	var err glib.Error
	ret0 := C.g_subprocess_communicate_utf8_finish(subprocess.native(), result.native(), &stdout_buf0, &stderr_buf0, (**C.GError)(unsafe.Pointer(&err)))
	stdout_buf := C.GoString(stdout_buf0)
	defer C.g_free(C.gpointer(stdout_buf0))
	stderr_buf := C.GoString(stderr_buf0)
	defer C.g_free(C.gpointer(stderr_buf0))
	if err.Ptr != nil {
		defer err.Free()
		return false, "", "", err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, stdout_buf, stderr_buf, nil
}

// ForceExit is a wrapper around g_subprocess_force_exit().
func (subprocess Subprocess) ForceExit() {
	C.g_subprocess_force_exit(subprocess.native())
}

// GetExitStatus is a wrapper around g_subprocess_get_exit_status().
func (subprocess Subprocess) GetExitStatus() int {
	ret0 := C.g_subprocess_get_exit_status(subprocess.native())
	return int(ret0)
}

// GetIdentifier is a wrapper around g_subprocess_get_identifier().
func (subprocess Subprocess) GetIdentifier() string {
	ret0 := C.g_subprocess_get_identifier(subprocess.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetIfExited is a wrapper around g_subprocess_get_if_exited().
func (subprocess Subprocess) GetIfExited() bool {
	ret0 := C.g_subprocess_get_if_exited(subprocess.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetIfSignaled is a wrapper around g_subprocess_get_if_signaled().
func (subprocess Subprocess) GetIfSignaled() bool {
	ret0 := C.g_subprocess_get_if_signaled(subprocess.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetStatus is a wrapper around g_subprocess_get_status().
func (subprocess Subprocess) GetStatus() int {
	ret0 := C.g_subprocess_get_status(subprocess.native())
	return int(ret0)
}

// GetStderrPipe is a wrapper around g_subprocess_get_stderr_pipe().
func (subprocess Subprocess) GetStderrPipe() InputStream {
	ret0 := C.g_subprocess_get_stderr_pipe(subprocess.native())
	return wrapInputStream(ret0)
}

// GetStdinPipe is a wrapper around g_subprocess_get_stdin_pipe().
func (subprocess Subprocess) GetStdinPipe() OutputStream {
	ret0 := C.g_subprocess_get_stdin_pipe(subprocess.native())
	return wrapOutputStream(ret0)
}

// GetStdoutPipe is a wrapper around g_subprocess_get_stdout_pipe().
func (subprocess Subprocess) GetStdoutPipe() InputStream {
	ret0 := C.g_subprocess_get_stdout_pipe(subprocess.native())
	return wrapInputStream(ret0)
}

// GetSuccessful is a wrapper around g_subprocess_get_successful().
func (subprocess Subprocess) GetSuccessful() bool {
	ret0 := C.g_subprocess_get_successful(subprocess.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetTermSig is a wrapper around g_subprocess_get_term_sig().
func (subprocess Subprocess) GetTermSig() int {
	ret0 := C.g_subprocess_get_term_sig(subprocess.native())
	return int(ret0)
}

// SendSignal is a wrapper around g_subprocess_send_signal().
func (subprocess Subprocess) SendSignal(signal_num int) {
	C.g_subprocess_send_signal(subprocess.native(), C.gint(signal_num))
}

// Wait is a wrapper around g_subprocess_wait().
func (subprocess Subprocess) Wait(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_subprocess_wait(subprocess.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// WaitAsync is a wrapper around g_subprocess_wait_async().
func (subprocess Subprocess) WaitAsync(cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_subprocess_wait_async(subprocess.native(), cancellable.native(), callback0)
}

// WaitCheck is a wrapper around g_subprocess_wait_check().
func (subprocess Subprocess) WaitCheck(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_subprocess_wait_check(subprocess.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// WaitCheckAsync is a wrapper around g_subprocess_wait_check_async().
func (subprocess Subprocess) WaitCheckAsync(cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_subprocess_wait_check_async(subprocess.native(), cancellable.native(), callback0)
}

// WaitCheckFinish is a wrapper around g_subprocess_wait_check_finish().
func (subprocess Subprocess) WaitCheckFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_subprocess_wait_check_finish(subprocess.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// WaitFinish is a wrapper around g_subprocess_wait_finish().
func (subprocess Subprocess) WaitFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_subprocess_wait_finish(subprocess.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Object SubprocessLauncher
type SubprocessLauncher struct {
	gobject.Object
}

func (v SubprocessLauncher) native() *C.GSubprocessLauncher {
	return (*C.GSubprocessLauncher)(v.Ptr)
}
func wrapSubprocessLauncher(p *C.GSubprocessLauncher) (v SubprocessLauncher) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapSubprocessLauncher(p unsafe.Pointer) (v SubprocessLauncher) {
	v.Ptr = p
	return
}
func (v SubprocessLauncher) IsNil() bool {
	return v.Ptr == nil
}
func IWrapSubprocessLauncher(p unsafe.Pointer) interface{} {
	return WrapSubprocessLauncher(p)
}
func (v SubprocessLauncher) GetType() gobject.Type {
	return gobject.Type(C.g_subprocess_launcher_get_type())
}
func (v SubprocessLauncher) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapSubprocessLauncher(unsafe.Pointer(ptr)), nil
	}
}

// SubprocessLauncherNew is a wrapper around g_subprocess_launcher_new().
func SubprocessLauncherNew(flags SubprocessFlags) SubprocessLauncher {
	ret0 := C.g_subprocess_launcher_new(C.GSubprocessFlags(flags))
	return wrapSubprocessLauncher(ret0)
}

// Getenv is a wrapper around g_subprocess_launcher_getenv().
func (self SubprocessLauncher) Getenv(variable string) string {
	variable0 := (*C.gchar)(C.CString(variable))
	ret0 := C.g_subprocess_launcher_getenv(self.native(), variable0)
	C.free(unsafe.Pointer(variable0)) /*ch:<stdlib.h>*/
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// SetCwd is a wrapper around g_subprocess_launcher_set_cwd().
func (self SubprocessLauncher) SetCwd(cwd string) {
	cwd0 := (*C.gchar)(C.CString(cwd))
	C.g_subprocess_launcher_set_cwd(self.native(), cwd0)
	C.free(unsafe.Pointer(cwd0)) /*ch:<stdlib.h>*/
}

// SetEnviron is a wrapper around g_subprocess_launcher_set_environ().
func (self SubprocessLauncher) SetEnviron(env []string) {
	env0 := make([]*C.gchar, len(env))
	for idx, elemG := range env {
		elem := (*C.gchar)(C.CString(elemG))
		env0[idx] = elem
	}
	var env0Ptr **C.gchar
	if len(env0) > 0 {
		env0Ptr = &env0[0]
	}
	C.g_subprocess_launcher_set_environ(self.native(), env0Ptr)
	for _, elem := range env0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
}

// SetFlags is a wrapper around g_subprocess_launcher_set_flags().
func (self SubprocessLauncher) SetFlags(flags SubprocessFlags) {
	C.g_subprocess_launcher_set_flags(self.native(), C.GSubprocessFlags(flags))
}

// SetStderrFilePath is a wrapper around g_subprocess_launcher_set_stderr_file_path().
func (self SubprocessLauncher) SetStderrFilePath(path string) {
	path0 := (*C.gchar)(C.CString(path))
	C.g_subprocess_launcher_set_stderr_file_path(self.native(), path0)
	C.free(unsafe.Pointer(path0)) /*ch:<stdlib.h>*/
}

// SetStdinFilePath is a wrapper around g_subprocess_launcher_set_stdin_file_path().
func (self SubprocessLauncher) SetStdinFilePath(path string) {
	path0 := (*C.gchar)(C.CString(path))
	C.g_subprocess_launcher_set_stdin_file_path(self.native(), path0)
	C.free(unsafe.Pointer(path0)) /*ch:<stdlib.h>*/
}

// SetStdoutFilePath is a wrapper around g_subprocess_launcher_set_stdout_file_path().
func (self SubprocessLauncher) SetStdoutFilePath(path string) {
	path0 := (*C.gchar)(C.CString(path))
	C.g_subprocess_launcher_set_stdout_file_path(self.native(), path0)
	C.free(unsafe.Pointer(path0)) /*ch:<stdlib.h>*/
}

// Setenv is a wrapper around g_subprocess_launcher_setenv().
func (self SubprocessLauncher) Setenv(variable string, value string, overwrite bool) {
	variable0 := (*C.gchar)(C.CString(variable))
	value0 := (*C.gchar)(C.CString(value))
	C.g_subprocess_launcher_setenv(self.native(), variable0, value0, C.gboolean(util.Bool2Int(overwrite)) /*go:.util*/)
	C.free(unsafe.Pointer(variable0)) /*ch:<stdlib.h>*/
	C.free(unsafe.Pointer(value0))    /*ch:<stdlib.h>*/
}

// Spawnv is a wrapper around g_subprocess_launcher_spawnv().
func (self SubprocessLauncher) Spawnv(argv []string) (Subprocess, error) {
	argv0 := make([]*C.gchar, len(argv))
	for idx, elemG := range argv {
		elem := (*C.gchar)(C.CString(elemG))
		argv0[idx] = elem
	}
	var argv0Ptr **C.gchar
	if len(argv0) > 0 {
		argv0Ptr = &argv0[0]
	}
	var err glib.Error
	ret0 := C.g_subprocess_launcher_spawnv(self.native(), argv0Ptr, (**C.GError)(unsafe.Pointer(&err)))
	for _, elem := range argv0 {
		C.free(unsafe.Pointer(elem)) /*ch:<stdlib.h>*/
	}
	if err.Ptr != nil {
		defer err.Free()
		return Subprocess{}, err.GoValue()
	}
	return wrapSubprocess(ret0), nil
}

// TakeFd is a wrapper around g_subprocess_launcher_take_fd().
func (self SubprocessLauncher) TakeFd(source_fd int, target_fd int) {
	C.g_subprocess_launcher_take_fd(self.native(), C.gint(source_fd), C.gint(target_fd))
}

// TakeStderrFd is a wrapper around g_subprocess_launcher_take_stderr_fd().
func (self SubprocessLauncher) TakeStderrFd(fd int) {
	C.g_subprocess_launcher_take_stderr_fd(self.native(), C.gint(fd))
}

// TakeStdinFd is a wrapper around g_subprocess_launcher_take_stdin_fd().
func (self SubprocessLauncher) TakeStdinFd(fd int) {
	C.g_subprocess_launcher_take_stdin_fd(self.native(), C.gint(fd))
}

// TakeStdoutFd is a wrapper around g_subprocess_launcher_take_stdout_fd().
func (self SubprocessLauncher) TakeStdoutFd(fd int) {
	C.g_subprocess_launcher_take_stdout_fd(self.native(), C.gint(fd))
}

// Unsetenv is a wrapper around g_subprocess_launcher_unsetenv().
func (self SubprocessLauncher) Unsetenv(variable string) {
	variable0 := (*C.gchar)(C.CString(variable))
	C.g_subprocess_launcher_unsetenv(self.native(), variable0)
	C.free(unsafe.Pointer(variable0)) /*ch:<stdlib.h>*/
}

// Object Task
type Task struct {
	gobject.Object
}

func (v Task) native() *C.GTask {
	return (*C.GTask)(v.Ptr)
}
func wrapTask(p *C.GTask) (v Task) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTask(p unsafe.Pointer) (v Task) {
	v.Ptr = p
	return
}
func (v Task) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTask(p unsafe.Pointer) interface{} {
	return WrapTask(p)
}
func (v Task) GetType() gobject.Type {
	return gobject.Type(C.g_task_get_type())
}
func (v Task) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTask(unsafe.Pointer(ptr)), nil
	}
}
func (v Task) AsyncResult() AsyncResult {
	return WrapAsyncResult(v.Ptr)
}

// TaskNew is a wrapper around g_task_new().
func TaskNew(source_object gobject.Object, cancellable Cancellable, callback AsyncReadyCallback) Task {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	ret0 := C._g_task_new(C.gpointer((C.gpointer)(source_object.Ptr)), cancellable.native(), callback0)
	return wrapTask(ret0)
}

// GetCancellable is a wrapper around g_task_get_cancellable().
func (task Task) GetCancellable() Cancellable {
	ret0 := C.g_task_get_cancellable(task.native())
	return wrapCancellable(ret0)
}

// GetCheckCancellable is a wrapper around g_task_get_check_cancellable().
func (task Task) GetCheckCancellable() bool {
	ret0 := C.g_task_get_check_cancellable(task.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetCompleted is a wrapper around g_task_get_completed().
func (task Task) GetCompleted() bool {
	ret0 := C.g_task_get_completed(task.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetContext is a wrapper around g_task_get_context().
func (task Task) GetContext() glib.MainContext {
	ret0 := C.g_task_get_context(task.native())
	return glib.WrapMainContext(unsafe.Pointer(ret0)) /*gir:GLib*/
}

// GetPriority is a wrapper around g_task_get_priority().
func (task Task) GetPriority() int {
	ret0 := C.g_task_get_priority(task.native())
	return int(ret0)
}

// GetReturnOnCancel is a wrapper around g_task_get_return_on_cancel().
func (task Task) GetReturnOnCancel() bool {
	ret0 := C.g_task_get_return_on_cancel(task.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetSourceObject is a wrapper around g_task_get_source_object().
func (task Task) GetSourceObject() gobject.Object {
	ret0 := C.g_task_get_source_object(task.native())
	return gobject.WrapObject(unsafe.Pointer(ret0)) /*gir:GObject*/
}

// GetSourceTag is a wrapper around g_task_get_source_tag().
func (task Task) GetSourceTag() unsafe.Pointer {
	ret0 := C.g_task_get_source_tag(task.native())
	return unsafe.Pointer(ret0)
}

// GetTaskData is a wrapper around g_task_get_task_data().
func (task Task) GetTaskData() unsafe.Pointer {
	ret0 := C.g_task_get_task_data(task.native())
	return unsafe.Pointer(ret0)
}

// HadError is a wrapper around g_task_had_error().
func (task Task) HadError() bool {
	ret0 := C.g_task_had_error(task.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// PropagateBoolean is a wrapper around g_task_propagate_boolean().
func (task Task) PropagateBoolean() (bool, error) {
	var err glib.Error
	ret0 := C.g_task_propagate_boolean(task.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// PropagateInt is a wrapper around g_task_propagate_int().
func (task Task) PropagateInt() (int, error) {
	var err glib.Error
	ret0 := C.g_task_propagate_int(task.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// PropagatePointer is a wrapper around g_task_propagate_pointer().
func (task Task) PropagatePointer() (unsafe.Pointer, error) {
	var err glib.Error
	ret0 := C.g_task_propagate_pointer(task.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return unsafe.Pointer(nil), err.GoValue()
	}
	return unsafe.Pointer(ret0), nil
}

// ReturnBoolean is a wrapper around g_task_return_boolean().
func (task Task) ReturnBoolean(result bool) {
	C.g_task_return_boolean(task.native(), C.gboolean(util.Bool2Int(result)) /*go:.util*/)
}

// ReturnError is a wrapper around g_task_return_error().
func (task Task) ReturnError(error glib.Error) {
	C.g_task_return_error(task.native(), (*C.GError)(error.Ptr))
}

// ReturnErrorIfCancelled is a wrapper around g_task_return_error_if_cancelled().
func (task Task) ReturnErrorIfCancelled() bool {
	ret0 := C.g_task_return_error_if_cancelled(task.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ReturnInt is a wrapper around g_task_return_int().
func (task Task) ReturnInt(result int) {
	C.g_task_return_int(task.native(), C.gssize(result))
}

// SetCheckCancellable is a wrapper around g_task_set_check_cancellable().
func (task Task) SetCheckCancellable(check_cancellable bool) {
	C.g_task_set_check_cancellable(task.native(), C.gboolean(util.Bool2Int(check_cancellable)) /*go:.util*/)
}

// SetPriority is a wrapper around g_task_set_priority().
func (task Task) SetPriority(priority int) {
	C.g_task_set_priority(task.native(), C.gint(priority))
}

// SetReturnOnCancel is a wrapper around g_task_set_return_on_cancel().
func (task Task) SetReturnOnCancel(return_on_cancel bool) bool {
	ret0 := C.g_task_set_return_on_cancel(task.native(), C.gboolean(util.Bool2Int(return_on_cancel)) /*go:.util*/)
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetSourceTag is a wrapper around g_task_set_source_tag().
func (task Task) SetSourceTag(source_tag unsafe.Pointer) {
	C.g_task_set_source_tag(task.native(), C.gpointer(source_tag))
}

// TaskIsValid is a wrapper around g_task_is_valid().
func TaskIsValid(result AsyncResult, source_object gobject.Object) bool {
	ret0 := C.g_task_is_valid(C.gpointer(result.native()), C.gpointer((C.gpointer)(source_object.Ptr)))
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// TaskReportError is a wrapper around g_task_report_error().
func TaskReportError(source_object gobject.Object, callback AsyncReadyCallback, source_tag unsafe.Pointer, error glib.Error) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_task_report_error(C.gpointer((C.gpointer)(source_object.Ptr)), callback0, C.gpointer(source_tag), (*C.GError)(error.Ptr))
}

// Object TcpConnection
type TcpConnection struct {
	SocketConnection
}

func (v TcpConnection) native() *C.GTcpConnection {
	return (*C.GTcpConnection)(v.Ptr)
}
func wrapTcpConnection(p *C.GTcpConnection) (v TcpConnection) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTcpConnection(p unsafe.Pointer) (v TcpConnection) {
	v.Ptr = p
	return
}
func (v TcpConnection) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTcpConnection(p unsafe.Pointer) interface{} {
	return WrapTcpConnection(p)
}
func (v TcpConnection) GetType() gobject.Type {
	return gobject.Type(C.g_tcp_connection_get_type())
}
func (v TcpConnection) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTcpConnection(unsafe.Pointer(ptr)), nil
	}
}

// GetGracefulDisconnect is a wrapper around g_tcp_connection_get_graceful_disconnect().
func (connection TcpConnection) GetGracefulDisconnect() bool {
	ret0 := C.g_tcp_connection_get_graceful_disconnect(connection.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// SetGracefulDisconnect is a wrapper around g_tcp_connection_set_graceful_disconnect().
func (connection TcpConnection) SetGracefulDisconnect(graceful_disconnect bool) {
	C.g_tcp_connection_set_graceful_disconnect(connection.native(), C.gboolean(util.Bool2Int(graceful_disconnect)) /*go:.util*/)
}

// Object TcpWrapperConnection
type TcpWrapperConnection struct {
	TcpConnection
}

func (v TcpWrapperConnection) native() *C.GTcpWrapperConnection {
	return (*C.GTcpWrapperConnection)(v.Ptr)
}
func wrapTcpWrapperConnection(p *C.GTcpWrapperConnection) (v TcpWrapperConnection) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTcpWrapperConnection(p unsafe.Pointer) (v TcpWrapperConnection) {
	v.Ptr = p
	return
}
func (v TcpWrapperConnection) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTcpWrapperConnection(p unsafe.Pointer) interface{} {
	return WrapTcpWrapperConnection(p)
}
func (v TcpWrapperConnection) GetType() gobject.Type {
	return gobject.Type(C.g_tcp_wrapper_connection_get_type())
}
func (v TcpWrapperConnection) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTcpWrapperConnection(unsafe.Pointer(ptr)), nil
	}
}

// TcpWrapperConnectionNew is a wrapper around g_tcp_wrapper_connection_new().
func TcpWrapperConnectionNew(base_io_stream IOStream, socket Socket) SocketConnection {
	ret0 := C.g_tcp_wrapper_connection_new(base_io_stream.native(), socket.native())
	return wrapSocketConnection(ret0)
}

// GetBaseIoStream is a wrapper around g_tcp_wrapper_connection_get_base_io_stream().
func (conn TcpWrapperConnection) GetBaseIoStream() IOStream {
	ret0 := C.g_tcp_wrapper_connection_get_base_io_stream(conn.native())
	return wrapIOStream(ret0)
}

// Object TestDBus
type TestDBus struct {
	gobject.Object
}

func (v TestDBus) native() *C.GTestDBus {
	return (*C.GTestDBus)(v.Ptr)
}
func wrapTestDBus(p *C.GTestDBus) (v TestDBus) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapTestDBus(p unsafe.Pointer) (v TestDBus) {
	v.Ptr = p
	return
}
func (v TestDBus) IsNil() bool {
	return v.Ptr == nil
}
func IWrapTestDBus(p unsafe.Pointer) interface{} {
	return WrapTestDBus(p)
}
func (v TestDBus) GetType() gobject.Type {
	return gobject.Type(C.g_test_dbus_get_type())
}
func (v TestDBus) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapTestDBus(unsafe.Pointer(ptr)), nil
	}
}

// TestDBusNew is a wrapper around g_test_dbus_new().
func TestDBusNew(flags TestDBusFlags) TestDBus {
	ret0 := C.g_test_dbus_new(C.GTestDBusFlags(flags))
	return wrapTestDBus(ret0)
}

// AddServiceDir is a wrapper around g_test_dbus_add_service_dir().
func (self TestDBus) AddServiceDir(path string) {
	path0 := (*C.gchar)(C.CString(path))
	C.g_test_dbus_add_service_dir(self.native(), path0)
	C.free(unsafe.Pointer(path0)) /*ch:<stdlib.h>*/
}

// Down is a wrapper around g_test_dbus_down().
func (self TestDBus) Down() {
	C.g_test_dbus_down(self.native())
}

// GetBusAddress is a wrapper around g_test_dbus_get_bus_address().
func (self TestDBus) GetBusAddress() string {
	ret0 := C.g_test_dbus_get_bus_address(self.native())
	ret := C.GoString((*C.char)(ret0))
	return ret
}

// GetFlags is a wrapper around g_test_dbus_get_flags().
func (self TestDBus) GetFlags() TestDBusFlags {
	ret0 := C.g_test_dbus_get_flags(self.native())
	return TestDBusFlags(ret0)
}

// Stop is a wrapper around g_test_dbus_stop().
func (self TestDBus) Stop() {
	C.g_test_dbus_stop(self.native())
}

// Up is a wrapper around g_test_dbus_up().
func (self TestDBus) Up() {
	C.g_test_dbus_up(self.native())
}

// TestDBusUnset is a wrapper around g_test_dbus_unset().
func TestDBusUnset() {
	C.g_test_dbus_unset()
}

// Object ThemedIcon
type ThemedIcon struct {
	gobject.Object
}

func (v ThemedIcon) native() *C.GThemedIcon {
	return (*C.GThemedIcon)(v.Ptr)
}
func wrapThemedIcon(p *C.GThemedIcon) (v ThemedIcon) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapThemedIcon(p unsafe.Pointer) (v ThemedIcon) {
	v.Ptr = p
	return
}
func (v ThemedIcon) IsNil() bool {
	return v.Ptr == nil
}
func IWrapThemedIcon(p unsafe.Pointer) interface{} {
	return WrapThemedIcon(p)
}
func (v ThemedIcon) GetType() gobject.Type {
	return gobject.Type(C.g_themed_icon_get_type())
}
func (v ThemedIcon) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapThemedIcon(unsafe.Pointer(ptr)), nil
	}
}
func (v ThemedIcon) Icon() Icon {
	return WrapIcon(v.Ptr)
}

// AppendName is a wrapper around g_themed_icon_append_name().
func (icon ThemedIcon) AppendName(iconname string) {
	iconname0 := C.CString(iconname)
	C.g_themed_icon_append_name(icon.native(), iconname0)
	C.free(unsafe.Pointer(iconname0)) /*ch:<stdlib.h>*/
}

// PrependName is a wrapper around g_themed_icon_prepend_name().
func (icon ThemedIcon) PrependName(iconname string) {
	iconname0 := C.CString(iconname)
	C.g_themed_icon_prepend_name(icon.native(), iconname0)
	C.free(unsafe.Pointer(iconname0)) /*ch:<stdlib.h>*/
}

// Object ThreadedSocketService
type ThreadedSocketService struct {
	SocketService
}

func (v ThreadedSocketService) native() *C.GThreadedSocketService {
	return (*C.GThreadedSocketService)(v.Ptr)
}
func wrapThreadedSocketService(p *C.GThreadedSocketService) (v ThreadedSocketService) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapThreadedSocketService(p unsafe.Pointer) (v ThreadedSocketService) {
	v.Ptr = p
	return
}
func (v ThreadedSocketService) IsNil() bool {
	return v.Ptr == nil
}
func IWrapThreadedSocketService(p unsafe.Pointer) interface{} {
	return WrapThreadedSocketService(p)
}
func (v ThreadedSocketService) GetType() gobject.Type {
	return gobject.Type(C.g_threaded_socket_service_get_type())
}
func (v ThreadedSocketService) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapThreadedSocketService(unsafe.Pointer(ptr)), nil
	}
}

// ThreadedSocketServiceNew is a wrapper around g_threaded_socket_service_new().
func ThreadedSocketServiceNew(max_threads int) SocketService {
	ret0 := C.g_threaded_socket_service_new(C.int(max_threads))
	return wrapSocketService(ret0)
}

// Object UnixConnection
type UnixConnection struct {
	SocketConnection
}

func (v UnixConnection) native() *C.GUnixConnection {
	return (*C.GUnixConnection)(v.Ptr)
}
func wrapUnixConnection(p *C.GUnixConnection) (v UnixConnection) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapUnixConnection(p unsafe.Pointer) (v UnixConnection) {
	v.Ptr = p
	return
}
func (v UnixConnection) IsNil() bool {
	return v.Ptr == nil
}
func IWrapUnixConnection(p unsafe.Pointer) interface{} {
	return WrapUnixConnection(p)
}
func (v UnixConnection) GetType() gobject.Type {
	return gobject.Type(C.g_unix_connection_get_type())
}
func (v UnixConnection) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapUnixConnection(unsafe.Pointer(ptr)), nil
	}
}

// ReceiveCredentials is a wrapper around g_unix_connection_receive_credentials().
func (connection UnixConnection) ReceiveCredentials(cancellable Cancellable) (Credentials, error) {
	var err glib.Error
	ret0 := C.g_unix_connection_receive_credentials(connection.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return Credentials{}, err.GoValue()
	}
	return wrapCredentials(ret0), nil
}

// ReceiveCredentialsAsync is a wrapper around g_unix_connection_receive_credentials_async().
func (connection UnixConnection) ReceiveCredentialsAsync(cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_unix_connection_receive_credentials_async(connection.native(), cancellable.native(), callback0)
}

// ReceiveCredentialsFinish is a wrapper around g_unix_connection_receive_credentials_finish().
func (connection UnixConnection) ReceiveCredentialsFinish(result AsyncResult) (Credentials, error) {
	var err glib.Error
	ret0 := C.g_unix_connection_receive_credentials_finish(connection.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return Credentials{}, err.GoValue()
	}
	return wrapCredentials(ret0), nil
}

// ReceiveFd is a wrapper around g_unix_connection_receive_fd().
func (connection UnixConnection) ReceiveFd(cancellable Cancellable) (int, error) {
	var err glib.Error
	ret0 := C.g_unix_connection_receive_fd(connection.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return 0, err.GoValue()
	}
	return int(ret0), nil
}

// SendCredentials is a wrapper around g_unix_connection_send_credentials().
func (connection UnixConnection) SendCredentials(cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_unix_connection_send_credentials(connection.native(), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SendCredentialsAsync is a wrapper around g_unix_connection_send_credentials_async().
func (connection UnixConnection) SendCredentialsAsync(cancellable Cancellable, callback AsyncReadyCallback) {
	callback0 := (*C.GClosure)(gobject.ClosureNew(callback).Ptr) /*gir:GObject*/
	C._g_unix_connection_send_credentials_async(connection.native(), cancellable.native(), callback0)
}

// SendCredentialsFinish is a wrapper around g_unix_connection_send_credentials_finish().
func (connection UnixConnection) SendCredentialsFinish(result AsyncResult) (bool, error) {
	var err glib.Error
	ret0 := C.g_unix_connection_send_credentials_finish(connection.native(), result.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// SendFd is a wrapper around g_unix_connection_send_fd().
func (connection UnixConnection) SendFd(fd int, cancellable Cancellable) (bool, error) {
	var err glib.Error
	ret0 := C.g_unix_connection_send_fd(connection.native(), C.gint(fd), cancellable.native(), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// Object UnixCredentialsMessage
type UnixCredentialsMessage struct {
	SocketControlMessage
}

func (v UnixCredentialsMessage) native() *C.GUnixCredentialsMessage {
	return (*C.GUnixCredentialsMessage)(v.Ptr)
}
func wrapUnixCredentialsMessage(p *C.GUnixCredentialsMessage) (v UnixCredentialsMessage) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapUnixCredentialsMessage(p unsafe.Pointer) (v UnixCredentialsMessage) {
	v.Ptr = p
	return
}
func (v UnixCredentialsMessage) IsNil() bool {
	return v.Ptr == nil
}
func IWrapUnixCredentialsMessage(p unsafe.Pointer) interface{} {
	return WrapUnixCredentialsMessage(p)
}
func (v UnixCredentialsMessage) GetType() gobject.Type {
	return gobject.Type(C.g_unix_credentials_message_get_type())
}
func (v UnixCredentialsMessage) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapUnixCredentialsMessage(unsafe.Pointer(ptr)), nil
	}
}

// UnixCredentialsMessageNew is a wrapper around g_unix_credentials_message_new().
func UnixCredentialsMessageNew() SocketControlMessage {
	ret0 := C.g_unix_credentials_message_new()
	return wrapSocketControlMessage(ret0)
}

// UnixCredentialsMessageNewWithCredentials is a wrapper around g_unix_credentials_message_new_with_credentials().
func UnixCredentialsMessageNewWithCredentials(credentials Credentials) SocketControlMessage {
	ret0 := C.g_unix_credentials_message_new_with_credentials(credentials.native())
	return wrapSocketControlMessage(ret0)
}

// GetCredentials is a wrapper around g_unix_credentials_message_get_credentials().
func (message UnixCredentialsMessage) GetCredentials() Credentials {
	ret0 := C.g_unix_credentials_message_get_credentials(message.native())
	return wrapCredentials(ret0)
}

// UnixCredentialsMessageIsSupported is a wrapper around g_unix_credentials_message_is_supported().
func UnixCredentialsMessageIsSupported() bool {
	ret0 := C.g_unix_credentials_message_is_supported()
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object UnixFDMessage
type UnixFDMessage struct {
	SocketControlMessage
}

func (v UnixFDMessage) native() *C.GUnixFDMessage {
	return (*C.GUnixFDMessage)(v.Ptr)
}
func wrapUnixFDMessage(p *C.GUnixFDMessage) (v UnixFDMessage) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapUnixFDMessage(p unsafe.Pointer) (v UnixFDMessage) {
	v.Ptr = p
	return
}
func (v UnixFDMessage) IsNil() bool {
	return v.Ptr == nil
}
func IWrapUnixFDMessage(p unsafe.Pointer) interface{} {
	return WrapUnixFDMessage(p)
}
func (v UnixFDMessage) GetType() gobject.Type {
	return gobject.Type(C.g_unix_fd_message_get_type())
}
func (v UnixFDMessage) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapUnixFDMessage(unsafe.Pointer(ptr)), nil
	}
}

// UnixFDMessageNew is a wrapper around g_unix_fd_message_new().
func UnixFDMessageNew() SocketControlMessage {
	ret0 := C.g_unix_fd_message_new()
	return wrapSocketControlMessage(ret0)
}

// UnixFDMessageNewWithFdList is a wrapper around g_unix_fd_message_new_with_fd_list().
func UnixFDMessageNewWithFdList(fd_list UnixFDList) SocketControlMessage {
	ret0 := C.g_unix_fd_message_new_with_fd_list(fd_list.native())
	return wrapSocketControlMessage(ret0)
}

// AppendFd is a wrapper around g_unix_fd_message_append_fd().
func (message UnixFDMessage) AppendFd(fd int) (bool, error) {
	var err glib.Error
	ret0 := C.g_unix_fd_message_append_fd(message.native(), C.gint(fd), (**C.GError)(unsafe.Pointer(&err)))
	if err.Ptr != nil {
		defer err.Free()
		return false, err.GoValue()
	}
	return util.Int2Bool(int(ret0)) /*go:.util*/, nil
}

// GetFdList is a wrapper around g_unix_fd_message_get_fd_list().
func (message UnixFDMessage) GetFdList() UnixFDList {
	ret0 := C.g_unix_fd_message_get_fd_list(message.native())
	return wrapUnixFDList(ret0)
}

// StealFds is a wrapper around g_unix_fd_message_steal_fds().
func (message UnixFDMessage) StealFds() []int {
	var length0 C.gint
	ret0 := C.g_unix_fd_message_steal_fds(message.native(), &length0)
	var ret0Slice []C.gint
	util.SetSliceDataLen(unsafe.Pointer(&ret0Slice), unsafe.Pointer(ret0), int(length0)) /*go:.util*/
	ret := make([]int, len(ret0Slice))
	for idx, elem := range ret0Slice {
		ret[idx] = int(elem)
	}
	C.g_free(C.gpointer(ret0))
	return ret
}

// Object UnixInputStream
type UnixInputStream struct {
	InputStream
}

func (v UnixInputStream) native() *C.GUnixInputStream {
	return (*C.GUnixInputStream)(v.Ptr)
}
func wrapUnixInputStream(p *C.GUnixInputStream) (v UnixInputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapUnixInputStream(p unsafe.Pointer) (v UnixInputStream) {
	v.Ptr = p
	return
}
func (v UnixInputStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapUnixInputStream(p unsafe.Pointer) interface{} {
	return WrapUnixInputStream(p)
}
func (v UnixInputStream) GetType() gobject.Type {
	return gobject.Type(C.g_unix_input_stream_get_type())
}
func (v UnixInputStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapUnixInputStream(unsafe.Pointer(ptr)), nil
	}
}
func (v UnixInputStream) FileDescriptorBased() FileDescriptorBased {
	return WrapFileDescriptorBased(v.Ptr)
}
func (v UnixInputStream) PollableInputStream() PollableInputStream {
	return WrapPollableInputStream(v.Ptr)
}

// UnixInputStreamNew is a wrapper around g_unix_input_stream_new().
func UnixInputStreamNew(fd int, close_fd bool) InputStream {
	ret0 := C.g_unix_input_stream_new(C.gint(fd), C.gboolean(util.Bool2Int(close_fd)) /*go:.util*/)
	return wrapInputStream(ret0)
}

// GetCloseFd is a wrapper around g_unix_input_stream_get_close_fd().
func (stream UnixInputStream) GetCloseFd() bool {
	ret0 := C.g_unix_input_stream_get_close_fd(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetFd is a wrapper around g_unix_input_stream_get_fd().
func (stream UnixInputStream) GetFd() int {
	ret0 := C.g_unix_input_stream_get_fd(stream.native())
	return int(ret0)
}

// SetCloseFd is a wrapper around g_unix_input_stream_set_close_fd().
func (stream UnixInputStream) SetCloseFd(close_fd bool) {
	C.g_unix_input_stream_set_close_fd(stream.native(), C.gboolean(util.Bool2Int(close_fd)) /*go:.util*/)
}

// Object UnixMountMonitor
type UnixMountMonitor struct {
	gobject.Object
}

func (v UnixMountMonitor) native() *C.GUnixMountMonitor {
	return (*C.GUnixMountMonitor)(v.Ptr)
}
func wrapUnixMountMonitor(p *C.GUnixMountMonitor) (v UnixMountMonitor) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapUnixMountMonitor(p unsafe.Pointer) (v UnixMountMonitor) {
	v.Ptr = p
	return
}
func (v UnixMountMonitor) IsNil() bool {
	return v.Ptr == nil
}
func IWrapUnixMountMonitor(p unsafe.Pointer) interface{} {
	return WrapUnixMountMonitor(p)
}
func (v UnixMountMonitor) GetType() gobject.Type {
	return gobject.Type(C.g_unix_mount_monitor_get_type())
}
func (v UnixMountMonitor) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapUnixMountMonitor(unsafe.Pointer(ptr)), nil
	}
}

// UnixMountMonitorGet is a wrapper around g_unix_mount_monitor_get().
func UnixMountMonitorGet() UnixMountMonitor {
	ret0 := C.g_unix_mount_monitor_get()
	return wrapUnixMountMonitor(ret0)
}

// Object UnixOutputStream
type UnixOutputStream struct {
	OutputStream
}

func (v UnixOutputStream) native() *C.GUnixOutputStream {
	return (*C.GUnixOutputStream)(v.Ptr)
}
func wrapUnixOutputStream(p *C.GUnixOutputStream) (v UnixOutputStream) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapUnixOutputStream(p unsafe.Pointer) (v UnixOutputStream) {
	v.Ptr = p
	return
}
func (v UnixOutputStream) IsNil() bool {
	return v.Ptr == nil
}
func IWrapUnixOutputStream(p unsafe.Pointer) interface{} {
	return WrapUnixOutputStream(p)
}
func (v UnixOutputStream) GetType() gobject.Type {
	return gobject.Type(C.g_unix_output_stream_get_type())
}
func (v UnixOutputStream) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapUnixOutputStream(unsafe.Pointer(ptr)), nil
	}
}
func (v UnixOutputStream) FileDescriptorBased() FileDescriptorBased {
	return WrapFileDescriptorBased(v.Ptr)
}
func (v UnixOutputStream) PollableOutputStream() PollableOutputStream {
	return WrapPollableOutputStream(v.Ptr)
}

// UnixOutputStreamNew is a wrapper around g_unix_output_stream_new().
func UnixOutputStreamNew(fd int, close_fd bool) OutputStream {
	ret0 := C.g_unix_output_stream_new(C.gint(fd), C.gboolean(util.Bool2Int(close_fd)) /*go:.util*/)
	return wrapOutputStream(ret0)
}

// GetCloseFd is a wrapper around g_unix_output_stream_get_close_fd().
func (stream UnixOutputStream) GetCloseFd() bool {
	ret0 := C.g_unix_output_stream_get_close_fd(stream.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// GetFd is a wrapper around g_unix_output_stream_get_fd().
func (stream UnixOutputStream) GetFd() int {
	ret0 := C.g_unix_output_stream_get_fd(stream.native())
	return int(ret0)
}

// SetCloseFd is a wrapper around g_unix_output_stream_set_close_fd().
func (stream UnixOutputStream) SetCloseFd(close_fd bool) {
	C.g_unix_output_stream_set_close_fd(stream.native(), C.gboolean(util.Bool2Int(close_fd)) /*go:.util*/)
}

// Object UnixSocketAddress
type UnixSocketAddress struct {
	SocketAddress
}

func (v UnixSocketAddress) native() *C.GUnixSocketAddress {
	return (*C.GUnixSocketAddress)(v.Ptr)
}
func wrapUnixSocketAddress(p *C.GUnixSocketAddress) (v UnixSocketAddress) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapUnixSocketAddress(p unsafe.Pointer) (v UnixSocketAddress) {
	v.Ptr = p
	return
}
func (v UnixSocketAddress) IsNil() bool {
	return v.Ptr == nil
}
func IWrapUnixSocketAddress(p unsafe.Pointer) interface{} {
	return WrapUnixSocketAddress(p)
}
func (v UnixSocketAddress) GetType() gobject.Type {
	return gobject.Type(C.g_unix_socket_address_get_type())
}
func (v UnixSocketAddress) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapUnixSocketAddress(unsafe.Pointer(ptr)), nil
	}
}
func (v UnixSocketAddress) SocketConnectable() SocketConnectable {
	return WrapSocketConnectable(v.Ptr)
}

// UnixSocketAddressNew is a wrapper around g_unix_socket_address_new().
func UnixSocketAddressNew(path string) SocketAddress {
	path0 := (*C.gchar)(C.CString(path))
	ret0 := C.g_unix_socket_address_new(path0)
	C.free(unsafe.Pointer(path0)) /*ch:<stdlib.h>*/
	return wrapSocketAddress(ret0)
}

// UnixSocketAddressNewWithType is a wrapper around g_unix_socket_address_new_with_type().
func UnixSocketAddressNewWithType(path []int8, type_ UnixSocketAddressType) SocketAddress {
	path0 := make([]C.gchar, len(path))
	for idx, elemG := range path {
		path0[idx] = C.gchar(elemG)
	}
	var path0Ptr *C.gchar
	if len(path0) > 0 {
		path0Ptr = &path0[0]
	}
	ret0 := C.g_unix_socket_address_new_with_type(path0Ptr, C.gint(len(path)), C.GUnixSocketAddressType(type_))
	return wrapSocketAddress(ret0)
}

// GetAddressType is a wrapper around g_unix_socket_address_get_address_type().
func (address UnixSocketAddress) GetAddressType() UnixSocketAddressType {
	ret0 := C.g_unix_socket_address_get_address_type(address.native())
	return UnixSocketAddressType(ret0)
}

// GetPath is a wrapper around g_unix_socket_address_get_path().
func (address UnixSocketAddress) GetPath() string {
	ret0 := C.g_unix_socket_address_get_path(address.native())
	ret := C.GoString(ret0)
	return ret
}

// GetPathLen is a wrapper around g_unix_socket_address_get_path_len().
func (address UnixSocketAddress) GetPathLen() uint {
	ret0 := C.g_unix_socket_address_get_path_len(address.native())
	return uint(ret0)
}

// UnixSocketAddressAbstractNamesSupported is a wrapper around g_unix_socket_address_abstract_names_supported().
func UnixSocketAddressAbstractNamesSupported() bool {
	ret0 := C.g_unix_socket_address_abstract_names_supported()
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// Object Vfs
type Vfs struct {
	gobject.Object
}

func (v Vfs) native() *C.GVfs {
	return (*C.GVfs)(v.Ptr)
}
func wrapVfs(p *C.GVfs) (v Vfs) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapVfs(p unsafe.Pointer) (v Vfs) {
	v.Ptr = p
	return
}
func (v Vfs) IsNil() bool {
	return v.Ptr == nil
}
func IWrapVfs(p unsafe.Pointer) interface{} {
	return WrapVfs(p)
}
func (v Vfs) GetType() gobject.Type {
	return gobject.Type(C.g_vfs_get_type())
}
func (v Vfs) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapVfs(unsafe.Pointer(ptr)), nil
	}
}

// GetFileForPath is a wrapper around g_vfs_get_file_for_path().
func (vfs Vfs) GetFileForPath(path string) File {
	path0 := C.CString(path)
	ret0 := C.g_vfs_get_file_for_path(vfs.native(), path0)
	C.free(unsafe.Pointer(path0)) /*ch:<stdlib.h>*/
	return wrapFile(ret0)
}

// GetFileForUri is a wrapper around g_vfs_get_file_for_uri().
func (vfs Vfs) GetFileForUri(uri string) File {
	uri0 := C.CString(uri)
	ret0 := C.g_vfs_get_file_for_uri(vfs.native(), uri0)
	C.free(unsafe.Pointer(uri0)) /*ch:<stdlib.h>*/
	return wrapFile(ret0)
}

// IsActive is a wrapper around g_vfs_is_active().
func (vfs Vfs) IsActive() bool {
	ret0 := C.g_vfs_is_active(vfs.native())
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// ParseName is a wrapper around g_vfs_parse_name().
func (vfs Vfs) ParseName(parse_name string) File {
	parse_name0 := C.CString(parse_name)
	ret0 := C.g_vfs_parse_name(vfs.native(), parse_name0)
	C.free(unsafe.Pointer(parse_name0)) /*ch:<stdlib.h>*/
	return wrapFile(ret0)
}

// UnregisterUriScheme is a wrapper around g_vfs_unregister_uri_scheme().
func (vfs Vfs) UnregisterUriScheme(scheme string) bool {
	scheme0 := C.CString(scheme)
	ret0 := C.g_vfs_unregister_uri_scheme(vfs.native(), scheme0)
	C.free(unsafe.Pointer(scheme0)) /*ch:<stdlib.h>*/
	return util.Int2Bool(int(ret0)) /*go:.util*/
}

// VfsGetDefault is a wrapper around g_vfs_get_default().
func VfsGetDefault() Vfs {
	ret0 := C.g_vfs_get_default()
	return wrapVfs(ret0)
}

// VfsGetLocal is a wrapper around g_vfs_get_local().
func VfsGetLocal() Vfs {
	ret0 := C.g_vfs_get_local()
	return wrapVfs(ret0)
}

// Object ZlibCompressor
type ZlibCompressor struct {
	gobject.Object
}

func (v ZlibCompressor) native() *C.GZlibCompressor {
	return (*C.GZlibCompressor)(v.Ptr)
}
func wrapZlibCompressor(p *C.GZlibCompressor) (v ZlibCompressor) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapZlibCompressor(p unsafe.Pointer) (v ZlibCompressor) {
	v.Ptr = p
	return
}
func (v ZlibCompressor) IsNil() bool {
	return v.Ptr == nil
}
func IWrapZlibCompressor(p unsafe.Pointer) interface{} {
	return WrapZlibCompressor(p)
}
func (v ZlibCompressor) GetType() gobject.Type {
	return gobject.Type(C.g_zlib_compressor_get_type())
}
func (v ZlibCompressor) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapZlibCompressor(unsafe.Pointer(ptr)), nil
	}
}
func (v ZlibCompressor) Converter() Converter {
	return WrapConverter(v.Ptr)
}

// ZlibCompressorNew is a wrapper around g_zlib_compressor_new().
func ZlibCompressorNew(format ZlibCompressorFormat, level int) ZlibCompressor {
	ret0 := C.g_zlib_compressor_new(C.GZlibCompressorFormat(format), C.int(level))
	return wrapZlibCompressor(ret0)
}

// GetFileInfo is a wrapper around g_zlib_compressor_get_file_info().
func (compressor ZlibCompressor) GetFileInfo() FileInfo {
	ret0 := C.g_zlib_compressor_get_file_info(compressor.native())
	return wrapFileInfo(ret0)
}

// SetFileInfo is a wrapper around g_zlib_compressor_set_file_info().
func (compressor ZlibCompressor) SetFileInfo(file_info FileInfo) {
	C.g_zlib_compressor_set_file_info(compressor.native(), file_info.native())
}

// Object ZlibDecompressor
type ZlibDecompressor struct {
	gobject.Object
}

func (v ZlibDecompressor) native() *C.GZlibDecompressor {
	return (*C.GZlibDecompressor)(v.Ptr)
}
func wrapZlibDecompressor(p *C.GZlibDecompressor) (v ZlibDecompressor) {
	v.Ptr = unsafe.Pointer(p)
	return
}
func WrapZlibDecompressor(p unsafe.Pointer) (v ZlibDecompressor) {
	v.Ptr = p
	return
}
func (v ZlibDecompressor) IsNil() bool {
	return v.Ptr == nil
}
func IWrapZlibDecompressor(p unsafe.Pointer) interface{} {
	return WrapZlibDecompressor(p)
}
func (v ZlibDecompressor) GetType() gobject.Type {
	return gobject.Type(C.g_zlib_decompressor_get_type())
}
func (v ZlibDecompressor) GetGValueGetter() gobject.GValueGetter {
	return func(p unsafe.Pointer) (interface{}, error) {
		ptr := C.g_value_get_object((*C.GValue)(p))
		return WrapZlibDecompressor(unsafe.Pointer(ptr)), nil
	}
}
func (v ZlibDecompressor) Converter() Converter {
	return WrapConverter(v.Ptr)
}

// ZlibDecompressorNew is a wrapper around g_zlib_decompressor_new().
func ZlibDecompressorNew(format ZlibCompressorFormat) ZlibDecompressor {
	ret0 := C.g_zlib_decompressor_new(C.GZlibCompressorFormat(format))
	return wrapZlibDecompressor(ret0)
}

// GetFileInfo is a wrapper around g_zlib_decompressor_get_file_info().
func (decompressor ZlibDecompressor) GetFileInfo() FileInfo {
	ret0 := C.g_zlib_decompressor_get_file_info(decompressor.native())
	return wrapFileInfo(ret0)
}

type AsyncReadyCallback func(source_object gobject.Object, res AsyncResult)
type FileMeasureProgressCallback func(reporting bool, current_size uint64, num_dirs uint64, num_files uint64)
type FileProgressCallback func(current_num_bytes int64, total_num_bytes int64)
type BusType int

const (
	BusTypeStarter BusType = -1
	BusTypeNone            = 0
	BusTypeSystem          = 1
	BusTypeSession         = 2
)

type ConverterResult int

const (
	ConverterResultError     ConverterResult = 0
	ConverterResultConverted                 = 1
	ConverterResultFinished                  = 2
	ConverterResultFlushed                   = 3
)

type CredentialsType int

const (
	CredentialsTypeInvalid             CredentialsType = 0
	CredentialsTypeLinuxUcred                          = 1
	CredentialsTypeFreebsdCmsgcred                     = 2
	CredentialsTypeOpenbsdSockpeercred                 = 3
	CredentialsTypeSolarisUcred                        = 4
	CredentialsTypeNetbsdUnpcbid                       = 5
)

type DBusError int

const (
	DBusErrorFailed                        DBusError = 0
	DBusErrorNoMemory                                = 1
	DBusErrorServiceUnknown                          = 2
	DBusErrorNameHasNoOwner                          = 3
	DBusErrorNoReply                                 = 4
	DBusErrorIoError                                 = 5
	DBusErrorBadAddress                              = 6
	DBusErrorNotSupported                            = 7
	DBusErrorLimitsExceeded                          = 8
	DBusErrorAccessDenied                            = 9
	DBusErrorAuthFailed                              = 10
	DBusErrorNoServer                                = 11
	DBusErrorTimeout                                 = 12
	DBusErrorNoNetwork                               = 13
	DBusErrorAddressInUse                            = 14
	DBusErrorDisconnected                            = 15
	DBusErrorInvalidArgs                             = 16
	DBusErrorFileNotFound                            = 17
	DBusErrorFileExists                              = 18
	DBusErrorUnknownMethod                           = 19
	DBusErrorTimedOut                                = 20
	DBusErrorMatchRuleNotFound                       = 21
	DBusErrorMatchRuleInvalid                        = 22
	DBusErrorSpawnExecFailed                         = 23
	DBusErrorSpawnForkFailed                         = 24
	DBusErrorSpawnChildExited                        = 25
	DBusErrorSpawnChildSignaled                      = 26
	DBusErrorSpawnFailed                             = 27
	DBusErrorSpawnSetupFailed                        = 28
	DBusErrorSpawnConfigInvalid                      = 29
	DBusErrorSpawnServiceInvalid                     = 30
	DBusErrorSpawnServiceNotFound                    = 31
	DBusErrorSpawnPermissionsInvalid                 = 32
	DBusErrorSpawnFileInvalid                        = 33
	DBusErrorSpawnNoMemory                           = 34
	DBusErrorUnixProcessIdUnknown                    = 35
	DBusErrorInvalidSignature                        = 36
	DBusErrorInvalidFileContent                      = 37
	DBusErrorSelinuxSecurityContextUnknown           = 38
	DBusErrorAdtAuditDataUnknown                     = 39
	DBusErrorObjectPathInUse                         = 40
	DBusErrorUnknownObject                           = 41
	DBusErrorUnknownInterface                        = 42
	DBusErrorUnknownProperty                         = 43
	DBusErrorPropertyReadOnly                        = 44
)

type DBusMessageByteOrder int

const (
	DBusMessageByteOrderBigEndian    DBusMessageByteOrder = 66
	DBusMessageByteOrderLittleEndian                      = 108
)

type DBusMessageHeaderField int

const (
	DBusMessageHeaderFieldInvalid     DBusMessageHeaderField = 0
	DBusMessageHeaderFieldPath                               = 1
	DBusMessageHeaderFieldInterface                          = 2
	DBusMessageHeaderFieldMember                             = 3
	DBusMessageHeaderFieldErrorName                          = 4
	DBusMessageHeaderFieldReplySerial                        = 5
	DBusMessageHeaderFieldDestination                        = 6
	DBusMessageHeaderFieldSender                             = 7
	DBusMessageHeaderFieldSignature                          = 8
	DBusMessageHeaderFieldNumUnixFds                         = 9
)

type DBusMessageType int

const (
	DBusMessageTypeInvalid      DBusMessageType = 0
	DBusMessageTypeMethodCall                   = 1
	DBusMessageTypeMethodReturn                 = 2
	DBusMessageTypeError                        = 3
	DBusMessageTypeSignal                       = 4
)

type DataStreamByteOrder int

const (
	DataStreamByteOrderBigEndian    DataStreamByteOrder = 0
	DataStreamByteOrderLittleEndian                     = 1
	DataStreamByteOrderHostEndian                       = 2
)

type DataStreamNewlineType int

const (
	DataStreamNewlineTypeLf   DataStreamNewlineType = 0
	DataStreamNewlineTypeCr                         = 1
	DataStreamNewlineTypeCrLf                       = 2
	DataStreamNewlineTypeAny                        = 3
)

type DriveStartStopType int

const (
	DriveStartStopTypeUnknown   DriveStartStopType = 0
	DriveStartStopTypeShutdown                     = 1
	DriveStartStopTypeNetwork                      = 2
	DriveStartStopTypeMultidisk                    = 3
	DriveStartStopTypePassword                     = 4
)

type EmblemOrigin int

const (
	EmblemOriginUnknown      EmblemOrigin = 0
	EmblemOriginDevice                    = 1
	EmblemOriginLivemetadata              = 2
	EmblemOriginTag                       = 3
)

type FileAttributeStatus int

const (
	FileAttributeStatusUnset        FileAttributeStatus = 0
	FileAttributeStatusSet                              = 1
	FileAttributeStatusErrorSetting                     = 2
)

type FileAttributeType int

const (
	FileAttributeTypeInvalid    FileAttributeType = 0
	FileAttributeTypeString                       = 1
	FileAttributeTypeByteString                   = 2
	FileAttributeTypeBoolean                      = 3
	FileAttributeTypeUint32                       = 4
	FileAttributeTypeInt32                        = 5
	FileAttributeTypeUint64                       = 6
	FileAttributeTypeInt64                        = 7
	FileAttributeTypeObject                       = 8
	FileAttributeTypeStringv                      = 9
)

type FileMonitorEvent int

const (
	FileMonitorEventChanged          FileMonitorEvent = 0
	FileMonitorEventChangesDoneHint                   = 1
	FileMonitorEventDeleted                           = 2
	FileMonitorEventCreated                           = 3
	FileMonitorEventAttributeChanged                  = 4
	FileMonitorEventPreUnmount                        = 5
	FileMonitorEventUnmounted                         = 6
	FileMonitorEventMoved                             = 7
	FileMonitorEventRenamed                           = 8
	FileMonitorEventMovedIn                           = 9
	FileMonitorEventMovedOut                          = 10
)

type FileType int

const (
	FileTypeUnknown      FileType = 0
	FileTypeRegular               = 1
	FileTypeDirectory             = 2
	FileTypeSymbolicLink          = 3
	FileTypeSpecial               = 4
	FileTypeShortcut              = 5
	FileTypeMountable             = 6
)

type FilesystemPreviewType int

const (
	FilesystemPreviewTypeIfAlways FilesystemPreviewType = 0
	FilesystemPreviewTypeIfLocal                        = 1
	FilesystemPreviewTypeNever                          = 2
)

type IOErrorEnum int

const (
	IOErrorEnumFailed             IOErrorEnum = 0
	IOErrorEnumNotFound                       = 1
	IOErrorEnumExists                         = 2
	IOErrorEnumIsDirectory                    = 3
	IOErrorEnumNotDirectory                   = 4
	IOErrorEnumNotEmpty                       = 5
	IOErrorEnumNotRegularFile                 = 6
	IOErrorEnumNotSymbolicLink                = 7
	IOErrorEnumNotMountableFile               = 8
	IOErrorEnumFilenameTooLong                = 9
	IOErrorEnumInvalidFilename                = 10
	IOErrorEnumTooManyLinks                   = 11
	IOErrorEnumNoSpace                        = 12
	IOErrorEnumInvalidArgument                = 13
	IOErrorEnumPermissionDenied               = 14
	IOErrorEnumNotSupported                   = 15
	IOErrorEnumNotMounted                     = 16
	IOErrorEnumAlreadyMounted                 = 17
	IOErrorEnumClosed                         = 18
	IOErrorEnumCancelled                      = 19
	IOErrorEnumPending                        = 20
	IOErrorEnumReadOnly                       = 21
	IOErrorEnumCantCreateBackup               = 22
	IOErrorEnumWrongEtag                      = 23
	IOErrorEnumTimedOut                       = 24
	IOErrorEnumWouldRecurse                   = 25
	IOErrorEnumBusy                           = 26
	IOErrorEnumWouldBlock                     = 27
	IOErrorEnumHostNotFound                   = 28
	IOErrorEnumWouldMerge                     = 29
	IOErrorEnumFailedHandled                  = 30
	IOErrorEnumTooManyOpenFiles               = 31
	IOErrorEnumNotInitialized                 = 32
	IOErrorEnumAddressInUse                   = 33
	IOErrorEnumPartialInput                   = 34
	IOErrorEnumInvalidData                    = 35
	IOErrorEnumDbusError                      = 36
	IOErrorEnumHostUnreachable                = 37
	IOErrorEnumNetworkUnreachable             = 38
	IOErrorEnumConnectionRefused              = 39
	IOErrorEnumProxyFailed                    = 40
	IOErrorEnumProxyAuthFailed                = 41
	IOErrorEnumProxyNeedAuth                  = 42
	IOErrorEnumProxyNotAllowed                = 43
	IOErrorEnumBrokenPipe                     = 44
	IOErrorEnumConnectionClosed               = 44
	IOErrorEnumNotConnected                   = 45
	IOErrorEnumMessageTooLarge                = 46
)

type IOModuleScopeFlags int

const (
	IOModuleScopeFlagsNone            IOModuleScopeFlags = 0
	IOModuleScopeFlagsBlockDuplicates                    = 1
)

type MountOperationResult int

const (
	MountOperationResultHandled   MountOperationResult = 0
	MountOperationResultAborted                        = 1
	MountOperationResultUnhandled                      = 2
)

type NetworkConnectivity int

const (
	NetworkConnectivityLocal   NetworkConnectivity = 1
	NetworkConnectivityLimited                     = 2
	NetworkConnectivityPortal                      = 3
	NetworkConnectivityFull                        = 4
)

type NotificationPriority int

const (
	NotificationPriorityNormal NotificationPriority = 0
	NotificationPriorityLow                         = 1
	NotificationPriorityHigh                        = 2
	NotificationPriorityUrgent                      = 3
)

type PasswordSave int

const (
	PasswordSaveNever       PasswordSave = 0
	PasswordSaveForSession               = 1
	PasswordSavePermanently              = 2
)

type ResolverError int

const (
	ResolverErrorNotFound         ResolverError = 0
	ResolverErrorTemporaryFailure               = 1
	ResolverErrorInternal                       = 2
)

type ResolverRecordType int

const (
	ResolverRecordTypeSrv ResolverRecordType = 1
	ResolverRecordTypeMx                     = 2
	ResolverRecordTypeTxt                    = 3
	ResolverRecordTypeSoa                    = 4
	ResolverRecordTypeNs                     = 5
)

type ResourceError int

const (
	ResourceErrorNotFound ResourceError = 0
	ResourceErrorInternal               = 1
)

type SocketClientEvent int

const (
	SocketClientEventResolving        SocketClientEvent = 0
	SocketClientEventResolved                           = 1
	SocketClientEventConnecting                         = 2
	SocketClientEventConnected                          = 3
	SocketClientEventProxyNegotiating                   = 4
	SocketClientEventProxyNegotiated                    = 5
	SocketClientEventTlsHandshaking                     = 6
	SocketClientEventTlsHandshaked                      = 7
	SocketClientEventComplete                           = 8
)

type SocketFamily int

const (
	SocketFamilyInvalid SocketFamily = 0
	SocketFamilyUnix                 = 1
	SocketFamilyIpv4                 = 2
	SocketFamilyIpv6                 = 10
)

type SocketListenerEvent int

const (
	SocketListenerEventBinding   SocketListenerEvent = 0
	SocketListenerEventBound                         = 1
	SocketListenerEventListening                     = 2
	SocketListenerEventListened                      = 3
)

type SocketProtocol int

const (
	SocketProtocolUnknown SocketProtocol = -1
	SocketProtocolDefault                = 0
	SocketProtocolTcp                    = 6
	SocketProtocolUdp                    = 17
	SocketProtocolSctp                   = 132
)

type SocketType int

const (
	SocketTypeInvalid   SocketType = 0
	SocketTypeStream               = 1
	SocketTypeDatagram             = 2
	SocketTypeSeqpacket            = 3
)

type TlsAuthenticationMode int

const (
	TlsAuthenticationModeNone      TlsAuthenticationMode = 0
	TlsAuthenticationModeRequested                       = 1
	TlsAuthenticationModeRequired                        = 2
)

type TlsCertificateRequestFlags int

const (
	TlsCertificateRequestFlagsNone TlsCertificateRequestFlags = 0
)

type TlsDatabaseLookupFlags int

const (
	TlsDatabaseLookupFlagsNone    TlsDatabaseLookupFlags = 0
	TlsDatabaseLookupFlagsKeypair                        = 1
)

type TlsError int

const (
	TlsErrorUnavailable         TlsError = 0
	TlsErrorMisc                         = 1
	TlsErrorBadCertificate               = 2
	TlsErrorNotTls                       = 3
	TlsErrorHandshake                    = 4
	TlsErrorCertificateRequired          = 5
	TlsErrorEof                          = 6
)

type TlsInteractionResult int

const (
	TlsInteractionResultUnhandled TlsInteractionResult = 0
	TlsInteractionResultHandled                        = 1
	TlsInteractionResultFailed                         = 2
)

type TlsRehandshakeMode int

const (
	TlsRehandshakeModeNever    TlsRehandshakeMode = 0
	TlsRehandshakeModeSafely                      = 1
	TlsRehandshakeModeUnsafely                    = 2
)

type UnixSocketAddressType int

const (
	UnixSocketAddressTypeInvalid        UnixSocketAddressType = 0
	UnixSocketAddressTypeAnonymous                            = 1
	UnixSocketAddressTypePath                                 = 2
	UnixSocketAddressTypeAbstract                             = 3
	UnixSocketAddressTypeAbstractPadded                       = 4
)

type ZlibCompressorFormat int

const (
	ZlibCompressorFormatZlib ZlibCompressorFormat = 0
	ZlibCompressorFormatGzip                      = 1
	ZlibCompressorFormatRaw                       = 2
)

type AppInfoCreateFlags int

const (
	AppInfoCreateFlagsNone                        AppInfoCreateFlags = 0
	AppInfoCreateFlagsNeedsTerminal                                  = 1
	AppInfoCreateFlagsSupportsUris                                   = 2
	AppInfoCreateFlagsSupportsStartupNotification                    = 4
)

type ApplicationFlags int

const (
	ApplicationFlagsFlagsNone          ApplicationFlags = 0
	ApplicationFlagsIsService                           = 1
	ApplicationFlagsIsLauncher                          = 2
	ApplicationFlagsHandlesOpen                         = 4
	ApplicationFlagsHandlesCommandLine                  = 8
	ApplicationFlagsSendEnvironment                     = 16
	ApplicationFlagsNonUnique                           = 32
	ApplicationFlagsCanOverrideAppId                    = 64
)

type AskPasswordFlags int

const (
	AskPasswordFlagsNeedPassword       AskPasswordFlags = 1
	AskPasswordFlagsNeedUsername                        = 2
	AskPasswordFlagsNeedDomain                          = 4
	AskPasswordFlagsSavingSupported                     = 8
	AskPasswordFlagsAnonymousSupported                  = 16
)

type BusNameOwnerFlags int

const (
	BusNameOwnerFlagsNone             BusNameOwnerFlags = 0
	BusNameOwnerFlagsAllowReplacement                   = 1
	BusNameOwnerFlagsReplace                            = 2
)

type BusNameWatcherFlags int

const (
	BusNameWatcherFlagsNone      BusNameWatcherFlags = 0
	BusNameWatcherFlagsAutoStart                     = 1
)

type ConverterFlags int

const (
	ConverterFlagsNone       ConverterFlags = 0
	ConverterFlagsInputAtEnd                = 1
	ConverterFlagsFlush                     = 2
)

type DBusCallFlags int

const (
	DBusCallFlagsNone                          DBusCallFlags = 0
	DBusCallFlagsNoAutoStart                                 = 1
	DBusCallFlagsAllowInteractiveAuthorization               = 2
)

type DBusCapabilityFlags int

const (
	DBusCapabilityFlagsNone          DBusCapabilityFlags = 0
	DBusCapabilityFlagsUnixFdPassing                     = 1
)

type DBusConnectionFlags int

const (
	DBusConnectionFlagsNone                         DBusConnectionFlags = 0
	DBusConnectionFlagsAuthenticationClient                             = 1
	DBusConnectionFlagsAuthenticationServer                             = 2
	DBusConnectionFlagsAuthenticationAllowAnonymous                     = 4
	DBusConnectionFlagsMessageBusConnection                             = 8
	DBusConnectionFlagsDelayMessageProcessing                           = 16
)

type DBusInterfaceSkeletonFlags int

const (
	DBusInterfaceSkeletonFlagsNone                            DBusInterfaceSkeletonFlags = 0
	DBusInterfaceSkeletonFlagsHandleMethodInvocationsInThread                            = 1
)

type DBusMessageFlags int

const (
	DBusMessageFlagsNone                          DBusMessageFlags = 0
	DBusMessageFlagsNoReplyExpected                                = 1
	DBusMessageFlagsNoAutoStart                                    = 2
	DBusMessageFlagsAllowInteractiveAuthorization                  = 4
)

type DBusObjectManagerClientFlags int

const (
	DBusObjectManagerClientFlagsNone           DBusObjectManagerClientFlags = 0
	DBusObjectManagerClientFlagsDoNotAutoStart                              = 1
)

type DBusPropertyInfoFlags int

const (
	DBusPropertyInfoFlagsNone     DBusPropertyInfoFlags = 0
	DBusPropertyInfoFlagsReadable                       = 1
	DBusPropertyInfoFlagsWritable                       = 2
)

type DBusProxyFlags int

const (
	DBusProxyFlagsNone                         DBusProxyFlags = 0
	DBusProxyFlagsDoNotLoadProperties                         = 1
	DBusProxyFlagsDoNotConnectSignals                         = 2
	DBusProxyFlagsDoNotAutoStart                              = 4
	DBusProxyFlagsGetInvalidatedProperties                    = 8
	DBusProxyFlagsDoNotAutoStartAtConstruction                = 16
)

type DBusSendMessageFlags int

const (
	DBusSendMessageFlagsNone           DBusSendMessageFlags = 0
	DBusSendMessageFlagsPreserveSerial                      = 1
)

type DBusServerFlags int

const (
	DBusServerFlagsNone                         DBusServerFlags = 0
	DBusServerFlagsRunInThread                                  = 1
	DBusServerFlagsAuthenticationAllowAnonymous                 = 2
)

type DBusSignalFlags int

const (
	DBusSignalFlagsNone               DBusSignalFlags = 0
	DBusSignalFlagsNoMatchRule                        = 1
	DBusSignalFlagsMatchArg0Namespace                 = 2
	DBusSignalFlagsMatchArg0Path                      = 4
)

type DBusSubtreeFlags int

const (
	DBusSubtreeFlagsNone                        DBusSubtreeFlags = 0
	DBusSubtreeFlagsDispatchToUnenumeratedNodes                  = 1
)

type DriveStartFlags int

const (
	DriveStartFlagsNone DriveStartFlags = 0
)

type FileAttributeInfoFlags int

const (
	FileAttributeInfoFlagsNone          FileAttributeInfoFlags = 0
	FileAttributeInfoFlagsCopyWithFile                         = 1
	FileAttributeInfoFlagsCopyWhenMoved                        = 2
)

type FileCopyFlags int

const (
	FileCopyFlagsNone               FileCopyFlags = 0
	FileCopyFlagsOverwrite                        = 1
	FileCopyFlagsBackup                           = 2
	FileCopyFlagsNofollowSymlinks                 = 4
	FileCopyFlagsAllMetadata                      = 8
	FileCopyFlagsNoFallbackForMove                = 16
	FileCopyFlagsTargetDefaultPerms               = 32
)

type FileCreateFlags int

const (
	FileCreateFlagsNone               FileCreateFlags = 0
	FileCreateFlagsPrivate                            = 1
	FileCreateFlagsReplaceDestination                 = 2
)

type FileMeasureFlags int

const (
	FileMeasureFlagsNone           FileMeasureFlags = 0
	FileMeasureFlagsReportAnyError                  = 2
	FileMeasureFlagsApparentSize                    = 4
	FileMeasureFlagsNoXdev                          = 8
)

type FileMonitorFlags int

const (
	FileMonitorFlagsNone           FileMonitorFlags = 0
	FileMonitorFlagsWatchMounts                     = 1
	FileMonitorFlagsSendMoved                       = 2
	FileMonitorFlagsWatchHardLinks                  = 4
	FileMonitorFlagsWatchMoves                      = 8
)

type FileQueryInfoFlags int

const (
	FileQueryInfoFlagsNone             FileQueryInfoFlags = 0
	FileQueryInfoFlagsNofollowSymlinks                    = 1
)

type IOStreamSpliceFlags int

const (
	IOStreamSpliceFlagsNone         IOStreamSpliceFlags = 0
	IOStreamSpliceFlagsCloseStream1                     = 1
	IOStreamSpliceFlagsCloseStream2                     = 2
	IOStreamSpliceFlagsWaitForBoth                      = 4
)

type MountMountFlags int

const (
	MountMountFlagsNone MountMountFlags = 0
)

type MountUnmountFlags int

const (
	MountUnmountFlagsNone  MountUnmountFlags = 0
	MountUnmountFlagsForce                   = 1
)

type OutputStreamSpliceFlags int

const (
	OutputStreamSpliceFlagsNone        OutputStreamSpliceFlags = 0
	OutputStreamSpliceFlagsCloseSource                         = 1
	OutputStreamSpliceFlagsCloseTarget                         = 2
)

type ResourceFlags int

const (
	ResourceFlagsNone       ResourceFlags = 0
	ResourceFlagsCompressed               = 1
)

type ResourceLookupFlags int

const (
	ResourceLookupFlagsNone ResourceLookupFlags = 0
)

type SettingsBindFlags int

const (
	SettingsBindFlagsDefault       SettingsBindFlags = 0
	SettingsBindFlagsGet                             = 1
	SettingsBindFlagsSet                             = 2
	SettingsBindFlagsNoSensitivity                   = 4
	SettingsBindFlagsGetNoChanges                    = 8
	SettingsBindFlagsInvertBoolean                   = 16
)

type SocketMsgFlags int

const (
	SocketMsgFlagsNone      SocketMsgFlags = 0
	SocketMsgFlagsOob                      = 1
	SocketMsgFlagsPeek                     = 2
	SocketMsgFlagsDontroute                = 4
)

type SubprocessFlags int

const (
	SubprocessFlagsNone          SubprocessFlags = 0
	SubprocessFlagsStdinPipe                     = 1
	SubprocessFlagsStdinInherit                  = 2
	SubprocessFlagsStdoutPipe                    = 4
	SubprocessFlagsStdoutSilence                 = 8
	SubprocessFlagsStderrPipe                    = 16
	SubprocessFlagsStderrSilence                 = 32
	SubprocessFlagsStderrMerge                   = 64
	SubprocessFlagsInheritFds                    = 128
)

type TestDBusFlags int

const (
	TestDBusFlagsNone TestDBusFlags = 0
)

type TlsCertificateFlags int

const (
	TlsCertificateFlagsUnknownCa    TlsCertificateFlags = 1
	TlsCertificateFlagsBadIdentity                      = 2
	TlsCertificateFlagsNotActivated                     = 4
	TlsCertificateFlagsExpired                          = 8
	TlsCertificateFlagsRevoked                          = 16
	TlsCertificateFlagsInsecure                         = 32
	TlsCertificateFlagsGenericError                     = 64
	TlsCertificateFlagsValidateAll                      = 127
)

type TlsDatabaseVerifyFlags int

const (
	TlsDatabaseVerifyFlagsNone TlsDatabaseVerifyFlags = 0
)

type TlsPasswordFlags int

const (
	TlsPasswordFlagsNone      TlsPasswordFlags = 0
	TlsPasswordFlagsRetry                      = 2
	TlsPasswordFlagsManyTries                  = 4
	TlsPasswordFlagsFinalTry                   = 8
)
